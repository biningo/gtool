package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCurrentTimeMillis(t *testing.T) {
	// default provider
	{
		currentMs := CurrentTimeMillis()
		assert.Equal(t, time.Now().UnixMilli(), currentMs)
	}

	// system provider
	{
		SetCurrentMillisSystem()
		currentMs := CurrentTimeMillis()
		assert.Equal(t, time.Now().UnixMilli(), currentMs)
	}

	// fixed provider
	{
		var fixedMs int64 = 10
		SetCurrentMillisFixed(fixedMs)
		currentMs := CurrentTimeMillis()
		assert.Equal(t, fixedMs, currentMs)
	}

	// offset provider
	{
		var offsetMs int64 = -10
		SetCurrentMillisOffset(offsetMs)
		currentMs := CurrentTimeMillis()
		assert.Equal(t, time.Now().UnixMilli()+offsetMs, currentMs)
	}
}
