# RouteVars

Easily get variables from any url.

## Example

```go
	var path = "/users/1234/john"
	var ok, vars = Match("/users/<<id:int>>/<<name:string>>", path)
	if ok {
		fmt.Println("Matched!")
		for k, v := range vars {
			fmt.Println(k, v)
		}
	} else {
		fmt.Println("No match")
	}
```

## Installation

```bash
go get github.com/Nigel2392/routevars
```
