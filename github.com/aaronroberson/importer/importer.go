package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	contactsFile, err := ioutil.ReadFile(pwd + "/contacts.csv")
	if err != nil {
		fmt.Print(err)
	}

	contacts := string(contactsFile)

	reader := csv.NewReader(strings.NewReader(contacts))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

