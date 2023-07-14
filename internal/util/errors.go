package util

import "errors"

func NoneExistentEnvErr() error {
	return errors.New("Looks like you don't have a .env\nLet's fix that")
}

func EmptyEnvErr() error {
	return errors.New("Your .env is empty\nLet's fix that")
}

func EnvCreationErr(err error) error {
	envCreationErr := errors.New("Could not create .env\n")
	e := errors.Join(envCreationErr, err)
	return e
}

func GetInputErr() error {
	return errors.New("Could not read your input")
}

func EmptyInputReadErr() error {
	return errors.New("No input was entered")
}

func FileCreationErr(f string, err error) error {
	msg := "Could not create " + f + "\n"
	fileCreationErr := errors.New(msg)
	e := errors.Join(fileCreationErr, err)
	return e
}

func DirCreationErr(err error) error {
	dirCreationErr := errors.New("Could not create .rgn\n")
	e := errors.Join(dirCreationErr, err)
	return e
}

func GetHomeErr(err error) error {
	getHomeErr := errors.New("HOME DIR SHOULD EXIST\n")
	e := errors.Join(getHomeErr, err)
	return e
}
