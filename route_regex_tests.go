package match_maker

import (
	"net/http"
	"reflect"
	"regexp"
	"testing"
)

func TestNewRegexURLMatcher(t *testing.T) {
	tests := []struct {
		name string
		want *RegexURLMatcher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegexURLMatcher(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegexURLMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexURLMatcher_Add(t *testing.T) {
	type fields struct {
		Patterns map[string]*regexp.Regexp
		Handlers map[string]http.HandlerFunc
	}
	type args struct {
		regex   string
		handler http.HandlerFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rg := &RegexURLMatcher{
				Patterns: tt.fields.Patterns,
				Handlers: tt.fields.Handlers,
			}
			if err := rg.Add(tt.args.regex, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("RegexURLMatcher.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegexURLMatcher_ServeHTTP(t *testing.T) {
	type fields struct {
		Patterns map[string]*regexp.Regexp
		Handlers map[string]http.HandlerFunc
	}
	type args struct {
		response http.ResponseWriter
		request  *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rg := &RegexURLMatcher{
				Patterns: tt.fields.Patterns,
				Handlers: tt.fields.Handlers,
			}
			rg.ServeHTTP(tt.args.response, tt.args.request)
		})
	}
}
