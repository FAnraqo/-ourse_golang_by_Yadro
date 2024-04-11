package calculations

import (
	"github.com/sirupsen/logrus"
)

func Calculate(n int64, log bool) int64 {
	if log {
		logrus.Infof("Start calculations...\nCalculate %d!\n", n)
	}

	var rez int64 = 1

	for i := n; i >= 1; i-- {
		rez *= i
	}

	if log {
		logrus.Info("Calculations complete!\n")
	}

	return rez
}
