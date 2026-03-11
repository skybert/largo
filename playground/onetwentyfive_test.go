package playground

import (
	"testing"
	"testing/synctest"

	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	val, err := Echo(ch)
	require.Nil(t, err)
	require.NotNil(t, val)
}

func TestEchoTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan int)
		_, err := Echo(ch)
		require.NotNil(t, err)
	})
}
