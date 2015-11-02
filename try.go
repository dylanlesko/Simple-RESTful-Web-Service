package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"flag"
	"encoding/json"
	"bytes"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"io"
	"log"

)


func main() {
	
	urlPtr	:=	flag.String("url", "", "url")
	methodPtr	:=	flag.String("method", "", "operation")
	dataPtr	:=	flag.String("data", "{}", "JSON Student Object")
	idPtr		:=	flag.String("NetID", "", "User's Net ID")
	namePtr	:=	flag.String("Name", "", "operation")
	majorPtr	:=	flag.String("Major", "", "operation")
	yearPtr	:=	flag.Int("Year", -1, "User's Net ID")
	gradePtr	:=	flag.Int("Grade", -1, "User's Net ID")
	ratingPtr	:=	flag.String("Rating", "", "operation")

	flag.Parse()


	fmt.Println("URL: ", *urlPtr)
	fmt.Println("Method: ", *methodPtr)
	fmt.Println(*dataPtr)
	fmt.Println("NetID: ", *idPtr)
	fmt.Println("Name: ", *namePtr)	
	fmt.Println("Major: ", *majorPtr)
	fmt.Println("Year: ", *yearPtr)
	fmt.Println("Grade: ", *gradePtr)
	fmt.Println("Rating: ", *ratingPtr)






	var s []Student
	fmt.Println("dsfs", *dataPtr)
	var jsonStr = []byte(*dataPtr)
	fmt.Println("sdf4",jsonStr)
	err := json.Unmarshal(jsonStr, &s)
	if err != nil {

	}
	fmt.Printf("%+v", s)











	newS := &Student{bson.NewObjectId(), *namePtr, *majorPtr, *yearPtr, *gradePtr, *ratingPtr}
	if newS != nil {
	}
	if dataPtr != nil {

	}

	//temp := "`[" + (*dataPtr) + "]`"









	//var jsonBlob = []byte(temp)

	type Student2 struct {
		NetID string `json: "id"`
		Name string `json: "name"`
		Major string `json: "major"`
		Year int `json: "year"`
		Grade int `json: "grade"`
		Rating string `json: "rating"`
	}
	dec := json.NewDecoder(strings.NewReader(*dataPtr))
			var m Student2

	for {
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Major)
	}

	fmt.Printf("%+v", m)















/*
	if *dataPtr != "{}" {

		data := *dataPtr

		byt := []byte(data)
		
		if err := json.Unmarshal(byt, &s); err != nil {
			panic(err)
   		}
	}
*/

	if *methodPtr == "Create" {
		Create( *urlPtr, newS );
	}
	if *methodPtr == "list" {
		Read( *urlPtr )
	}
	if *methodPtr == "remove" {
		fmt.Println("DELETE")
	}
	if *methodPtr == "update" {
		fmt.Println("PUT")
	}

	//Read()


}

func Create( url string, s *Student ) {
	//var jsonStr = []byte(s)
	//response, err := http.NewRequest("POST", url, bytes.NewBuffer(s))
	//response, err := http.NewRequest("POST", url, jsonStr)

	//if err != nil {
//
	//	panic(err)
	//}
	//fmt.Println("here: %s", jsonStr)
	//response, err := http.Post(url, s, &buf)

	buf, _ := json.Marshal(s)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post(url, "text/json", body)
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))

	r.Header.Set("X-Custom-Header", "myvalue")
	r.Header.Set("Content-Type", "application/json")


	//client := &http.Client{}
	//resp, err := client.Do(response)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

}































func Read( url string ) {
	//response, err := http.Get("http://localhost:1234")
	response, err := http.Get(url)

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }


/*

	resp, err := http.Get(url)

	if err != nil {
	  panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
*/
	
}

/* students.go */
type Student struct {
	NetID bson.ObjectId `json: "id" bson:"_id"`
	Name string `json: "name"bson:"name"`
	Major string `json: "major"bson:"major"`
	Year int `json: "year"bson:"year"`
	Grade int `json: "grade"bson:"grade"`
	Rating string `json: "rating"bson:"rating"`
}

type Students []Student
/* End students.go */