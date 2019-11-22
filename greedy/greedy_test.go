package greedy

import (
    "fmt"
    "testing"
    "github.com/deckarep/golang-set"
)

func producersToString(producers []*Producer) string {
    var producersList []string
    for _, producer := range producers {
        producersList = append(producersList, producer.name)
    }
    return fmt.Sprint(producersList)
}

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

        stations := make(MapOfActorSets)
        stations[&kone] = mapset.NewSetFromSlice([]interface{}{ &id, &nv, &ut })
        stations[&ktwo] = mapset.NewSetFromSlice([]interface{}{ &wa, &id, &mt })
        stations[&kthree] = mapset.NewSetFromSlice([]interface{}{ &or, &nv, &ca })
        stations[&kfour] = mapset.NewSetFromSlice([]interface{}{ &nv, &ut })
        stations[&kfive] = mapset.NewSetFromSlice([]interface{}{ &ca, &az })

        result, ok := FindCoverage(desiredState, stations)

        expectedOk := true
        if ok != expectedOk {
            t.Fatalf("FindCoverage() should return ok=`%v`, got \"%v\" instead", expectedOk, ok)
        }

        resultStr := producersToString(result)
        expectedResult := "[A C D]"
        if resultStr != expectedResult {
            t.Fatalf("FindCoverage() should return path=`%v`, got \"%v\" instead", expectedResult, resultStr)
        }
    })
}
