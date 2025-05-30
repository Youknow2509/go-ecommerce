package model

// VO: Get ticketItems returns
type TicketItemsOutput struct {
	TicketId       int    `json:"ticket_Id"`
	TicketName     string `json:"ticket_Name"`
	StockAvailable int    `json:"stock_Available"`
	StockInitial   int    `json:"stock_Initial"`
}

// DTO
type TicketItemRequest struct {
	TicketId string `json:"ticket_Id"`
}

// ticket decrease request
type TicketItemDecreaseRequest struct {
	TicketId        int `json:"ticket_Id"`
	Quantity        int `json:"quantity"`
	TicketInventory int `json:"ticket_Inventory"`
}
