package main

import (
	"github.com/golang-collections/collections/stack"
//	"fmt"
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
		DELIM_VAL		  = ','

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

var closeParenOpenMap = map[byte]byte {
	']': '[',
	')': '(',
}

type Routable interface {

}

type Group struct {
	ParentGroup *Group
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
	Path		string		`json:"path,omitempty"`
	Verbs		[]string	`json:"verbs,omitempty"`
	Middlewares	[]string	`json:"middlewares,omitempty"`
	Group		[]Routable	`json:"group,omitempty"`
}

var TopGroup = new(Group)
var ActiveGroup = TopGroup
var ActiveRoute Route
var WriteChannel *string
var parenStack = stack.New()

func main() {
	tokenize()
}

func tokenize() {
	f, err := os.Open("/tmp/dat")
	check(err)

	defer f.Close()

	for {
		b := make([]byte, 1)
		_, err := f.Read(b)
		if err == io.EOF {
			break
		}
		check(err)

		processByte(b)
	}
}

func processByte(b []byte) {
	switch b[0] {
	case ' ':
		return
	case DELIM_ROUTE:
		ActiveRoute = Route{}
		ActiveGroup.Routables = append(ActiveGroup.Routables, ActiveRoute)
		return
	case DELIM_PATH:
		ActiveRoute.Path = *new(string)
		WriteChannel = &ActiveRoute.Path
		ActiveRoute.Path += string(b)
		return
	case DELIM_MID_START:
		parenStack.Push(b)

		return
	case DELIM_MID_END:
		if parenStack.Pop() != DELIM_MID_START {
			panic(1)
		}
		return
	case DELIM_NAME:
		return
	case DELIM_CONTROLLER:
		return
	case DELIM_METHOD:
		return
	case DELIM_GROUP_START:
		parenStack.Push(b)
		group := new(Group)
		group.ParentGroup = ActiveGroup
		ActiveGroup = group
		return
	case DELIM_GROUP_END:
		if parenStack.Pop() != DELIM_GROUP_START || ActiveGroup.ParentGroup == nil {
			panic(1)
		}
		ActiveGroup = ActiveGroup.ParentGroup
		return
	default:
		break;
	}

	return
}
