package debug

import (
    "encoding/json"
    "fmt"
)

func Json(items ...interface{}){
    j, _:=json.MarshalIndent(items,"","    ")
    fmt.Println(string(j))

}


