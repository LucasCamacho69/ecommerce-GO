-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProcutsByID :many
SELECT * FROM products WHERE id = $1;