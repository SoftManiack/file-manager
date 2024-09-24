package service

import (
	files "file-manager/dto"
	"file-manager/pkg/repository"
	"fmt"
)

type FilesService struct {
	repo  repository.Files
	trash repository.Trash
}

func NewFilesService(repo repository.Files, trash repository.Trash) *FilesService {
	return &FilesService{repo: repo, trash: trash}
}

func (s *FilesService) CreateFile(newFile files.NewFile) (files.File, error) {

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

func (s *FilesService) MoveTrashFile(uidUser, uidFile string) error {

	// записать в базу

	fmt.Println(uidUser)
	fmt.Println(uidFile)

	err := s.trash.MoveTrashFile(uidUser, uidFile)

	if err != nil {
		return err
	}

	return nil
}

/* func (s *FilesService) CreateTextFile(uid string, input files.NewTextFile) (files.TextFile, error) {

}
*/
