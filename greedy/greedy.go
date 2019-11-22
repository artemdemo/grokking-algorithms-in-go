package greedy

import "github.com/deckarep/golang-set"

// Actor is an item of the set
type Actor struct {
    name string
}

type Producer struct {
    name string
}

type mapOfActorSets map[*Producer]mapset.Set

func FindCoverage(desiredActors mapset.Set, actorGroups mapOfActorSets) []*Producer {
    var set mapset.Set
    set := NewSetFromSlice([]interface{}{
        &YourType{Name: "Alise"},
        &YourType{Name: "Bob"},
        &YourType{Name: "John"},
        &YourType{Name: "Nick"},
    })
}
