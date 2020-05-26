package logger

import (
	"log"
	"os"
)

var logger *log.Logger
var f *os.File
var err error

func Loggers(name interface{}) {
	f, err = os.OpenFile("./logger/err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatal(err)
	}
	logger = log.New(f, "", log.LstdFlags)
	logger.SetPrefix("TEST LOG ")
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Hello %s", name)
	defer f.Close()
}
