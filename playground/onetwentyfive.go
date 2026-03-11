package playground

import (
	"errors"
	"time"
)

func Echo(in chan int) (int, error) {
	select {
	case v := <-in:
		return v, nil
	case <-time.After(60 * time.Second):
		return 0, errors.New("timed out waiting for an int")
	}
}
