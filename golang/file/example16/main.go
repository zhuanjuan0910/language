package main

import (
	"fmt"
	"log"
	"os"
)

func debug(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("file open is :", err)
		return
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}
	debugLog := log.New(logFile, "[debug]", log.Ldate)
	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Debug log")

}
func warn(logName string) {
	logFile, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open logfile is", err)
		return
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}
	warnLog := log.New(logFile, "[warn]", log.Ldate)
	warnLog.SetPrefix("[warn]")
	warnLog.SetFlags(log.Llongfile)
	warnLog.Println("this is warn log")

}
func main() {
	logName := "./test.log"
	debug(logName)
	warn(logName)
}
