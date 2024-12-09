-- name: CreateSchool :one
INSERT INTO Schools (School_Name, Subdomain, Status)
VALUES (?, ?, ?)
RETURNING School_ID, School_Name, Subdomain, Created_At, Status;

-- name: GetSchoolByID :one
SELECT School_ID, School_Name, Subdomain, Created_At, Status
FROM Schools
WHERE School_ID = ?;

-- name: GetSchoolBySubdomain :one
SELECT School_ID, School_Name, Subdomain, Created_At, Status
FROM Schools
WHERE Subdomain = ?;

-- name: UpdateSchoolStatus :exec
UPDATE Schools
SET Status = ?
WHERE School_ID = ?;

-- name: DeleteSchool :exec
DELETE FROM Schools
WHERE School_ID = ?;

-- name: ListSchools :many
SELECT School_ID, School_Name, Subdomain, Created_At, Status
FROM Schools
ORDER BY Created_At DESC;
