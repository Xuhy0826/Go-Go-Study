package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全地在多个goroutine间共享的资源。且资源必须实现了io.Closer接口
type Pool struct {
	//互斥锁
	m sync.Mutex
	//保存共享的资源
	resources chan io.Closer
	//工厂函数：当池需要一个新资源且池中没有可用资源时，用这个函数创建。
	factory func() (io.Closer, error)
	//标志量，Pool 是否已经被关闭
	closed bool
}

// ErrPoolClosed 描述一个请求已关闭池的错误
var ErrPoolClosed = errors.New("Pool has been closed")

// New 创建一个用来管理资源的池。
// 需要一个工厂函数并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 先检查池中是否还有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	// 因为没有空闲资源，创建新的资源返回
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 释放资源，将使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	//上锁，保证本操作与Close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//若池已关闭，则直接关闭资源即可
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	//如果池已满，则直接关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close 会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	//上锁，保证本操作与Release的安全
	p.m.Lock()
	defer p.m.Unlock()

	//若池已关闭，直接返回
	if p.closed {
		return
	}

	//关闭池
	p.closed = true

	// 在清空通道里的资源之前，将通道关闭
	// !!!如果不这样做，会发生死锁!!!
	close(p.resources)

	// 关闭所有资源
	for r := range p.resources {
		r.Close()
	}
}
