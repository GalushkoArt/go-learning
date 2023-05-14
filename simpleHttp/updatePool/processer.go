package updatePool

import (
	"time"
)

type processor struct {
}

func newProcessor() *processor {
	return &processor{}
}

func (w *processor) process(j Job) Result {
	result := Result{}
	now := time.Now()

	resp, err := j.Execute()
	if err != nil {
		result.Error = err

		return result
	}
	result.Result = resp
	result.ResponseTime = time.Since(now)

	return result
}
