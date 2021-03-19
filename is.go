package utils

import (
    "os"
    "syscall"
)

// 检查是否是目录
func IsDir(dir string) (bool, error) {
    fd, err := os.Stat(dir)
    if err != nil {
        return false, err
    }
    fm := fd.Mode()
    return fm.IsDir(), nil
}

// 检查是否是文件
func IsFile(filename string) bool {
    _, err := os.Stat(filename)
    if err != nil && os.IsNotExist(err) {
        return false
    }
    return true
}

// 验证文件或目录是否可读
func IsReadable(filename string) bool {
    _, err := syscall.Open(filename, syscall.O_RDONLY, 0)
    if err != nil {
        return false
    }
    return true
}

// 验证文件或目录是否可写
// TODO 目录权限判断异常
func IsWriteable(filename string) bool {
    _, err := syscall.Open(filename, syscall.O_WRONLY, 0)
    if err != nil {
        return false
    }
    return true
}
