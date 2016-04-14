package logger

import (
	"fmt"
	"log"
	"os"
)

const log_path string = ""
const log_file string = "APP_NAME.log"

func Log(m string) {
	f, err := os.OpenFile(log_path+log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logToStd(fmt.Sprintf("error opening file: %v", err))
		return
	}
	defer f.Close()
	logToFile(&f, m)
}

func logToFile(f *os.File, msg string) {
	if f != nil {
		log.SetOutput(f)
	}
	log.Println(m)
}

func logToStd(msg string) {
	logToFile(nil, msg)
}
