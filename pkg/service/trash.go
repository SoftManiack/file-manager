package service

import (
	directories "file-manager/dto"
	files "file-manager/dto"
	"file-manager/pkg/repository"
)

type TrashService struct {
	repo repository.Trash
}

func NewTrashService(repo repository.Trash) *TrashService {
	return &TrashService{repo: repo}
}

func (s *TrashService) GetTrash(uidUser string) ([]directories.Directories, []files.File, error) {

	var getDrectories []directories.Directories
	var getFiles []files.File

	getDrectories, getFiles, err := s.repo.GetTrash(uidUser)

	if err != nil {
		return getDrectories, getFiles, err
	}

	return getDrectories, getFiles, nil
}

func (s *TrashService) DeleteTrashFile(uidFile, uidUser string) error {

	err := s.repo.DeleteTrashFile(uidFile)

	if err != nil {
		return err
	}

	//path := os.Getenv("PATH_FILES") + "/" + userUid + "/" + input.Name

	return nil
}

func (s *TrashService) DeleteTrashDirectory(uidDir, uidUser string) error {

	return nil
}
