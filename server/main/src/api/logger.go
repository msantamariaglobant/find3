package api

import "github.com/msantamariaglobant/find3server/main/src/logging"

var logger *logging.SeelogWrapper

func init() {
	var err error
	logger, err = logging.New()
	if err != nil {
		panic(err)
	}
	Debug(false)
}

func Debug(debugMode bool) {
	if debugMode {
		logger.SetLevel("debug")
	} else {
		logger.SetLevel("info")
	}
}
