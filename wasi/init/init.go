package init

import (
	"github.com/pingcap/log"
	"go.uber.org/zap"
)

func init() {
	log.SetLevel(zap.FatalLevel)
}
