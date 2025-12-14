package threadpool

import (
	"context"
	"log"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
type Runnable interface {
	Run(context.Context) error
}

type ThreadPool interface {
	Run(Runnable)
	Close()
}

type threadPool struct {
	ctx       context.Context    //egész threadpool élettartamát és futását koordinálja
	cancel    context.CancelFunc //futó munka megállítása
	wg        sync.WaitGroup
	tasks     chan Runnable //feladat queue
	errCh     chan error
	closeOnce sync.Once
}

func (p *threadPool) Run(task Runnable) {
	//nem blokkoló küldés
	select {
	case <-p.ctx.Done():
	case p.tasks <- task:
	}
}

func (p *threadPool) Close() {
	p.closeOnce.Do(func() {
		p.cancel()     //megszakítjuk a contextet
		close(p.tasks) //lezárjuk a task csatornát
		p.wg.Wait()    //workerek befejezik a munkát
		close(p.errCh)
	})
}

func (p *threadPool) worker() {
	defer p.wg.Done()

	for {
		select {
		case <-p.ctx.Done():
			//leállítás
			return
		case task, ok := <-p.tasks:
			if !ok {
				//task csatorna lezárult
				return
			}
			err := task.Run(p.ctx)
			if err != nil {
				select {
				case p.errCh <- err:
				default:
					//ha tele az error csatorna, a standard outputra logolunk
					log.Println(err)
				}
			}
		}
	}
}

func NewThreadPool(n int) (ThreadPool, chan error) {
	ctx, cancel := context.WithCancel(context.Background())
	p := &threadPool{
		ctx:    ctx,
		cancel: cancel,
		tasks:  make(chan Runnable, n*10),
		//error channel legyen buffered
		errCh: make(chan error, 100),
	}

	//elindítjuk az n workert
	for i := 0; i < n; i++ {
		p.wg.Add(1)
		go p.worker()
	}

	return p, p.errCh
}
