package files

import "os"

func ZipDirExists() error {
	_, err := os.ReadDir("./clients/")
	if err != nil {
		err = os.Mkdir("./clients/", 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
