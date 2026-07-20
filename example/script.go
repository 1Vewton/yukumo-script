package example

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/1Vewton/yukumo-script/generator/generatorwin"
	"github.com/1Vewton/yukumo-script/language/all2jap"
	"github.com/1Vewton/yukumo-script/utils/logger"
	"golang.org/x/sync/errgroup"
)

var scriptLogger = logger.NewLogger("Example", nil)

// GenerateExampleWin generates examples for phont file in win64
func GenerateExampleWin(
	ctx context.Context,
	targetDir string,
	phontDir string,
	files []os.DirEntry,
) error {
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(runtime.NumCPU() * 2)
	for _, i := range files {
		if !i.IsDir() {
			name := i.Name()
			group.Go(
				func() error {
					phontFile := fmt.Sprintf(
						"%s/%s",
						phontDir,
						name,
					)
					targetFile := fmt.Sprintf(
						"%s/example_%s.wav",
						targetDir,
						strings.TrimSuffix(
							name,
							filepath.Ext(name),
						),
					)
					scriptLogger.Info(
						fmt.Sprintf(
							"Phont: %s, Target: %s",
							phontFile,
							targetFile,
						),
					)
					_, errStat := os.Stat(targetFile)
					if errStat == nil {
						return nil
					} else if os.IsNotExist(errStat) {
						generatorW := generatorwin.NewGeneratorWin(
							100,
							phontFile,
							targetFile,
							all2jap.AllToKana("僕はGopherです。"),
						)
						err := generatorW.GenerateWav()
						return err
					} else {
						return errStat
					}
				},
			)
		}
	}
	err := group.Wait()
	return err
}
