package main

import "fmt"

func main() {
  myMap := make(map[string]int)

  myMap["key1"] = 7
  myMap["key2"] = 13

  fmt.Println("map: ", myMap)

  version1 := myMap["key1"]
  fmt.Println("v1: ", version1)

  fmt.Println("length: ", len(myMap))

  delete(myMap, "key2")
  fmt.Println("map: ", myMap)

  _, item := myMap["key2"]
  fmt.Println("item: ", item)

  length := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", length)
}
