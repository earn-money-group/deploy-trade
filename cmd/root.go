package cmd

import (
	"os"

	"github.com/earn-money-group/deploy-trade/share"
	"github.com/urfave/cli/v2"
	"github.com/xiaobaiskill/kit/mlog"
	"go.uber.org/zap"
)

func MainRun() {
	svr := cli.NewApp()
	svr.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  share.Log,
			Value: "info",
			Usage: "set log level: debug info warn error fatal panic",
		},
	}
	svr.Before = func(c *cli.Context) error {
		setLogLevel(c.String(share.Log))
		return nil
	}
	appendCmdList(svr, watch, mint)

	err := svr.Run(os.Args)
	if err != nil {
		log.Fatal("cli fatal", zap.Error(err))
	}
}

func appendCmdList(app *cli.App, subcmd ...*cli.Command) {
	app.Commands = append(app.Commands, subcmd...)
}

func setLogLevel(l string) {
	switch l {
	case "info":
		mlog.SetLevel(zap.InfoLevel)
	case "debug":
		mlog.SetLevel(zap.DebugLevel)
	case "warn":
		mlog.SetLevel(zap.WarnLevel)
	case "error":
		mlog.SetLevel(zap.ErrorLevel)
	case "fatal":
		mlog.SetLevel(zap.FatalLevel)
	case "panic":
		mlog.SetLevel(zap.PanicLevel)
	default:
		mlog.SetLevel(zap.InfoLevel)
	}
}
