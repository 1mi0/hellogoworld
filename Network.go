package main

import (
	"fmt"
	"net/http"
)

func initializeHttp() {
    http.HandleFunc("/print", printHandler)
    http.HandleFunc("/add", addHandler)
    http.HandleFunc("/delete", deleteHandler)

    http.ListenAndServe(":8090", nil)
}

func printHandler(w http.ResponseWriter, req *http.Request) {
    headNeighbour.Print(w)
}

func addHandler(w http.ResponseWriter, req *http.Request) {
    ans := parseArguments(req.RequestURI)
    if (len(ans) == 0) {
        fmt.Fprintf(w, "Error, could not find the proper arguments: \"%v\"", req.RequestURI)
        return
    }

    name, hasName := ans["name"]
    occupancy, hasOccupancy := ans["occupancy"]

    if hasName && hasOccupancy {
        n := &NeighbourHood{User{name, occupancy}, nil}
        Add(n)
        fmt.Fprintf(w, "Successfully added %v", n)
    } else {
        fmt.Fprintf(w, "Error, could not find the proper arguments: \"%v\"", req.RequestURI)
    }
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
    ans := parseArguments(req.RequestURI)
    if len(ans) == 0 {
        fmt.Fprintf(w, "Error, could not find \"name\" argument: %s", req.RequestURI)
        return
    }

    name, hasName := ans["name"]
    if hasName {
        if Remove(name) {
            fmt.Fprintf(w, "You've successfully removed \"%s\" from the list.", name)
        } else {
            fmt.Fprintf(w, "Error, could not find anyone in the list with the name: \"%s\"", name)
        }
    } else {
        fmt.Fprintf(w, "Error, could not find \"name\" argument: %s", req.RequestURI)
    }
}
