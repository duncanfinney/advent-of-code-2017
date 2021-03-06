package lib

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInput(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(bytes))
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Atois(items []string) []int {
	ret := []int{}
	for _, item := range items {
		ret = append(ret, Atoi(item))
	}
	return ret
}
