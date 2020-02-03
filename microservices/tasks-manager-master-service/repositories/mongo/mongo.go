package mongo

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/config"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type logs interface {
	GetLogs(req requests.GetLogsRequest) (logs []*entities.TaskLog, err error)
}

type LogDB struct {
	client *mongo.Client
	logCollection *mongo.Collection
}

var (
	SgtLogDB *LogDB
	once      sync.Once
)

func newMongoDB(client *mongo.Client, collection *mongo.Collection) *LogDB {
	once.Do(func() {
		SgtLogDB = &LogDB{
			client:        client,
			logCollection: collection,
		}
	})

	return SgtLogDB
}

func InitMongo() (err error)  {
	var (
		client *mongo.Client
		opt *options.ClientOptions
	)

	opt = options.Client().ApplyURI(myconfig.ConfigSingleton.MongoEndpoint)
	opt.SetConnectTimeout(time.Duration(myconfig.ConfigSingleton.MongoTimeout) * time.Millisecond)
	// 建立mongodb连接
	if client, err = mongo.Connect(context.TODO(), opt); err != nil {
		return
	}

	newMongoDB(client, client.Database(myconfig.ConfigSingleton.MongoDatabase).Collection(myconfig.ConfigSingleton.MongoCollection))
	return
}