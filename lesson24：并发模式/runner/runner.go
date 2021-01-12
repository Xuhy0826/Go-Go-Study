package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 表示一个任务执行者，需要在给定的超时时间内执行一组任务，如果接收到操作系统发送来的中断信号会结束这些任务
type Runner struct {
	// tasks表示要执行的任务集合
	tasks []func(int)

	// interrupt 通道获取从操作系统发来的信号
	interrupt chan os.Signal

	// timeout 报告处理任务已经超时，标记为“<-chan”表示此通道为一个单向只读通道，因为向通道中传入数据是Go运行时
	timeout <-chan time.Time

	// complete 通道报告处理任务已经完成
	complete chan error
}

// ErrTimeout 超时的错误，这个错误值会在收到超时事件时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 中断的错误，会在收到操作系统的中断事件时返回
var ErrInterrupt = errors.New("received interrupt")

//New 设定一个超时时间，返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 为 Runner 添加任务。任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并且监视通道事件
func (r *Runner) Start() error {
	//设置我们希望获取哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	// goroutine 执行任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 依次执行已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//执行前先判断是否有中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		//执行任务
		task(id)
	}

	//任务执行完成
	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当接收到中断信号时
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// 防止阻塞，使其继续正常运行
	default:
		return false
	}
}
