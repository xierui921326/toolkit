package worker_pool

// Job 任务
type Job interface {
	Do() // 执行任务...
}

// Worker 协程
type Worker struct {
	JobQueue chan Job  //任务队列
	Quit     chan bool //停止当前任务
}

// NewWorker 新建一个 worker(协程)通道实例 新建一个协程
func NewWorker() *Worker {
	return &Worker{
		JobQueue: make(chan Job), //初始化工作队列为null
		Quit:     make(chan bool),
	}
}

/*
整个过程中 每个Worker(协程)都会被运行在一个协程中，
在整个WorkerPool(线程池)中就会有num个可空闲的Worker(协程)，
当来一条数据的时候，领导就会小组中取一个空闲的Worker(协程)去执行该Job，
当工作池中没有可用的worker(协程)时，就会阻塞等待一个空闲的worker(协程)。
每读到一个通道参数 运行一个 worker
*/

func (w *Worker) Run(wq chan chan Job) {
	//这是一个独立的协程 循环读取通道内的数据，
	//保证 每读到一个通道参数就 去做这件事，没读到就阻塞
	go func() {
		for {
			wq <- w.JobQueue //注册工作通道  到 线程池
			select {
			case job := <-w.JobQueue: //读到参数
				job.Do()
			case <-w.Quit: //终止当前任务
				return
			}
		}
	}()
}

// Stop 停止当前worker(协程)
func (w *Worker) Stop() {
	go func() {
		w.Quit <- true //发送停止信号
	}()
}

// WorkerPool 线程池
type WorkerPool struct {
	workerLen   int      //线程池中  worker(协程)的数量
	JobQueue    chan Job //线程池的  job 通道
	WorkerQueue chan chan Job
	workers     []*Worker //线程池中  worker(协程)的实例
	Quit        chan bool //停止信号
}

// NewWorkerPool 初始化worker(协程)
func NewWorkerPool(workerLen int) *WorkerPool {
	return &WorkerPool{
		workerLen:   workerLen,                      //开始建立workerLen个worker(协程)
		JobQueue:    make(chan Job),                 //工作队列 通道
		WorkerQueue: make(chan chan Job, workerLen), //最大通道参数设为 最大协程数 workerLen协程的数量最大值
		workers:     make([]*Worker, 0, workerLen),
		Quit:        make(chan bool),
	}
}

// Run 运行线程池
func (wp *WorkerPool) Run() {
	//初始化时会按照传入的num，启动num个后台协程，然后循环读取Job通道里面的数据，
	//读到一个数据时，再获取一个可用的Worker，并将Job对象传递到该Worker的chan通道
	for i := 0; i < wp.workerLen; i++ {
		//新建 workerLen个worker(协程) 并发执行，每个协程可处理一个请求
		worker := NewWorker() //运行一个协程 将线程池 通道的参数  传递到 worker协程的通道中 进而处理这个请求
		worker.Run(wp.WorkerQueue)
		wp.workers = append(wp.workers, worker)
	}

	// 循环获取可用的worker,往worker中写job
	go func() { //这是一个单独的协程 只负责保证 不断获取可用的worker
		for {
			select {
			case job := <-wp.JobQueue: //读取任务
				//尝试获取一个可用的worker作业通道。
				//这将阻塞，直到一个worker空闲
				worker := <-wp.WorkerQueue
				worker <- job //将任务 分配给该协程
			case <-wp.Quit:

				// 通知所有 worker 停止
				for _, worker := range wp.workers {
					worker.Stop()
				}
				return
			}
		}
	}()
}

// Add 添加任务
func (wp *WorkerPool) Add(job Job) {
	wp.JobQueue <- job
}

// Stop 停止 WorkerPool
func (wp *WorkerPool) Stop() {
	wp.Quit <- true
}
