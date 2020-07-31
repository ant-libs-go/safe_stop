/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-03-25 14:48:58
# File Name: safe_stop.go
# Description: 本质上就是一个wg，但封装成pkg使用起来更方便
####################################################################### */

package safe_stop

import "sync"

type SafeStop struct {
	wg *sync.WaitGroup
}

func NewSafeStop() *SafeStop {
	o := &SafeStop{}
	o.wg = &sync.WaitGroup{}
	return o
}

func (this *SafeStop) Lock(in int) {
	this.wg.Add(in)
}

func (this *SafeStop) Unlock() {
	this.wg.Done()
}

func (this *SafeStop) Wait() {
	this.wg.Wait()
}

var Default = NewSafeStop()

func Lock(in int) {
	Default.Lock(in)
}

func Unlock() {
	Default.Unlock()
}

func Wait() {
	Default.Wait()
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
