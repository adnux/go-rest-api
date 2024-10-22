-- name: GetUserByEmail :one
SELECT id, password
FROM users
WHERE email = ?
;

-- name: InsertUser :one
INSERT INTO users (email, password, first_name, last_name)
VALUES (?, ?, ?, ?)
RETURNING id, email, password, first_name, last_name
;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
;

-- name: GetEvents :many
SELECT id, name, description, location, date_time, user_id
FROM events
;

-- name: GetEvent :one
SELECT id, name, description, location, date_time, user_id
FROM events
WHERE id = ?
;

-- name: InsertEvent :one
INSERT INTO events (name, description, date_time, location, user_id)
VALUES (?, ?, ?, ?, ?)
RETURNING id, name, description, location, date_time, user_id

;

-- name: UpdateEvent :exec
UPDATE events
SET name = ?, description = ?, location = ?, date_time = ?, user_id = ?
WHERE id = ?
;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = ?
;

-- name: RegisterUserForEvent :one
INSERT INTO registrations(event_id, user_id)
VALUES (?, ?)
RETURNING id, event_id, user_id, active
;

-- name: CancelRegistration :one
UPDATE registrations
SET active = false
WHERE event_id = ? AND user_id = ? AND active = true
RETURNING id, event_id, user_id, active
;

-- name: GetRegistrations :many
SELECT sqlc.embed(reg), sqlc.embed(ev), u.id, u.email, u.first_name, u.last_name
	FROM registrations reg
	LEFT JOIN users u
		ON u.id = reg.user_id
	LEFT JOIN events ev
		ON reg.event_id = ev.id
	WHERE event_id = ?
		AND ev.user_id = ?
;