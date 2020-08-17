package common

import "time"

const (
	ObserverMaxBlockNumber = 10000
	ObserverPruneInterval  = 10 * time.Second
	ObserverAlertInterval  = 5 * time.Second
	ObserverFetchInterval  = 2 * time.Second

	PackageDelayAlertInterval = 5 * time.Second
)

const (
	DBDialectMysql   = "mysql"
	DBDialectSqlite3 = "sqlite3"
)

type BlockAndPackageLogs struct {
	Height          int64
	BlockHash       string
	ParentBlockHash string
	BlockTime       int64
	Packages        []interface{}
}
