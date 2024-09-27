package service

import (
	directories "file-manager/dto"
	files "file-manager/dto"
	"file-manager/pkg/repository"
	"os"
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

func (s *TrashService) DeleteTrashFile(deleteFile files.DeleteFile, uidUser string) error {

	err := s.repo.DeleteTrashFile(deleteFile.Uid)
	var path string

	if err != nil {
		return err
	}

	if deleteFile.RootUid == uidUser {
		path = os.Getenv("PATH_FILES") + "/" + uidUser + "/" + deleteFile.Name
	} else {
		path = os.Getenv("PATH_FILES") + "/" + uidUser + "/" + deleteFile.RootUid + "/" + deleteFile.Name
	}

	os.Remove(path)

	return nil
}

func (s *TrashService) DeleteTrashDirectory(uidDir, uidUser string) error {

	return nil
}

func (s *TrashService) RemoveTrashFile(uidFile string) error {

	err := s.repo.RemoveTrashFile(uidFile)

	if err != nil {

		return err
	}

	return nil
}
