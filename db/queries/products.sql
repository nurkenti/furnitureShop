/* Chair */

-- name: CreateChair :one
INSERT INTO chair(
  id,model, material, price
)VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetChair :one
SELECT * FROM chair 
WHERE id = $1 LIMIT 1;

-- name: GetChairByModel :one
SELECT * FROM chair
WHERE model = $1 LIMIT 1;


-- name: DeleteChair :exec
DELETE FROM chair
WHERE id = $1;

-- name: UpdateChair :one
UPDATE chair
SET
  model = $2,
  material = $3,
  price = $4
WHERE id = $1
RETURNING *;

-- name: ListChairs :many
SELECT * FROM chair
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;


/* Wardrobe*/

-- name: CreateWardrobe :one
INSERT INTO wardrobe(
  id, model, material, price
)Values($1,$2,$3,$4)
RETURNING *;

-- name: GetWardrobe :one
SELECT * FROM wardrobe
WHERE id = $1 LIMIT 1; 

-- name: GetWardrobeByModel :one
SELECT * From wardrobe
where model = $1 LIMIT 1;

-- name: DeleteWardrobe :exec
DELETE FROM wardrobe
WHERE id = $1; 

-- name: UpdateWardrobe :one
UPDATE wardrobe
SET
  model = $2,
  material = $3,
  price=  $4
Where id = $1
RETURNING *;


-- name: ListWardrobe :many
SELECT * FROM wardrobe
ORDER By created_at DESC
LIMIT $1 OFFSET $2;

/* Warehouse*/

-- name: GetWarhouse :one
SELECT 
   product_model, 
   product_type, 
   quantity, 
   updated_at 
FROM warehouse
Where product_model = $1 LIMIT 1;

-- name: UpdateWarehouseQuantity :exec
UPDATE warehouse
SET quantity = $2, updated_at = now()
WHERE product_model = $1; 

-- name: ListWarehouse :many
SELECT * FROM warehouse
ORDER BY product_type, product_model;



