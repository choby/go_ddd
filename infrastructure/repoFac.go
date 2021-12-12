package infrastructure

import (
	"github.com/choby/go_ddd/infrastructure/repos_mysql"
	"github.com/choby/go_ddd/infrastructure/repos_redis"
	"github.com/choby/go_ddd/interfaces"
)

func init() {
	RepoFac.CaptchaRepo = repos_redis.NewcaptchaRepo()

	RepoFac.FilesRepo = repos_mysql.NewfileRepo()
}

var (
	RepoFac *RepoFactory = &RepoFactory{}
	Empty   interface{}  = struct{}{}
)

type RepoFactory struct {
	CaptchaRepo interfaces.ICaptchaRepo
	FilesRepo   interfaces.IFileInfoRepo
}
