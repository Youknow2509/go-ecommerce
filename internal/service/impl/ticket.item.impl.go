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
	r                *database.Queries
	localCache       service.ILocalCache
	distributedCache service.IRedisCache
}

func NewTicketItemImpl(r *database.Queries, localCache service.ILocalCache, distributedCache service.IRedisCache) *sTicketItem {
	return &sTicketItem{
		r:                r,
		localCache:       localCache,
		distributedCache: distributedCache,
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
		// global.Logger.Info("get data from local cache: %s", zap.String("data", localData.(string)))
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
		// global.Logger.Info("get data from distributed cache: %s", zap.String("data", data))
		return &ticketItem, nil
	}

	// ctxx := context.Background()
	// dataOut, err := s.getTicketItemByIdFromDB(ctxx, ticketId)
	dataOut, err := s.getTicketItemByLock(ctx, ticketId)
	if err != nil {
		global.Logger.Error("get lock failed", zap.Error(err))
		return nil, err
	}
	return dataOut, nil
}

// get ticket from lock database
func (s *sTicketItem) getTicketItemByLock(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error) {
	lockKey := fmt.Sprintf("lock_ticket_item_%d", ticketId)
	cacheKey := fmt.Sprintf("ticket_item_%d", ticketId)
	// lock
	err = s.distributedCache.WithDistributedLock(ctx, lockKey, 5, func(ctx context.Context) error {
		global.Logger.Info("get data from database")
		// get data from database
		ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
		if err != nil {
			return err
		}

		// set data to cache
		res, ok := cache.SetCache(
			ctx,
			cacheKey,
			ticketItem,
			(int64)(consts.TIME_TTL_LOCAL_CACHE),
			service.GetRedisCache(),
			s.localCache,
		)
		if !ok {
			return fmt.Errorf("set cache failed - %s", res)
		}
		// write data to out
		out = &model.TicketItemsOutput{
			TicketId:       int(ticketItem.ID),
			TicketName:     ticketItem.Name,
			StockAvailable: int(ticketItem.StockAvailable),
			StockInitial:   int(ticketItem.StockInitial),
		}
		return nil
	})

	return
}

// get data from database
func (s *sTicketItem) getTicketItemByIdFromDB(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error) {
	cacheKey := fmt.Sprintf("ticket_item_%d", ticketId)
	//
	global.Logger.Info("get data from database")
	// get data from database
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
	// write data to out
	out = &model.TicketItemsOutput{
		TicketId:       int(ticketItem.ID),
		TicketName:     ticketItem.Name,
		StockAvailable: int(ticketItem.StockAvailable),
		StockInitial:   int(ticketItem.StockInitial),
	}
	return out, nil
}
