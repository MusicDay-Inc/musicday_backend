package repository

import (
	"errors"
	"server/internal/core"
)

type ReleaseRepo struct {
	data []core.Song
}

func NewReleaseRepo() *ReleaseRepo {
	var data []core.Song
	return &ReleaseRepo{data: data}
}

func (rep *ReleaseRepo) GetById(releaseId int) (core.Song, error) {
	//err := ((releaseId < 0 || releaseId > len(rep.users)  "error": "ok")
	//var err error
	var item core.Song
	if releaseId < 0 || releaseId >= len(rep.data) {
		err := errors.New("incorrect release id")
		return item, err
	}
	err := error(nil)
	return rep.data[releaseId], err
}
