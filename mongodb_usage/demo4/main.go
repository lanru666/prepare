package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
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

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		condition  *FindByJobName
		cursor     mongo.Cursor
		record     *LogRecord
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
	//4、按照jobName字段过滤,想找出jobName=10的记录,找出5条  collection的find
	condition = &FindByJobName{
		JobName: "job10",
	}
	//5、查询(过滤+翻页参数),
	limit := int64(2)
	if cursor, err = collection.Find(context.TODO(), condition, &options.FindOptions{Limit: &limit,}); err != nil {
		fmt.Println(err)
		return
	}
	//6、遍历结果集
	// 释放游标
	defer cursor.Close(context.TODO())
	
	for cursor.Next(context.TODO()) {
		//定义一个日志对象
		record = &LogRecord{}
		//反序列化bson到对象
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}
		//把日志行打印
		fmt.Println(*record)
	}
	
}
