package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/consts"
	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/utils/cache"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
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
		global.Logger.Info("get data from local cache: %s", zap.String("data", localData.(string)))
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
		// set data to local cache
		ok = s.localCache.SetWithTTL(ctx, cacheKey, data)
		if !ok {
			return nil, fmt.Errorf("set local cache failed")
		}
		// unmarshal data
		var ticketItem model.TicketItemsOutput
		err = json.Unmarshal([]byte(data), &ticketItem)
		if err != nil {
			return nil, err
		}
		global.Logger.Info("get data from distributed cache: %s", zap.String("data", data))
		return &ticketItem, nil
	}

	// get data dfrom database
	ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
	if err != nil {
		return out, err
	}

	res, ok := cache.SetCache(
		ctx,
		cacheKey,
		ticketItem,
		(int64)(consts.TIME_TTL_LOCAL_CACHE),
		service.GetRedisCache(),
		s.localCache,
	)
	if !ok {
		return nil, fmt.Errorf("set local cache failed - %s", res)
	}
	
	global.Logger.Info("get data from db: %s", zap.String("data", fmt.Sprintf("%+v", ticketItem)))
	return &model.TicketItemsOutput{
		TicketId:       int(ticketItem.ID),
		TicketName:     ticketItem.Name,
		StockAvailable: int(ticketItem.StockAvailable),
		StockInitial:   int(ticketItem.StockInitial),
	}, nil
}
