package fileutil

import "io/ioutil"

// write a file to disk using 777 perm
func WriteFile(path string, contents string) {
	err := ioutil.WriteFile(path, []byte(contents), 0777)
	if err != nil {
		panic(err)
	}
}
