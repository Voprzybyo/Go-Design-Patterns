package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var recipesCount = 0

type Cookbook struct {
	recipes []string
}

//Main Cookbook methods...

func (cb *Cookbook) String() string {
	return strings.Join(cb.recipes, "\n")
}

func (cb *Cookbook) AddRecipe(text string) int {

	recipesCount++
	recipe := fmt.Sprintf("%d: %s", recipesCount, text)
	cb.recipes = append(cb.recipes, recipe)
	return recipesCount
}

func (cb *Cookbook) RemoveRecipe(index int) int {
	recipesCount--
	cb.recipes = append(cb.recipes[:index], cb.recipes[index+1:]...)
	return recipesCount
}

// Breaks Single Responsibility Principle...

func (cb *Cookbook) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(cb.String()), 0644)
}

func (cb *Cookbook) Load(filename string) {
	//...
}

func (cb *Cookbook) LoadFromWeb(url *url.URL) {
	//...
}

// Better solution below...

var lineSeparator = "\n"

func SaveToFile(cb *Cookbook, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(cb.recipes, lineSeparator)), 0644)
}

// ... or ...

type FileSave struct {
	lineSeparator string
}

func (fs *FileSave) saveToFile(cb *Cookbook, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(cb.recipes, fs.lineSeparator)), 0644)
}

func main() {

	cb := Cookbook{}
	cb.AddRecipe("Steak")
	cb.AddRecipe("Pizza")
	cb.AddRecipe("Sushi")
	fmt.Println(strings.Join(cb.recipes, "\n"))

	//Separate function
	SaveToFile(&cb, "Recipes.txt")

	fs := FileSave{"\n"}
	fs.saveToFile(&cb, "Recipes.txt")
}
