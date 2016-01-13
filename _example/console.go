package main

import (
	"github.com/arteev/logger"
	"bytes"
	"os"
)

func main() {
	logger.InitToConsole(logger.LevelError)
	logger.Info.Println("The package Logger!")
	logger.Error.Println("This is error message!!!")
	logger.Trace.Println("This is trace out")

	buf := &bytes.Buffer{}
	logger.InitToWriter(logger.LevelTrace,buf)
	logger.Info.Println("Out to writer!")
	buf.WriteTo(os.Stdout)


}
