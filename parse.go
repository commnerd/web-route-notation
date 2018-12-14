package main

import (
	"fmt"
    "io"
    "os"
)

const (
		// Delimiter constants
		DELIM_ROUTE       = '+'
		DELIM_PATH        = '/'
		DELIM_MID_START   = '['
		DELIM_MID_END     = ']'
		DELIM_NAME        = '*'
		DELIM_CONTROLLER  = '@'
		DELIM_METHOD	  = '#'
		DELIM_GROUP_START = '('
		DELIM_GROUP_END	  = ')'

		// Verb labels
		VERB_GET_CHAR		= 'G'
		VERB_POST_CHAR		= 'P'
		VERB_PUT_CHAR		= 'U'
		VERB_PATCH_CHAR		= 'A'
		VERB_DELETE_CHAR	= 'D'
		VERB_HEAD_CHAR		= 'H'
)

var verbMap = map[byte]string {
	VERB_GET_CHAR: "GET",
	VERB_POST_CHAR: "POST",
	VERB_PUT_CHAR: "PUT",
	VERB_PATCH_CHAR: "PATCH",
	VERB_DELETE_CHAR: "DELETE",
	VERB_HEAD_CHAR: "HEAD",
}

type Routable interface {

}

type Group struct {
	Routables []Routable
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Route struct {
	Name		string		`json:"name,omitempty"`
	Controller	string		`json:"controller,omitempty"`
	Method		string		`json:"method,omitempty"`
	Route		string		`json:"route,omitempty"`
	Verbs		[]string	`json:"verbs,omitempty"`
	Middlewares	[]string	`json:"middlewares,omitempty"`
	Group		[]Routable	`json:"group,omitempty"`
}

func main() {
	tokenize()
}

func tokenize() {
	f, err := os.Open("/tmp/dat")
	check(err)

	defer f.Close()

	for {
		b := make([]byte, 1)
		char, err := f.Read(b)
		if err == io.EOF {
			break
		}
		check(err)

		fmt.Print(char)
	}

}
