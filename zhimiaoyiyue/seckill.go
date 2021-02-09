package zhimiaoyiyue

import (
	"context"
	"fmt"
	"sync"
	"zmyy_seckill/model"
)

var mutex sync.Mutex

//一次完整的秒杀流程
func (e *ZMYYEngine) SecKill(ctx context.Context, cancel func(), dateDetail model.DateDetail, productId string) {
	picOk := make(chan bool, 1)
	picOk <- true
	verifyOk := make(chan string, 100)
	guidOk := make(chan string, 1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-picOk:
			mutex.Lock()
			//1.获取验证码图片
			fmt.Printf("正在获取验证码图片：%s  %s-%s\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
			path, err := e.GetVerifyPic(dateDetail)
			if err != nil || path == "" {
				mutex.Unlock()
				picOk <- true
			} else {
				verifyOk <- path
			}
		case path := <-verifyOk:
			//2.验证码识别
			fmt.Printf("正在识别验证码图片：%s  %s-%s\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
			guid, _ := e.CaptchaVerify(path)
			//2.1 若识别失败，则重新获取并识别
			if guid == "" {
				mutex.Unlock()
				picOk <- true
			} else { //若识别成功,则提交订单
				guidOk <- guid
			}
		case guid := <-guidOk:
			//3.下单
			fmt.Printf("验证码图片:%s  %s-%s 识别成功，尝试下单中...\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
			ok, _ := e.SaveOrder(dateDetail, productId, guid)
			mutex.Unlock()
			//若未成功下单，则重新走获取验证码流程
			if !ok {
				picOk <- true
			} else { //若成功，则退出Seckill
				wg.Done()
				cancel()
				return
			}

		default:
		}
	}
}
