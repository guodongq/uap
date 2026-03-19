package main

import (
	uapfx "github.com/guodongq/uap/adapters/fx"
	"github.com/guodongq/uap/logging/logrus"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		uapfx.Module(),
		logrus.Module(),
	).Run()
}
