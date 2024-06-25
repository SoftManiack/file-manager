package service

import (
	files "file-manager/dto"
	"file-manager/pkg/repository"
	"fmt"
)

type FilesService struct {
	repo repository.Files
}

func NewFilesService(repo repository.Files) *FilesService {
	return &FilesService{repo: repo}
}

func (s *FilesService) UploadFile(rootUid, name string, size int64) (files.File, error) {

	var newFile files.File

	fmt.Println(rootUid)
	fmt.Println(name)
	fmt.Println(size)

	return newFile, nil
}

/* func (s *FilesService) CreateTextFile(uid string, input files.NewTextFile) (files.TextFile, error) {

}
*/
