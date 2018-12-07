package fileutil

import "io/ioutil"

func WriteFile(path string, contents string) {
	err := ioutil.WriteFile(path, []byte(contents), 0666)
	if err != nil {
		panic(err)
	}
}
