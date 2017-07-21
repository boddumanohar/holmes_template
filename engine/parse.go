package main

import (
	"os"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	// rename files
	os.Rename("template/service.conf.tpl", "template/service.conf")

	service_name := "helloworld"
// replace a word other name 
	input, err := ioutil.ReadFile("template/service.conf") // or anyfile
	if err != nil {
	  fmt.Println(err)
	  os.Exit(1)
   }
   output := bytes.Replace(input, []byte("{$name}"), []byte(service_name), -1)

   if err = ioutil.WriteFile("template/service.conf", output, 0666); err != nil {
			fmt.Println(err)
			 os.Exit(1)
	  }

}

