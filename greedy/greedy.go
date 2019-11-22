package greedy

import "github.com/deckarep/golang-set"

// Actor is an item of the set
type Actor struct {
    name string
}

type Producer struct {
    name string
}

type MapOfActorSets map[*Producer]mapset.Set

func FindCoverage(desiredActors mapset.Set, actorGroups MapOfActorSets) ([]*Producer, bool) {
    test := make(MapOfActorSets)
    return test, true
}
