package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
	//"github.com/mongodb/mongo-go-driver/bson"
)

//记录任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

//一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   //任务名
	Command   string    `bson:"command"`   //shell命令
	Err       string    `bson:"err"`       //错误
	Content   string    `bson:"content"`   //脚本输出
	TimePoint TimePoint `bson:"timePoint"` //执行时间点
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		logArr     []interface{} //类似C语言的void*
		insertId   interface{}
		result     *mongo.InsertManyResult
	)
	//1、 建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017"); err != nil {
		fmt.Println(err)
		return
	}
	//2、 选择数据库my_db
	database = client.Database("cron")
	//3、 选择表my_collection
	collection = database.Collection("log")
	//4 插入记录
	record = &LogRecord{
		JobName: "job10",
		Command: "echo hello",
		Err:     "",
		Content: "Hello",
		TimePoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10,
		},
	}
	// 批量插入多条document
	logArr = []interface{}{record, record, record}
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}
	
	for _, insertId = range result.InsertedIDs {
		fmt.Println(insertId)
	}
}
