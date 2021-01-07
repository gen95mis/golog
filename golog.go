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

// Panic эквивалентно вызову panic() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panic(v ...interface{}) {
	const lvl = LvlPanic
	const plvl = "[Panic]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s", plvl, l.prefix)

	l.logger.Panic(prefix, " ", v)
	log.Panic(prefix, " ", v)
}

// Panicf эквивалентно вызову panicf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panicf(format string, v ...interface{}) {
	const lvl = LvlPanic
	const plvl = "[Panic]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)

	l.logger.Panicf(prefix, v)
	log.Panicf(prefix, v)
}

// Panicln эквивалентно вызову panicln() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panicln(v ...interface{}) {
	const lvl = LvlPanic
	const plvl = "[Panic]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s", plvl, l.prefix)

	l.logger.Panicln(prefix, " ", v)
	log.Panicln(prefix, " ", v)
}

// Fatal эквивалентно вызову fatal() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatal(v ...interface{}) {
	const lvl = LvlFatal
	const plvl = "[FATAL]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s", plvl, l.prefix)

	l.logger.Fatal(prefix, " ", v)
	log.Fatal(prefix, " ", v)
}

// Fatalf эквивалентно вызову fatalf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatalf(format string, v ...interface{}) {
	const lvl = LvlFatal
	const plvl = "[FATAL]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)

	l.logger.Printf(prefix, v)
	log.Fatalf(prefix, v)
}

// Fatalln эквивалентно вызову fatalln() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatalln(v ...interface{}) {
	const lvl = LvlFatal
	const plvl = "[FATAL]"
	if !l.checkLevel(lvl) {
		return
	}

	prefix := fmt.Sprintf("%s %s", plvl, l.prefix)

	l.logger.Fatalln(prefix, v)
	log.Fatalln(prefix, v)
}

// Error эквивалентно вызову Print() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Error(v ...interface{}) {
	go func() {
		const lvl = LvlError
		const plvl = "[ERROR]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Print(prefix, " ", v)
	}()
}

// Errorf эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Errorf(format string, v ...interface{}) {
	go func() {
		const lvl = LvlError
		const plvl = "[ERROR]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, v)
	}()
}

// Errorln эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Errorln(v ...interface{}) {
	go func() {
		const lvl = LvlError
		const plvl = "[ERROR]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(prefix, v)
	}()
}

// Warn эквивалентно вызову Print() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warn(v ...interface{}) {
	go func() {
		const lvl = LvlWarn
		const plvl = "[WARN]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Printf(prefix, " ", v)
	}()
}

// Warnf эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warnf(format string, v ...interface{}) {
	go func() {
		const lvl = LvlWarn
		const plvl = "[WARN]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, v)
	}()
}

// Warn эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warnln(v ...interface{}) {
	go func() {
		const lvl = LvlWarn
		const plvl = "[WARN]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(prefix, v)
	}()
}

// Info эквивалентно вызову Print() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Info(v ...interface{}) {
	go func() {
		const lvl = LvlInfo
		const plvl = "[INFO]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Print(prefix, " ", v)
	}()
}

// Infof эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Infof(format string, v ...interface{}) {
	go func() {
		const lvl = LvlInfo
		const plvl = "[INFO]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, v)
	}()
}

// Infoln эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Infoln(v ...interface{}) {
	go func() {
		const lvl = LvlInfo
		const plvl = "[INFO]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(prefix, v)
	}()
}

// Debug эквивалентно вызову Print() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debug(v ...interface{}) {
	go func() {
		const lvl = LvlDebug
		const plvl = "[DEBUG]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(prefix, v)
	}()
}

// Debugf эквивалентно вызову Printf() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debugf(format string, v ...interface{}) {
	go func() {
		const lvl = LvlDebug
		const plvl = "[DEBUG]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s %s", plvl, l.prefix, format)
		l.logger.Printf(prefix, v)
	}()
}

// Debugln эквивалентно вызову Println() из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debugln(v ...interface{}) {
	go func() {
		const lvl = LvlDebug
		const plvl = "[DEBUG]"
		if !l.checkLevel(lvl) {
			return
		}

		prefix := fmt.Sprintf("%s %s", plvl, l.prefix)
		l.logger.Println(prefix, v)
	}()
}

func (l *Logger) Print(v ...interface{}) {
	l.logger.Print(v)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Panicf(format, v)
}

func (l *Logger) Println(v ...interface{}) {
	l.logger.Println(v)
}

func (l *Logger) checkLevel(lvl Level) bool {
	return l.lvl >= lvl
}
