package utils

import (
	"io/ioutil"
	"os"
)

// Проверка на существование папки
func CheckFolder(argPath string) error {
	_, error := ioutil.ReadDir(argPath)
	if error != nil {
		return error
	}
	return nil
}

// создание вложенных папок
func CreateFolder(argPathFolder string) error {
	err := os.MkdirAll(argPathFolder, 0777)
	if err != nil {
		return err
	}
	return nil
}
