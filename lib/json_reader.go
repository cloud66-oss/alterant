package lib

import "io/ioutil"

type JsonReader struct {
}

func NewJsonReader(file string) string {
	jsonFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(jsonFile)
}
