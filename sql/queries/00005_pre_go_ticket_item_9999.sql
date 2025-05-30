-- name: GetTicketItemById :one
SELECT id, name, stock_initial, stock_available
FROM ticket_item
WHERE id = ?;

-- name: DecreaseTicketV1 :execresult
UPDATE ticket_item
SET stock_available = stock_available - ?
WHERE id = ? AND stock_available >= ?;

-- name: DecreaseTicketV2 :execresult
UPDATE ticket_item
SET stock_available = stock_available - ?
WHERE id = ? AND stock_available = ?;

-- name: CheckTicketItemExists :one
SELECT id
FROM ticket_item
WHERE id = ?;