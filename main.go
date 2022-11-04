/*
This program first checks the included file "utensils.json" for a "spoon" value, and if it finds one
it continues on to make an API GET to a random dessert generator API, then formats the response into a sentence.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Dessert struct {
	Id      int    `json:"id"`
	Uid     string `json:"uid"`
	Variety string `json:"variety"`
	Topping string `json:"topping"`
	Flavor  string `json:"flavor"`
}

type Utensils struct {
	Utensil1 string `json:"utensil1`
	Utensil2 string `json:"utensil2`
	Utensil3 string `json:"utensil3`
}

func main() {

	// open the included json file
	jsonFile, err := os.Open("./utensils.json")

	if err != nil {
		log.Fatalf("Error during os.Open(): %v", err)
	}

	fmt.Println("Successfully Opened utensils.json")
	defer jsonFile.Close()

	// read the json into a []byte
	byteValue, _ := io.ReadAll(jsonFile)

	var utensils Utensils

	// unmarshal the []byte into a Utensils struct
	err = json.Unmarshal(byteValue, &utensils)
	if err != nil {
		log.Fatalf("Error during json Unmarshal: %v", err)
	}

	fmt.Println("Looking for a spoon...")

	// check for a "spoon" value in the Utensils struct
	if utensils.Utensil1 != "spoon" && utensils.Utensil2 != "spoon" && utensils.Utensil3 != "spoon" {
		log.Printf("What? There's no spoon?")
		return
	}

	fmt.Println("Found a spoon, that's all I need...")

	url := "https://random-data-api.com/api/dessert/random_dessert"

	// create http client
	client := &http.Client{}

	// create GET request for desserts API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// execute request with client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	var dessert Dessert

	// create a json decoder and decode response body into a Dessert struct
	if err := json.NewDecoder(resp.Body).Decode(&dessert); err != nil {
		log.Println(err)
	}

	fmt.Printf("This dessert looks tasty...%v flavored %v topped with %v!\n", dessert.Flavor, dessert.Variety, dessert.Topping)
}
