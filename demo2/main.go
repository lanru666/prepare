package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)
	cmd = exec.Command("/bin/bash", "-c", "sleep 5;ls -l")
	// 执行了
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}
	// 打印子进程的输出
	
	fmt.Println(output)
}
