package instrumentation

import (
	"fmt"
	myconfig "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/config"
	"time"
)

func (mw *ImplInstrumentation) SaveOneTask(name string) (err error)  {
	defer func(begin time.Time) {
		lvs := []string{"Method", "SaveOneTask", "Name", myconfig.JOB_SAVE_DIR + name, "Error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.SaveOneTask(name)
	return
}

func (mw *ImplInstrumentation) DeleteOneTask(name string) (err error)  {
	tn := myconfig.JOB_SAVE_DIR + name
	defer func(begin time.Time) {

		lvs := []string{"Method", "DeleteOneTask", "Name", tn, "Error",
			fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.DeleteOneTask(tn)
	return
}

func (mw *ImplInstrumentation) GetOneTask(name string) (err error)  {
	tn := myconfig.JOB_SAVE_DIR + name
	defer func(begin time.Time) {
		lvs := []string{"Method", "GetOneTask", "Name", tn, "Error",
			fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.GetOneTask(tn)
	return
}