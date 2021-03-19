package utils

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "syscall"
)

// 检查目录是否存在
func DirExists(dir string) (bool, error) {
    return IsDir(dir)
}

// 创建目录
func DirCreate(dir string, mode os.FileMode, recursive bool) error {
    if recursive {
        return os.MkdirAll(dir, mode)
    }
    return os.Mkdir(dir, mode)
}

// 删除目录
func DirRemove(dir string) error {
    return os.RemoveAll(dir)
}

// 返回父目录的路径
func DirName(dir string) string {
    return filepath.Dir(dir)
}

// 查找与模式匹配的路径名
func DirGlob(pattern string) ([]string, error) {
    return filepath.Glob(pattern)
}

// 获取真实路径
func DirRealpath(dir string) (string, error) {
    return filepath.Abs(dir)
}

// 遍历目录
type dirScan struct {
    Name       string
    Path       string
    FullPath   string
    RealPath   string
    ParentPath string
    Mode       os.FileMode
    Size       int64
    CTime      syscall.Timespec
    MTime      syscall.Timespec
    ATime      syscall.Timespec
}

// 遍历目录
func DirScan(dir string, recursion bool, pattern []string) ([]dirScan, error) {
    if ok, err := IsDir(dir); !ok {
        return []dirScan{}, err
    }
    trans := func(v os.FileInfo, filename string) dirScan {
        realPath, _ := DirRealpath(filename)
        sysInfo := v.Sys().(*syscall.Stat_t)
        return dirScan{
            Name:       v.Name(),
            Path:       dir,
            FullPath:   filename,
            RealPath:   realPath,
            ParentPath: DirName(filename),
            Mode:       v.Mode(),
            Size:       v.Size(),
            CTime:      sysInfo.Ctimespec,
            MTime:      sysInfo.Mtimespec,
            ATime:      sysInfo.Atimespec,
        }
    }
    dir = strings.TrimRight(dir, "/")
    var res []dirScan
    var filename string
    fh, err := ioutil.ReadDir(dir)
    for _, v := range fh {
        filename = dir + "/" + v.Name()
        if v.IsDir() && recursion {
            res = append(res, trans(v, filename))
            subRes, _ := DirScan(filename, recursion, pattern)
            res = append(res, subRes...)
            continue
        }
        if len(pattern) > 0 {
            for _, vv := range pattern {
                if matched, _ := regexp.MatchString(vv, v.Name()); matched {
                    res = append(res, trans(v, filename))
                }
            }
        } else {
            res = append(res, trans(v, filename))
        }
    }
    return res, err
}
