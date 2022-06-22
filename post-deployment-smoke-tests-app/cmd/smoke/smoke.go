package smoke

import (
	"os"
	"post-deployment-smoke-tests/types"
	"post-deployment-smoke-tests/internal/config"
	"post-deployment-smoke-tests/internal/http"
)

func StartHttpCheck() {
	if _, err := os.Stat(types.HttpConfigName); err == nil {
		http.DoHttpCheck(config.GetHttpConfig())
	}
}