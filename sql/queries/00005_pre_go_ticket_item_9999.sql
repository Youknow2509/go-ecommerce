-- name: GetTicketItemById :one
SELECT id, name, stock_initial, stock_available
FROM ticket_item
WHERE id = ?;