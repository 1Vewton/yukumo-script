package example

import (
	"github.com/1Vewton/yukumo-script/utils/logger"
	"github.com/1Vewton/yukumo-script/utils/syncutils"
)

var scriptLogger = logger.NewLogger("Example", nil)
var examplesMap = syncutils.NewMap()
