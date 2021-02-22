module github.com/binance-chain/oracle-relayer

go 1.13

require (
	github.com/PagerDuty/go-pagerduty v1.3.0
	github.com/aws/aws-sdk-go v1.25.48
	github.com/binance-chain/go-sdk v0.0.0-00010101000000-000000000000
	github.com/ethereum/go-ethereum v1.9.12
	github.com/golang/mock v1.4.3
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.12
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/prometheus/tsdb v0.7.1 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.0.0
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/tendermint v0.32.3
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/binance-chain/go-sdk => github.com/binance-chain/go-sdk v1.2.3-bscAlpha.0
	github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1
	github.com/zondax/hid => github.com/binance-chain/hid v0.9.1-0.20190807012304-e1ffd6f0a3cc
)
