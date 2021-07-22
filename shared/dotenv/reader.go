package dotenv

import (
	"strings"

	"github.com/kamontat/fthelper/shared/fs"
)

// Load will read your env file(s) and load them into ENV for this process.
//
// Call this function as close as possible to the start of your program (ideally in main)
//
// If you call Load without any args it will default to loading .env in the current path
//
// You can otherwise tell it which files to load (there can be more than one) like
//
//		godotenv.Load("fileone", "filetwo")
//
// It's important to note that it WILL NOT OVERRIDE an env variable that already exists - consider the .env file to set dev vars or sensible defaults
func Load(files ...fs.FileSystem) error {
	var resolveFiles, err = ResolveFiles(files)
	if err != nil {
		return err
	}

	for _, file := range resolveFiles {
		err = loadFile(file, false)
		if err != nil {
			return err
		}
	}

	return nil
}

// Overload will read your env file(s) and load them into ENV for this process.
//
// Call this function as close as possible to the start of your program (ideally in main)
//
// If you call Overload without any args it will default to loading .env in the current path
//
// You can otherwise tell it which files to load (there can be more than one) like
//
//		godotenv.Overload("fileone", "filetwo")
//
// It's important to note this WILL OVERRIDE an env variable that already exists - consider the .env file to forcefilly set all vars.
func Overload(files ...fs.FileSystem) error {
	var resolveFiles, err = ResolveFiles(files)
	if err != nil {
		return err
	}

	for _, file := range resolveFiles {
		err = loadFile(file, true)
		if err != nil {
			return err
		}
	}

	return nil
}

// Read all env (with same file loading semantics as Load) but return values as
// a map rather than automatically writing values into env
func Read(files ...fs.FileSystem) (envMap map[string]string, err error) {
	envMap = make(map[string]string)

	for _, file := range files {
		individualEnvMap, individualErr := readFile(file)

		if individualErr != nil {
			err = individualErr
			return // return early on a spazout
		}

		for key, value := range individualEnvMap {
			envMap[key] = value
		}
	}

	return
}

//Unmarshal reads an env file from a string, returning a map of keys and values.
func Unmarshal(str string) (envMap map[string]string, err error) {
	return Parse(strings.NewReader(str))
}
