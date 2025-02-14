package main

import (
	"fmt"

	"github.com/eugen-bondarev/backup-tool/config"
	"github.com/eugen-bondarev/backup-tool/dump"
	"github.com/eugen-bondarev/backup-tool/router"
	"github.com/eugen-bondarev/backup-tool/storage"
	"github.com/eugen-bondarev/backup-tool/util"
	"github.com/gin-gonic/gin"
)

type DumpConfig struct {
	Database string `json:"database"`
	Type     string `json:"type"`
}

type OutputConfig struct {
	Bucket string `json:"bucket"`
	Path   string `json:"path"`
}

type BackupRequest struct {
	Dump   DumpConfig   `json:"dump"`
	Output OutputConfig `json:"output"`
}

type BackupResponse struct {
	Message string `json:"message"`
}

func main() {
	util.EnsureTmpDir()

	cfg := util.LoadEnv[config.Config]()

	s, err := storage.NewGCPStorage(cfg.GCPConfig)
	util.CheckErr(err)

	r := gin.Default()

	r.POST("/", router.GinWrap(func(ctx router.Ctx) (any, error) {
		var req BackupRequest
		err := ctx.GetBody(&req)
		if err != nil {
			return nil, router.NewHttpError(err.Error(), 400)
		}

		tmpInPath := util.CreateTmpSqlFilePath()

		d := dump.NewDump(req.Dump.Database, tmpInPath)
		err = d.Create(cfg, dump.DBType(req.Dump.Type))
		if err != nil {
			return nil, router.NewHttpError(err.Error(), 500)
		}

		outPath := fmt.Sprintf("%s/%s", req.Output.Bucket, req.Output.Path)

		err = s.Push(tmpInPath, outPath)
		if err != nil {
			return nil, router.NewHttpError(err.Error(), 500)
		}

		return BackupResponse{
			Message: "success",
		}, nil
	}))

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
