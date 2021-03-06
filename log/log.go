package log

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Flag byte

const (
	DebugFlag Flag = iota
	NoteFlag
	WarningFlag
	ErrorFlag
)

var (
	LogFlag       = ErrorFlag
	ErrorFormat   = "\033[31mERROR"
	WarningFormat = "\033[33mWARNING"
	NoteFormat    = "\033[36mNOTE"
	DebugFormat   = "\033[34mDEBUG"
	PanicFormat   = "\033[35mPANIC"
)

func SetLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		LogFlag = DebugFlag
	case "note":
		LogFlag = NoteFlag
	case "warning":
		LogFlag = WarningFlag
	case "error":
		LogFlag = ErrorFlag
	default:
		LogFlag = ErrorFlag
	}
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func prepareLog(raw string) string {
	logTime := time.Now()

	functionName := "???"
	packageName := "???"
	// Get filename + line
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)

		if f := runtime.FuncForPC(pc); f != nil {
			i := strings.LastIndex(f.Name(), "/")
			j := strings.Index(f.Name()[i+1:], ".")
			if j >= 1 {
				pkg, fun := f.Name()[:i+j+1], f.Name()[i+j+2:]
				functionName = fun
				packageName = path.Base(pkg)
			}
		}
	}
	v := fmt.Sprintf("%s %s %s:%d %s.%s:   ", rightPad2Len(raw, " ", 13), logTime.Format("2006-01-02 15:04:05.0000"), file, line, packageName, functionName)

	return v
}

func Error(args ...interface{}) {
	if LogFlag <= ErrorFlag {
		args = append([]interface{}{prepareLog(ErrorFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Print(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if LogFlag <= ErrorFlag {
		var buffer bytes.Buffer
		buffer.WriteString("%s ")
		buffer.WriteString(format)

		args = append([]interface{}{prepareLog(ErrorFormat)}, args...)
		args = append(args, "\033[0m\n")
		fmt.Printf(buffer.String(), args...)
	}
}

func Errorln(args ...interface{}) {
	if LogFlag <= ErrorFlag {
		args = append([]interface{}{prepareLog(ErrorFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Println(args...)
	}
}

func Warning(args ...interface{}) {
	if LogFlag <= WarningFlag {
		args = append([]interface{}{prepareLog(WarningFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Print(args...)
	}
}

func Warningf(format string, args ...interface{}) {
	if LogFlag <= WarningFlag {
		var buffer bytes.Buffer
		buffer.WriteString("%s ")
		buffer.WriteString(format)
		buffer.WriteString("%s")

		args = append([]interface{}{prepareLog(WarningFormat)}, args...)
		args = append(args, "\033[0m\n")
		fmt.Printf(buffer.String(), args...)
	}
}

func Warningln(args ...interface{}) {
	if LogFlag <= WarningFlag {
		args = append([]interface{}{prepareLog(WarningFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Println(args...)
	}
}

func Note(args ...interface{}) {
	if LogFlag <= NoteFlag {
		args = append([]interface{}{prepareLog(NoteFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Print(args...)
	}
}

func Notef(format string, args ...interface{}) {
	if LogFlag <= NoteFlag {
		var buffer bytes.Buffer
		buffer.WriteString("%s ")
		buffer.WriteString(format)
		buffer.WriteString("%s")

		args = append([]interface{}{prepareLog(NoteFormat)}, args...)
		args = append(args, "\033[0m\n")
		fmt.Printf(buffer.String(), args...)
	}
}

func Noteln(args ...interface{}) {
	if LogFlag <= NoteFlag {
		args = append([]interface{}{prepareLog(NoteFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Println(args...)
	}
}

func Debug(args ...interface{}) {
	if LogFlag <= NoteFlag {
		args = append([]interface{}{prepareLog(DebugFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Print(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if LogFlag <= DebugFlag {
		var buffer bytes.Buffer
		buffer.WriteString("%s ")
		buffer.WriteString(format)
		buffer.WriteString("%s")

		args = append([]interface{}{prepareLog(DebugFormat)}, args...)
		args = append(args, "\033[0m\n")
		fmt.Printf(buffer.String(), args...)
	}
}

func Debugln(args ...interface{}) {
	if LogFlag <= DebugFlag {
		args = append([]interface{}{prepareLog(DebugFormat)}, args...)
		args = append(args, "\033[0m")
		fmt.Println(args...)
	}
}

func Panic(args ...interface{}) {
	args = append([]interface{}{prepareLog(PanicFormat)}, args...)
	args = append(args, "\033[0m")
	fmt.Print(args...)
	panic(fmt.Sprintf("%v", args))
}

func Panicf(format string, args ...interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("%s ")
	buffer.WriteString(format)
	buffer.WriteString("%s")

	args = append([]interface{}{prepareLog(PanicFormat)}, args...)
	args = append(args, "\033[0m\n")
	fmt.Printf(buffer.String(), args...)
	panic(fmt.Sprintf("%v", args))
}

func Panicln(args ...interface{}) {
	args = append([]interface{}{prepareLog(PanicFormat)}, args...)
	args = append(args, "\033[0m")
	fmt.Println(args...)
	panic(fmt.Sprintf("%v", args))
}
