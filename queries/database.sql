-- name: CreateDatabaseDetail :one
INSERT INTO DatabaseDetails (Foreign_Database_ID, School_ID, Database_Name, Connection_URI)
VALUES (?, ?, ?, ?)
RETURNING Foreign_Database_ID, Database_ID, School_ID, Database_Name, Connection_URI, Created_At;

-- name: GetDatabaseDetailByID :one
SELECT Database_ID, Foreign_Database_ID, School_ID, Database_Name, Connection_URI, Created_At
FROM DatabaseDetails
WHERE Database_ID = ?;

-- name: GetDatabaseDetailsBySchool :many
SELECT Database_ID, Foreign_Database_ID, School_ID, Database_Name, Connection_URI, Created_At
FROM DatabaseDetails
WHERE School_ID = ?
ORDER BY Created_At DESC;

-- name: GetDatabaseDetailsByForeignID :many
SELECT Database_ID, Foreign_Database_ID, School_ID, Database_Name, Connection_URI, Created_At
FROM DatabaseDetails
WHERE Foreign_Database_ID = ?;

-- name: DeleteDatabaseDetail :exec
DELETE FROM DatabaseDetails
WHERE Database_ID = ?;
