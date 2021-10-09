# SafeStop

safe\_stop是基于sync.WaitGroup抽象的服务安全退出模块

[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/ant-libs-go/safe_stop?status.png)](http://godoc.org/github.com/ant-libs-go/safe_stop)
[![Go Report Card](https://goreportcard.com/badge/github.com/ant-libs-go/safe_stop)](https://goreportcard.com/report/github.com/ant-libs-go/safe_stop)

## 特性

* 支持多个任务间隔指定时间运行
* 支持多实例部署情况下，避免并发运行

## 安装

	go get github.com/ant-libs-go/safe_stop

## 快速开始

```golang
// lock
safe_stop.Lock(1)

// safe exit
safe_stop.Wait()
```
