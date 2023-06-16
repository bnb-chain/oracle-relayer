package main

import (
	"flag"
	"fmt"

	"github.com/bnb-chain/go-sdk/common/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/binance-chain/oracle-relayer/admin"
	"github.com/binance-chain/oracle-relayer/executor/bbc"
	"github.com/binance-chain/oracle-relayer/executor/bsc"
	"github.com/binance-chain/oracle-relayer/model"
	"github.com/binance-chain/oracle-relayer/observer"
	"github.com/binance-chain/oracle-relayer/relayer"
	"github.com/binance-chain/oracle-relayer/util"
)

const (
	flagConfigType         = "config-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagConfigPath         = "config-path"
	flagBBCNetwork         = "bbc-network"
)

const (
	ConfigTypeLocal = "local"
	ConfigTypeAws   = "aws"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config path")
	flag.String(flagConfigType, "", "config type, local or aws")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")
	flag.Int(flagBBCNetwork, int(types.TestNetwork), "bbc chain network type")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("bind flags error, err=%s", err))
	}
}

func printUsage() {
	fmt.Print("usage: ./relayer --bbc-network [0 for testnet, 1 for mainnet] --config-type [local or aws] --config-path config_file_path\n")
}

func main() {
	initFlags()

	bbcNetwork := viper.GetInt(flagBBCNetwork)
	if bbcNetwork != int(types.TestNetwork) &&
		bbcNetwork != int(types.ProdNetwork) &&
		bbcNetwork != int(types.TmpTestNetwork) &&
		bbcNetwork != int(types.GangesNetwork) {
		printUsage()
		return
	}

	types.SetNetwork(types.ChainNetwork(bbcNetwork))

	configType := viper.GetString(flagConfigType)
	if configType == "" {
		printUsage()
		return
	}

	if configType != ConfigTypeAws && configType != ConfigTypeLocal {
		printUsage()
		return
	}

	var config *util.Config
	if configType == ConfigTypeAws {
		awsSecretKey := viper.GetString(flagConfigAwsSecretKey)
		if awsSecretKey == "" {
			printUsage()
			return
		}

		awsRegion := viper.GetString(flagConfigAwsRegion)
		if awsRegion == "" {
			printUsage()
			return
		}

		configContent, err := util.GetSecret(awsSecretKey, awsRegion)
		if err != nil {
			fmt.Printf("get aws config error, err=%s", err.Error())
			return
		}
		config = util.ParseConfigFromJson(configContent)
	} else {
		configFilePath := viper.GetString(flagConfigPath)
		if configFilePath == "" {
			printUsage()
			return
		}
		config = util.ParseConfigFromFile(configFilePath)
	}
	config.Validate()

	// init logger
	util.InitLogger(*config.LogConfig)
	util.InitAlert(config.AlertConfig)

	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.DBPath)
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	defer db.Close()
	model.InitTables(db)

	bscExecutor := bsc.NewExecutor(config.ChainConfig.BSCProviders, config)
	ob := observer.NewObserver(db, config, bscExecutor)
	go ob.Start()

	bbcExecutor, err := bbc.NewExecutor(config.ChainConfig.BBCRpcAddrs, types.ChainNetwork(bbcNetwork), config)
	if err != nil {
		fmt.Printf("new bbc executor error, err=%s\n", err.Error())
		return
	}
	oracleRelayer := relayer.NewRelayer(db, bbcExecutor, config)
	go oracleRelayer.Main()

	adm := admin.NewAdmin(config, bbcExecutor)
	go adm.Serve()

	select {}
}
