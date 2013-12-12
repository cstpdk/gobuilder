package main

import(
    j "encoding/json"
    "github.com/codegangsta/martini"
    "net/http"
)

const (
    ContentType   = "Content-Type"
    ContentJSON   = "application/json"
)

/*
Json function type for converting anything to json string
*/
type Json func(v interface{}) string

/*
JsonEncoder injects a Json function into requests that handles the encoding of
anything to Json, setting http headers and sending a 500 http error if
encoding fails
*/
func JsonEncoder(c martini.Context, w http.ResponseWriter, r *http.Request){
    var jfun Json = func(v interface{}) string {
        result, err := j.Marshal(v)

        if err != nil {
            http.Error(w, err.Error(), 500)
        }

        w.Header().Set(ContentType, ContentJSON)
        return string(result)
    }

    c.Map(jfun)
}
