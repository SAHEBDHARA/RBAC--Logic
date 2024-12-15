package logger


import "log"

func Debug(args ...interface{}) {
    log.Println(args...)
}

func Error(args ...interface{}) {
    log.Println(args...)
}


func Warn(args ...interface{}){
    log.Println(args...)
}

func Info(args ...interface{}){
    log.Println(args...)
}