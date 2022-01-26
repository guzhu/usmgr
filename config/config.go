package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
)

// USmgr Config represents USmgr configuration dir env var.
const USMgrConfig = "USMGRCONFIG"

var (
	// USMgrConfigFile represents usmgr config file location.
	USMgrConfigFile = filepath.Join(USMgrHome(), "config.yml")
	// USMgrLogs represents usmgr log.
	USMgrLogs = filepath.Join(os.TempDir(), fmt.Sprintf("usmgr-%s.log", MustUSmgrUser()))
	// USMgrDumpDir represents a directory where usmgr screen dumps will be persisted.
	USMgrDumpDir = filepath.Join(os.TempDir(), fmt.Sprintf("usmgr-screens-%s", MustUSmgrUser()))
)

// USMgrHome returns usmgr configs home directory.
func USMgrHome() string {
	if env := os.Getenv(USMgrConfig); env != "" {
		return env
	}
	xdgUSmgrHome, err := xdg.ConfigFile("usmgr")
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create configuration directory for usmgr")
	}

	return xdgUSmgrHome
}
