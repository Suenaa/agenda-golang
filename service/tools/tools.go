package tools

import (
	"fmt"
	"os"
	"github.com/Suenaa/agenda-golang/logs"
)

//Report 输出错误
func Report(err error) {
	if err != nil {
		logs.ErrLog(err)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
