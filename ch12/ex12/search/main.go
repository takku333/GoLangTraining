// Search is a demo of the params.Unpack function.
package main

import (
	"GoLangTraining/ch12/ex12/myparams2"
	"fmt"
	"gopl-master/CH12/ex11/myparams"
	"log"
	"net/http"
)

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		MailAddr         string `http:"maddr,mail"`
		CreditCardNumber string `http:"ccn,credit"`
	}
	if err := myparams2.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
	fmt.Println(myparams.Pack("http://localhost:12345/search", &data))

}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
//!+output
$ go build gopl.io/ch12/search
$ ./search &
$ ./fetch 'http://localhost:12345/search'
Search: {Labels:[] MaxResults:10 Exact:false}
$ ./fetch 'http://localhost:12345/search?l=golang&l=programming'
Search: {Labels:[golang programming] MaxResults:10 Exact:false}
$ ./fetch 'http://localhost:12345/search?l=golang&l=programming&max=100'
Search: {Labels:[golang programming] MaxResults:100 Exact:false}
$ ./fetch 'http://localhost:12345/search?x=true&l=golang&l=programming'
Search: {Labels:[golang programming] MaxResults:10 Exact:true}
$ ./fetch 'http://localhost:12345/search?q=hello&x=123'
x: strconv.ParseBool: parsing "123": invalid syntax
$ ./fetch 'http://localhost:12345/search?q=hello&max=lots'
max: strconv.ParseInt: parsing "lots": invalid syntax
//!-output
*/
