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

//jobName的过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"` //jobName赋值为job10
}

//一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   //任务名
	Command   string    `bson:"command"`   //shell命令
	Err       string    `bson:"err"`       //错误
	Content   string    `bson:"content"`   //脚本输出
	TimePoint TimePoint `bson:"timePoint"` //执行时间点
}

// startTime小于某时间
// {"$lt":timestamp}
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}
type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		delCond    *DeleteCond
		delResult  *mongo.DeleteResult
		//condition  *FindByJobName
		//cursor     mongo.Cursor
		//record     *LogRecord
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
	//4、要删除开始时间早于当前时间的所有日志($lt是less than)
	// delete({"timePoint.startTime":{"$lt":当前时间}})
	delCond = &DeleteCond{
		beforeCond: TimeBeforeCond{
			Before: time.Now().Unix(),
		},
	}
	// 执行删除
	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("删除的行数", delResult.DeletedCount)
}
