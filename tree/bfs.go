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
// I want to keep pointers to friends.
// Not actual objects.
func (person *Person) addFriend(friend *Person) {
    person.friends = append(person.friends, friend)
}

type PersonsMap map[string]Person

// Check whether slice (haystack) has given item (needle) in it
func contains(haystack []*Person, needle *Person) bool {
    for idx, _ := range haystack {
        // Here I'm using index reference, and not item.
        // Because item will always be a copy of an object.
        if haystack[idx] == needle {
            return true
        }
    }
    return false
}

/**
 * Breadth-first search
 */
func SearchForSeller(personsMap PersonsMap, name string) (*Person, bool) {
    searchQueue := deque.New()
    firstInLine := personsMap[name]
    searchQueue.PushRight(&firstInLine)

    // Here I'm using index reference, and not item.
    // Because item will always be a copy of an object.
    for idx, _ := range firstInLine.friends {
        searchQueue.PushRight(firstInLine.friends[idx])
    }
    var searched []*Person
    searched = append(searched, &firstInLine)

    for {
       if searchQueue.Size() == 0 {
           break
       }
       personPointer, ok := searchQueue.PopLeft().(*Person)

       if ok && contains(searched, personPointer) == false {
           if personPointer.isSeller {
               return personPointer, true
           } else {
               for _, friend := range personPointer.friends {
                   searchQueue.PushRight(friend)
               }
               searched = append(searched, personPointer)
           }
       }
    }
    return &Person{}, false
}

