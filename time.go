package utils

import (
    "fmt"
    "github.com/golang-module/carbon"
)

// 主要为了保存carbon库
// 具体使用方式可以查看 https://github.com/golang-module/carbon
func timeCarbon() {
    fmt.Println(carbon.Now())
}
