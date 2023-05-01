package main

import "fmt"

func doThing1(val1 int) (int, error) {
	// 何か処理
	return 0, nil
}

func doThing2(val2 string) (string, error) {
	// 何か処理
	return "", nil
}

func doThing3(val3 int, val4 string) (string, error) {
	// 何か処理
	return "", nil
}

func DoSomeThings(val1 int, val2 string) (string, error) {
	val3, err := doThing1(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	result, err := doThing3(val3, val4)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	return result, nil
}

func DoSomeThingsDefer(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThingsDefer: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}
