package zhimiaoyiyue

import (
	"zmyy_seckill/model"
)

//一次完整的秒杀流程
func (e *ZMYYEngine) SecKill(dateDetail model.DateDetail, productId string) {
	defer wg.Done()
	for {
		if stop {
			return
		}
		mutex.Lock()
		flag = true
		ok, _ := e.Bingo(dateDetail, productId)
		if ok {
			stop = true
			flag = false
			mutex.Unlock()
			return
		}
		flag = false
		mutex.Unlock()
	}
}
