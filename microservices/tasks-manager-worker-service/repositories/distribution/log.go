package distribution

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-worker-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

// MongoDB存储日志
type Sink struct {
	client  *mongo.Client
	collection *mongo.Collection
	logChan chan *entities.TaskLog
	autoCommitChan chan *entities.LogBatch
}

var (
	SgtSink *Sink
	onceSink sync.Once
)

func InitSink() *Sink {
	onceSink.Do(func() {
		var (
			client *mongo.Client
			opts *options.ClientOptions
			err error
		)

		opts = options.Client().ApplyURI(myconfig.ConfigSingleton.MongoEndpoint)
		opts.SetConnectTimeout(time.Duration(myconfig.ConfigSingleton.MongoTimeout) * time.Millisecond)

		// 建立mongodb连接
		if client, err = mongo.Connect(context.TODO(), opts); err != nil {
			log.Println(err)
			return
		}

		// 选择db和collection
		SgtSink = &Sink{
			client: client,
			collection: client.Database(myconfig.ConfigSingleton.MongoDatabase).Collection(myconfig.ConfigSingleton.MongoCollection),
			logChan: make(chan *entities.TaskLog, myconfig.ConfigSingleton.LogChanSize),
			autoCommitChan: make(chan *entities.LogBatch, myconfig.ConfigSingleton.LogCommitChanSize),
		}
	})
	go SgtSink.WriteLoop()
	return SgtSink
}

// 日志存储协程
func (s *Sink) WriteLoop()  {
	var (
		log *entities.TaskLog
		batch *entities.LogBatch // 当前的批次
		timeoutBatch *entities.LogBatch // 超时批次
		commitTimer *time.Timer
	)

	for {
		select {
		case log = <-s.logChan:
			// 把这条log写到mongodb中
			if batch == nil {
				 batch = &entities.LogBatch{}
				 // 让这个批次超时自动提交
				 commitTimer = time.AfterFunc(
				 	time.Duration(myconfig.ConfigSingleton.LogCommitTimeout) * time.Millisecond,
				 	func(batch *entities.LogBatch) func() {
					 	return func() {
							s.autoCommitChan <- batch
						}
				 	}(batch),
				 )
			}

			// 把新的日志追加到批次中
			batch.Logs = append(batch.Logs,log)

			// 如果batch 满了就立即发送
			if len(batch.Logs) >= myconfig.ConfigSingleton.LogBatchSize {
				s.SaveLogs(batch)
				// 清空batch
				batch = nil
				// 取消定时器
				commitTimer.Stop()
			}
		case timeoutBatch = <-s.autoCommitChan: // 过去的批次
			// 判断过期批次是否仍然是当前批次
			if timeoutBatch != batch {
				continue // 跳过
			}
			// 把批次写入到mongo 中
			s.SaveLogs(timeoutBatch)
			// 清空batch
			batch = nil
			//
		}
	}
}

// 批量储存日志
func (s *Sink) SaveLogs(batch *entities.LogBatch)  {

	_, err := s.collection.InsertMany(context.TODO(), batch.Logs)
	if err != nil {
		log.Println("err", err)
	}
}

// 发送日志
func (s *Sink) Append(taskLog *entities.TaskLog)  {
	select {
	case s.logChan <- taskLog:
	default:
		// 队列满了就丢弃
	}
}
