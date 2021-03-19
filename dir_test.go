package utils

import (
    "os"
    "reflect"
    "testing"
)

func TestDirExists(t *testing.T) {
    type args struct {
        dir string
    }
    tests := []struct {
        name    string
        args    args
        want    bool
        wantErr bool
    }{
        {name: "目录存在", args: args{dir: "../utils"}, want: true, wantErr: false},
        {name: "目录不存在", args: args{dir: "./utils"}, want: false, wantErr: true},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                got, err := DirExists(tt.args.dir)
                if (err != nil) != tt.wantErr {
                    t.Errorf("DirExists() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                if got != tt.want {
                    t.Errorf("DirExists() got = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestDirGlob(t *testing.T) {
    type args struct {
        pattern string
    }
    tests := []struct {
        name    string
        args    args
        want    []string
        wantErr bool
    }{
        {name: "匹配成功", args: args{"../utils/*.mod"}, want: []string{"../utils/go.mod"}, wantErr: false},
        {name: "匹配不成功", args: args{"../utils/*.g"}, want: []string{}, wantErr: false},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                got, err := DirGlob(tt.args.pattern)
                if (err != nil) != tt.wantErr {
                    t.Errorf("DirGlob() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                if len(tt.want) == 0 && len(got) == 0 {
                    return
                }
                if !reflect.DeepEqual(got, tt.want) {
                    t.Errorf("DirGlob() got = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestDirName(t *testing.T) {
    type args struct {
        dir string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {name: "存在目录", args: args{dir: "../utils"}, want: ".."},
        {name: "任意目录或文件", args: args{dir: "/tmp/a/b/c/d.go"}, want: "/tmp/a/b/c"},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := DirName(tt.args.dir); got != tt.want {
                    t.Errorf("DirName() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestDirCreate(t *testing.T) {
    type args struct {
        dir       string
        mode      os.FileMode
        recursive bool
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {name: "创建失败", args: args{dir: "tests/a/b/c", mode: os.ModePerm}, wantErr: true},
        {name: "创建成功", args: args{dir: "tests/a/b/c", mode: os.ModePerm, recursive: true}, wantErr: false},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if err := DirCreate(tt.args.dir, tt.args.mode, tt.args.recursive); (err != nil) != tt.wantErr {
                    t.Errorf("Mkdir() error = %v, wantErr %v", err, tt.wantErr)
                }
            },
        )
    }
    _ = DirRemove("tests")
}

func TestDirRealpath(t *testing.T) {
    type args struct {
        dir string
    }
    dir, _ := os.Getwd()
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {name: "当前目录", args: args{dir: "."}, want: dir, wantErr: false},
        {name: "上级目录", args: args{dir: "../utils"}, want: dir, wantErr: false},
        {name: "根目录", args: args{dir: "/tmp/aaa/../bbb"}, want: "/tmp/bbb", wantErr: false},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                got, err := DirRealpath(tt.args.dir)
                if (err != nil) != tt.wantErr {
                    t.Errorf("DirRealpath() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                if got != tt.want {
                    t.Errorf("DirRealpath() got = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestDirScan(t *testing.T) {
    type args struct {
        dir       string
        recursion bool
        pattern   []string
    }
    tests := []struct {
        name    string
        args    args
        want    []dirScan
        wantErr bool
    }{
        {
            name: "不递归", args: args{dir: "./tests"}, want: []dirScan{
            {Name: "a", Path: "./tests"},
            {Name: "b.txt", Path: "./tests"},
        }, wantErr: false,
        },
        {
            name: "递归", args: args{dir: "./tests", recursion: true, pattern: []string{}}, want: []dirScan{
            {Name: "a", Path: "./tests"}, {Name: "b.txt", Path: "./tests/a"}, {Name: "b.txt", Path: "./tests"},
        }, wantErr: false,
        },
    }
    _ = DirCreate("tests/a", os.ModePerm, true)
    _ = os.WriteFile("tests/a/b.txt", []byte{}, os.ModePerm)
    _ = os.WriteFile("tests/b.txt", []byte{}, os.ModePerm)
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                got, err := DirScan(tt.args.dir, tt.args.recursion, tt.args.pattern)
                if (err != nil) != tt.wantErr {
                    t.Errorf("DirScan() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                for k, v := range got {
                    if v.Name != tt.want[k].Name || v.Path != tt.want[k].Path {
                        t.Errorf(
                            "DirScan() got name: %s, path: %s, want: name: %s, path: %s", v.Name, v.Path,
                            tt.want[k].Name, tt.want[k].Path,
                        )
                    }
                }
            },
        )
    }
    _ = DirRemove("./tests")
}
