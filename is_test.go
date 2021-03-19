package utils

import (
    "os"
    "testing"
)

func TestIsDir(t *testing.T) {
    type args struct {
        dir string
    }
    dir, _ := os.Getwd()
    tests := []struct {
        name    string
        args    args
        want    bool
        wantErr bool
    }{
        {name: "目录存在", args: args{dir: dir}, want: true, wantErr: false},
        {name: "目录不存在", args: args{dir: "/utils"}, want: false, wantErr: true},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                got, err := IsDir(tt.args.dir)
                if (err != nil) != tt.wantErr {
                    t.Errorf("IsDir() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                if got != tt.want {
                    t.Errorf("IsDir() got = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestIsFile(t *testing.T) {
    type args struct {
        filename string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {name: "文件存在", args: args{filename: "is.go"}, want: true},
        {name: "文件不存在", args: args{filename: "file_not_exists.go"}, want: false},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := IsFile(tt.args.filename); got != tt.want {
                    t.Errorf("IsFile() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestIsReadable(t *testing.T) {
    type args struct {
        filename string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {name: "可读文件", args: args{filename: "./is.go"}, want: true},
        {name: "可读目录", args: args{filename: "./"}, want: true},
        {name: "不可读目录", args: args{filename: "/root"}, want: false},
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := IsReadable(tt.args.filename); got != tt.want {
                    t.Errorf("IsReadable() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}

func TestIsWriteable(t *testing.T) {
    type args struct {
        filename string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {name: "可写文件", args: args{filename: "./is.go"}, want: true},
        //{name: "可写目录", args: args{filename: "./tests"}, want: true},
        {name: "不可写目录", args: args{filename: "/root"}, want: false},
    }
    _ = DirCreate("tests", os.ModePerm, false)
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := IsWriteable(tt.args.filename); got != tt.want {
                    t.Errorf("IsWriteable() = %v, want %v", got, tt.want)
                }
            },
        )
    }
    _ = DirRemove("tests")
}
