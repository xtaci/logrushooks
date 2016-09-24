package logrushooks

import (
	"path"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

type LineNoHook struct{}

func (hook LineNoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook LineNoHook) Fire(entry *logrus.Entry) error {
	i := 0
	for {
		if pc, file, line, ok := runtime.Caller(i); ok {
			funcName := runtime.FuncForPC(pc).Name()
			if !strings.Contains(funcName, "github.com/Sirupsen/logrus") && !strings.Contains(funcName, "github.com/xtaci/logrushooks") {
				entry.Data["file"] = path.Base(file)
				entry.Data["func"] = path.Base(funcName)
				entry.Data["line"] = line
				return nil
			}
			i++
		} else {
			break
		}
	}

	return nil
}
