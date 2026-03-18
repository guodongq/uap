package fxevent

import (
	"fmt"
	"strings"

	logger "github.com/guodongq/uap/logging"
	"go.uber.org/fx/fxevent"
)

type FxEventLoggerAdapter struct {
	logger.Logger
}

func New(logger logger.Logger) *FxEventLoggerAdapter {
	return &FxEventLoggerAdapter{
		Logger: logger,
	}
}

func (l *FxEventLoggerAdapter) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Infof("HOOK OnStart\t\t%s executing (caller: %s)", e.FunctionName, e.CallerName)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Errorf("HOOK OnStart\t\t%s called by %s failed in %s: %+v", e.FunctionName, e.CallerName, e.Runtime, e.Err)
		} else {
			l.Infof("HOOK OnStart\t\t%s called by %s ran successfully in %s", e.FunctionName, e.CallerName, e.Runtime)
		}
	case *fxevent.OnStopExecuting:
		l.Infof("HOOK OnStop\t\t%s executing (caller: %s)", e.FunctionName, e.CallerName)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Errorf("HOOK OnStop\t\t%s called by %s failed in %s: %+v", e.FunctionName, e.CallerName, e.Runtime, e.Err)
		} else {
			l.Infof("HOOK OnStop\t\t%s called by %s ran successfully in %s", e.FunctionName, e.CallerName, e.Runtime)
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.Errorf("ERROR\tFailed to supply %v: %+v", e.TypeName, e.Err)
		} else if e.ModuleName != "" {
			l.Infof("SUPPLY\t%v from module %q", e.TypeName, e.ModuleName)
		} else {
			l.Infof("SUPPLY\t%v", e.TypeName)
		}
	case *fxevent.Provided:
		var privateStr string
		if e.Private {
			privateStr = " (PRIVATE)"
		}
		for _, rtype := range e.OutputTypeNames {
			if e.ModuleName != "" {
				l.Infof("PROVIDE%v\t%v <= %v from module %q", privateStr, rtype, e.ConstructorName, e.ModuleName)
			} else {
				l.Infof("PROVIDE%v\t%v <= %v", privateStr, rtype, e.ConstructorName)
			}
		}
		if e.Err != nil {
			l.Errorf("Error after options were applied: %+v", e.Err)
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			if e.ModuleName != "" {
				l.Infof("REPLACE\t%v from module %q", rtype, e.ModuleName)
			} else {
				l.Infof("REPLACE\t%v", rtype)
			}
		}
		if e.Err != nil {
			l.Errorf("ERROR\tFailed to replace: %+v", e.Err)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			if e.ModuleName != "" {
				l.Infof("DECORATE\t%v <= %v from module %q", rtype, e.DecoratorName, e.ModuleName)
			} else {
				l.Infof("DECORATE\t%v <= %v", rtype, e.DecoratorName)
			}
		}
		if e.Err != nil {
			l.Errorf("Error after options were applied: %+v", e.Err)
		}
	case *fxevent.BeforeRun:
		var moduleStr string
		if e.ModuleName != "" {
			moduleStr = fmt.Sprintf(" from module %q", e.ModuleName)
		}
		l.Infof("BEFORE RUN\t%s: %s%s", e.Kind, e.Name, moduleStr)
	case *fxevent.Run:
		var moduleStr string
		if e.ModuleName != "" {
			moduleStr = fmt.Sprintf(" from module %q", e.ModuleName)
		}
		l.Infof("RUN\t%v: %v in %s%v", e.Kind, e.Name, e.Runtime, moduleStr)
		if e.Err != nil {
			l.Errorf("Error returned: %+v", e.Err)
		}

	case *fxevent.Invoking:
		if e.ModuleName != "" {
			l.Infof("INVOKE\t\t%s from module %q", e.FunctionName, e.ModuleName)
		} else {
			l.Infof("INVOKE\t\t%s", e.FunctionName)
		}
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Errorf("ERROR\t\tfx.Invoke(%v) called from:\n%+vFailed: %+v", e.FunctionName, e.Trace, e.Err)
		}
	case *fxevent.Stopping:
		l.Infof("%v", strings.ToUpper(e.Signal.String()))
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Errorf("ERROR\t\tFailed to stop cleanly: %+v", e.Err)
		}
	case *fxevent.RollingBack:
		l.Infof("ERROR\t\tStart failed, rolling back: %+v", e.StartErr)
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Errorf("ERROR\t\tCouldn't roll back cleanly: %+v", e.Err)
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Errorf("ERROR\t\tFailed to start: %+v", e.Err)
		} else {
			l.Infof("RUNNING")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Errorf("ERROR\t\tFailed to initialize custom logger: %+v", e.Err)
		} else {
			l.Infof("LOGGER\tInitialized custom logger from %v", e.ConstructorName)
		}
	}
}
