package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	mockdb "github.com/nurkenti/furnitureShop/db/mock"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	"github.com/nurkenti/furnitureShop/db/util"
	"github.com/stretchr/testify/require"
)

// Все это нужно чтобы тестироватьс хэш паролем
type eqCreateUserParamsMatcher struct {
	arg      sqlc.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(sqlc.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.PasswordHash)
	if err != nil {
		return false
	}

	e.arg.PasswordHash = arg.PasswordHash
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("mathes arg %v and password %v", e.arg, e.password)
}

func randomUser() sqlc.User {
	return sqlc.User{
		ID:           pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
		FullName:     util.RandomName(),
		Age:          int32(util.RandomAge()),
		Role:         sqlc.NullUserRole{UserRole: "admin", Valid: true},
	}
}

func EqCreateUserParams(arg sqlc.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}
func TestCreatUserAPI(t *testing.T) {
	user := randomUser()
	password := user.PasswordHash
	testCase := []struct {
		name          string
		body          gin.H //Тела user в формате JSON
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":     user.Email,
				"password":  password,
				"full_name": user.FullName,
				"age":       user.Age,
			},
			buildStubs: func(store *mockdb.MockStore) {
				// Создаем параметры с хэш паролем
				arg := sqlc.CreateUserParams{
					ID:       user.ID,
					FullName: user.FullName,
					Email:    user.Email,
				}
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				fmt.Printf("Final Status: %d\n", recorder.Code)
				fmt.Printf("Final Body: %s\n", recorder.Body.String())
				require.Equal(t, http.StatusOK, recorder.Code)
				requestBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "PwToShort",
			body: gin.H{
				"email":     user.Email,
				"password":  "121",
				"full_name": user.FullName,
				"age":       user.Age,
			},
			buildStubs: func(store *mockdb.MockStore) {
				// ✅ При коротком пароле запрос НЕ должен доходить до БД
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(0). // - ноль вызовов
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//Ожидаем 400 Bad Request - клиент отправил плохие данные
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NameToShort",
			body: gin.H{
				"email":     user.Email,
				"password":  "secret",
				"full_name": "Sa",
				"age":       user.Age,
			},
			buildStubs: func(store *mockdb.MockStore) {
				// ✅ При коротком запрос НЕ должен доходить до БД
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(0). // - ноль вызовов
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "EmptyName",
			body: gin.H{
				"email":     user.Email,
				"password":  "secret",
				"full_name": "",
				"age":       user.Age,
			},
			buildStubs: func(store *mockdb.MockStore) {
				// ✅ При коротком запрос НЕ должен доходить до БД
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(0). // - ноль вызовов
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NumberInName",
			body: gin.H{
				"email":     user.Email,
				"password":  "secret",
				"full_name": "1231",
				"age":       user.Age,
			},
			buildStubs: func(store *mockdb.MockStore) {
				// ✅ При коротком запрос НЕ должен доходить до БД
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(0). // - ноль вызовов
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "AgeMinNumber",
			body: gin.H{
				"email":     user.Email,
				"password":  "secret",
				"full_name": "1231",
				"age":       155,
			},
			buildStubs: func(store *mockdb.MockStore) {
				// ✅ При коротком запрос НЕ должен доходить до БД
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(0). // - ноль вызовов
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recording := httptest.NewRecorder()

			data, err := json.Marshal(tc.body) //Конвертируем map в JSON
			require.NoError(t, err)

			url := "/users" // правильный post запрос
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json") //важно

			fmt.Printf("Making post request to: %s\n", url)

			server.router.ServeHTTP(recording, request)
			tc.checkResponse(t, recording)
		})
	}
}

func TestGetUserIDAPI(t *testing.T) {
	user := randomUser()

	testCase := []struct {
		name          string
		userID        string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: user.ID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requestBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name:   "NoteFound",
			userID: user.ID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(sqlc.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:   "InternalError",
			userID: user.ID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(sqlc.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:   "InvalidID",
			userID: "invalid-uuid-format",
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
	}
	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/users/id/%s", tc.userID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func requestBodyMatchUser(t *testing.T, body *bytes.Buffer, user sqlc.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser sqlc.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
