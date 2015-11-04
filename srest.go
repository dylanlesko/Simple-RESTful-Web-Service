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
func main() {
	router := NewRouter()
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
	/* Homepage */
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	/* POST Route for Create */
	Route{
		"StudentPost",
		"POST",
		"/Student",
		StudentPost,
	},
	/* POST Route for Read */
	Route{
		"StudentGet",
		"GET",
		"/Student/getstudent",
		StudentGet,
	},
	/* POST Route for Delete*/
	Route{
		"StudentDelete",
		"DELETE",
		"/Student",
		StudentDelete,
	},
	/* POST Route for Update */
	Route{
		"StudentPut",
		"PUT",
		"/Student",
		StudentPut,
	},
	/* Post Route for List All*/
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
	Id bson.ObjectId `json: "id" bson:"_id,omitempty"`
	NetID string `json: "nid" bson:"netid"`
	Name string `json: "name" bson:"name"`
	Major string `json: "major" bson:"major"`
	Year int `json: "year" bson:"year"`
	Grade int `json: "grade" bson:"grade"`
	Rating string `json: "rating" bson:"rating"`
}
type Students []Student
type StudentString struct {
	NetID string `json: "nid"`
	Name string `json: "name"`
	Major string `json: "major"`
	Year int `json: "year"`
	Grade int `json: "grade"`
	Rating string `json: "rating"`
}
/* End students.go */


/* handlers.go */
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
func StudentPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()
	//fmt.Printf("body\n", r.PostForm   )

	dec := json.NewDecoder(r.Body)
	var data StudentString
	dec.Decode(&data)

//	newS := &Student{ ID:bson.NewObjectId(), NetID:data.NetID, Name:data.Name, Major:data.Major, Year:data.Year, Grade:data.Grade, Rating:data.Rating}



	uri := "mongodb://localhost:27017"

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	} 
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("test").C("student")

	newS := Student{}
	newS.Id = bson.NewObjectId()
	newS.NetID = data.NetID
	newS.Major = data.Major
	newS.Year = data.Year
	newS.Grade = data.Grade
	newS.Rating = data.Rating



	collection.Insert(&newS)

	fmt.Fprintln(w, "Inserted: ")
	fmt.Fprintln(w, newS)
	fmt.Fprintln(w, "into test.Student")
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

	result := []StudentString{}
	err = collection.Find(bson.M{"name": search}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	resLen := len(result)
	output := strconv.Itoa(resLen) + " Results Found\n"

	a := 0
	for a < resLen {
		output = output + "\nResult " + strconv.Itoa(a+1)
		output = output + "\n\tNetID:\t"
		output = output + result[a].NetID
		output = output + "\n\tName:\t"
		output = output + result[a].Name
		output = output + "\n\tMajor:\t"
		output = output + result[a].Major
		output = output + "\n\tYear:\t"
		output = output + strconv.Itoa(result[a].Year)
		output = output + "\n\tGrade:\t"
		output = output + strconv.Itoa(result[a].Grade)
		output = output + "\n\tRating:\t"
		output = output + result[a].Rating
		a++
	}
	
	fmt.Fprintf(w, output)
}

func StudentDelete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r.ParseForm()

	dec := json.NewDecoder(r.Body)
	var data StudentString
	dec.Decode(&data)

	fmt.Println(data.Year)

	year := data.Year

	uri := "mongodb://localhost:27017"

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	} 
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("test").C("student")

	collection.RemoveAll( bson.M{ "year": bson.M{"$lt": year}  }  )

	fmt.Fprintln(w, "Removed Years Before")
	fmt.Fprint(w, year)
}
func StudentPut(w http.ResponseWriter, r *http.Request) {
	uri := "mongodb://localhost:27017"
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	} 
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("test").C("student")

	result := []Student{}
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	resLen := len(result)
	total := 0
	a := 0
	for a < resLen {
		total = total + result[a].Grade
		a++
	}
	average := total/resLen
	
	A_avg := average + 10
	B_avg := average - 10
	C_avg := average - 20

	changeA := bson.M{"$set": bson.M{"rating": "A"}}
	changeB := bson.M{"$set": bson.M{"rating": "B"}}
	changeC := bson.M{"$set": bson.M{"rating": "C"}}

	searchA := bson.M{
		"grade": bson.M{"$gt": A_avg},
	}
	searchB := bson.M{
		"grade": bson.M{"$gt": B_avg},
	}
	searchC := bson.M{
		"grade": bson.M{"$gt": C_avg},
	}

	collection.UpdateAll(searchC, changeC)
	collection.UpdateAll(searchB, changeB)
	collection.UpdateAll(searchA, changeA)
}
func StudentList(w http.ResponseWriter, r *http.Request) {
	uri := "mongodb://localhost:27017"
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	} 
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("test").C("student")
	
	result := []StudentString{}
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	resLen := len(result)
	output := strconv.Itoa(resLen) + " Results Found\n"

	a := 0
	for a < resLen {
		output = output + "\nResult " + strconv.Itoa(a+1)
		output = output + "\n\tNetID:\t"
		output = output + result[a].NetID
		output = output + "\n\tName:\t"
		output = output + result[a].Name
		output = output + "\n\tMajor:\t"
		output = output + result[a].Major
		output = output + "\n\tYear:\t"
		output = output + strconv.Itoa(result[a].Year)
		output = output + "\n\tGrade:\t"
		output = output + strconv.Itoa(result[a].Grade)
		output = output + "\n\tRating:\t"
		output = output + result[a].Rating
		a++
	}
	
	fmt.Fprintf(w, output)
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

