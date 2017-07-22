package main

import (
	"os"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	servicename := "helloworld"
	// create a new directory for the service.
	createDir(servicename)
	dest := createFile(servicename, "service.conf")
	parseAndReplace(servicename, dest)

}

func createDir(service_name string) {
	name := service_name
	os.Mkdir(name,0700)
}

func createFile(service_name string, filename string) string {
	dest := service_name
	src := "template" + "/" + filename + ".tpl"
	err := os.Rename(src, dest)
	if err != nil {
		panic(err)
}
	return dest
}
func parseAndReplace(servicename string, dest string) {
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
