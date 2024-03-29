package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/Kristianfriis/go-user-secrets/pkg/config"
	filehandler "github.com/Kristianfriis/go-user-secrets/pkg/file-handler"
	"github.com/joho/godotenv"
)

var (
	conf        config.Config           = config.NewConfig()
	fileHandler filehandler.FileHandler = filehandler.NewFileHandler(&conf)
)

func main() {
	setup := flag.Bool("setup", false, "creates the user-secret files in the default location")
	flag.Parse()
	if *setup {
		setupSecretFiles()
	}
}

func setupSecretFiles() {
	err := fileHandler.HandleLocalFile()
	if err != nil {
		return
	}

	fileHandler.HandleUserFile()
}

func AddUserSecretsIfApplicable() error {
	conf := config.NewConfig()
	fh := filehandler.NewFileHandler(&conf)

	secret, err := fh.ReadUserSecretId(nil)

	if err != nil {
		fmt.Println("could not get secrets id")
		fmt.Println("please run setup in the folder root, to setup secrets for this project")
		return err
	}

	userDirPath, _ := os.UserHomeDir()
	dirPath := path.Join(userDirPath, conf.UserSecretsFolder, secret.String(), conf.UserSecretFileName)
	err = godotenv.Load(dirPath)
	if err != nil {
		fmt.Println("error loading user secrets")
		fmt.Println(err)
	}
	return nil
}
