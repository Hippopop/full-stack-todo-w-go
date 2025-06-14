package logging

type LogOptionsType int

const (
	LogInfo LogOptionsType = iota
	LogWarn
	LogError
	LogFatal // This will stop the program!
	LogDebug // Only print when the program is running on debug mode!
)

func (l LogOptionsType) String() string {
	switch l {
	case LogInfo:
		return "INFO"
	case LogWarn:
		return "WARN"
	case LogError:
		return "ERROR"
	case LogFatal:
		return "FATAL-ERROR"
	default:
		return "UNKNOWN"
	}
}
