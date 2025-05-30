package cron

import (
	"context"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// StockReleaseCron sets up a job to update stock count in Redis every 5 minutes.
// It also runs once immediately at startup.
func StockReleaseCron(cr *cron.Cron, q *database.Queries) error {
	const schedule = "@every 5m"

	// Run immediately
	runStockUpdate(q)

	// Schedule cron job
	_, err := cr.AddFunc(schedule, func() {
		runStockUpdate(q)
	})
	if err != nil {
		global.Logger.Error("Failed to add stock release cron job", zap.Error(err))
		return err
	}

	global.Logger.Info("Stock release cron job set", zap.String("schedule", schedule))
	return nil
}

// runStockUpdate queries stock and updates Redis cache.
func runStockUpdate(q *database.Queries) {
	ctx := context.Background()
	const ttlSeconds = 8 * 60 * 60 // 8 hours

	stockItem, err := q.GetTicketItemById(ctx, 1)
	if err != nil {
		global.Logger.Warn("Failed to get stock from DB", zap.Error(err))
		return
	}

	err = global.Rdb.Set(ctx, "stock_release_count", stockItem.StockAvailable, time.Duration(ttlSeconds)*time.Second).Err()
	if err != nil {
		global.Logger.Warn("Failed to set stock in Redis", zap.Error(err))
		return
	}

	global.Logger.Info("Stock count cached", zap.Int("stock_release_count", int(stockItem.StockAvailable)))
}
