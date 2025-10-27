package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	"github.com/nurkenti/furnitureShop/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomChair(t *testing.T) sqlc.Chair {
	m := util.RandomModel("sonyx", "kurumi")
	mater := util.RandomMaterial("wood", "metal", "fabric")
	/*materials := []sqlc.ChairMaterial{
		sqlc.ChairMaterialWood,
		sqlc.ChairMaterialMetal,
		sqlc.ChairMaterialFabric,
	}*/
	//randomMat := materials[rand.Intn(len(materials))]
	arg := sqlc.CreateChairParams{
		ID:       pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Model:    sqlc.ChairModel(m),
		Material: sqlc.NullChairMaterial{ChairMaterial: sqlc.ChairMaterial(mater), Valid: true},
		Price:    pgtype.Float8{Float64: 5000, Valid: true},
	}
	chair, err := testQueries.CreateChair(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, chair)

	require.Equal(t, arg.ID, chair.ID)
	require.Equal(t, arg.Model, chair.Model)
	require.Equal(t, arg.Material, chair.Material)
	require.Equal(t, arg.Price, chair.Price)
	require.NotZero(t, chair.ID)
	require.NotZero(t, chair.CreatedAt)

	return chair
}

func TestCreatChair(t *testing.T) {
	createRandomChair(t)
}

func TestGetChair(t *testing.T) {
	chair1 := createRandomChair(t)

	chair2, err := testQueries.GetChair(context.Background(), chair1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, chair2)

	require.Equal(t, chair1.ID, chair2.ID)
	require.Equal(t, chair1.Model, chair2.Model)
	require.Equal(t, chair1.Material, chair2.Material)
	require.Equal(t, chair1.Price, chair2.Price)
	require.WithinDuration(t, chair1.CreatedAt.Time, chair2.CreatedAt.Time, time.Second)
}

func TestGetChairByModel(t *testing.T) {
	// Создаем стул с конкретными данными

	chair1 := createRandomChair(t)
	chair2, err := testQueries.GetChairByModel(context.Background(), chair1.Model)
	require.NoError(t, err)
	require.NotEmpty(t, chair2)

	require.Equal(t, chair1.Model, chair2.Model)
	require.Equal(t, chair1.Material, chair2.Material)
	require.Equal(t, chair1.Price, chair2.Price)

}

func TestListChair(t *testing.T) {
	arg := sqlc.ListChairsParams{
		Limit:  5,
		Offset: 0,
	}
	chairs, err := testQueries.ListChairs(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, chairs, 5)
	for _, chair := range chairs {
		require.NotEmpty(t, chair)
	}
}

func TestDeleteChair(t *testing.T) {
	chair1 := createRandomChair(t)

	err := testQueries.DeleteChair(context.Background(), chair1.ID)
	require.NoError(t, err)

	chair2, err := testQueries.GetChair(context.Background(), chair1.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, pgx.ErrNoRows)
	require.Empty(t, chair2)
}

func TestUpdateChair(t *testing.T) {
	chair1 := createRandomChair(t)
	arg := sqlc.UpdateChairParams{
		ID:       chair1.ID,
		Model:    chair1.Model,
		Material: chair1.Material,
		Price:    chair1.Price,
	}
	chair2, err := testQueries.UpdateChair(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, chair2)

	require.Equal(t, chair1.ID, chair2.ID)
	require.Equal(t, chair1.Model, chair2.Model)
	require.Equal(t, chair1.Material, chair2.Material)
	require.Equal(t, chair1.Price, chair2.Price)
	require.WithinDuration(t, chair1.CreatedAt.Time, chair2.CreatedAt.Time, time.Second)
}

// Waredrobe
func createRandomWardrobe(t *testing.T) sqlc.Wardrobe {
	m := util.RandomModel("unibi", "facito")
	mater := util.RandomMaterial("mdf", "dsp", "dsp")
	/*materials := []sqlc.ChairMaterial{
		sqlc.ChairMaterialWood,
		sqlc.ChairMaterialMetal,
		sqlc.ChairMaterialFabric,
	}*/
	//randomMat := materials[rand.Intn(len(materials))]
	arg := sqlc.CreateWardrobeParams{
		ID:       pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Model:    sqlc.WardrobeModel(m),
		Material: sqlc.NullWardrobeMaterial{WardrobeMaterial: sqlc.WardrobeMaterial(mater), Valid: true},
		Price:    float64(5000),
	}
	wardrobe, err := testQueries.CreateWardrobe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wardrobe)

	require.Equal(t, arg.ID, wardrobe.ID)
	require.Equal(t, arg.Model, wardrobe.Model)
	require.Equal(t, arg.Material, wardrobe.Material)
	require.Equal(t, arg.Price, wardrobe.Price)
	require.NotZero(t, wardrobe.ID)
	require.NotZero(t, wardrobe.CreatedAt)

	return wardrobe
}

