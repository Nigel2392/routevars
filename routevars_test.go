package routevars_test

import (
	"testing"

	"github.com/Nigel2392/routevars"
)

func TestFormatter(t *testing.T) {
	var path = "/users/<<id:int>>/<<name:string>>"
	var formatter = routevars.URLFormatter(path)
	var url = formatter.Format(1234, "john")
	if url != "/users/1234/john" {
		t.Error("URLFormatter failed")
	}
	t.Log("URLFormatter passed: " + url)
}

func TestMatch(t *testing.T) {
	var path = "/users/<<id:int>>/<<name:string>>"
	var other = "/users/1234/john"
	var ok, vars = routevars.Match(path, other)
	if !ok {
		t.Error("Match failed")
	}
	t.Log("Match passed: " + other)
	for k, v := range vars {
		t.Log(k, v)
	}

	other = "/users/1234/johna/a"
	ok, vars = routevars.Match(path, other)
	if ok {
		t.Error("Match failed")
	}
	t.Log("Match passed: " + other)
	for k, v := range vars {
		t.Log(k, v)
	}
}
