package types

import (
	"os"
	"path/filepath"
)

const (
	DefaultHttpListenAddr = "0.0.0.0:8008"
	DefaultDatasourceName = "mysql://root:123456@127.0.0.1:3306/coin-ant?charset=utf8mb4"
)

var (
	DefaultImagesHome = os.ExpandEnv(filepath.Join("$HOME", ".admin-system/images"))
	DefaultConfigHome = os.ExpandEnv(filepath.Join("$HOME", ".admin-system/config"))
)
