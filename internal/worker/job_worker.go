package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/l10-bhushan/mini_worker_queue/internal/model"
	"github.com/l10-bhushan/mini_worker_queue/internal/repository"
)

func StartPoolWorker(n int, jobs <-chan model.Job, repo *repository.PostgresDb) {
	for i := 0; i < n; i++ {
		go worker(i, jobs, repo)
	}
}

func worker(i int, jobs <-chan model.Job, repo *repository.PostgresDb) {
	for job := range jobs {
		ctx := context.Background()

		fmt.Println("Worker", i, "started job", job.Id)

		repo.UpdateStatus(ctx, job.Id, model.JobStatusProcessing)

		// Simulate work
		time.Sleep(2 * time.Second)

		// Simulate success/failure
		if job.Id[len(job.Id)-1]%2 == 0 {
			repo.UpdateStatus(ctx, job.Id, model.JobStatusCompleted)
			fmt.Println("Worker", i, "completed job", job.Id)
		} else {
			repo.UpdateStatus(ctx, job.Id, model.JobStatusFailed)
			fmt.Println("Worker", i, "failed job", job.Id)
		}
	}

}
