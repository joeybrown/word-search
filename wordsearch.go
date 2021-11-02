package main

import (
  "fmt"
  "strings"
)

type Direction int64

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
  North Direction = iota
  NorthEast
  East
  SouthEast
  South
  SouthWest
  West
  NorthWest
)

func (direction Direction) String() string {
  switch direction {
  case North:
    return "north"
  case NorthEast:
    return "northeast"
  case East:
    return "est"
  case SouthEast:
    return "southeast"
  case South:
    return "south"
  case SouthWest:
    return "southwest"
  case West:
    return "west"
  case NorthWest:
    return "northwest"
  }
  return "unknown"
}

type Word struct {
  Value string
  OriginX int
  OriginY int
  Direction Direction
}

func (word Word) String() string {
  return fmt.Sprintf("%s: origin at (%d, %d) direction %s", word.Value, word.OriginX, word.OriginY, word.Direction)
}

type WordSearch struct {
  Graph Graph
  Words []Word
}

type Graph [][]string

func (graph Graph) String() string {
  var rowsJoined []string
  for _, row := range graph {
    joined := strings.Join(row, " ")
    rowsJoined = append(rowsJoined, joined)
  }
  graphJoined := strings.Join(rowsJoined, "\n")
  return graphJoined
}

func BuildWordSearch(inputWords []string) WordSearch{
  graph := Graph{
    {"H", "B", "C", "D", "D"},
    {"D", "E", "F", "G", "H"},
    {"G", "H", "L", "M", "N"},
    {"H", "B", "C", "L", "L"},
    {"D", "E", "F", "O", "O"},
  }

  var words []Word
  for i := range inputWords{
    words = append(words, Word{Value: inputWords[i]})
  }

  return WordSearch{graph, words}
}

func (graph Graph) GetNextPosition(cursorX int, cursorY int, direction Direction) (inBounds bool,  int, newPosY int){
  switch direction {
  case North:
    if cursorX >= len(graph) - 1 || cursorX < -1 {
      return false, cursorX, cursorY
    }
    return true, cursorX + 1, cursorY

  //case NorthEast:
  //  return "northeast"
  //case East:
  //  return "est"
  //case SouthEast:
  //  return "southeast"
  //case South:
  //  return "south"
  //case SouthWest:
  //  return "southwest"
  //case West:
  //  return "west"
  //case NorthWest:
  //  return "northwest"
  }
  return false, -1, -1
}

func FindWord(graph Graph, word string, cursorX int, cursorY int) (wasFound bool, xOrigin int, yOrigin int, direction Direction){
  return
}

func (graph Graph) FindWord(word string) (wasFound bool, xOrigin int, yOrigin int, direction Direction){
  return FindWord(graph, word, 0, 0);
}

func main() {
  words := []string {"Hello"}
  wordSearch := BuildWordSearch(words)
  fmt.Println(wordSearch.Graph)

  for i := range words{
    fmt.Println(words[i])
  }
}
