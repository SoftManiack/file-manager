package service

import (
	files "file-manager/dto"
	"file-manager/pkg/repository"
	"fmt"
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

func (s *RecentService) UpdateRecent(uidUser, uidFile string) error {

	fmt.Println("recet services 12112")
	fmt.Println(uidUser)
	fmt.Println(uidFile)

	err := s.repo.UpdateRecent(uidUser, uidFile)

	if err != nil {
		return err
	}

	return nil
}
