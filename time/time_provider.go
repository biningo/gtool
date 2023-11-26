package main

import "time"

type MillisProvider interface {
	GetMillis() int64
}

type SystemMillisProvider struct {
}

func (provider SystemMillisProvider) GetMillis() int64 {
	return time.Now().UnixMilli()
}

type FixedMillisProvider struct {
	millis int64
}

func (provider FixedMillisProvider) GetMillis() int64 {
	return provider.millis
}

type OffsetMillisProvider struct {
	offset int64
}

func (provider OffsetMillisProvider) GetMillis() int64 {
	return time.Now().UnixMilli() + provider.offset
}

var systemMillisProvider = SystemMillisProvider{}
var fixedMillisProvider = FixedMillisProvider{}
var offsetMillisOffset = OffsetMillisProvider{}

var millisProvider MillisProvider = systemMillisProvider

func SetMillisProvider(provider MillisProvider) {
	millisProvider = provider
}

func SetCurrentMillisSystem() {
	millisProvider = systemMillisProvider
}

func SetCurrentMillisFixed(millis int64) {
	fixedMillisProvider.millis = millis
	millisProvider = fixedMillisProvider
}

func SetCurrentMillisOffset(offset int64) {
	offsetMillisOffset.offset = offset
	millisProvider = offsetMillisOffset
}

func CurrentTimeMillis() int64 {
	return millisProvider.GetMillis()
}
