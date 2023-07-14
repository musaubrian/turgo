package util

import (
	"os"
)

func CreateEnv(envPath string) error {
	var envDets []string

	err := createEnvDir()
	if err != nil {
		return err
	}
	env, err := os.Create(envPath)
	if err != nil {
		return EnvCreationErr(err)
	}
	defer env.Close()
	HeaderMsg("Let's set you up\nvisit (https://docs.turso.tech/tutorials/get-started-turso-cli/) \n")
	URL := MustGetInput("Turso url")
	URL = "url=" + URL
	envDets = append(envDets, URL)
	token := MustGetInput("Your turso token")
	token = "token=" + token

	envDets = append(envDets, token)

	for _, v := range envDets {
		env.WriteString(v + "\n")
	}
	bold.Println("\n// All set, re-run to access the commands")
	os.Exit(1)

	return nil
}

func createEnvDir() error {
	h, err := os.UserHomeDir()
	if err != nil {
		return GetHomeErr(err)
	}
	envDir := h + string(os.PathSeparator) + ".turgo"
	err = os.Mkdir(envDir, 0o770)
	if err != nil {
		return DirCreationErr(err)
	}
	return nil
}

func GetEnvLoc() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return "", GetHomeErr(err)
	}
	envLoc := h + string(os.PathSeparator) + ".turgo" + string(os.PathSeparator) + ".env"
	return envLoc, nil
}
