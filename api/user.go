package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	"github.com/nurkenti/furnitureShop/db/util"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,contains=@"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required,min=3,max=20,alpha"`
	Age      int32  `json:"age" binding:"required,gte=18,lte=100"`
}

// Убираем пароль чтобы было безопасно
type createUserResponse struct {
	ID        pgtype.UUID       `json:"id"`
	Email     string            `json:"email"`
	FullName  string            `json:"full_name"`
	Age       int32             `json:"age"`
	Role      sqlc.NullUserRole `json:"role"`
	CreatedAt pgtype.Timestamp  `json:"created_at"`
	UpdateAt  pgtype.Timestamp  `json:"update_at"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//Hash pw for security
	hashedPw, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create uuid id
	userUUID := uuid.New()
	pgUUID := pgtype.UUID{}
	copy(pgUUID.Bytes[:], userUUID[:])
	pgUUID.Valid = true

	// Create role
	role := sqlc.NullUserRole{
		UserRole: sqlc.UserRoleAdmin,
		Valid:    true,
	}

	arg := sqlc.CreateUserParams{
		ID:           pgUUID,
		Email:        req.Email,
		PasswordHash: hashedPw,
		FullName:     req.FullName,
		Age:          req.Age,
		Role:         role,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := createUserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Age:       user.Age,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type getUserIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getUserID(ctx *gin.Context) {
	var req getUserIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Printf("🔍 Searching user with ID: %s\n", req.ID) // ← Логируем ID
	userUUID := pgtype.UUID{}
	err := userUUID.Scan(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	user, err := server.store.GetUserByID(ctx, userUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		// для ЛЮБОЙ другой ошибки БД возвращаем 500
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) getUser(ctx *gin.Context) {
	email := ctx.Param("email") // Берем из URL параметра
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	user, err := server.store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type listUsersRequest struct {
	PageID   int `form:"page_id" json:"user"  binding:"required"`
	PageSize int `form:"page_size"  binding:"required,min=5,max=10"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // internet server error
		return
	}
	arg := sqlc.ListUsersParams{
		Limit:  int32(req.PageSize),
		Offset: (int32(req.PageID-1) * int32(req.PageSize)),
	}
	user, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // internet server error
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (server *Server) deleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	err := server.store.DeleteUserByEmail(context.Background(), email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"succses": "User has been delete"})
	}

	ctx.JSON(http.StatusOK, err)

}

type numbs struct {
	A int `json:"number1"`
	B int `json:"number2"`
}
type result struct {
	Plus   int `json:"plus"`
	Minus  int `json:"minus"`
	OwerHi int `json:"*"`
}

func (server *Server) numb(ctx *gin.Context) {
	var req numbs
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	plus := req.A + req.B
	minus := req.A - req.B
	oh := req.A * req.B

	arg := result{
		Plus:   plus,
		Minus:  minus,
		OwerHi: oh,
	}
	ctx.JSON(http.StatusOK, arg)
}
