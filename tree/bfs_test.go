package tree

import (
    "testing"
)

func Test_Bfs(t *testing.T) {
    t.Run("findASeller", func(t *testing.T) {
        friendsList := make(PersonsMap)

        alex := Person{name: "Alex", isSeller: false}
        alice := Person{name: "Alice", isSeller: false}
        bob := Person{name: "Bob", isSeller: false}
        claire := Person{name: "Claire", isSeller: false}

        alex.addFriend(&alice)
        alex.addFriend(&bob)
        alex.addFriend(&claire)

        peggy := Person{name: "Peggy", isSeller: false}

        bob.addFriend(&peggy)

        thom := Person{name: "Thom", isSeller: false}
        johny := Person{name: "Johny", isSeller: true}

        claire.addFriend(&thom)
        claire.addFriend(&johny)

        friendsList[alex.name] = alex
        friendsList[bob.name] = bob
        friendsList[alice.name] = alice
        friendsList[peggy.name] = peggy
        friendsList[claire.name] = claire
        friendsList[thom.name] = thom
        friendsList[johny.name] = johny

        person, ok := SearchForSeller(friendsList, "Alex")

        if ok != true {
            t.Fatalf("SearchForSeller() should return `%v`, got \"%v\" instead", true, ok)
        }
        if person != &johny {
            t.Fatalf("SearchForSeller() should return pointer to %v, got %v instead", &johny, person)
        }
    })
}
