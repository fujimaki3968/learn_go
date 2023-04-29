package main

import (
	"errors"
	"fmt"
)

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("resource %s not found", re.Resource)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

func main() {
	err1 := ResourceErr{Resource: "Database", Code: 100}
	err2 := ResourceErr{Resource: "Function", Code: 200}

	if errors.Is(err1, ResourceErr{Resource: "Database"}) {
		fmt.Println("err1 occur Database resource error")
	}

	if errors.Is(err1, ResourceErr{Code: 100}) {
		fmt.Println("err1 occur Code 100 error")
	}

	if errors.Is(err2, ResourceErr{Resource: "Function"}) {
		fmt.Println("err2 occur Function error")
	}

	if errors.Is(err2, ResourceErr{Code: 200}) {
		fmt.Println("err2 occur Code 200 error")
	}

	if errors.Is(err1, ResourceErr{Resource: "Function"}) {
		// これは呼ばれない
		fmt.Println("err1 occur Function error (呼ばれない)")
	}
}
