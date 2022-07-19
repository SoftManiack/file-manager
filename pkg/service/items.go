package service

import (
	item "file-manager"
	"file-manager/pkg/repository"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type ItemService struct {
	repo repository.Items
}

// Получить элементы +
// Удалить элемент +
// Переименовать +

// Добавть ссылку
// Редактировать текстовый файл
// Загрузка
// Скачивание
// Пермещение

func NewItemService(repo repository.Items) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) GetItems(rootuid string) ([]item.Directory, []item.File, error) {

	var directories []item.Directory
	var files []item.File

	directories, files, err := s.repo.GetItems(rootuid)

	if err != nil {
		return directories, files, err
	}

	return directories, files, err

}

func (s *ItemService) CreateDirectory(input item.NewDirectory, uid string) (item.Directory, error) {

	var directory item.Directory
	var path string

	pathRoot, _ := findRoot(uid, input.RootUid)

	directory, err := s.repo.CreateDirectory(input, uid)

	if err != nil {
		return directory, err
	}

	path = pathRoot + "/" + directory.Uid

	err = os.Mkdir(path, 0777)

	err = s.repo.SetPath(directory.Uid, "Облачный диск"+"/"+directory.Name, "directory")

	//f, _ := os.Open(path)
	//fi, _ := f.Stat()
	//fmt.Printf("Size: %d\n", fi.Size())

	if err != nil {
		return directory, err
	}

	return directory, err
}

func (s *ItemService) CreateTextFile(input item.NewFile, uid string) (item.File, error) {
	fmt.Println("CreateTextFile")

	var file item.File

	pathRoot, _ := findRoot(uid, input.RootUid)
	fmt.Println(pathRoot)
	file, err := s.repo.CreateTextFile(input)

	if err != nil {
		return file, err
	}

	var pathFile string = pathRoot + "/" + file.Uid + ".txt"

	_, err = os.Create(pathFile)
	err = os.WriteFile(pathFile, []byte(input.Text), 0666)
	err = s.repo.SetPath(file.Uid, "Облачный диск"+"/"+file.Name, "file")

	if err != nil {
		logrus.Error(err)
		return file, err
	}

	return file, err
}

func (s *ItemService) Rename(uid string, input item.Rename) error {

	err := s.repo.Rename(uid, input)
	return err
}

/*
func (s *ItemService) Upload(file *multipart.FileHeader, uid string) (item.Item, error) {

	return item.Item{}, nil
}

func (s *ItemService) Delete(uid, userUid string) error {

	err := s.repo.Delete(uid)

	pathRoot, _ := findRoot(userUid, uid)

	fmt.Println(pathRoot)
	err = os.RemoveAll(pathRoot)

	return err
}
*/
func findRoot(uid, rootuid string) (string, error) {

	root := fmt.Sprintf("C:/Програмирование/Проекты/file-manager/cloud/%s", uid)
	var pathRoot string

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, rootuid) {
				pathRoot = path
			}
			return nil
		})

	if err != nil {
		logrus.Println(err)
	}

	return pathRoot, nil
}
