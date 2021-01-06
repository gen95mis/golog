package golog

import (
	"errors"
	"fmt"
	"io"
	"log"
)

type Level int

const (
	LvlOff   Level = 0
	LvlPanic Level = 1 << iota
	LvlFatal
	LvlError
	LvlWarn
	LvlInfo
	LvlDebug
)

const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	Lmsgprefix    = log.Lmsgprefix
	LstdFlags     = log.LstdFlags
)

type Logger struct {
	prefix string
	lvl    Level
	logger *log.Logger
}

// Создание нового журнала для регистрации
func NewLogger(out io.Writer, prefix, level string, flag int) (*Logger, error) {
	lvl, err := parsLevel(level)
	if err != nil {
		return nil, err
	}

	l := new(Logger)
	l.prefix = prefix
	l.lvl = lvl
	l.logger = log.New(out, "", flag)

	return l, nil
}

func parsLevel(level string) (Level, error) {
	switch level {
	case "Off":
		return LvlOff, nil
	case "Panic":
		return LvlPanic, nil
	case "Fatal":
		return LvlFatal, nil
	case "Error":
		return LvlError, nil
	case "Warn":
		return LvlWarn, nil
	case "Info":
		return LvlInfo, nil
	case "Debug":
		return LvlDebug, nil
	default:
		return 0, errors.New("logger's level is not exist")
	}
}

// Panicf эквивалентно вызову panicf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panicf(format string, args ...interface{}) {
	const lvl = LvlPanic
	const plvl = "[Panic]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)

	l.logger.Panicf(prefix, args)
	log.Panicf(prefix, args)
}

// Panicln эквивалентно вызову panicln() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panic(args ...interface{}) {
	const lvl = LvlPanic
	const plvl = "[Panic]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s", plvl, l.prefix)

	l.logger.Panicln(prefix, args)
	log.Panicln(prefix, args)
}

// Fatalf эквивалентно вызову fatalf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatalf(format string, args ...interface{}) {
	const lvl = LvlFatal
	const plvl = "[FATAL]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)

	l.logger.Printf(prefix, args)
	log.Fatalf(prefix, args)
}

// Fatal эквивалентно вызову fatalln() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatal(args ...interface{}) {
	const lvl = LvlFatal
	const plvl = "[FATAL]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s", plvl, l.prefix)

	l.logger.Fatalln(prefix, args)
	log.Fatalln(prefix, args)
}

// Errorf  эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Errorf(format string, args ...interface{}) {
	go func() {
		const lvl = LvlError
		const plvl = "[ERROR]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, args)
	}()
}

// Error  эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Error(args ...interface{}) {
	go func() {
		const lvl = LvlError
		const plvl = "[ERROR]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Printf(prefix, args)
	}()
}

// Warnf эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warnf(format string, args ...interface{}) {
	go func() {
		const lvl = LvlWarn
		const plvl = "[WARN]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, args)
	}()
}

// Warn эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warn(args ...interface{}) {
	go func() {
		const lvl = LvlWarn
		const plvl = "[WARN]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(prefix, args)
	}()
}

// Infof эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Infof(format string, args ...interface{}) {
	go func() {
		const lvl = LvlInfo
		const plvl = "[INFO]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, args)
	}()
}

// Info эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Info(args ...interface{}) {
	go func() {
		const lvl = LvlInfo
		const plvl = "[INFO]"
		if !l.checkLevel(lvl) {
			return
		}

		newFormat := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(newFormat, args)
	}()
}

// Debugf эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debugf(format string, args ...interface{}) {
	go func() {
		const lvl = LvlDebug
		const plvl = "[DEBUG]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, args)
	}()
}

// Debug эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debug(args ...interface{}) {
	go func() {
		const lvl = LvlDebug
		const plvl = "[DEBUG]"
		if !l.checkLevel(lvl) {
			return
		}

		newFormat := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(newFormat, args)
	}()
}

func (l *Logger) checkLevel(lvl Level) bool {
	return l.lvl >= lvl
}
