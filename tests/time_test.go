package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBeforeAfter(t *testing.T) {
	time1 := time.Now()
	time.Sleep(time.Microsecond * 10)
	time2 := time.Now()
	assert.True(t, time1.Before(time2))
	// time1.Before(time2) 代表 time1 的時候在 time2 之前
}
