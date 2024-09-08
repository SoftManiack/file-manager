package service

import (
	files "file-manager/dto"
	"file-manager/pkg/repository"
)

type FilesService struct {
	repo repository.Files
}

func NewFilesService(repo repository.Files) *FilesService {
	return &FilesService{repo: repo}
}

func (s *FilesService) UploadFile(newFile files.NewFile) (files.File, error) {

	var file files.File

	file, err := s.repo.CreateFile(newFile)

	if err != nil {
		return file, err
	}

	return file, nil
}

func (s *FilesService) UpdateFile(updateFile files.UpdateFile) (files.File, error) {
	var file files.File

	file, err := s.repo.UpdateFile(updateFile)

	if err != nil {
		return file, err
	}

	return file, nil
}

/* func (s *FilesService) CreateTextFile(uid string, input files.NewTextFile) (files.TextFile, error) {

}
*/
