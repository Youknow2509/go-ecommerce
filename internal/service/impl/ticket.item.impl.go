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
	"github.com/Youknow2509/go-ecommerce/internal/utils/cache"
	"github.com/Youknow2509/go-ecommerce/pkg/redis/script"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type sTicketItem struct {
	// implementation interface here
	r                *database.Queries
	localCache       service.ILocalCache
	distributedCache service.IRedisCache
}

// DecreaseTicketItemRelease implements service.ITicketItem.
func (s *sTicketItem) DecreaseTicketItemRelease(ctx context.Context, ticketId int, quantity int) (err error) {
	/**
	 * use distributed cache to decrease ticket item after change database
	 */
	key := fmt.Sprintf("ticket_%d_release_count", ticketId)
	// lua script to decrease ticket item
	keysLua := []string{key}
	argsLua := []string{fmt.Sprintf("%d", quantity)}
	rLua, err := global.Rdb.Eval(ctx, script.DecreaseTicketItemLua, keysLua, argsLua).Result()
	if err != nil {
		global.Logger.Error("decrease ticket item release failed", zap.Error(err))
		return err
	}
	// fmt.Println("decrease ticket item release result:", rLua)
	if 0 == int(rLua.(int64)) {
		global.Logger.Warn("ticket item release not found or stock is not enough", zap.Int("ticketId", ticketId), zap.Int("quantity", quantity))
		return fmt.Errorf("ticket item release not found or stock is not enough")
	}

	// change data db with distributed lock
	// if err := s.decreaseTicketItemReleaseByLock(ctx, ticketId, quantity); err != nil {
	// 	global.Logger.Error("decrease ticket item release by lock failed", zap.Error(err))
	// 	return err
	// }

	// decrease ticket item release from database
	res, err := s.r.DecreaseTicketV1(ctx, database.DecreaseTicketV1Params{
		ID:               int64(ticketId),
		StockAvailable:   int32(quantity),
		StockAvailable_2: int32(quantity),
	})
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		global.Logger.Error("get rows affected failed", zap.Error(err))
		return fmt.Errorf("get rows affected failed: %w", err)
	}
	if rowsAffected == 0 {
		global.Logger.Warn("ticket item release not found or stock is not enough", zap.Int("ticketId", ticketId), zap.Int("quantity", quantity))
		return fmt.Errorf("ticket item release not found or stock is not enough")
	}

	return nil
}

// DecreaseTicketItem implements service.ITicketItem.
func (s *sTicketItem) DecreaseTicketItem(ctx context.Context, ticketId int, ticketInventory int, quantity int) (err error) {
	// check ticket after decrease
	ticketAfter := ticketInventory - quantity
	if ticketAfter < 0 {
		return fmt.Errorf("ticket inventory is not enough, current: %d, decrease: %d", ticketInventory, quantity)
	}
	// distribute lock key
	err = s.decreaseTicketItemByLock(ctx, ticketId, ticketInventory, quantity)
	if err != nil {
		global.Logger.Error("decrease ticket item failed", zap.Error(err))
		return err
	}
	cacheKey := fmt.Sprintf("ticket_item_%d", ticketId)
	dataJson, _ := json.Marshal(model.TicketItemsOutput{
		TicketId:       ticketId,
		StockAvailable: ticketAfter,
		StockInitial:   ticketInventory,
	})
	// update distributed cache
	err = global.Rdb.Set(ctx, cacheKey, dataJson, time.Duration(consts.TIME_TTL_LOCAL_CACHE)*time.Minute).Err()
	if err != nil {
		global.Logger.Error("set distributed cache failed", zap.Error(err))
		return fmt.Errorf("set distributed cache failed: %w", err)
	}
	// update local cache inventory ticket
	if s.localCache.Set(ctx, cacheKey, string(dataJson)) {
		global.Logger.Warn("set local cache failed", zap.String("cacheKey", cacheKey))
	}
	return nil
}

