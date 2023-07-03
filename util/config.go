package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	ethcmm "github.com/ethereum/go-ethereum/common"

	"github.com/binance-chain/oracle-relayer/common"
)

const (
	KeyTypeMnemonic    = "mnemonic"
	KeyTypeAWSMnemonic = "aws_mnemonic"
)

type Config struct {
	DBConfig    *DBConfig    `json:"db_config"`
	ChainConfig *ChainConfig `json:"chain_config"`
	LogConfig   *LogConfig   `json:"log_config"`
	AlertConfig *AlertConfig `json:"alert_config"`
	AdminConfig *AdminConfig `json:"admin_config"`
}

func (cfg *Config) Validate() {
	cfg.DBConfig.Validate()
	cfg.ChainConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.AlertConfig.Validate()
}

type AlertConfig struct {
	Moniker string `json:"moniker"`

	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`

	PagerDutyAuthToken string `json:"pager_duty_auth_token"`

	BlockUpdateTimeOut         int64 `json:"block_update_time_out"`
	PackageDelayAlertThreshold int64 `json:"package_delay_alert_threshold"`
}

func (cfg *AlertConfig) Validate() {
	if cfg.Moniker == "" {
		panic("moniker should not be empty")
	}

	if cfg.BlockUpdateTimeOut <= 0 {
		panic("block_update_time_out should be larger than 0")
	}

	if cfg.PackageDelayAlertThreshold <= 0 {
		panic("package_delay_alert_threshold should be larger than 0")
	}
}

type DBConfig struct {
	Dialect string `json:"dialect"`
	DBPath  string `json:"db_path"`
}

func (cfg *DBConfig) Validate() {
	if cfg.Dialect != common.DBDialectMysql && cfg.Dialect != common.DBDialectSqlite3 {
		panic(fmt.Sprintf("only %s and %s supported", common.DBDialectMysql, common.DBDialectSqlite3))
	}
	if cfg.DBPath == "" {
		panic("db path should not be empty")
	}
}

type ChainConfig struct {
	BSCStartHeight               int64          `json:"bsc_start_height"`
	BSCProviders                 []string       `json:"bsc_providers"`
	BSCConfirmNum                int64          `json:"bsc_confirm_num"`
	BSCChainId                   uint16         `json:"bsc_chain_id"`
	BSCCrossChainContractAddress ethcmm.Address `json:"bsc_cross_chain_contract_address"`

	BBCRpcAddrs      []string `json:"bbc_rpc_addrs"`
	BBCMnemonic      string   `json:"bbc_mnemonic"`
	BBCKeyType       string   `json:"bbc_key_type"`
	BBCAWSRegion     string   `json:"bbc_aws_region"`
	BBCAWSSecretName string   `json:"bbc_aws_secret_name"`

	RelayInterval int64 `json:"relay_interval"`
}

func (cfg *ChainConfig) Validate() {
	if cfg.BSCStartHeight < 0 {
		panic("bsc_start_height should not be less than 0")
	}
	if len(cfg.BSCProviders) == 0 {
		panic("bsc_providers should not be empty")
	}
	if cfg.BSCConfirmNum <= 0 {
		panic("bsc_confirm_num should be larger than 0")
	}

	// replace bsc_confirm_num if it is less than DefaultConfirmNum
	if cfg.BSCConfirmNum <= common.DefaultConfirmNum {
		cfg.BSCConfirmNum = common.DefaultConfirmNum
	}

	var emptyAddr ethcmm.Address
	if cfg.BSCCrossChainContractAddress.String() == emptyAddr.String() {
		panic("bsc_token_hub_contract_address should not be empty")
	}

	if len(cfg.BBCRpcAddrs) == 0 {
		panic("bbc_rpc_addrs should not be empty")
	}
	if cfg.BBCKeyType != KeyTypeMnemonic && cfg.BBCKeyType != KeyTypeAWSMnemonic {
		panic(fmt.Sprintf("bbc_key_type of bnb beacon chain chain only supports %s and %s", KeyTypeMnemonic, KeyTypeAWSMnemonic))
	}
	if cfg.BBCKeyType == KeyTypeAWSMnemonic && cfg.BBCAWSRegion == "" {
		panic("bbc_aws_region of bnb beacon chain chain should not be empty")
	}
	if cfg.BBCKeyType == KeyTypeAWSMnemonic && cfg.BBCAWSSecretName == "" {
		panic("bbc_aws_secret_name of bnb beacon chain chain should not be empty")
	}
	if cfg.BBCKeyType == KeyTypeMnemonic && cfg.BBCMnemonic == "" {
		panic("bbc_mnemonic should not be empty")
	}

	if cfg.RelayInterval <= 0 {
		panic(fmt.Sprintf("relay interval should be larger than 0"))
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

type AdminConfig struct {
	ListenAddr string `json:"listen_addr"`
}

// ParseConfigFromFile returns the config from json file
func ParseConfigFromFile(filePath string) *Config {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	return &config
}

// ParseConfigFromJson returns the config from json string
func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	return &config
}
