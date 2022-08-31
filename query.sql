-- name: ListServices :many
SELECT
  *
FROM
  services;

-- name: ListServicesByIDs :many
SELECT
  *
FROM
  services
WHERE
  id IN ($1);

-- name: GetService :one
SELECT
  *
FROM
  services
WHERE
  id = $1;

-- name: CreateService :one
INSERT INTO
  services (backend)
VALUES
  ($1) RETURNING *;

-- name: UpdateService :one
UPDATE
  services
SET
  backend = $1
WHERE
  id = $2 RETURNING *;

-- name: DeleteService :exec
DELETE FROM
  services
WHERE
  id = $1;

-- name: ListRoutes :many
SELECT
  *
FROM
  routes;

-- name: ListRoutesForService :many
SELECT
  *
FROM
  routes
WHERE
  service_id = $1;

-- name: GetRoute :one
SELECT
  *
FROM
  routes
WHERE
  id = $1;

-- name: UpdateRoute :one
UPDATE
  routes
SET
  path = $1,
  service_id = $2
WHERE
  id = $3 RETURNING *;

-- name: CreateRoute :one
INSERT INTO
  routes (path, service_id)
VALUES
  ($1, $2) RETURNING *;

-- name: DeleteRoute :exec
DELETE FROM
  routes
WHERE
  id = $1;