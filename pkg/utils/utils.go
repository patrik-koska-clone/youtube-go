package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ChooseRandomIndex(list []string) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if len(list) == 0 {
		fmt.Println("List is empty.")
		return -1
	}

	return r.Intn(len(list))
}

func ConvertStrToInt64(s string) (int64, error) {
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
