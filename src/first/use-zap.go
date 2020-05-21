package main

import (
	"errors"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewDevelopment()
	err := errors.New("bad input")
	log.Sugar().Info("request failed", "err", err)
	log.Sugar().Infow("request failed ", "err", err)
	err = nil
	log.Sugar().Info("request failed ", "err", err)
	log.Sugar().Infow("request failed ", "err", err)
}
