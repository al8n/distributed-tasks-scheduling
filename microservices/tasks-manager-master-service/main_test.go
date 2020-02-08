package main

import (
	"context"
	"encoding/json"
	taskspb "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/pb/tasks"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/responses"
	"google.golang.org/grpc"
	"net/http"
	"sync"
	"testing"
)

func TestHTTP(t *testing.T) {


	var wg sync.WaitGroup

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var defaultTransport http.RoundTripper = &http.Transport{}
			var data responses.GetAllTasksResponse
			req, err := http.NewRequest( "POST","http://localhost:10010/api/v1/task/get/all", nil)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := defaultTransport.RoundTrip(req)
			if err != nil {
				t.Errorf("%d: %v", i, err)
			}

			if resp != nil {
				defer resp.Body.Close()
				err = json.NewDecoder(resp.Body).Decode(&data)

				if err != nil {
					t.Fatal(err)
				}

				if resp.StatusCode != http.StatusFound {
					t.Errorf("handler returned wrong status code: got %v want %v",
						resp.StatusCode, http.StatusFound)
				}
			}



		}(i)
	}
	wg.Wait()

}

func TestGRPC(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
			if err != nil {
				t.Error(err)
			}


			c := taskspb.NewTasksClient(conn)

			res, err := c.GetAllTasks(context.TODO(), &taskspb.GetAllTasksRequest{})
			if err != nil {
				t.Error(err)
			}

			if res.ErrorNumber != 0 {
				t.Errorf("%d: Error Number is %d, expected %d", i, res.ErrorNumber, 0)
			}

		}(i)
	}
	wg.Wait()

}
