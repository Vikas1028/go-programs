package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
description: If error occur then print error and teminate the program
param: error
*/
func errorHandling(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: Convert json object data into map
return: RecipeIngredientsInMap map[string]string (return map of json data)
*/
func jsonTomap() map[string]string {
	RecipeIngredientsInJson := `{
		"Thinly Sliced Peeled Apples" : "6 Cups",
        "sugar" : "3/4 cup",
        "flour" : "2 tablespoons",
        "cinamon" : "1/4 teaspoon",
        "nutmeg" : "1/8 tablespoon",
        "lemon juice": "1 tablespoon"
	}`

	var RecipeIngredientsInMap map[string]string
	err := json.Unmarshal([]byte(RecipeIngredientsInJson), &RecipeIngredientsInMap)
	errorHandling(err)

	for key, value := range RecipeIngredientsInMap {
		fmt.Println(key, " -> ", value)
	}

	return RecipeIngredientsInMap
}

/*
description: Convert map into json object and write these json object data in json file
param : RecipeIngredientsInMap map[string]string (map data)
*/
func mapToJson(RecipeIngredientsInMap map[string]string) {

	RecipeIngredientsInJson, err := json.MarshalIndent(RecipeIngredientsInMap, "", "  ")
	errorHandling(err)

	file, err := os.Create("RecipeIngredients.json")

	_, err = file.Write(RecipeIngredientsInJson)
	errorHandling(err)
	fmt.Println(string(RecipeIngredientsInJson))
}

func main() {

	RecipeIngredientsInMap := jsonTomap()

	mapToJson(RecipeIngredientsInMap)
}
