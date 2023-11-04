// Package utils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2023-05-19 10:50
package utils

import "cnic/fairman-gin/global"

type SEE struct {
	TraceId string `json:"traceId,omitempty"`

	seeChan chan global.ChanLogModal
}

func (s *SEE) InitSEE() error {
	seeChan, ok := global.BaseSEEChan.Load(s.TraceId)
	if !ok {
		return Errorf("No progress information was found for this task，No progress information was found for this task。")
	}

	s.seeChan = seeChan.(chan global.ChanLogModal)

	return nil
}

func (s *SEE) SendLog(message string, level string, progress float64) {
	s.seeChan <- global.ChanLogModal{
		Message:  message,
		Level:    level,
		Progress: progress,
	}
}
