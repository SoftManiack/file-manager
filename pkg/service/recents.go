package service

import (
	files "file-manager/dto"
	"file-manager/pkg/repository"
)

type RecentService struct {
	repo repository.Recent
}

func NewRecentsService(repo repository.Recent) *RecentService {
	return &RecentService{repo: repo}
}

func (s *RecentService) GetRecent(uidUser string) ([]files.File, error) {

	var files []files.File

	files, err := s.repo.GetRecent(uidUser)

	if err != nil {
		return files, err
	}

	return files, nil
}

func (s *RecentService) UpdateRecent(uidUser string) error {

	err := s.repo.UpdateRecent(uidUser)

	if err != nil {
		return err
	}

	return nil
}
