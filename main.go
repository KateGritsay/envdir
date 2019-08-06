package main

import (
"errors"
"fmt"
"io"
"io/ioutil"
"log"
"os"
"os/exec"
"path/filepath"
"strings"
)

func Envdir(in io.Reader, out io.Writer, errWriten io.Writer, args []string) error {
	if len(args) != 2{
		return errors.New("It has to have two arguments: fileName and path")
	}
	path := args[0]
	progName := args[1]

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	cmd := exec.Command(progName)
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = errWriten
	cmd.Env = envir(path, files)

	err = cmd.Run()
	if err != nil {
		fmt.Errorf("Command finished with error: %v", err)
	}
	return nil
}

func envir(path string, files []os.FileInfo) []string{

	envir := make ([]string,0,len(files))

	for _, file := range files {
		if file.IsDir(){
			continue
		}

		fileName := filepath.Join(path, file.Name())
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			continue
		}

		var b strings.Builder
		b.WriteString(file.Name())
		b.WriteString("=")
		b.WriteString(string(content))

		envir = append(envir,b.String())

	}
	return envir
}

func main (){

	err := Envdir(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

