package main

import (
	fxadapter "github.com/guodongq/uap/adapters/fx"
	"github.com/guodongq/uap/log/logrus"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fxadapter.Module(),
		logrus.Module(),
	).Run()
}
