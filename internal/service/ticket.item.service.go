package service

import (
	"context"
	"github.com/Youknow2509/go-ecommerce/internal/model"

)

type (
	ITicketHome interface {
	}
	ITicketItem interface {
		GetTicketItemById(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error)
	}
)

var (
	localTicketItem ITicketItem
	localTicketHome ITicketHome
)

func TicketHome() ITicketHome {
	if localTicketHome == nil {
		panic("implement localTicketHome not found for interface ITicketHome")
	}

	return localTicketHome
}

func InitTicketHome(i ITicketHome) {
	localTicketHome = i
}

func TicketItem() ITicketItem {
	if localTicketItem == nil {
		panic("implement localTicketItem not found for interface ITicketItem")
	}

	return localTicketItem
}

func InitTicketItem(i ITicketItem) {
	localTicketItem = i
}
