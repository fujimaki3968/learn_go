package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

func login(uid, pwd string) error {
	return nil
}

func getData(file string) ([]byte, error) {
	var dummy []byte
	return dummy, nil
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err,
		}
	}

	return data, nil
}

func GenerateError(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{Status: NotFound}
	}
	return genErr
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println("in main, wrappedErr:", wrappedErr)
		}
	}
}
