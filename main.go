package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"sync"
	"time"
)

type TimerManager struct {
	manager map[int64]*time.Timer
	mu      sync.Mutex
}

type AgeInt int

type UserInfo struct {
	Name string
	Age  AgeInt
	Tm   *TimerManager
}

func main() {
	u := new(UserInfo)

	u.Tm = new(TimerManager)
	fmt.Printf("hello : %v\n", u.Tm)
}

func mainHash() {
	for i := 0; i < 100; i++ {
		token := strconv.Itoa(i)
		hashCode := int(crc32.ChecksumIEEE([]byte(token)))
		if hashCode < 0 {
			hashCode = -hashCode
		}

		fmt.Println(hashCode % 20)
	}

}

func (tm *TimerManager) AddTimer(name int64) error {
	timeFunc := func() {
		fmt.Println("hello name: ", name)
	}

	timer := time.AfterFunc(5*time.Second, timeFunc)

	tm.mu.Lock()
	tm.manager[name] = timer
	tm.mu.Unlock()

	return nil
}

func (tm *TimerManager) DelTimer(name int64) error {
	if timer, ok := tm.manager[name]; ok {
		timer.Stop()
	}

	tm.mu.Lock()
	delete(tm.manager, name)
	tm.mu.Unlock()

	return nil
}

func (tm *TimerManager) GetAllTimer() []*time.Timer {
	allTimer := []*time.Timer{}

	for _, timer := range tm.manager {
		allTimer = append(allTimer, timer)
	}

	return allTimer
}
