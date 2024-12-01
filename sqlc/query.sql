-- name: CreateSchool :exec
INSERT INTO Schools (School_Name, Subdomain, Status)
VALUES (?, ?, ?);

-- name: GetCreatedSchool :one
SELECT School_ID, School_Name, Subdomain, Created_At, Status
FROM Schools
WHERE School_ID = last_insert_rowid();

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

-- name: CreateDatabaseDetail :exec
INSERT INTO DatabaseDetails (School_ID, Database_Name, Connection_URI, Region)
VALUES (?, ?, ?, ?);

-- name: GetCreatedDatabaseDetail :one
SELECT Database_ID, School_ID, Database_Name, Connection_URI, Region, Created_At
FROM DatabaseDetails
WHERE Database_ID = last_insert_rowid();

-- name: GetDatabaseDetailByID :one
SELECT Database_ID, School_ID, Database_Name, Connection_URI, Region, Created_At
FROM DatabaseDetails
WHERE Database_ID = ?;

-- name: GetDatabaseDetailsBySchool :many
SELECT Database_ID, School_ID, Database_Name, Connection_URI, Region, Created_At
FROM DatabaseDetails
WHERE School_ID = ?
ORDER BY Created_At DESC;

-- name: UpdateDatabaseRegion :exec
UPDATE DatabaseDetails
SET Region = ?
WHERE Database_ID = ?;

-- name: DeleteDatabaseDetail :exec
DELETE FROM DatabaseDetails
WHERE Database_ID = ?;
