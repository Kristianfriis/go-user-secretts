package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/google/uuid"
)

var localSecretFileName = "user-secret.config"
var userSecretFileName = "user-secret.env"
var userSecretsFolder = "go-user-secrets"
var userSecretId = uuid.New()

func main() {
	setup := flag.Bool("setup", false, "creates the user-secret files in the default location")
	flag.Parse()
	if *setup {
		setupSecretFiles()
	}
}

func setupSecretFiles() {
	err := handleLocalFile()
	if err != nil {
		return
	}

	handleUserFile()
}

func handleLocalFile() error {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	var filePath = path.Join(currentDir, localSecretFileName)

	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			newFile, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return err
			}

			newFile.Write([]byte(userSecretId.String()))
			defer newFile.Close()
		}
	}

	userid, err := readUserSecretId(filePath)
	if err != nil {
		fmt.Println("Could not get usersecretId", err)
		return err
	}

	userSecretId = userid

	return nil
}
func handleUserFile() error {
	userDirPath, err := checkAndCreateUserSecretsFolder()
	if err != nil {
		fmt.Println("Error creating user directory:", err)
		return err
	}

	dirPath := path.Join(userDirPath, userSecretId.String())
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

	var filePath = path.Join(dirPath, userSecretFileName)

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

func checkAndCreateUserSecretsFolder() (string, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return "", err
	}

	dirPath := path.Join(userPath, userSecretsFolder)
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

func readUserSecretId(filePath string) (uuid.UUID, error) {
	body, err := os.ReadFile(filePath)
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
