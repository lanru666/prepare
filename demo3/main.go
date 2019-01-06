package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)
	resultChan = make(chan *result, 1000)
	// contex chan byte
	// cancelFunc close(chan byte)
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello;")
		// 执行任务 输出结果
		output, err = cmd.CombinedOutput()
		// 把任务输出结果，传递给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()
	
	time.Sleep(1 * time.Second)
	// 取消上下文
	cancelFunc()
	// 在main协程里，等待子协程的退出 并打印任务执行结果
	res = <-resultChan
	// 打印任务执行结果
	fmt.Println(res.err, string(res.output))
}
