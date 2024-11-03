-- name: CreateCarro :one
INSERT INTO CARROS (marca, modelo, preco) VALUES ($1, $2, $3) RETURNING *;

-- name: GetCarro :one
SELECT * FROM CARROS WHERE id = $1 LIMIT 1;

-- name: ListCarros :many
SELECT * FROM CARROS ORDER BY marca;

-- name: GetCarroByMarca :many
SELECT * FROM CARROS WHERE marca = $1;