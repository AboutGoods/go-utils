package debug

import (
    "encoding/json"
    "fmt"
)

//Json print any argument in pretty json indented.
//Be aware that this will not print the non-exported properties of a structure for example.
func Json(items ...interface{}){
    j, _:=json.MarshalIndent(items,"","    ")
    fmt.Println(string(j))

}


