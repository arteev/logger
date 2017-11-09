package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//A Level for logger
type Level int

//Log - Logger by level
type Log struct {
	*log.Logger
	Level Level
	Name  string
}

//Pre Defined levels log
const (
	LevelNone = Level(iota)
	LevelInfo
	LevelWarn
	LevelError
	LevelDebug
	LevelTrace
)

var names = map[Level]string{
	LevelNone:  "None",
	LevelInfo:  "Info",
	LevelWarn:  "Warn",
	LevelError: "Error",
	LevelDebug: "Debug",
	LevelTrace: "Trace",
}

//Concrete loggers by level
var (
	Info         *Log
	Warn         *Log
	Error        *Log
	Debug        *Log
	Trace        *Log
	CurrentLevel Level
)

//InitLoggers Initialize the loggers
func InitLoggers(level Level, wInfo, wWarn, wError, wDebug, wTrace io.Writer) {
	CurrentLevel = level
	if level < LevelTrace {
		wTrace = ioutil.Discard
	}
	if level < LevelDebug {
		wDebug = ioutil.Discard
	}
	if level < LevelError {
		wError = ioutil.Discard
	}
	if level < LevelWarn {
		wWarn = ioutil.Discard
	}
	if level < LevelInfo {
		wInfo = ioutil.Discard
	}

	formatFlag := log.Ldate | log.Ltime | log.Lshortfile
	Info = mustLevel(LevelInfo, wInfo, formatFlag)
	Warn = mustLevel(LevelWarn, wWarn, formatFlag)
	Error = mustLevel(LevelError, wError, formatFlag)
	Debug = mustLevel(LevelDebug, wDebug, formatFlag)
	Trace = mustLevel(LevelTrace, wTrace, formatFlag)
}

//InitToConsole initialize the loggers for all levels with a output to console
func InitToConsole(level Level) {
	InitLoggers(level, os.Stdout, os.Stdout, os.Stderr, os.Stdout, os.Stdout)
}

//InitToWriter initialize the loggers for all levels with "w" io.Writer
func InitToWriter(level Level, w io.Writer) {
	InitLoggers(level, w, w, w, w, w)
}

//InitEmpty initialize the loggers for all levels with a output to ioutil.Discard.
//This default.
func InitEmpty() {
	InitLoggers(LevelNone, nil, nil, nil, nil, nil)
}

func init() {
	InitEmpty()
	Error = mustLevel(LevelError, os.Stderr, log.Ldate|log.Ltime|log.Lshortfile)
}

//String is Stringer
func (l Log) String() string {
	return fmt.Sprintf("%s (%d)", l.Name, l.Level)
}

//Name returns name for pre defined levels
func (level Level) Name() string {
	if name, exists := names[level]; exists {
		return name
	}
	return ""
}

//mustLevel returns logger by level
func mustLevel(level Level, w io.Writer, flag int) *Log {
	namelevel := level.Name()
	if namelevel == "unknow" {
		panic(fmt.Errorf("Unknow level: %d", level))
	}
	lgr := log.New(w, strings.ToUpper(namelevel)+":", flag)
	return &Log{lgr, level, namelevel}
}
