package logger

import (
	"flag"
	"go/build"
	"log"
	"os"
)

var Log *log.Logger

func init() {
	// set location of log file
	var path = build.Default.GOPATH + "./log.log"

	flag.Parse()
	var file, err1 = os.Create(path)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + path)
}
