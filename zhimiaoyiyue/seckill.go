package zhimiaoyiyue

import (
	"context"
	"fmt"
	"sync"
	"zmyy_seckill/consts"
	"zmyy_seckill/model"
)

var mutex sync.Mutex

//一次完整的秒杀流程
func (e *ZMYYEngine) SecKill(ctx context.Context, cancel func(), dateDetail model.DateDetail, productId string) {
	defer wg.Done()
	//picOk := make(chan bool, 1)
	//picOk <- true
	//verifyOk := make(chan string, 100)
	//guidOk := make(chan string, 1)
	//获取验证码
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if consts.Stop {
				return
			}
			mutex.Lock()
			fmt.Printf("正在获取验证码图片：%s  %s-%s\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
			path, _ := e.GetVerifyPic(dateDetail)
			if path != "" {
				wg2.Add(1)
				//识别验证码
				fmt.Printf("正在识别验证码图片：%s  %s-%s\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
				guid, _ := e.CaptchaVerify(path)
				if guid != "" {
					//下单
					fmt.Printf("验证码图片:%s  %s-%s 识别成功，尝试下单中...\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
					ok, _ := e.SaveOrder(dateDetail, productId, guid)
					if ok {
						consts.Stop = true
						cancel()
						mutex.Unlock()
						wg2.Done()
						return
					} else {
						wg2.Done()
					}
				} else {
					wg2.Done()
				}
			}
			mutex.Unlock()
		}

	}
}
