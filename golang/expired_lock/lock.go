package expired_lock

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ExpiredLock struct {
	mu    sync.Mutex         // 互斥锁
	owner string             // 持有者
	stop  context.CancelFunc // 结束方法
}

func (e *ExpiredLock) Lock(expired int) {
	e.mu.Lock()

	token := GetGoroutineToken()
	e.owner = token

	if expired <= 0 {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	e.stop = cancel
	go func() {
		select {
		case <-time.After(time.Second * time.Duration(expired)):
			e.unlock(token)
		case <-ctx.Done():
		}
	}()
}

func (e *ExpiredLock) Unlock() {
	e.unlock(GetGoroutineToken())
}

func (e *ExpiredLock) unlock(token string) {
	if e.owner != token {
		return
	}

	if e.stop != nil {
		e.stop()
	}
	e.owner = ""
	e.mu.Unlock()
}

func GetGoroutineToken() string {
	pid := os.Getpid()
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	goroutineID := strings.Split(strings.TrimPrefix(string(buf[:n]), "goroutine "), " ")[0]
	id, err := strconv.Atoi(goroutineID)
	if err != nil {
		panic(fmt.Sprintf("Can't get goroutine id: %s", err))
	}
	return fmt.Sprintf("%d_%d", pid, id)
}
