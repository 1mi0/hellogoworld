package main

import (
	"fmt"
	"net/http"
	"strings"
)

func initializeHttp() {
    srv := &http.Server { Addr: ":8090" }

    http.HandleFunc("/print", printHandler)
    http.HandleFunc("/add", addHandler)
    http.HandleFunc("/delete", deleteHandler)

    srv.ListenAndServe()
}

func printHandler(w http.ResponseWriter, req *http.Request) {
    ll_print(headNeighbour, w)
}

func addHandler(w http.ResponseWriter, req *http.Request) {
    parsedResult := parseArguments(req.RequestURI)
    if parsedResult.HasError() {
        fmt.Fprintf(w, "%s", parsedResult.Error())
        return
    }

    mappedArguments := *parsedResult.Get()
    name, hasName := mappedArguments["name"]
    occupancy, hasOccupancy := mappedArguments["occupancy"]

    if hasName && hasOccupancy {
        n := &NeighbourHood{User{name, occupancy}, nil}
        ll_add(headNeighbour, n)
        fmt.Fprintf(w, "Successfully added %v", n)
    } else {
        fmt.Fprintf(w, "Error, could not find the proper arguments: \"%v\"", req.RequestURI)
    }
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
    parsedResult := parseArguments(req.RequestURI)
    if parsedResult.HasError() {
        fmt.Fprintf(w, "%s", parsedResult.Error())
        return
    }

    name, hasName := (*parsedResult.Get())["name"]
    if hasName {
        if ll_remove(headNeighbour, name) {
            fmt.Fprintf(w, "You've successfully removed \"%s\" from the list.", name)
        } else {
            fmt.Fprintf(w, "Error, could not find anyone in the list with the name: \"%s\"", name)
        }
    } else {
        fmt.Fprintf(w, "Error, could not find \"name\" argument: %s", req.RequestURI)
    }
}

func parseArguments(URI string) Result[map[string]string] {
    split := strings.Split(URI, "?")
    if len(split) < 2 {
        return Error[map[string]string] { "No arguments passed" }
    }
    
    pairs := strings.Split(split[1], "&")
    var properlySplitPairs = make(map[string]string, 0)

    for _,element := range pairs {
        toSplit := strings.ReplaceAll(element, "%20", " ")
        t := strings.Split(toSplit, "=")
        properlySplitPairs[t[0]] = t[1]
    }

    return Success[map[string]string] { properlySplitPairs }
}
