package routevars_test

import (
	"fmt"
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

func TestFormatSafe(t *testing.T) {
	var path = "/users/<<id:int>>/<<name:string>>"
	var formatter = routevars.URLFormatter(path)
	var url, err = formatter.FormatSafe(1234, "john")
	if err != nil {
		t.Error("URLFormatter failed: " + err.Error())
	}
	if url != "/users/1234/john" {
		t.Error("URLFormatter failed, url is not correct.")
	}
	var newurl, newerr = formatter.FormatSafe("john", 1234)
	if newerr == nil {
		t.Error("URLFormatter failed: " + newerr.Error())
		return
	}
	if newurl != "" {
		t.Error("URLFormatter failed, url was returned.")
		return
	}

	t.Log("URLFormatter Safe passed: " + url)
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

func BenchmarkMatch(b *testing.B) {
	var path = "/users/<<id:int>>/<<name:string>>"
	var other = "/users/%d/john"
	for i := 0; i < b.N; i++ {
		routevars.Match(path, fmt.Sprintf(other, i))
	}
}
