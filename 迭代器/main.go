package main

import "fmt"
import iterator "p15/iterator"

func main() {
	user1 := &iterator.User{
		Name: "a",
		Age: 30,
	}
	user2 := &iterator.User{
		Name: "b",
		Age: 20,
	}

	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}
	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}