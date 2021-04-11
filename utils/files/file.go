package files

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadJsonFile read file and unmarshal to struct instance
func ReadJsonFile(path string, ptr interface{}) error {
	enc, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(enc, ptr)
}

// WriteJsonFile encode struct instance to bytes and persis in file
func WriteJsonFile(path string, ptr interface{}, indent bool) (err error) {
	var enc []byte

	if indent {
		enc, err = json.MarshalIndent(ptr, "", "    ")
	} else {
		enc, err = json.Marshal(ptr)
	}

	if err != nil {
		return
	}

	return ioutil.WriteFile(path, enc, os.ModePerm)
}
