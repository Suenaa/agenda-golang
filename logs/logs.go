package logs

import (
	"log"
	"os"
	"io"
)

var (
	out, e = os.OpenFile("logs/logs.log", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)
	outWriter = io.Writer(out)
	logger = log.New(outWriter, "[agenda] ", log.LstdFlags)
)

func ErrLog(err error) {
	if e != nil {
		log.Fatalln("log file error")
	}
	defer out.Close()
	logger.SetPrefix("[error]")
	logger.Println(err)
}

func EventLog(op string) {
	if e != nil {
		log.Fatalln("log file error")
	}
	defer out.Close()
	logger.SetPrefix("[event]")
	logger.Println(op)
}