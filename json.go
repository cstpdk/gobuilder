package main

import(
    j "encoding/json"
    "github.com/codegangsta/martini"
    "net/http"
    "reflect"
)

const (
    contenttype   = "Content-Type"
    contentJSON   = "application/json"
)

/*
JSON function type for converting anything to json string
*/
type JSON func(v interface{}) string

/*
JSONEncoder injects a JSON function into requests that handles the encoding of
anything to JSON, setting http headers and sending a 500 http error if
encoding fails
*/
func JSONEncoder(c martini.Context, w http.ResponseWriter, r *http.Request){
    jfun := func(v interface{}) string {
        result, err := j.Marshal(v)

        if err != nil {
            http.Error(w, err.Error(), 500)
        }

        w.Header().Set(contenttype, contentJSON)
        return string(result)
    }

    c.Map(jfun)
}

/*
JSONReturnHandler
*/
func JSONReturnHandler(w http.ResponseWriter, vals []reflect.Value) {
    if len(vals) > 1 && vals[0].Kind() == reflect.Int {
        result, err := j.Marshal(vals[1].Interface())

        if err != nil {
            http.Error(w, err.Error(), 500)
        }

        w.WriteHeader(int(vals[0].Int()))
        w.Write(result)
    } else if len(vals) > 0 {
        result, err := j.Marshal(vals[0].Interface())

        if err != nil {
            http.Error(w, err.Error(), 500)
        }

        w.Write(result)
    }

}
