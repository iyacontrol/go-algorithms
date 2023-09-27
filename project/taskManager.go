package project

import (
	"fmt"
	"math"
)

type ITask interface {
	Run()
}

type Worker struct {
	id        int
	taskQueue chan ITask
	stopChan  chan struct{}
}

func NewWorker(id int, taskQueque chan ITask) *Worker {
	worker := &Worker{
		id:        id,
		taskQueue: taskQueque,
		stopChan:  make(chan struct{}),
	}

	go worker.run()

	return worker
}

func (w *Worker) run() {
	for {
		select {
		case task := <-w.taskQueue:
			task.Run()
		case <-w.stopChan:
			return
		}
	}
}

func (w *Worker) Stop() {
	go func() {
		w.stopChan <- struct{}{}
	}()
}

type TaskManager struct {
	taskQueue  chan ITask
	workerPool []*Worker
	stopChan   chan struct{}
}

func New(maxWorker int) *TaskManager {
	if maxWorker < 1 {
		maxWorker = 1
	}

	workerPool := make([]*Worker, maxWorker)
	for i := 0; i < maxWorker; i++ {
		workerPool[i] = NewWorker(i, make(chan ITask, 10))

	}

	tm := &TaskManager{
		taskQueue:  make(chan ITask, 2*maxWorker*10),
		workerPool: workerPool,
		stopChan:   make(chan struct{}),
	}

	go tm.run()

	return tm
}

func (tm *TaskManager) run() {
	for {
		select {
		case task := <-tm.taskQueue:
			worker := tm.getWorker()
			worker.taskQueue <- task
		case <-tm.stopChan:
			for _, w := range tm.workerPool {
				w.stopChan <- struct{}{}
			}
			return
		}
	}
}

func (tm *TaskManager) Stop() {
	go func() {
		tm.stopChan <- struct{}{}
	}()
}

func (tm *TaskManager) getWorker() *Worker {
	var idleWorker *Worker
	minTaskCount := math.MaxInt32

	for _, worker := range tm.workerPool {
		select {
		case <-worker.stopChan:
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

func (tm *TaskManager) AddTask(task ITask) {
	tm.taskQueue <- task
}
