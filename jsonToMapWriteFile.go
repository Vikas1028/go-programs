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
func jsonToMap() map[string]string {

	jsonData := `{
		"Thinly Sliced Peeled Apples" : "6 Cups",
		"sugar" : "3/4 cup",
		"flour" : "2 tablespoons",
		"cinamon" : "1/4 teaspoon",
		"nutmeg" : "1/8 tablespoon",
		"lemon juice": "1 tablespoon"
		}`

	var recipeIngredientsMap map[string]string
	err := json.Unmarshal([]byte(jsonData), &recipeIngredientsMap)
	errorHandling(err)

	for key, value := range recipeIngredientsMap {
		fmt.Println(key, "->", value)
	}
	return recipeIngredientsMap
}

/*
description: Write map data in file
param : RecipeIngredientsInMap map[string]string (map data)
*/
func writeFile(recipeIngredientsMap map[string]string) {

	file, err := os.Create("ApplePieRecipeIngredients.txt")
	errorHandling(err)
	defer file.Close()

	for key, value := range recipeIngredientsMap {
		fileData := fmt.Sprintf("%s -> %s\n", key, value)
		_, err := file.WriteString(fileData)
		errorHandling(err)
	}
}

func main() {

	recipeIngredientsMap := jsonToMap()
	writeFile(recipeIngredientsMap)
}
