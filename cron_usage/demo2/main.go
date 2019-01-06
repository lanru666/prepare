package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type Cronjob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	// 需要一个调度协程 它定时检查所有的Cron任务，谁过期了就执行谁
	// 我们定义两个cronjob
	var (
		cronjob       *Cronjob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*Cronjob
	)
	scheduleTable = make(map[string]*Cronjob)
	// 当前时间
	now = time.Now()
	
	expr = cronexpr.MustParse("*/5 * * * * * * ")
	cronjob = &Cronjob{
		expr,
		expr.Next(now),
	}
	
	//任务注册到调度表
	scheduleTable["job1"] = cronjob
	
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronjob = &Cronjob{
		expr,
		expr.Next(now),
	}
	
	//任务注册到调度表
	scheduleTable["job2"] = cronjob
	// 启动一个调度协程
	go func() {
		var (
			jobName string
			cronjob *Cronjob
			now     time.Time
		)
		// 定时检查任务调度表
		for {
			now = time.Now()
			for jobName, cronjob = range scheduleTable {
				//判断是否过期
				if cronjob.nextTime.Before(now) || cronjob.nextTime.Equal(now) { //启动一个协程 执行这个任务
					go func(jobName string) {
						fmt.Println("执行:", jobName)
					}(jobName)
					
					// 计算下一次调度时间
					cronjob.nextTime = cronjob.expr.Next(now)
					fmt.Println(jobName, "下次执行时间", cronjob.nextTime)
				}
			}
			// 睡眠100毫秒
			select {
			case <-time.NewTimer(100 * time.Millisecond).C: //将在100毫秒内可读，返回
			}
		}
	}()
	
	time.Sleep(100 * time.Second)
}
