package main

import (
	"fmt"
	"os"
	"log"
	//"net"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"strconv"
	//"bytes"
	//"io/ioutil"
	//"encoding/hex"
	//"html"
)

/* main.go */
type msg struct {
	Id bson.ObjectId `bson:"_id"`
	Msg string `bson:"msg"`
	Count int `bson:"count"`
} 

func main() {




	router := NewRouter()


	//net.Listen("tcp6", "ip6-localhost:1234")

	log.Fatal(http.ListenAndServe(":1234", router))



}
/* end main.go */


/* router.go */
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}

/* End router.go */

/* Routes.go */
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	
    	Route{
    		"StudentPost",
    		"POST",
    		"/Student",
    		StudentPost,

	},
    	Route{
    		"StudentGet",
    		"GET",
    		"/Student/getstudent",
    		StudentGet,

	},
    	Route{
    		"listall",
    		"GET",
    		"/Student/listall",
    		StudentList,

	},
}
/* End Routes.go */


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


/* handlers.go */
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}




func StudentPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	//test := r.PostForm()
	fmt.Printf("%+v\n", r.Form)

	dec := json.NewDecoder(r.Body)
	var data Student
	dec.Decode(&data)

	fmt.Printf(data.Name, data.Major, data.Year, data.Grade, data.Rating)
    
	//var s Student
	//err := json.Unmarshal(r, &s)

//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	
	//byt := []byte(r.Body)
	
//	if err := json.Unmarshal(byt, &s); err != nil {
//		panic(err)
//		}




	//err := json.Unmarshal( (r.URL.Path), &s )


/*
	//req := r.Ctx.Request     //in beego this.Ctx.Request points to the Http#Request
	p := make([]byte, r.ContentLength)    
	_, err := r.Body.Read(p)
	if err == nil {
		var newStudent Student
		err1 := json.Unmarshal(p, &newStudent)
		if err1 == nil {
			fmt.Println(newStudent)
		} else {
			fmt.Println("Unable to unmarshall the JSON request", err1);

		}
	}






	fmt.Fprintf(w, "hello")
	//m, _  := url.ParseRequestURI(r.URL)

	decoder := json.NewDecoder(r.Body)
	var t Student   
	err = decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Name)
*/

	//fmt.Fprintf(w, m)
	//fmt.Printf("url: %s\n", r.URL.RawQuery)

/*	
	uri := "mongodb://localhost:27017"

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	} 
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})
*/

/*
	collection := sess.DB("test").C("student")

	//newS := Student{NetID: t.NetID, Name: t.Name, Major: t.Major, Year: t.Year, Grade: t.Grade, Rating: t.Rating}
	newS := Student{NetID: "dfg455367", Name: "Jayyy", Major: "Comp Sci", Year: 2016, Grade: 2, Rating: "D"}

	err = collection.Insert(&newS)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(w, newS.Name)
*/
}
func StudentGet(w http.ResponseWriter, r *http.Request) {
	
	uri := "mongodb://localhost:27017"

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	} 
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})


	collection := sess.DB("test").C("student")
	


	m, _  := url.ParseQuery(r.URL.RawQuery)
	search := (m["name"][0])

	result := Student{}
	err = collection.Find(bson.M{"name": search}).One(&result)
	if err != nil {
		log.Fatal(err)
	}






	output := "NetID:\t"
	//id := Hex(result.NetID)
	id := result.NetID.Hex
	//hex.Dump(result.NetID.Hex)
	fmt.Printf("%s", id)
	//output = output + id
	//output = output + strconv.Itoa(id)
	output = output + "\nName:\t"
	output = output + result.Name
	output = output + "\nMajor:\t"
	output = output + result.Major
	output = output + "\nYear:\t"
	year := result.Year
	output = output + strconv.Itoa(year)
	output = output + "\nGrade:\t"
	grade := result.Grade
	output = output + strconv.Itoa(grade)
	output = output + "\nRating:\t"
	output = output + result.Rating

	fmt.Fprintf(w, output)
}
func StudentPut(w http.ResponseWriter, r *http.Request) {
	

}
func StudentDelete(w http.ResponseWriter, r *http.Request) {
	

}


func StudentList(w http.ResponseWriter, r *http.Request) {









}
/* End handlers.go*/


/* logger.go */
func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
/* End logger.go */