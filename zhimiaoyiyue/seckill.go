package zhimiaoyiyue

import (
	"context"
	"sync"
	"zmyy_seckill/model"
)

//一次完整的秒杀流程
func (e *ZMYYEngine) SecKill(dateDetail model.DateDetail, productId string, wg sync.WaitGroup,
	ctx context.Context, stopGetDetail context.CancelFunc,
	stopSeckill context.CancelFunc, ip ...string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ok, _ := e.Bingo(dateDetail, productId, ctx, ip...)
			if ok {
				stopGetDetail()
				stopSeckill()
				return
			}
		}
	}
}
