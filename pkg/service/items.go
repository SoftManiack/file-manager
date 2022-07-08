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

	fmt.Println(input.RootUid)
	pathRoot, _ := findRoot(uid, input.RootUid)

	directory, err := s.repo.CreateDirectory(input, uid)

	if err != nil {
		return directory, err
	}

	err = os.Mkdir(pathRoot+"/"+directory.Uid, 0777)

	return directory, err
}

func (s *ItemService) CreateTextFile(input item.NewFile, uid string) (item.File, error) {
	fmt.Println("CreateTextFile")

	var item item.File

	pathRoot, _ := findRoot(uid, input.RootUid)
	fmt.Println(pathRoot)
	item, err := s.repo.CreateTextFile(input)

	if err != nil {
		return item, err
	}

	var file string = pathRoot + "/" + item.Uid + ".txt"

	_, err = os.Create(file)
	err = os.WriteFile(file, []byte(input.Text), 0666)
	if err != nil {
		logrus.Error(err)
		return item, err
	}

	return item, err
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
