package service

import (
	directories "file-manager/dto"
	files "file-manager/dto"
	"file-manager/pkg/repository"
	"fmt"
	"os"
)

type DirectoriesService struct {
	repo repository.Directories
}

func NewDirectoriesService(repo repository.Directories) *DirectoriesService {
	return &DirectoriesService{repo: repo}
}

func (s *DirectoriesService) GetDirectory(uid string) ([]directories.Directories, []files.File, error) {

	var getDrectories []directories.Directories
	var getFiles []files.File

	getDrectories, getFiles, err := s.repo.GetDirectory(uid)

	if err != nil {
		return getDrectories, getFiles, err
	}

	return getDrectories, getFiles, nil
}

func (s *DirectoriesService) CreateDirectory(userUid string, input directories.NewDirectory) (directories.Directories, error) {

	var directory directories.Directories
	var path string

	pathRoot := os.Getenv("PATH_FILES") + "/" + userUid

	directory, err := s.repo.CreateDirectory(userUid, input)

	fmt.Println(directory)

	if err != nil {
		return directory, err
	}

	path = pathRoot + "/" + directory.Uid

	err = os.Mkdir(path, 0777)

	//f, _ := os.Open(path)
	//fi, _ := f.Stat()
	//fmt.Printf("Size: %d\n", fi.Size())

	if err != nil {
		return directory, err
	}

	return directory, nil
}

func (s *DirectoriesService) UpdateDirectory(input directories.UpdateDirectory) (directories.Directories, error) {

	var directory directories.Directories

	directory, err := s.repo.UpdateDirectory(input)

	if err != nil {
		return directory, err
	}

	return directory, nil
}

func (s *DirectoriesService) DownloadDirectory(uid string) error {

	return nil
}

/*

func (s *DirectoriesService) setTreeDownloadDir(uid string) (directories.DirectoriesTree, error) {

	getDrectories, getFiles, err := s.repo.GetDirectory(uid)

	for _, dir := range getDrectories {

	}

}
*/
