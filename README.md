//go:generate go-user-secrets -setup

	userDirPath, _ := os.UserHomeDir()
	dirPath := path.Join(userDirPath, "go-user-secrets", "98ec119d-0e12-401e-8ae9-1a72ad327dc0", "user-secret.env")
	err := godotenv.Load(dirPath)
	if err != nil {
		fmt.Println("Error loading user secrets")
		fmt.Println(err)
	}