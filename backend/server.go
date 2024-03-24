package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Entity struct {
	videoUrl    string
	entityName  string
	marketValue int
}

//entityCode string is generated using uuid library

// defining an empty bid book
var bidBook map[string][]int

// defining an empty entity list
var entityList []Entity

func main() {
	bidBook = make(map[string][]int)
	// // Add values to the map
	// bidBook["key1"] = []int{1, 2, 3}
	// bidBook["key2"] = []int{4, 5, 6}
	// bidBook["key3"] = []int{7, 8, 9}
	router := gin.Default()
	router.Run(":8080")

	router.GET("/placeBet", func(c *gin.Context) {
		username := c.Query("username")
		betAmount := c.Query("betAmount")
		entityCode := c.Query("entityCode")
		placeBet(username, betAmount, entityCode)
	})

	router.GET("/getNewEntity", getNewEntity)

}

// this function places bets for the user
// the function takes in the username, bet amount, and the entity code (Each item has an entity code)
// this function also checks if the bid is already placed by other user for certain bid-price
// will be using in-memory data structure to store the bets to make it faster
// maybe use dictionary
// returns 1 if the bet is placed successfully
// returns -1 if the bet is already placed by another user
// whether 1 or -1 determines what to display in the frontend
// also use the return value to update the value of the bet in the frontend
// if returns 1, then update the value of the bet in the frontend by 50 cents
func placeBet(username string, betAmount int, entityCode string) int {
	res, err := bidBook[entityCode]
	if err == false {
		//if the bidBook for an entity with entityCode is already present
		lastBid := res[len(res)-1]
		if lastBid == betAmount {
			fmt.Println("Bid already placed by another user")
			return -1 // return -1 if the bid is already placed by another user
		}
	}
	//if the bidBook for an entity with entityCode is not present
	bidBook[entityCode] = append(bidBook[entityCode], betAmount)
	return 1
}

func getNewEntity() Entity {
	return entityList[0]
}

func postNewEntity(entity Entity) {
	entityList = append(entityList, entity)
}
