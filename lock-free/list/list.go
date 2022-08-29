package main

import (
	"sync"
	"unsafe"
)

var (
	nodePool = sync.Pool{}
	nodeGet  = nodePool.Get
	nodePut  = nodePool.Put
)

type List struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func init() {

}
