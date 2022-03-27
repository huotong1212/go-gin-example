package logging

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.LogSetting.LogSavePath)
}

//func getLogFileName() string {
//	return fmt.Sprintf("%s",  setting.LogSetting.LogSaveName)
//}

func getLogFileName() string {
	//prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", setting.LogSetting.LogSaveName, time.Now().Format(setting.LogSetting.TimeFormat), setting.LogSetting.LogFileExt)

	//return fmt.Sprintf("%s%s", prefixPath, suffixPath)
	return suffixPath
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err:%v", err)
	}
	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src:%s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile:%v", err)
	}
	return f, nil
}

//func openLogFile(filePath string) *os.File {
//	_, err := os.Stat(filePath) //os.Stat：返回文件信息结构描述文件
//	switch {
//	case os.IsNotExist(err):
//		mkDir()
//	case os.IsPermission(err):
//		log.Fatalf("Permission: %v", err)
//	}
//	/**
//	调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。
//
//	const (
//	    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
//	    O_RDONLY int = syscall.O_RDONLY // 以只读模式打开文件
//	    O_WRONLY int = syscall.O_WRONLY // 以只写模式打开文件
//	    O_RDWR   int = syscall.O_RDWR   // 以读写模式打开文件
//	    // The remaining values may be or'ed in to control behavior.
//	    O_APPEND int = syscall.O_APPEND // 在写入时将数据追加到文件中
//	    O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件
//	    O_EXCL   int = syscall.O_EXCL   // 使用O_CREATE时，文件必须不存在
//	    O_SYNC   int = syscall.O_SYNC   // 同步IO
//	    O_TRUNC  int = syscall.O_TRUNC  // 如果可以，打开时
//	)
//	*/
//	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatalf("Fail to OpenFile : %v", err)
//	}
//
//	return handle
//}
//
//func mkDir() {
//	dir, _ := os.Getwd() //返回与当前目录对应的根路径名
//	fmt.Printf("dir : %s", dir)
//
//	// const定义ModePerm FileMode = 0777
//	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}
