package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/conf"
	"poly-bridge/monitor/healthmonitor"
	"runtime"
	"syscall"
)

func main() {
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var RelayerConfigPathFlag = cli.StringFlag{
	Name:  "relayerConfig",
	Usage: "Relayer config file `<path>`",
}

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Name = "poly-bridge Health Monitor Service"
	app.Usage = "poly-bridge Health Monitor Service"
	app.Action = StartServer
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2019 The PolyNetwork Authors"
	app.Flags = []cli.Flag{
		conf.ConfigPathFlag,
		RelayerConfigPathFlag,
	}
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func StartServer(ctx *cli.Context) {
	configFile := ctx.GlobalString("config")
	relayerConfigFile := ctx.GlobalString("relayerConfig")
	config := conf.NewConfig(configFile)
	if config == nil {
		logs.Error("startServer - read server config failed!")
		return
	}
	relayerConfig := conf.NewRelayerConfig(relayerConfigFile)
	if relayerConfig == nil {
		logs.Error("startServer - read relayer config failed!")
		return
	}
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s"}`, config.LogFile))
	cacheRedis.Init()
	basedef.ConfirmEnv(config.Env)
	healthmonitor.StartHealthMonitor(config, relayerConfig)
	for true {
		sig := waitSignal()
		if sig != syscall.SIGHUP {
			break
		} else {
			continue
		}
	}
	return
}

func waitSignal() os.Signal {
	exit := make(chan os.Signal, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(sc)
	go func() {
		for sig := range sc {
			logs.Info("monitor received signal:(%s).", sig.String())
			exit <- sig
			close(exit)
			break
		}
	}()
	sig := <-exit
	return sig
}
