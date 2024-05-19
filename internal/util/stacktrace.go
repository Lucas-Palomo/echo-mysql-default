package util

import (
	"github.com/ztrue/tracerr"
)

func GetStackTrace(err error) string {
	return tracerr.Sprint(tracerr.Wrap(err))
}
