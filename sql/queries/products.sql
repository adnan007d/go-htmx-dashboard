-- name: CreateProduct :one
INSERT INTO products (id, name, description, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING *;

-- name: UpdatedProduct :one
UPDATE  products SET name = $2, description = $3, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: GetAllProducts :many
SELECT * FROM products LIMIT $1 OFFSET $2;

-- name: GetAProduct :one
SELECT * FROM products WHERE id = $1;

-- name: GetProductsCount :one
SELECT COUNT(*) FROM products;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;
