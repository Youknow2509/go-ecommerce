package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/consts"
	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/redis/go-redis/v9"
)

type sTicketItem struct {
	// implementation interface here
	r          *database.Queries
	localCache service.ILocalCache
}

func NewTicketItemImpl(r *database.Queries, localCache service.ILocalCache) *sTicketItem {
	return &sTicketItem{
		r:          r,
		localCache: localCache,
	}
}

func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error) {
	fmt.Println("CAL service GetTicketItemById...")
	cacheKey := fmt.Sprintf("ticket_item_%d", ticketId)
	// get in local cache
	localData, ok := s.localCache.Get(ctx, cacheKey)
	if ok {
		// unmarshal data
		var ticketItem model.TicketItemsOutput
		err = json.Unmarshal([]byte(localData.(string)), &ticketItem)
		if err != nil {
			return nil, err
		}
		return &ticketItem, nil
	}
	// get in distributed cache
	data, err := global.Rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
	}
	if data != "" {
		// unmarshal data
		var ticketItem model.TicketItemsOutput
		err = json.Unmarshal([]byte(data), &ticketItem)
		if err != nil {
			return nil, err
		}
		return &ticketItem, nil
	}

	// get data dfrom database
	ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
	if err != nil {
		return out, err
	}

	// set data to distributed cache
	jsonData, _ := json.Marshal(ticketItem)
	timeTTl := time.Duration(consts.TIME_TTL_LOCAL_CACHE) * time.Minute
	err = global.Rdb.Set(ctx, cacheKey, jsonData, timeTTl).Err()
	if err != nil {
		return nil, err
	}

	// set data to local cache
	ok = s.localCache.SetWithTTL(ctx, cacheKey, jsonData)
	if !ok {
		return nil, fmt.Errorf("set local cache failed")
	}

	return &model.TicketItemsOutput{
		TicketId:       int(ticketItem.ID),
		TicketName:     ticketItem.Name,
		StockAvailable: int(ticketItem.StockAvailable),
		StockInitial:   int(ticketItem.StockInitial),
	}, nil
}
