package fswatcher

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"tools.filewatcher",
		fx.Provide(
			fx.Annotate(
				NewWatcher,
				fx.As(new(FileWatcher)),
			),
		),
	)
}
