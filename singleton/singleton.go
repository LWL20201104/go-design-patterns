package singleton

import (
	"sync"
	"sync/atomic"
)

type Singleton struct {
	Index int
}

var (
	instance *Singleton
	lock     sync.Mutex
	rwLock   sync.RWMutex
	flag     int32
	once     = &sync.Once{}
)

// 单线程版本
func GetSingletonV1() *Singleton {
	if instance == nil {
		instance = &Singleton{Index: 1}
	}
	return instance
}

// 双检查锁
// C++ 中instance的初始化可能存在指令的reorder过程，导致返回空的未初始化为Singleton对象的堆指针
// golang 中第一次 check instance 和初始化 instance 存在并发读写冲突
func GetInstanceV2() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{Index: 1}
		}
	}
	return instance
}

// 读写锁版双检查锁
// 通过读写锁来避免golang中的并发读写问题
func GetInstanceV3() *Singleton {
	rwLock.RLock()
	ins := instance
	rwLock.RUnlock()
	if ins == nil {
		rwLock.Lock()
		defer rwLock.Unlock()
		instance = &Singleton{Index: 1}
	}
	return instance
}

// 原子操作版双检测锁
// 通过原子操作避免golang中的并发读写问题
func GetInstanceV4() *Singleton {
	if atomic.LoadInt32(&flag) == 0 {
		lock.Lock()
		defer lock.Unlock()
		if atomic.LoadInt32(&flag) == 0 {
			instance = &Singleton{Index: 1}
			atomic.StoreInt32(&flag, 1)
		}
	}
	return instance
}

// golang once版
func GetInstanceV5() *Singleton {
	once.Do(func() {
		instance = &Singleton{Index: 1}
	})
	return instance
}
