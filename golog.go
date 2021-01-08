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

var LevelToStrung = map[Level]string{
	LvlPanic: "Panic",
	LvlFatal: "Fatal",
	LvlError: "Error",
	LvlWarn:  "Warn",
	LvlInfo:  "Info",
	LvlDebug: "Debug",
}

var StringToLevel = map[string]Level{
	"Off":   LvlOff,
	"Panic": LvlPanic,
	"Fatal": LvlFatal,
	"Error": LvlError,
	"Warn":  LvlWarn,
	"Info":  LvlInfo,
	"Debug": LvlDebug,
}

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
	lvl, ok := StringToLevel[level]
	if !ok {
		return 0, errors.New("logger's level is not exist")
	}
	return lvl, nil
}

// Panic эквивалентно вызову Panic из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panic(v ...interface{}) {
	const lvl = LvlPanic
	l.write(lvl, v)
	Panic(v)
}

// Panicf эквивалентно вызову Panicf из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panicf(format string, v ...interface{}) {
	const lvl = LvlPanic
	l.writef(lvl, format, v)
	Panicf(format, v)
}

// Panicln эквивалентно вызову Panicln из стандартной библиотеки golang с записью в журнал
func (l *Logger) Panicln(v ...interface{}) {
	const lvl = LvlPanic
	l.writeln(lvl, v)
	Panicln(v)
}

// Fatal эквивалентно вызову Fatal из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatal(v ...interface{}) {
	const lvl = LvlFatal
	l.write(lvl, v)
	Fatal(v)
}

// Fatalf эквивалентно вызову Fatalf из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatalf(format string, v ...interface{}) {
	const lvl = LvlFatal
	l.writef(lvl, format, v)
	Fatalf(format, v)
}

// Fatalln эквивалентно вызову Fatalln из стандартной библиотеки golang с записью в журнал
func (l *Logger) Fatalln(v ...interface{}) {
	const lvl = LvlFatal
	l.writeln(lvl, v)
	Fatalln(v)
}

// Error эквивалентно вызову Print из стандартной библиотеки golang с записью в журнал
func (l *Logger) Error(v ...interface{}) {
	const lvl = LvlError
	go l.write(lvl, v)
}

// Errorf эквивалентно вызову Printf из стандартной библиотеки golang с записью в журнал
func (l *Logger) Errorf(format string, v ...interface{}) {
	const lvl = LvlError
	go l.writef(lvl, format, v)
}

// Errorln эквивалентно вызову Println из стандартной библиотеки golang с записью в журнал
func (l *Logger) Errorln(v ...interface{}) {
	const lvl = LvlError
	go l.writeln(lvl, v)
}

// Warn эквивалентно вызову Print из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warn(v ...interface{}) {
	const lvl = LvlWarn
	go l.write(lvl, v)
}

// Warnf эквивалентно вызову Printf из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warnf(format string, v ...interface{}) {
	const lvl = LvlWarn
	go l.writef(lvl, format, v)
}

// Warn эквивалентно вызову Println из стандартной библиотеки golang с записью в журнал
func (l *Logger) Warnln(v ...interface{}) {
	const lvl = LvlWarn
	go l.writeln(lvl, v)
}

// Info эквивалентно вызову Print из стандартной библиотеки golang с записью в журнал
func (l *Logger) Info(v ...interface{}) {
	const lvl = LvlInfo
	go l.write(lvl, v)
}

// Infof эквивалентно вызову Printf из стандартной библиотеки golang с записью в журнал
func (l *Logger) Infof(format string, v ...interface{}) {
	const lvl = LvlInfo
	go l.writef(lvl, format, v)
}

// Infoln эквивалентно вызову Println из стандартной библиотеки golang с записью в журнал
func (l *Logger) Infoln(v ...interface{}) {
	const lvl = LvlInfo
	go l.writeln(lvl, v)
}

// Debug эквивалентно вызову Print из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debug(v ...interface{}) {
	const lvl = LvlDebug
	go l.write(lvl, v)
}

// Debugf эквивалентно вызову Printf из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debugf(format string, v ...interface{}) {
	const lvl = LvlDebug
	go l.writef(lvl, format, v)
}

// Debugln эквивалентно вызову Println из стандартной библиотеки golang с записью в журнал
func (l *Logger) Debugln(v ...interface{}) {
	const lvl = LvlDebug
	go l.writeln(lvl, v)
}

func (l *Logger) write(level Level, v ...interface{}) {
	if !l.checkLevel(level) {
		return
	}

	prefix := ""
	if l.prefix != "" {
		prefix = fmt.Sprintf("%s %s", LevelToStrung[level], l.prefix)
	} else {
		prefix = LevelToStrung[level]
	}

	l.logger.Print(prefix, " ", v)
}

func (l *Logger) writef(level Level, format string, v ...interface{}) {
	if !l.checkLevel(level) {
		return
	}

	prefix := ""
	if l.prefix != "" {
		prefix = fmt.Sprintf("%s %s %s", LevelToStrung[level], l.prefix, format)
	} else {
		prefix = fmt.Sprintf("%s %s", LevelToStrung[level], format)
	}

	l.logger.Printf(prefix, v)
}

func (l *Logger) writeln(level Level, v ...interface{}) {
	if !l.checkLevel(level) {
		return
	}

	prefix := ""
	if l.prefix != "" {
		prefix = fmt.Sprintf("%s %s", LevelToStrung[level], l.prefix)
	} else {
		prefix = LevelToStrung[level]
	}

	l.logger.Println(prefix, v)
}

func (l *Logger) checkLevel(lvl Level) bool {
	return l.lvl >= lvl
}

// Panic эквивалентно вызову Panic из стандартной библиотеки golang
func Panic(v ...interface{}) {
	log.Panic(v)
}

// Panicf эквивалентно вызову Panicf из стандартной библиотеки golang
func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v)
}

// Panicln эквивалентно вызову Panicln из стандартной библиотеки golang с записью в журнал
func Panicln(v ...interface{}) {
	log.Panicln(v)
}

// Fatal эквивалентно вызову Fatal из стандартной библиотеки golang с записью в журнал
func Fatal(v ...interface{}) {
	log.Fatal(v)
}

// Fatalf эквивалентно вызову Fatalf из стандартной библиотеки golang с записью в журнал
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

// Fatalln эквивалентно вызову Fatalln из стандартной библиотеки golang с записью в журнал
func Fatalln(v ...interface{}) {
	log.Fatalln(v)
}

// Print эквивалентно вызову Print из стандартной библиотеки golang с записью в журнал
func Print(v ...interface{}) {
	log.Print(v)
}

// Printf эквивалентно вызову Printf из стандартной библиотеки golang с записью в журнал
func Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

// Println эквивалентно вызову Println из стандартной библиотеки golang с записью в журнал
func Println(v ...interface{}) {
	log.Println(v)
}
