package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/roylee0704/gron"
)

// 实现Job接口
// type GreetingJob struct {
// 	Name string
// }

// func (g GreetingJob) Run() {
// 	fmt.Println("Hello ", g.Name)
// }

type ExponentialBackOffSchedule struct {
	last int
}

func (e *ExponentialBackOffSchedule) Next(t time.Time) time.Time {
	interval := time.Duration(math.Pow(2.0, float64(e.last))) * time.Second
	e.last += 1
	return t.Truncate(time.Second).Add(interval)
}

func main() {
	// 1.快速开始
	// var wg sync.WaitGroup
	// wg.Add(1)

	// c := gron.New()
	// c.AddFunc(gron.Every(5*time.Second), func() {
	// 	fmt.Println("runs every 5 seconds.")
	// })
	// c.Start()

	// wg.Wait()

	// 2.时间格式
	// var wg sync.WaitGroup
	// wg.Add(1)

	// c := gron.New()
	// c.AddFunc(gron.Every(1*time.Second), func() {
	// 	fmt.Println("runs every second.")
	// })
	// c.AddFunc(gron.Every(1*time.Minute), func() {
	// 	fmt.Println("runs every minute.")
	// })
	// c.AddFunc(gron.Every(1*time.Hour), func() {
	// 	fmt.Println("runs every hour.")
	// })
	// c.AddFunc(gron.Every(1*xtime.Day), func() {
	// 	fmt.Println("runs every day.")
	// })
	// c.AddFunc(gron.Every(1*xtime.Week), func() {
	// 	fmt.Println("runs every week.")
	// })
	// t, _ := time.ParseDuration("4m10s")
	// c.AddFunc(gron.Every(t), func() {
	// 	fmt.Println("runs every 4 minutes 10 seconds.")
	// })
	// c.Start()

	// wg.Wait()

	// 3.隔时间执行
	// var wg sync.WaitGroup
	// wg.Add(1)

	// c := gron.New()
	// c.AddFunc(gron.Every(1*xtime.Day).At("22:00"), func() {
	//   fmt.Println("runs every second.")
	// })
	// c.Start()

	// wg.Wait()

	// 4.自定义任务
	// var wg sync.WaitGroup
	// wg.Add(1)

	// g1 := GreetingJob{Name: "dj"}
	// g2 := GreetingJob{Name: "dajun"}

	// c := gron.New()
	// c.Add(gron.Every(5*time.Second), g1)
	// c.Add(gron.Every(10*time.Second), g2)
	// c.Start()

	// wg.Wait()

	// 5.自定义时间策略
	var wg sync.WaitGroup
	wg.Add(1)

	c := gron.New()
	c.AddFunc(&ExponentialBackOffSchedule{}, func() {
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), "hello")
	})
	c.Start()

	wg.Wait()

}
