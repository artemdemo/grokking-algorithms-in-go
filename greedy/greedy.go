package greedy

import (
    "github.com/deckarep/golang-set"
)

// Actor is an item of the set
type Actor struct {
    name string
}

type Producer struct {
    name string
}

type MapOfActorSets map[*Producer]mapset.Set

func FindCoverage(desiredActors mapset.Set, actorGroups MapOfActorSets) ([]*Producer, bool) {
    var resultState []*Producer
    attempts := desiredActors.Cardinality()

    for {
        if desiredActors.Cardinality() == 0 || attempts < 1 {
            break
        }
        var bestProducer *Producer
        var actorsCovered mapset.Set
        for producer, producerActors := range actorGroups {
            // We need to compare only against actually covered actors.
            covered := desiredActors.Intersect(producerActors)
            // At each time we want to find out the group that will cover most of actors.
            if actorsCovered == nil || covered.Cardinality() > actorsCovered.Cardinality() {
                bestProducer = producer
                actorsCovered = covered
            }
        }

        // If there is a group that we covered, we'll remove it from the goal list.
        // So next time we'll try to cover only what is left.
        if actorsCovered != nil {
            desiredActors = desiredActors.Difference(actorsCovered)
        }
        if bestProducer != nil {
            resultState = append(resultState, bestProducer)
        }
        attempts--
    }

    return resultState, attempts > 0
}
