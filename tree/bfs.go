package tree

/**
 * `package deque`
 * Package deque implements a double ended queue supporting arbitrary types (even a mixture).
 * Internally it uses a dynamically growing circular slice of blocks,
 * resulting in faster resize than a simple dynamic array/slice would allow, yet less gc overhead.
 * @source https://godoc.org/gopkg.in/karalabe/cookiejar.v1/collections/deque
 */

import (
    "gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

type Person struct {
    name string
    // `isSeller` is a custom property, just to look something for
    isSeller bool
    friends []*Person
}
func (person *Person) addFriend(friend Person) {
    person.friends = append(person.friends, &friend)
}

type PersonsMap map[string]Person

// Check whether slice (haystack) has an given item (needle) in it
func contains(haystack []Person, needle Person) bool {
    for _, item := range haystack {
        if &item == &needle {
            return true
        }
    }
    return false
}

/**
 * Breadth-first search
 */
func SearchForSeller(personsMap PersonsMap, name string) (Person, bool) {
    searchQueue := deque.New()
    for _, friend := range personsMap[name].friends {
        searchQueue.PushRight(friend)
    }
    var searched []Person
    searched = append(searched, personsMap[name])

    for {
        if searchQueue.Size() == 0 {
            break
        }
        person, ok := searchQueue.PopLeft().(Person)
        if ok && contains(searched, person) == false {
            if person.isSeller {
                return person, true
            } else {
                for _, friend := range person.friends {
                    searchQueue.PushRight(friend)
                }
                searched = append(searched, person)
            }
        }
    }
    return Person{}, false
}

