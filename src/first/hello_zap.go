package main

import (
	"go.uber.org/zap"
)

func main() {
	var sugared = zap.NewExample().Sugar()
	sugared.Infof("out zap!")
	sugared.Errorf("err zap!")
}
