package match_maker

import (
	"fmt"
	"net/http"
	"regexp"
)

type RegexURLMatcher struct {
	Patterns map[string]*regexp.Regexp
	Handlers map[string]http.HandlerFunc
}

// Returns a new RegexURLMatcher empty struct containing two string maps:
// - Patterns: stores the compiled regex url patterns
// - Handler: holds the handler functions
func NewRegexURLMatcher() *RegexURLMatcher {
	return &RegexURLMatcher{
		Patterns: make(map[string]*regexp.Regexp),
		Handlers: make(map[string]http.HandlerFunc),
	}
}

// Associate a regex pattern to a particular http.HandlerFunc.
// First the regex pattern will be compiled, if it's not valid, an error will be raised,
// else, the compiled regex and the http.HandlerFunc are added to the RegexURLMatcher struct.
//
// Exemple:
//
//  package main
//
//  import "net/http"
//
//  func main() {
//	  r := match_maker.NewRegexURLMatcher()
//	  r.Add("(GET) /howdy(/?[A-Za-z0-9]*)?", howdyHandler)
//       http.ListenAndServe(":8080", r)
//  }
//
//  func howdyHandler(res http.ResponseWriter, req *http.Request) {
//	  // Do whatever you need...
//  }
func (rg *RegexURLMatcher) Add(regex string, handler http.HandlerFunc) error {
	compiled, err := regexp.Compile(regex)

	if err != nil {
		return fmt.Errorf("regex string cannot compile with err: %s", compiled)
	}

	rg.Handlers[regex] = handler
	rg.Patterns[regex] = compiled

	return nil
}

func (rg *RegexURLMatcher) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	patternToMatch := request.Method + " " + request.URL.Path

	for regexString, handlerFunc := range rg.Handlers {
		if rg.Patterns[regexString].MatchString(patternToMatch) {
			handlerFunc(response, request)
			return
		}
	}

	http.NotFound(response, request)
}
