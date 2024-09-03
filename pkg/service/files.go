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

func (s *FilesService) UploadFile(rootUid string, newFile files.NewFile) (files.File, error) {

	fmt.Println(newFile)

	return files.File{}, nil
}

/* func (s *FilesService) CreateTextFile(uid string, input files.NewTextFile) (files.TextFile, error) {

}
*/
