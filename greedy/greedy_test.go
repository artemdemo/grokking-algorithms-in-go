package greedy

import (
    "testing"
    "github.com/deckarep/golang-set"
)

func Test_FindCoverage(t *testing.T) {
    t.Run("radioCoverage", func(t *testing.T) {
        mt := Actor{name: "mt"}
        wa := Actor{name: "wa"}
        or := Actor{name: "or"}
        id := Actor{name: "id"}
        nv := Actor{name: "nv"}
        ut := Actor{name: "ut"}
        ca := Actor{name: "ca"}
        az := Actor{name: "az"}

        desiredState := mapset.NewSetFromSlice([]interface{}{ &mt, &wa, &or, &id, &nv, &ut, &ca, &az })

        kone := Producer{name: "kone"}
        ktwo := Producer{name: "ktwo"}
        kthree := Producer{name: "kthree"}
        kfour := Producer{name: "kfour"}
        kfive := Producer{name: "kfive"}

        stations := make(mapOfActorSets)
        stations[&kone] = mapset.NewSetFromSlice([]interface{}{ &id, &nv, &ut })
        stations[&ktwo] = mapset.NewSetFromSlice([]interface{}{ &wa, &id, &mt })
        stations[&kthree] = mapset.NewSetFromSlice([]interface{}{ &or, &nv, &ca })
        stations[&kfour] = mapset.NewSetFromSlice([]interface{}{ &nv, &ut })
        stations[&kfive] = mapset.NewSetFromSlice([]interface{}{ &ca, &az })
    })
}
