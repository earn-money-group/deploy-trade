package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/earn-money-group/deploy-trade/app"
	"github.com/earn-money-group/deploy-trade/share"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var watch = &cli.Command{
	Name:  "watch",
	Usage: "watch to deployed vmpx",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    share.Rpc,
			Aliases: []string{"r"},
			Usage:   "get data rpc",
		},
	},
	Before: func(c *cli.Context) error {
		setLogLevel(c.String(share.Log))
		return nil
	},
	Action: watchAction,
}

func watchAction(c *cli.Context) (err error) {
	client, err := ethclient.Dial(c.String(share.Rpc))
	if err != nil {
		log.Fatal("dial rpc failed", zap.Error(err))
	}

	watch := app.NewWatch(client)

	go watch.Watch(context.Background())

	// watch singal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-stopChan
	log.Info("end...")

	return nil
}
