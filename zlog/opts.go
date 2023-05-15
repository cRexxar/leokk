package zlog

import "github.com/rs/zerolog"

type options func(l *Logger)

func WithMaxMBytes(maxMBytes int) options {
	return func(l *Logger) {
		l.maxMBytes = maxMBytes
	}
}

func WithBackupCount(backupCount int) options {
	return func(l *Logger) {
		l.backupCount = backupCount
	}
}

func WithLevel(lv level) options {
	return func(l *Logger) {
		l.log = l.log.Level(zerolog.Level(lv))
	}
}
