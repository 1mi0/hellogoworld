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

func (n *NeighbourHood) Add(o *NeighbourHood) {
    o.next = n.next
    n.next = o
}

func ll_remove(head *NeighbourHood, name string) bool {
    if head.person.name == name {
        head = head.next
        return true
    }

    priv := head
    next := head.next
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

func ll_add(head *NeighbourHood, o *NeighbourHood) {
    o.next = head
    head = o
}

func ll_print(n *NeighbourHood, w io.Writer) {
    if n.next != nil {
        ll_print(n.next, w)
    }
    fmt.Fprintf(w, "%v\n", n.String())
}

var headNeighbour *NeighbourHood = nil

func ll_populate() {
    n1 := NeighbourHood{User{"Ivan Petrov", "Bezraboten"}, nil}
    n2 := NeighbourHood{User{"Georgi Ivanov", "Obsht rabotnik"}, &n1}
    n3 := NeighbourHood{User{"Ivanka Petrova", "Servitiorka"}, &n2}

    headNeighbour = &n3
}
