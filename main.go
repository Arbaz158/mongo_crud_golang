package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("main is running")
	sess, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("error in getting connection :", err)
	}
	defer sess.Close()
	fmt.Println("Connected to MongoDB!", sess)
}
