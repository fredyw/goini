# goini
A Go library for INI parsing.

### How to Get
    go get -u github.com/fredyw/gomerge

### Usage
**Reading an INI file**

```go
// setting it to true to preserve the order
import "github.com/fredy/goini"

ini, err := goini.ReadFile("test.ini"), true)
if err != nil {
    log.Fatal(err)
}
val, found := ini.GetOption("section", "option")
if found {
    fmt.Println(val)
}
````

**Writing an INI file**
```go
import "github.com/fredy/goini"

// setting it to true to preserve the order
ini := goini.NewINI(true)
ini.AddOption("section", "option", "value")
```