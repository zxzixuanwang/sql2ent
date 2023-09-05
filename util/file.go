package util

import (
	"os"
	"path/filepath"
)

func CreateFiles(modelList map[string]string, dir string) error {
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	getDir := filepath.Dir(dir)
	if dir != getDir {
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			return err
		}
	}
	closeF := make([]*os.File, 0, len(modelList))
	defer func() {
		for _, v := range closeF {
			v.Close()
		}
	}()

	for fileName, code := range modelList {
		filename := filepath.Join(dirAbs, fileName)
		f, err := os.Open(filename)
		if err != nil {
			if os.IsNotExist(err) {
				closeF = append(closeF, f)

				err = os.WriteFile(filename, []byte(code), os.ModePerm)
				if err != nil {
					return err
				}
			}

			continue
		}

	}

	return nil
}
