# RouteVars

Easily get variables from any url.

## Example
```go
	var path = "/users/1234/john"
	var ok, vars = routevars.Match("/users/<<id:int>>/<<name:string>>", path)
	if ok {
		fmt.Println("Match!")
		for k, v := range vars {
			fmt.Printf("%s: %s\n", k, v)
		}
	} else {
		fmt.Println("No match")
	}
```

## Installation
```bash
go get github.com/Nigel2392/routevars
```
