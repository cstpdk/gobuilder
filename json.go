package main

import(
    j "encoding/json"
    "net/http"
    "reflect"
)

const (
    contenttype   = "Content-Type"
    contentJSON   = "application/json"
)

/*
JSONReturnHandler converts return values into JSON
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
