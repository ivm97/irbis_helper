package files

import (
	"archive/zip"
	"encoding/json"
	"io"
	"os"
)

func (f *Files) LastVersion() (string, error) {
	b, err := os.ReadFile("./settings/settings.json")
	if err != nil {
		return "", err
	}
	var fv Files
	err = json.Unmarshal(b, &fv)
	if err != nil {
		return "", err
	}

	return fv.ClientVersion(), nil
}

func (f *Files) GetSourceFiles() (*string, error) {
	files, err := os.ReadDir(f.clientPath + "/IRBIS64")
	if err != nil {
		return nil, err
	}

	if ZipDirExists() != nil {
		return nil, err
	}
	zipName := "./clients/" + f.clientVersion + ".zip"
	if _, err := os.OpenFile(zipName, os.O_RDONLY, 0777); err == nil {
		return &zipName, nil
	}

	zipFile, err := os.Create(zipName)
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	for _, file := range files {
		fileToZip, err := os.Open(f.clientPath + "/IRBIS64/" + file.Name())
		if err != nil {
			return nil, err
		}
		defer fileToZip.Close()

		// Получаем информацию о файле
		info, err := fileToZip.Stat()
		if err != nil {
			return nil, err
		}

		// Создаем заголовок файла в ZIP-архиве
		zipEntry, err := zipWriter.Create(info.Name())
		if err != nil {
			return nil, err
		}

		// Копируем содержимое файла в ZIP-архив
		_, err = io.Copy(zipEntry, fileToZip)
		if err != nil {
			return nil, err
		}
	}

	return &zipName, err
}
