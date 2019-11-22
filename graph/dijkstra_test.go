package graph

import (
    "fmt"
    "testing"
)

func pathToString(path []*Node) string {
    var pathNames []string
    for _, node := range path {
        pathNames = append(pathNames, node.name)
    }
    return fmt.Sprint(pathNames)
}

func Test_Dijkstra(t *testing.T) {
    t.Run("calcSimplePath", func(t *testing.T) {
        nodeA := NewNode("A")
        nodeB := NewNode("B")
        nodeC := NewNode("C")
        nodeD := NewNode("D")

        nodeA.addConnection(nodeB, 5)
        nodeA.addConnection(nodeC, 1)

        nodeB.addConnection(nodeD, 2)
        nodeC.addConnection(nodeD, 3)

        path, price, ok := Dijkstra(nodeA, nodeD)

        expectedOk := true
        if ok != expectedOk {
            t.Fatalf("Dijkstra() should return ok=`%v`, got \"%v\" instead", expectedOk, ok)
        }

        expectedPrice := 4
        if price != float64(expectedPrice) {
            t.Fatalf("Dijkstra() should return price=`%v`, got \"%v\" instead", expectedPrice, price)
        }

        resultStr := pathToString(path)
        expectedResult := "[A C D]"
        if resultStr != expectedResult {
            t.Fatalf("Dijkstra() should return path=`%v`, got \"%v\" instead", expectedResult, resultStr)
        }
    })

    t.Run("calcBookExample", func(t *testing.T) {
        nodeStart := NewNode("Start")
        nodeA := NewNode("A")
        nodeB := NewNode("B")
        nodeFin := NewNode("Fin")

        nodeStart.addConnection(nodeA, 6)
        nodeStart.addConnection(nodeB, 2)
        nodeB.addConnection(nodeA, 3)
        nodeB.addConnection(nodeFin, 5)
        nodeA.addConnection(nodeFin, 1)

        path, price, ok := Dijkstra(nodeStart, nodeFin)

        expectedOk := true
        if ok != expectedOk {
            t.Fatalf("Dijkstra() should return ok=`%v`, got \"%v\" instead", expectedOk, ok)
        }

        expectedPrice := 6
        if price != float64(expectedPrice) {
            t.Fatalf("Dijkstra() should return price=`%v`, got \"%v\" instead", expectedPrice, price)
        }

        resultStr := pathToString(path)
        expectedResult := "[Start B A Fin]"
        if resultStr != expectedResult {
            t.Fatalf("Dijkstra() should return path=`%v`, got \"%v\" instead", expectedResult, resultStr)
        }
    })
}
