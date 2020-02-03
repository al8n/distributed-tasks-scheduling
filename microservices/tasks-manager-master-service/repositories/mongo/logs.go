package mongo

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)



func (db *LogDB) GetLogs(req requests.GetLogsRequest) (logs []*entities.TaskLog, err error)  {
	var (
		opts *options.FindOptions
		cur *mongo.Cursor

	)

	opts = options.Find()
	opts.SetProjection(bson.D{{"_id",0}})
	opts.SetSkip(int64(req.Page - 1) * int64(req.Limit))
	opts.SetSort(bson.D{{"start_time", -1}})
	opts.SetLimit(int64(req.Limit))

	if req.After == 0 {
		req.After = time.Now().UnixNano() / 1000 /1000
	}

	if req.Start == 0 {
		cur, err = db.logCollection.Find(context.TODO(), bson.D{{}}, opts)
	} else {
		cur, err = db.logCollection.Find(context.TODO(), bson.D{
			{
				"start_time", bson.D{
				{
					"$lte", req.After,
				},
				{
					"$gte", req.Start,
				},
			},
			},
		}, opts)
	}

	if cur != nil {
		for cur.Next(context.TODO())  {
			var log *entities.TaskLog
			err = cur.Decode(&log)
			if err != nil {
				return
			}

			logs = append(logs, log)
		}
	}

	return
}

