package graph

import (
    "fmt"
    "testing"
)

func Test_Dijkstra(t *testing.T) {
    t.Run("calcSimplePath", func(t *testing.T) {
        nodeA := NewNode("A")
        nodeB := NewNode("B")
        nodeC := NewNode("C")
        nodeD := NewNode("D")

        nodeA.addConnection(nodeB, 5)
        nodeA.addConnection(nodeC, 1)

        nodeB.addConnection(nodeD, 5)
        nodeC.addConnection(nodeD, 1)

        path, price, ok := Dijkstra(nodeA, nodeD)
        var pathNames []string

        for _, node := range path {
            pathNames = append(pathNames, node.name)
        }

        fmt.Println(pathNames)

        expectedOk := true
        if ok != expectedOk {
            t.Fatalf("Dijkstra() should return ok=`%v`, got \"%v\" instead", expectedOk, ok)
        }

        expectedPrice := 2
        if price != float64(expectedPrice) {
            t.Fatalf("Dijkstra() should return price=`%v`, got \"%v\" instead", expectedPrice, price)
        }
    })
}
