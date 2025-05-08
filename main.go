package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eugen-bondarev/backup-tool/config"
	"github.com/eugen-bondarev/backup-tool/dump"
	"github.com/eugen-bondarev/backup-tool/router"
	"github.com/eugen-bondarev/backup-tool/storage"
	"github.com/eugen-bondarev/backup-tool/util"
	"github.com/gin-gonic/gin"
)

const (
	modeRelease = "release"
	modeDebug   = "debug"
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
	Async  bool         `json:"async"`
}

type BackupResponse struct {
	Message string `json:"message"`
}

type App struct {
	storage storage.Storage
	cfg     config.Config
}

func NewApp(cfg config.Config, storage storage.Storage) *App {
	return &App{
		cfg:     cfg,
		storage: storage,
	}
}

func (a *App) dump(dumpCfg DumpConfig, outputCfg OutputConfig) (any, error) {
	tmpInPath := util.CreateTmpSqlFilePath()

	d := dump.NewDump(dumpCfg.Database, tmpInPath)
	err := d.Create(a.cfg, dump.DBType(dumpCfg.Type))
	if err != nil {
		return nil, router.NewHttpError(err.Error(), 500)
	}

	outPath := fmt.Sprintf("%s/%s", outputCfg.Bucket, outputCfg.Path)

	err = a.storage.Push(tmpInPath, outPath)
	defer os.Remove(tmpInPath)
	if err != nil {
		return nil, router.NewHttpError(err.Error(), 500)
	}

	return BackupResponse{
		Message: "success",
	}, nil
}

func (a *App) DumpController(ctx router.Ctx) (any, error) {
	var req BackupRequest
	err := ctx.GetBody(&req)
	if err != nil {
		return nil, router.NewHttpError(err.Error(), 400)
	}

	if req.Async {
		go a.dump(req.Dump, req.Output)
	} else {
		res, err := a.dump(req.Dump, req.Output)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return BackupResponse{
		Message: "success",
	}, nil
}

func main() {
	util.EnsureTmpDir()

	cfg := util.LoadEnv[config.Config]()

	s, err := storage.New(cfg)
	util.CheckErr(err)

	app := NewApp(cfg, s)

	if cfg.Mode == modeRelease {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	if cfg.BasicAuth.User != "" && cfg.BasicAuth.Password != "" {
		r.Use(gin.BasicAuth(gin.Accounts{
			cfg.BasicAuth.User: cfg.BasicAuth.Password,
		}))
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/dump", router.GinWrap(app.DumpController))
	}

	log.Printf("listening on 0.0.0.0:%d", cfg.Port)
	err = r.Run(fmt.Sprintf(":%d", cfg.Port))
	util.CheckErr(err)
}
