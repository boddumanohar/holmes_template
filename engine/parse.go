package main

import (
	"os"
	"bytes"
	"fmt"
	"io"
	"log"
	"io/ioutil"
)

func main() {
	servicename := "helloworld"
	// create a new directory for the service.
	filenames := []string{"service.conf", "Dockerfile", "README.md", "serviceREST.scala", "service.go"}
	createDir(servicename)
	for i:=0; i<4; i++ {
		dest := createFile(servicename, filenames[i])
		parseAndReplace(servicename, dest)
	}
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createDir(service_name string) {
	name := service_name
	os.Mkdir(name,0700)
}

func createFile(service_name, filename string) string {
	src := "template" + "/" + filename + ".tpl"
	dest := service_name + "/" + filename
	sFile,err := os.Open(src)
	Check(err)

	eFile, err := os.Create(dest)
	Check(err)

	_, err = io.Copy(eFile, sFile)
	if err != nil {
		log.Fatal(err)
	}
	err = eFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

	//err := os.Rename(, dest)
	//if err != nil {
	//	panic(err)
	//}

	return dest
}

func parseAndReplace(servicename, dest string) {
// replace a word other name

	input, err := ioutil.ReadFile(dest) // or anyfile
	if err != nil {
	  fmt.Println(err)
	  os.Exit(1)
   }

   output := bytes.Replace(input, []byte("{$name}"), []byte(servicename), -1)
   if err = ioutil.WriteFile(dest, output, 0666); err != nil {
			fmt.Println(err)
			 os.Exit(1)
	  }
 }
