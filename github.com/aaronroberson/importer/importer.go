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

func publisher(pub chan<- []string, msg []string) {
	pub <- msg
	fmt.Println(msg)
}

/*
func subscriber(pub chan<- []string, sub chan<- []string) {
	msg := <-pub
	sub <- msg
}
*/

func main() {
	pwd, _ := os.Getwd()
	contactsFile, err := ioutil.ReadFile(pwd + "/contacts.csv")
	if err != nil {
		fmt.Print(err)
	}

	contacts := string(contactsFile)

	reader := csv.NewReader(strings.NewReader(contacts))

	// TODO: remove hard-coded buffer
	pub := make(chan []string, 200)
	// TODO: implement channels subscriber
	// sub := make(chan []string, 200)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}


		publisher(pub, record)
		// TODO: implement channels subscriber
		// subscriber(pub, sub)
	}
}
