package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func zipFileRecursively(zipWriter *zip.Writer, dir string, absInpPath string, num int) error {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	relPath, err := filepath.Rel(absInpPath, dir)
	if err != nil {
		return err
	}

	dirHeader := &zip.FileHeader{
		Name:   filepath.ToSlash(relPath + "/"),
		Method: zip.Store,
	}

	_, err = zipWriter.CreateHeader(dirHeader)
	if err != nil {
		return err
	}

	for _, dirEntry := range dirEntries {
		fullPath := filepath.Join(dir, dirEntry.Name())

		if dirEntry.IsDir() {
			err = zipFileRecursively(zipWriter, dir+"/"+dirEntry.Name(), absInpPath, num+1)
			if err != nil {
				return err
			}
		} else {
			file, err := os.Open(dir + "/" + dirEntry.Name())
			if err != nil {
				return err
			}

			fileInfo, err := os.Stat(dir + "/" + dirEntry.Name())
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(fileInfo)
			if err != nil {
				return err
			}

			header.Name, err = filepath.Rel(absInpPath, fullPath)
			if err != nil {
				return err
			}

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ZipFile(zipName, dirName string) error {
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return zipFileRecursively(zipWriter, dirName, dirName, 0)
}

func UnZipFile(zipPath string, outputDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		desPath := filepath.Join(outputDir, f.Name)

		if strings.HasSuffix(f.Name, "/") {
			if err := os.MkdirAll(desPath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(desPath), os.ModePerm); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		outFile, err := os.Create(desPath)
		if err != nil {
			rc.Close()
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
