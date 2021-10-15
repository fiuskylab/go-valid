# go-valid
A struct validator

```golang
package main

import (
  "github.com/fiuskylab/go-valid/validator"
)

type Foo struct {
  Field1 string `validator:""`
  Field2 int
}

var fooRules map[string][]string = {
  "field1": [
    "required", 
    "min:2", 
    "max:10", 
    "type:string"
  ],
  "field2": [
    "required", 
    "min:2", 
    "max:10", 
    "regex:[a-zA-Z]{1,10}", 
    "type:int"
  ],
}

func main() {
  foo := Foo{
    Field1: "something",
    Field2: 20,
  }
  res, err := validator.Check(&foo, fooRules)
  if err != nil {
    // handles error
  }

  if !res.IsValid() {
    // not valid
  }
}
```
