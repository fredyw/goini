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
ini.DeleteSection("section3")

err := goini.WriteFile(ini, "test.ini")
if err != nil {
    log.Fatal(err)
}
```