package MyLog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

// color funcs
// var bold = color.New(color.Bold).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

// Log is a type for structured message logging
type Log struct {
	stdVar           *log.Logger
	infoVar          *log.Logger
	debugVar         *log.Logger
	warningVar       *log.Logger
	errorVar         *log.Logger
	panicVar         *log.Logger
	flagEnableBuffer bool
	bufferData       []string
	flagVerbose      bool
	flagDebug        bool
	flagColor        bool
}

// LogInit is a member function for Log
// Inits all logging to given file handle except panic
// Default mode is "no color"
func (l *Log) Init(stdOut, stdErr io.Writer) {
	stdFlags := log.Ldate | log.Ltime | log.Lmsgprefix

	l.stdVar = log.New(stdOut, "       ", stdFlags)
	l.infoVar = log.New(stdOut, "INFO:  ", stdFlags)
	l.warningVar = log.New(stdOut, "WARN:  ", stdFlags)
	l.debugVar = log.New(stdErr, "DEBUG: ", stdFlags)
	l.errorVar = log.New(stdErr, "ERROR: ", stdFlags)
	l.panicVar = log.New(os.Stderr, "PANIC: ", stdFlags)

	l.flagEnableBuffer = false
}

func (l *Log) SetFlags(flags int) {
	l.stdVar.SetFlags(flags)
	l.infoVar.SetFlags(flags)
	l.warningVar.SetFlags(flags)
	l.debugVar.SetFlags(flags)
	l.errorVar.SetFlags(flags)
	l.panicVar.SetFlags(flags)
}

func (l *Log) SetColorPrefix() {
	if l.flagColor {
		l.infoVar.SetPrefix(green("INFO:  "))
		l.warningVar.SetPrefix(yellow("WARN:  "))
		l.debugVar.SetPrefix(red("DEBUG: "))
		l.errorVar.SetPrefix(red("ERROR: "))
		l.panicVar.SetPrefix(red("PANIC: "))
	}
}

func (l *Log) SetOutput(stdOut, stdErr io.Writer) {
	l.stdVar.SetOutput(stdOut)
	l.infoVar.SetOutput(stdOut)
	l.warningVar.SetOutput(stdOut)
	l.debugVar.SetOutput(stdErr)
	l.errorVar.SetOutput(stdErr)
	l.panicVar.SetOutput(stdErr)
}

func (l *Log) SetInteractive() {
	l.SetFlags(log.Lmsgprefix)
	l.SetColorPrefix()
}

func (l *Log) EnableBuffer() {
	l.flagEnableBuffer = true
}

func (l *Log) DisableBuffer() {
	l.flagEnableBuffer = false
}

func (l *Log) SetVerbose(b bool) {
	l.flagVerbose = b
}

func (l *Log) SetDebug(b bool) {
	l.flagDebug = b
}

func (l *Log) SetColor(b bool) {
	l.flagColor = b
}

// Buffer Handling
func (l *Log) AddBuffer(format string, v ...interface{}) {
	if l.flagEnableBuffer {
		l.bufferData = append(l.bufferData, fmt.Sprintf(format, v...))
	}
}

func (l *Log) GetBuffer() string {
	return strings.Join(l.bufferData, "\n")
}

// Intrinsic functions
func (l *Log) log(format string, v ...interface{}) {
	l.stdVar.Printf(format, v...)
	l.AddBuffer(format, v...)
}

func (l *Log) info(format string, v ...interface{}) {
	l.infoVar.Printf(green(format), v...)
	l.AddBuffer(format, v...)
}

func (l *Log) warn(format string, v ...interface{}) {
	l.warningVar.Printf(yellow(format), v...)
	l.AddBuffer(format, v...)
}

func (l *Log) debug(format string, v ...interface{}) {
	l.debugVar.Printf(red(format), v...)
	l.AddBuffer(format, v...)
}

func (l *Log) error(format string, v ...interface{}) {
	l.errorVar.Printf(red(format), v...)
	l.AddBuffer(format, v...)
}

// User functions
func (l *Log) Panic(format string, v ...interface{}) {
	l.panicVar.Printf(red(format), v...)
}

func (l *Log) Standard(format string, v ...interface{}) {
	l.log(format, v...)
}

func (l *Log) StandardInfo(format string, v ...interface{}) {
	l.info(format, v...)
}

func (l *Log) Verbose(format string, v ...interface{}) {
	if l.flagVerbose {
		l.log(format, v...)
	}
}

func (l *Log) VerboseInfo(format string, v ...interface{}) {
	if l.flagVerbose {
		l.info(format, v...)
	}
}

func (l *Log) Debug(format string, v ...interface{}) {
	if l.flagDebug {
		l.debug(format, v...)
	}
}

func (l *Log) Warn(format string, v ...interface{}) {
	l.warn(format, v...)
}

func (l *Log) Error(format string, v ...interface{}) {
	l.error(format, v...)
}
