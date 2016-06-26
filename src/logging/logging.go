package logging

import (
    "log"
    "os"
    "fmt"
)

var logger *log.Logger

func Init(loggingFile string, prefix string) {
  f, err := os.OpenFile(loggingFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
      panic("error opening file: ")
  }
  //defer f.Close()
  logger = log.New(f, "logger: ", log.Lshortfile)
  logger.SetPrefix(prefix)
}

func Info(args ...interface{}) {
  fmt.Println(args...)
  fmt.Println("logger", logger)
  logger.Println(args...)
}

func Println(args ...interface{}) {
  logger.Println(args...)
}

func Printf(format string, args ...interface{}) {
  logger.Printf(format, args...)
}
