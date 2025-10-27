package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	user "github.com/nurkenti/furnitureShop/internal/service/user"
	"github.com/nurkenti/furnitureShop/menu"
)

func AddChair(q *sqlc.Queries) error {
	var p float64
	model, _, err := menu.NewMenuTemplate("Select model:", []string{"sonyx", "kurumi"})
	material, i, err := menu.NewMenuTemplate("Select material:", []string{"wood", "metal", "fabric"})
	if model == "sonyx" {
		p = 5000
	} else {
		p = 10000
	}
	switch i {
	case 0:
		p = p + 250
	case 1:
		p = p + 2150
	case 2:
		p = p + 3230
	}

	arg := sqlc.CreateChairParams{
		ID:       pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Model:    sqlc.ChairModel(model),
		Material: sqlc.NullChairMaterial{ChairMaterial: sqlc.ChairMaterial(material), Valid: true},
		Price:    pgtype.Float8{Float64: p, Valid: true},
	}
	chair, err := q.CreateChair(context.Background(), arg)
	if err != nil {
		return err
	}
	FormatProduct(chair)

	return err
}

func GetChair(q *sqlc.Queries) error {
	ans, err := user.AddInfo("Search Chair \nChair ID:")
	if err != nil {
		return err
	}

	parsedUUID, err := uuid.Parse(ans)
	if err != nil {
		return err
	}
	chair, err := q.GetChair(context.Background(), pgtype.UUID{Bytes: parsedUUID, Valid: true})
	if err != nil {
		return err
	}
	FormatProduct(chair)
	return nil
}

func DeleteChair(q *sqlc.Queries) error {
	ans, err := user.AddInfo("Delete Chair \nChair ID: ")
	if err != nil {
		return err
	}
	parseUUID, err := uuid.Parse(ans)
	if err != nil {
		return err
	}
	chair, err := q.GetChair(context.Background(), pgtype.UUID{Bytes: parseUUID, Valid: true})
	if err != nil {
		return err
	}
	FormatProduct(chair)

	err = q.DeleteChair(context.Background(), pgtype.UUID{Bytes: parseUUID, Valid: true})
	if err != nil {
		return err
	}
	return nil
}

func ListChair(q *sqlc.Queries) error {
	arg := sqlc.ListChairsParams{
		Limit:  5,
		Offset: 0,
	}
	chairs, err := q.ListChairs(context.Background(), arg)
	if err != nil {
		return err
	}
	for _, chair := range chairs {
		fmt.Println()
		FormatProduct(chair)
	}
	return nil
}

func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func UpdateChair(q *sqlc.Queries) (err error) {
	// Обрабатываем панику
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	ansID := must(user.AddInfo("Update chair \nChair ID: "))
	parseUUID := must(uuid.Parse(ansID))

	oldChair := must(q.GetChair(context.Background(), pgtype.UUID{Bytes: parseUUID, Valid: true}))
	FormatProduct(oldChair)

	arg := sqlc.UpdateChairParams{
		ID:       oldChair.ID,
		Model:    oldChair.Model,
		Material: oldChair.Material,
		Price:    oldChair.Price,
	}

	for {
		menu.ClearInputBuffer()
		_, i, err := menu.NewMenuTemplate("Select update:", []string{"ID", "Model", "Material", "Price", "end"})
		if err != nil {
			return err
		}
		if i == 4 {
			break
		}
		switch i {
		case 0:
			ansNew := must(user.AddInfo("New ID: "))
			parseUUID := must(uuid.Parse(ansNew))
			arg.ID = pgtype.UUID{Bytes: parseUUID, Valid: true}
			continue

		case 1:
			ansNew := must(user.AddInfo("New Model: "))
			arg.Model = sqlc.ChairModel(ansNew)
			continue

		case 2:
			ansMaterial := must(user.AddInfo("New Material: "))
			arg.Material = sqlc.NullChairMaterial{ChairMaterial: sqlc.ChairMaterial(ansMaterial), Valid: true}
			continue

		case 3:
			ansPrice := must(user.AddInfo("New Price: "))

			price := must(strconv.ParseFloat(ansPrice, 64))
			arg.Price = pgtype.Float8{Float64: price, Valid: true}
			continue
		}

	}
	NewChair := must(q.UpdateChair(context.Background(), arg))
	fmt.Println("Result Update:")
	FormatProduct(NewChair)

	return nil
}

func FormatProduct(chair sqlc.Chair) {
	fmt.Printf("Chair: %s\n", uuid.UUID(chair.ID.Bytes))
	fmt.Printf("   Model: %s \n   Material: %s\n", chair.Model, chair.Material.ChairMaterial)
	fmt.Printf("   Price: %d\n", int(chair.Price.Float64))
	fmt.Printf("   Создан: %s\n", chair.CreatedAt.Time.Format("2006-01-02 15:04:05"))
}
