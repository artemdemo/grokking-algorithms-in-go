package graph

/**
 * `package deque`
 * Package deque implements a double ended queue supporting arbitrary types (even a mixture).
 * Internally it uses a dynamically growing circular slice of blocks,
 * resulting in faster resize than a simple dynamic array/slice would allow, yet less gc overhead.
 * @source https://godoc.org/gopkg.in/karalabe/cookiejar.v1/collections/deque
 */

import (
    "gopkg.in/karalabe/cookiejar.v1/collections/deque"
    "math"
)

/**
 * Node struct.
 * Shouldn't be created directly only via `NewNode`
 */
type Node struct {
    name string
    edges []Edge
}
func NewNode(name string) *Node {
    node := new(Node)
    node.name = name
    return node
}
func (node *Node) addConnection(toNode *Node, price int) {
    edge := Edge{price:price, fromNode:node, toNode:toNode}
    node.edges = append(node.edges, edge)
    // I'm adding the same edge to `toNode` as well
    // this way I'll be able to walk back from the endNode to the start
    toNode.edges = append(toNode.edges, edge)
}

type Edge struct {
    price int
    fromNode *Node
    toNode *Node
}

/**
 * Map of costs for each node.
 * Will be filled as a result of Dijkstra's algorithm.
 */
type Costs map[*Node]float64

func addNodeEdgesToWalkQueue(queue *deque.Deque, costs Costs, node *Node) {
    for _, edge := range node.edges {
        if edge.toNode != node {
            if toNodeCurrentPrice, ok := costs[edge.toNode]; ok {
                // In order to prevent circular dependencies,
                // I'm checking whether edge will help te reduce price of the next node.
                if toNodeCurrentPrice > float64(edge.price) + costs[edge.fromNode]  {
                    queue.PushRight(edge)
                }
            } else {
                queue.PushRight(edge)
            }
        }
    }
}

/**
 * Find path between two nodes.
 * This function is assuming that all nodes are already priced.
 */
func findPathBetweenNodes(path []*Node, costs Costs, startNode *Node, endNode *Node) []*Node {
    var nextNode *Node

    for _, edge := range endNode.edges {
        // The idea is to take endNode and find edge that has been used.
        // If "endNode cost" minus "the edge" equals to the "fromNode",
        // then this exact path was taken.
        if edge.toNode == endNode && costs[endNode] - float64(edge.price) == costs[edge.fromNode] {
            nextNode = edge.fromNode
        }
    }
    var newPath []*Node
    if nextNode != startNode {
        if len(path) == 0 {
            newPath = append(path, endNode)
        }
        newPath = append(newPath, append(path, nextNode)...)
        return findPathBetweenNodes(newPath, costs, startNode, nextNode)
    }
    return append(path, startNode)
}

/**
 * Dijkstra's algorithm
 * Will find shortest path between 2 given nodes
 * @wiki https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
 */
func Dijkstra(startNode *Node, endNode *Node) ([]*Node, float64, bool) {
    walkQueue := deque.New()
    costs := make(Costs)

    // Start node will always have cost of 0
    costs[startNode] = 0

    addNodeEdgesToWalkQueue(walkQueue, costs, startNode)

    for {
        if walkQueue.Size() == 0 {
            break
        }
        edge, ok := walkQueue.PopLeft().(Edge)

        if ok {
            var oldCost = math.Inf(+1)
            if val, ok := costs[edge.toNode]; ok {
                oldCost = val
            }
            newCost := costs[edge.fromNode] + float64(edge.price)

            if newCost < oldCost {
                costs[edge.toNode] = newCost
            }
            addNodeEdgesToWalkQueue(walkQueue, costs, edge.toNode)
        }
    }

    var resultPath []*Node

    endNodePrice, endNodeReached := costs[endNode]

    if endNodeReached {
        resultPath = findPathBetweenNodes(resultPath, costs, startNode, endNode)
        // reversing result slice
        // So order will appear from the left, to the right (from the start, to the end)
        for i := len(resultPath)/2-1; i >= 0; i-- {
            opp := len(resultPath)-1-i
            resultPath[i], resultPath[opp] = resultPath[opp], resultPath[i]
        }
    }

    return resultPath, endNodePrice, endNodeReached
}
