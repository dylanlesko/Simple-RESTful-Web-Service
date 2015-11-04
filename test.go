package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"flag"
	"encoding/json"
	"bytes"
	//"gopkg.in/mgo.v2/bson"
	"strings"
	"io"
	"log"

)
	type Student2 struct {
		NetID string `json: "nid"`
		Name string `json: "name"`
		Major string `json: "major"`
		Year int `json: "year"`
		Grade int `json: "grade"`
		Rating string `json: "rating"`
	}

	type Year struct {
		Year int `json: "year"`
	}

func main() {
	urlPtr	:=	flag.String("url", "", "url")
	methodPtr	:=	flag.String("method", "", "operation")
	dataPtr	:=	flag.String("data", "{}", "JSON Student Object")
	idPtr		:=	flag.String("netid", "", "User's Net ID")
	namePtr	:=	flag.String("name", "", "operation")
	majorPtr	:=	flag.String("major", "", "operation")
	yearPtr	:=	flag.Int("year", -1, "User's Net ID")
	gradePtr	:=	flag.Int("grade", -1, "User's Net ID")
	ratingPtr	:=	flag.String("rating", "", "operation")

	flag.Parse()
	if 1 == 0 {
		fmt.Println("URL: ", *urlPtr)
		fmt.Println("Method: ", *methodPtr)
		fmt.Println(*dataPtr)
		fmt.Println("NetID: ", *idPtr)
		fmt.Println("Name: ", *namePtr)	
		fmt.Println("Major: ", *majorPtr)
		fmt.Println("Year: ", *yearPtr)
		fmt.Println("Grade: ", *gradePtr)
		fmt.Println("Rating: ", *ratingPtr)
	}

	dec := json.NewDecoder(strings.NewReader(*dataPtr))
	var m Student2
	for {
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s: %s\n", m.Name, m.Major)
	}
	//fmt.Printf("%+v", m)
	//fmt.Println(m.Name)

	if strings.ToLower(*methodPtr) == "create" {
		Create( *urlPtr, m );
	}
	if strings.ToLower(*methodPtr) == "list" {
		List( *urlPtr )
	}
	if strings.ToLower(*methodPtr) == "remove" {
		Remove( *urlPtr, *yearPtr )
	}
	if strings.ToLower(*methodPtr) == "update" {
		Update( *urlPtr )
	}
}


func Create( url string, s Student2 ) {



//var jsonStr = []byte(s)

	jsonStr, err := json.Marshal(s)

	//fmt.Println("\nSending: ", jsonStr)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

//    fmt.Println("response Status:", resp.Status)
//    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))









}

func List( url string ) {
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
}


func Remove( url string, year int ) {
	//var jsonStr = []byte(year)

	s := Year{}
	s.Year = year

	jsonStr, err := json.Marshal(s)

    req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

//    fmt.Println("response Status:", resp.Status)
//    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func Update( url string ) {
	req, err := http.NewRequest("PUT", url, nil)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

//    fmt.Println("response Status:", resp.Status)
//    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

}