func TestCreatWardrobe(t *testing.T) {
	createRandomWardrobe(t)
}

func TestGetWardrobe(t *testing.T) {
	wardrobe1 := createRandomWardrobe(t)

	wardrobe2, err := testQueries.GetWardrobe(context.Background(), wardrobe1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, wardrobe2)

	require.Equal(t, wardrobe1.ID, wardrobe2.ID)
	require.Equal(t, wardrobe1.Model, wardrobe2.Model)
	require.Equal(t, wardrobe1.Material, wardrobe2.Material)
	require.Equal(t, wardrobe1.Price, wardrobe2.Price)
	require.WithinDuration(t, wardrobe1.CreatedAt.Time, wardrobe2.CreatedAt.Time, time.Second)
}

func TestGetWardrobeByModel(t *testing.T) {
	wardrobe1 := createRandomWardrobe(t)
	wardrobe2, err := testQueries.GetWardrobeByModel(context.Background(), wardrobe1.Model)
	require.NoError(t, err)
	require.NotEmpty(t, wardrobe2)

	require.Equal(t, wardrobe1.Model, wardrobe2.Model)
	require.Equal(t, wardrobe1.Price, wardrobe2.Price)

}

func TestListWardrobe(t *testing.T) {
	// Сначала создаем 5 записей
	for i := 0; i < 5; i++ {
		createRandomChair(t)
	}

	arg := sqlc.ListWardrobeParams{
		Limit:  5,
		Offset: 0,
	}
	wardrobs, err := testQueries.ListWardrobe(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, wardrobs, 5)
	for _, wardrobe := range wardrobs {
		require.NotEmpty(t, wardrobe)
	}
}

func TestDeleteWardrobe(t *testing.T) {
	wardrobe1 := createRandomWardrobe(t)

	err := testQueries.DeleteWardrobe(context.Background(), wardrobe1.ID)
	require.NoError(t, err)

	wardrobe2, err := testQueries.GetChair(context.Background(), wardrobe1.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, pgx.ErrNoRows)
	require.Empty(t, wardrobe2)
}

func TestUpdateWardrobe(t *testing.T) {
	wardrobe1 := createRandomWardrobe(t)
	arg := sqlc.UpdateWardrobeParams{
		ID:       wardrobe1.ID,
		Model:    wardrobe1.Model,
		Material: wardrobe1.Material,
		Price:    wardrobe1.Price,
	}
	wardrobe2, err := testQueries.UpdateWardrobe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wardrobe2)

	require.Equal(t, wardrobe1.ID, wardrobe2.ID)
	require.Equal(t, wardrobe1.Model, wardrobe2.Model)
	require.Equal(t, wardrobe1.Material, wardrobe2.Material)
	require.Equal(t, wardrobe1.Price, wardrobe2.Price)
	require.WithinDuration(t, wardrobe1.CreatedAt.Time, wardrobe2.CreatedAt.Time, time.Second)
}

func TestGetWarehouse(t *testing.T) {
	// Просто сделали без creat.
	productModels := []string{"sonyx", "kurumi", "unibi", "facito"}

	for _, model := range productModels {
		t.Run(model, func(t *testing.T) {
			warehouse, err := testQueries.GetWarhouse(context.Background(), model)
			require.NoError(t, err)
			require.NotEmpty(t, warehouse)
			require.Equal(t, model, warehouse.ProductModel)
		})
	}
}

func TestUpdateWarehouse(t *testing.T) {
	warehouse, err := testQueries.GetWarhouse(context.Background(), "sonyx")
	require.NoError(t, err)

	newQuantity := int32(25)
	err = testQueries.UpdateWarehouseQuantity(context.Background(), sqlc.UpdateWarehouseQuantityParams{
		ProductModel: "sonyx",
		Quantity:     newQuantity,
	})
	require.NoError(t, err)

	warehouse2, err := testQueries.GetWarhouse(context.Background(), "sonyx")
	require.NoError(t, err)
	require.NotEmpty(t, warehouse2)
	require.Equal(t, newQuantity, warehouse2.Quantity)
	require.NotEqual(t, warehouse.Quantity, warehouse2.Quantity)
}
