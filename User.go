package main

import (
	"fmt"
	"io"
)

type User struct {
    name        string
    occupation  string
}

type NeighbourHood struct {
    person  User
    next    *NeighbourHood
}

func (n NeighbourHood) String() string {
    return fmt.Sprintf("%v", n.person)
}

func (n NeighbourHood) Print(w io.Writer) {
    if n.next != nil {
        n.next.Print(w)
    }
    fmt.Fprintf(w, "%v\n", n.String())
}

func (n *NeighbourHood) Add(o *NeighbourHood) {
    o.next = n.next
    n.next = o
}

func Remove(name string) bool {
    if headNeighbour.person.name == name {
        headNeighbour = headNeighbour.next
        return true
    }

    priv := headNeighbour
    next := headNeighbour.next
    for next != nil {
        if next.person.name == name {
            priv.next = next.next
            return true
        }
        priv = next
        next = next.next
    }
    
    return false
}

func Add(o *NeighbourHood) {
    o.next = headNeighbour
    headNeighbour = o
}

var headNeighbour *NeighbourHood = nil

func populate() {
    n1 := NeighbourHood{User{"Ivan Petrov", "Bezraboten"}, nil}
    n2 := NeighbourHood{User{"Georgi Ivanov", "Obsht rabotnik"}, &n1}
    n3 := NeighbourHood{User{"Ivanka Petrova", "Servitiorka"}, &n2}

    headNeighbour = &n3
}
