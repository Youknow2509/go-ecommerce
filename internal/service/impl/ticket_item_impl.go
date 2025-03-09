package impl

import (
	"context"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/model"
)

type sTicketItem struct {
	// implementation interface here
	r *database.Queries
}

func NewTicketItemImpl(r *database.Queries) *sTicketItem {
	return &sTicketItem{
		r: r,
	}
}

func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error) {
	fmt.Println("CAL service GetTicketItemById...")
	// get data dfrom database
	ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
	if err != nil {
		return out, err
	}
	// mapper

	return &model.TicketItemsOutput{
		TicketId:       int(ticketItem.ID),
		TicketName:     ticketItem.Name,
		StockAvailable: int(ticketItem.StockAvailable),
		StockInitial:   int(ticketItem.StockInitial),
	}, nil
}
