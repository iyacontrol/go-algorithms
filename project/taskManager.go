package project

import (
	"fmt"
	"math"
)

type TaskI interface {
	Do()
}

type Worker struct {
	id        int
	taskQueue chan TaskI
	quitChan  chan struct{}
}

func NewWorker(id int, taskQueue chan TaskI) *Worker {
	worker := &Worker{
		id:        id,
		taskQueue: taskQueue,
		quitChan:  make(chan struct{}),
	}

	go worker.run()

	return worker
}

func (j *Worker) run() {
	for {
		select {
		case task := <-j.taskQueue:
			task.Do()
		case <-j.quitChan:
			return
		}
	}
}

func (j *Worker) Stop() {
	go func() {
		j.quitChan <- struct{}{}
	}()
}

type TaskManager struct {
	taskQueue  chan TaskI
	workerPool []*Worker
	stopChan   chan struct{}
}

func New(maxWorkers int) TaskManager {
	if maxWorkers < 1 {
		maxWorkers = 1
	}

	workerPool := make([]*Worker, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		workerPool[i] = NewWorker(i, make(chan TaskI, 10))
	}

	taskQueue := make(chan TaskI, 2*maxWorkers*10)

	tm := TaskManager{
		taskQueue:  taskQueue,
		workerPool: workerPool,
		stopChan:   make(chan struct{}),
	}

	go tm.run()
	return tm
}

func (t *TaskManager) run() {
	for {
		select {
		case task := <-t.taskQueue:
			worker := t.getWorker()
			worker.taskQueue <- task
		case <-t.stopChan:
			for _, worker := range t.workerPool {
				worker.Stop()
			}
			return
		}
	}
}

func (t *TaskManager) Stop() {
	go func() {
		t.stopChan <- struct{}{}
	}()
}

func (t *TaskManager) AddTask(task TaskI) {
	t.taskQueue <- task
}

func (t *TaskManager) getWorker() *Worker {
	var idleWorker *Worker
	minTaskCount := math.MaxInt32

	for _, worker := range t.workerPool {
		select {
		case <-worker.quitChan:
			continue
		default:
			if len(worker.taskQueue) < minTaskCount {
				minTaskCount = len(worker.taskQueue)
				fmt.Printf("%d 号 worker 本地任务队列大小是:%d\n", worker.id, minTaskCount)
				idleWorker = worker
			}
		}
	}
	return idleWorker
}
