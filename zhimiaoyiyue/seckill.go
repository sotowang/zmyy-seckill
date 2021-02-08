package zhimiaoyiyue

import (
	"context"
	"fmt"
	"zmyy_seckill/model"
)

//一次完整的秒杀流程
func (e *ZMYYEngine) Seckill(ctx context.Context, cancel func(), dateDetail model.DateDetail, productId string, index string) {
	defer wg.Done()
	picOk := make(chan bool, 1)
	picOk <- true
	verifyOk := make(chan bool, 1)
	guidOk := make(chan string, 1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-picOk:
			//1.获取验证码图片
			fmt.Printf("正在获取验证码图片：%s-%s-%s\n", productId, dateDetail.Date, index)
			err := e.GetVerifyPic(dateDetail.Date, productId, index)
			if err != nil {
				picOk <- true
			} else {
				verifyOk <- true
			}
		case <-verifyOk:
			//2.验证码识别
			fmt.Printf("正在识别验证码图片：%s-%s-%s\n", productId, dateDetail.Date, index)
			guid, _ := e.CaptchaVerify(dateDetail.Date, productId, index)
			//2.1 若识别失败，则重新获取并识别
			if guid == "" {
				picOk <- true
			} else { //若识别成功,则提交订单
				guidOk <- guid
			}
		case guid := <-guidOk:
			//3.下单
			fmt.Printf("尝试下单中...：guid= %s\n", guid)
			ok, _ := e.SaveOrder(dateDetail.Date, productId, guid, dateDetail.Mxid)
			//若未成功下单，则重新走获取验证码流程
			if !ok {
				picOk <- true
			} else { //若成功，则退出Seckill
				cancel()
				return
			}
		default:
		}
	}
}
