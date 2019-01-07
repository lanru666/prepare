package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	// lease 实现锁自动过期
	// op操作
	// txn事务 if else then
	
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)
	//客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	//建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	client = client;
	// 1 上锁(创建租约 自动续租 拿着租约去抢占一个key)
	
	// 2 处理业务
	
	// 3 释放锁(取消自动续租，释放租约)
	
	
}
