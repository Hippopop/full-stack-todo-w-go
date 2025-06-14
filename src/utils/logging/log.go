package logging

import (
	coreLog "log"
	system "os"
)

func BaseLog(options LogOptions) {
	if options.Tag != "" {
		coreLog.Printf("#%s(%s): %s", options.LogOptionsType.String(), options.Tag, options.Msg)
	} else {
		coreLog.Printf("#%s: %s", options.LogOptionsType.String(), options.Msg)
	}
}

func IfError(options LogOptions, error error) (exists bool) {
	exists = error != nil
	if exists && (options.Show || options.LogOptionsType == LogFatal) {
		if options.Tag != "" {
			coreLog.Printf("#%s(%s): %s -> %s", options.LogOptionsType.String(), options.Tag, options.Msg, error.Error())
		} else {
			coreLog.Printf("#%s: %s -> %s", options.LogOptionsType.String(), options.Msg, error.Error())
		}
	}

	if options.LogOptionsType == LogFatal && exists {
		system.Exit(1)
	}

	return exists
}
