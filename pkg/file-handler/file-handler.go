package filehandler

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/Kristianfriis/go-user-secrets/pkg/config"
	"github.com/google/uuid"
)

type FileHandler struct {
	conf *config.Config
}

func NewFileHandler(c *config.Config) FileHandler {
	return FileHandler{
		conf: c,
	}
}

func (f *FileHandler) HandleLocalFile() error {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	var filePath = path.Join(currentDir, f.conf.LocalSecretFileName)

	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			newFile, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return err
			}

			newFile.Write([]byte(f.conf.UserSecretId.String()))
			defer newFile.Close()
		}
	}

	userid, err := f.ReadUserSecretId(&filePath)
	if err != nil {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return err
		}

		var filePath = path.Join(currentDir, f.conf.LocalSecretFileName)

		_, err = os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				newFile, err := os.Create(filePath)
				if err != nil {
					fmt.Println("Error creating file:", err)
					return err
				}

				newFile.Write([]byte(f.conf.UserSecretId.String()))
				defer newFile.Close()
			}
		}
		fmt.Println("Could not get usersecretId", err)
		return err
	}

	f.conf.UserSecretId = userid

	return nil
}
func (f *FileHandler) HandleUserFile() error {
	userDirPath, err := f.checkAndCreateUserSecretsFolder()
	if err != nil {
		fmt.Println("Error creating user directory:", err)
		return err
	}

	dirPath := path.Join(userDirPath, f.conf.UserSecretId.String())
	_, err = os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the directory if it doesn't exist
			err := os.Mkdir(dirPath, 0755) // Adjust permissions as needed
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return err
			}
		} else {
			fmt.Println("Error checking path existence:", err)
			return err
		}
	}

	var filePath = path.Join(dirPath, f.conf.UserSecretFileName)

	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			newFile, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return err
			}
			defer newFile.Close()
		}
	}
	return nil
}

func (f *FileHandler) checkAndCreateUserSecretsFolder() (string, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return "", err
	}

	dirPath := path.Join(userPath, f.conf.UserSecretsFolder)
	_, err = os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the directory if it doesn't exist
			err := os.Mkdir(dirPath, 0755) // Adjust permissions as needed
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	return dirPath, nil
}

func (f *FileHandler) ReadUserSecretId(filePath *string) (uuid.UUID, error) {
	var pathToGetFrom string

	if filePath == nil {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return uuid.New(), err
		}

		var filePathToCheck = path.Join(currentDir, f.conf.LocalSecretFileName)

		_, err = os.Stat(filePathToCheck)
		if err != nil {
			if os.IsNotExist(err) {
				if err != nil {
					fmt.Println("Error reading file:", err)
					return uuid.New(), err
				}
			} else {
				pathToGetFrom = filePathToCheck
			}
		}
	} else {
		pathToGetFrom = *filePath
	}

	body, err := os.ReadFile(pathToGetFrom)
	if err != nil {
		log.Printf("unable to read file: %v/n", err)
		return uuid.New(), err
	}

	parsedUuid, err := uuid.Parse(string(body))
	if err != nil {
		log.Printf("unable to read file: %v/n", err)
		return uuid.New(), err
	}

	return parsedUuid, nil
}
