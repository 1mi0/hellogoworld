package main

import (
	"strings"
)

func parseArguments(URI string) map[string]string {
    split := strings.Split(URI, "?")
    if len(split) < 2 {
        return make(map[string]string, 0)
    }
    
    pairs := strings.Split(split[1], "&")
    var properlySplitPairs = make(map[string]string, 0)

    for _,element := range pairs {
        toSplit := strings.ReplaceAll(element, "%20", " ")
        t := strings.Split(toSplit, "=")
        properlySplitPairs[t[0]] = t[1]
    }

    return properlySplitPairs
}

func main() {
    populate()

    initializeHttp()
}
