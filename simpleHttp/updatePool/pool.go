package updatePool

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Job struct {
	Execute func() (interface{}, error)
}

type Result struct {
	Result       interface{}
	ResponseTime time.Duration
	Error        error
}

type Pool struct {
	context      context.Context
	processor    *processor
	workersCount int
	timeout      time.Duration

	jobs    chan Job
	results chan<- Result

	wg      *sync.WaitGroup
	stopped bool
}

func New(ctx context.Context, workersCount int, timeout time.Duration, results chan<- Result) *Pool {
	p := &Pool{
		context:      ctx,
		processor:    newProcessor(),
		timeout:      timeout,
		workersCount: workersCount,
		jobs:         make(chan Job, workersCount),
		results:      results,
		wg:           new(sync.WaitGroup),
	}
	p.Init()
	go p.gracefulShutdown()
	return p
}

func (p *Pool) gracefulShutdown() {
	stopWg := p.context.Value("stopWg").(*sync.WaitGroup)
	stopWg.Add(1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	p.Stop()
	stopWg.Add(1)
}

func (p *Pool) Init() {
	for i := 1; i <= p.workersCount; i++ {
		go p.initWorker(i)
	}
}

func (p *Pool) Push(j Job) {
	if p.stopped {
		return
	}

	p.jobs <- j
	p.wg.Add(1)
}

func (p *Pool) Stop() {
	p.stopped = true
	p.wg.Wait()
}

func (p *Pool) initWorker(id int) {
	for job := range p.jobs {
		p.results <- p.processor.process(job)
		p.wg.Done()
	}

	log.Printf("[processe ID %d] finished proccesing\n", id)
}
