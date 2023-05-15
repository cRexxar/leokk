package zlog

import (
	"fmt"
	"os"

	"github.com/cRexxar/leokk/files"

	"github.com/rs/zerolog"
)

type level zerolog.Level

const (
	logPath = "log"

	DEBUG level = level(zerolog.DebugLevel)
	INFO  level = level(zerolog.InfoLevel)
	WARN  level = level(zerolog.WarnLevel)
	ERROR level = level(zerolog.ErrorLevel)
	FATAL level = level(zerolog.FatalLevel)
	PANIC level = level(zerolog.PanicLevel)
)

func init() {
	files.CreateDir(logPath)
}

type Logger struct {
	log         zerolog.Logger
	file        *os.File
	fileName    string
	maxMBytes   int
	backupCount int
}

func NewLogger(fileName string, opts ...options) (*Logger, error) {
	fileName = logPath + "/" + fileName
	if out, err := files.OpenFile(fileName); err != nil {
		return nil, err
	} else {
		l := &Logger{
			log:         zerolog.New(out).Level(zerolog.InfoLevel).With().Timestamp().Logger(),
			file:        out,
			fileName:    fileName,
			maxMBytes:   500,
			backupCount: 2,
		}
		for _, opt := range opts {
			opt(l)
		}
		return l, nil
	}
}

func (l *Logger) rollover() error {
	f, err := l.file.Stat()
	if err != nil {
		return err
	}
	// not touch maxMBytes
	if f.Size() < int64(l.maxMBytes*1024*1024) {
		return nil
	}
	// touch maxMBytes and backupCount > 0
	if l.backupCount > 0 {
		l.file.Close()
		for i := l.backupCount - 1; i > 0; i-- {
			oldName := fmt.Sprintf("%s.%d", l.fileName, i)
			newName := fmt.Sprintf("%s.%d", l.fileName, i+1)
			os.Rename(oldName, newName)
		}
		os.Rename(l.fileName, l.fileName+".1")
		if out, err := files.OpenFile(l.fileName); err != nil {
			return err
		} else {
			l.file = out
			l.log = l.log.Output(out)
		}
	} else {
		l.file.Truncate(0)
		l.file.Seek(0, 0)
	}
	return nil
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.rollover()
	l.log.Debug().Msgf(msg, args...)
}

func (l *Logger) DebugWith(msg string, fields map[string]interface{}) {
	l.rollover()
	l.log.Debug().Fields(fields).Msg(msg)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.rollover()
	l.log.Info().Msgf(msg, args...)
}

func (l *Logger) InfoWith(msg string, fields map[string]interface{}) {
	l.rollover()
	l.log.Info().Fields(fields).Msg(msg)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.rollover()
	l.log.Warn().Msgf(msg, args...)
}

func (l *Logger) WarnWith(msg string, fields map[string]interface{}) {
	l.rollover()
	l.log.Warn().Fields(fields).Msg(msg)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.rollover()
	l.log.Error().Msgf(msg, args...)
}

func (l *Logger) ErrorWith(msg string, fields map[string]interface{}) {
	l.rollover()
	l.log.Error().Fields(fields).Msg(msg)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.rollover()
	l.log.Fatal().Msgf(msg, args...)
}

func (l *Logger) FatalWith(msg string, fields map[string]interface{}) {
	l.rollover()
	l.log.Fatal().Fields(fields).Msg(msg)
}

func (l *Logger) Panic(msg string, args ...interface{}) {
	l.rollover()
	l.log.Panic().Msgf(msg, args...)
}

func (l *Logger) PanicWith(msg string, fields map[string]interface{}) {
	l.rollover()
	l.log.Panic().Fields(fields).Msg(msg)
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func (l *Logger) WithLevel(lv level) {
	l.log = l.log.Level(zerolog.Level(lv))
}
