package initialization

import (
	"path"
	"v3board/global"
	"v3board/lib/logrotate"

	"github.com/rs/zerolog"
)

func InitLog(debug bool) error {
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	fileName := path.Join(global.WorkDir, "logs", "v3board.log")

	fileLog := &logrotate.Logger{
		Filename:   fileName,
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false,
	}

	logger := zerolog.New(fileLog).With().Timestamp().Caller().Logger()

	global.Log = &logger

	return nil
}
