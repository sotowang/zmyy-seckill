package limit

import (
	"sync"
	"time"
)

type LR struct {
	LRMap map[string]*LimitRate
}

func (lr *LR) SetRate(r int, second int) {
	for _, l := range lr.LRMap {
		l.SetRate(r, second)
	}
}

//LimitRate 限速
type LimitRate struct {
	rate       int
	interval   time.Duration
	lastAction time.Time
	lock       sync.Mutex
}

func (l *LimitRate) Limit() bool {
	result := false
	for {
		l.lock.Lock()
		//判断最后一次执行的时间与当前的时间间隔是否大于限速速率
		if time.Now().Sub(l.lastAction) > l.interval {
			l.lastAction = time.Now()
			result = true
		}
		l.lock.Unlock()
		if result {
			return result
		}
		time.Sleep(l.interval)
	}
}

//SetRate 设置Rate
func (l *LimitRate) SetRate(r int, second int) {
	l.rate = r
	l.interval = time.Microsecond * time.Duration(second*100*1000/l.rate)
}

//GetRate 获取Rate
func (l *LimitRate) GetRate() int {
	return l.rate
}
