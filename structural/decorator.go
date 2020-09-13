package structural

import (
	"fmt"
	"time"
)

type (
	obj struct{}

	objWithLogger struct {
		obj
	}
)

func (o obj) method() {
	// do some work
}

// decorated method with logger
func (o objWithLogger) method() {
	start := time.Now()
	defer func() {
		fmt.Printf("method took: %s", time.Now().Sub(start).String())
	}()
}