// var mu sync.Mutex
func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error) {
	fmt.Println("CAL service GetTicketItemById...")
	cacheKey := fmt.Sprintf("ticket_item_%d", ticketId)

	// mu.Lock()
	// defer mu.Unlock()

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
	// dataOut, err := s.getTicketItemByIdFromDB(ctx, ticketId)
	dataOut, err := s.getTicketItemByLock(ctx, ticketId)
	if err != nil {
		global.Logger.Error("get lock failed", zap.Error(err))
		return nil, err
	}
	// set data to distributed cache
	dataJson, _ := json.Marshal(dataOut)
	_, err = global.Rdb.Set(ctx, cacheKey, dataJson, time.Duration(consts.TIME_TTL_LOCAL_CACHE)*time.Minute).Result()
	if err != nil {
		global.Logger.Error("set distributed cache failed", zap.Error(err))
		return nil, fmt.Errorf("set distributed cache failed: %w", err)
	}
	// set data to local cache
	ok = s.localCache.SetWithTTL(ctx, cacheKey, string(dataJson))
	if !ok {
		global.Logger.Warn("set local cache failed", zap.String("cacheKey", cacheKey))
		return nil, fmt.Errorf("set local cache failed")
	}

	return dataOut, nil
}

func NewTicketItemImpl(r *database.Queries, localCache service.ILocalCache, distributedCache service.IRedisCache) *sTicketItem {
	return &sTicketItem{
		r:                r,
		localCache:       localCache,
		distributedCache: distributedCache,
	}
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

// decrease ticket item lock database
func (s *sTicketItem) decreaseTicketItemByLock(ctx context.Context, ticketId int, ticketInventory int, quantity int) (err error) {
	lockKey := fmt.Sprintf("lock_ticket_item_decrease_%d", ticketId)
	// lock
	err = s.distributedCache.WithDistributedLock(ctx, lockKey, 2, func(ctx context.Context) error {
		global.Logger.Info("decrease ticket item from database")
		// decrease ticket item from database
		res, err := s.r.DecreaseTicketV2(ctx, database.DecreaseTicketV2Params{
			ID:               int64(ticketId),
			StockAvailable:   int32(quantity),
			StockAvailable_2: int32(ticketInventory),
		})
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			global.Logger.Error("get rows affected failed", zap.Error(err))
			return fmt.Errorf("get rows affected failed: %w", err)
		}
		if rowsAffected == 0 {
			global.Logger.Warn("ticket item not found or stock is not enough", zap.Int("ticketId", ticketId), zap.Int("quantity", quantity))
			return fmt.Errorf("ticket item not found or stock is not enough")
		}
		return nil
	})
	if err != nil {
		global.Logger.Error("decrease ticket item failed", zap.Error(err))
		return err
	}
	return nil
}

// decrease ticket release item lock database
func (s *sTicketItem) decreaseTicketItemReleaseByLock(ctx context.Context, ticketId int, quantity int) (err error) {
	lockKey := fmt.Sprintf("lock_ticket_item_release_%d", ticketId)
	// lock
	err = s.distributedCache.WithDistributedLock(ctx, lockKey, 2, func(ctx context.Context) error {
		global.Logger.Info("decrease ticket item release from database")
		// decrease ticket item release from database
		res, err := s.r.DecreaseTicketV1(ctx, database.DecreaseTicketV1Params{
			ID:               int64(ticketId),
			StockAvailable:   int32(quantity),
			StockAvailable_2: int32(quantity),
		})
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			global.Logger.Error("get rows affected failed", zap.Error(err))
			return fmt.Errorf("get rows affected failed: %w", err)
		}
		if rowsAffected == 0 {
			global.Logger.Warn("ticket item release not found or stock is not enough", zap.Int("ticketId", ticketId), zap.Int("quantity", quantity))
			return fmt.Errorf("ticket item release not found or stock is not enough")
		}
		return nil
	})
	if err != nil {
		global.Logger.Error("decrease ticket item release failed", zap.Error(err))
		return err
	}
	return nil
}
