package logger

// Level is a log level.
type Level int

// Log levels.
const (
	Debug Level = iota + 1
	Info
	Warn
	Error
)

func GetLvStr(level Level) string {
	if level == Debug {
		return "DEB:WmpLogLv"
	} else if level == Info {
		return "INF:WmpLogLv"
	} else if level == Warn {
		return "WAR:WmpLogLv"
	}
	return "ERR:WmpLogLv"
}
