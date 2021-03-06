# goini
A Go library for INI parsing.

### How to Get
    go get -u github.com/fredyw/goini

### Doumentation
[https://godoc.org/github.com/fredyw/goini](https://godoc.org/github.com/fredyw/goini)

### Usage
**Reading an INI file**

```go
import "github.com/fredy/goini"

// setting it to true to preserve the order
ini, err := goini.ReadFile("test.ini"), true)
if err != nil {
    log.Fatal(err)
}
val, found := ini.GetOption("section", "option")
if found {
    fmt.Println(val)
}
for _, section := range ini.Sections() {
    for _, option := range ini.Options(section) {
        val, value := ini.GetOption(section, option)
        if found {
            fmt.Println(option, "-->", value)
        }
    }
}
```

**Writing an INI file**
```go
import "github.com/fredy/goini"

// setting it to true to preserve the order
ini := goini.NewINI(true)
ini.AddOption("section1", "option1", "value1")
ini.AddOption("section2", "option2", "value2")
ini.AddSection("section3")
ini.RemoveSection("section3")
ini.RemoveOption("section2", "option2")

err := goini.WriteFile(ini, "test.ini")
if err != nil {
    log.Fatal(err)
}
```