package main

import (
  "github.com/waikiApp/dice"
  "strconv"
  "os"
  "fmt"
  "time"
  mrand "math/rand"
)

func genRandRange(start int, end int) (int) {
  mrand.Seed(time.Now().UnixNano())
  return mrand.Intn(end - start + 1) + start
}

func genDiceToken(size int) string {
  dl := dice.List()
  var token string
  for q := 0; q < size; q++ {
    var rolls string
    for i := 1; i < 6; i++ {
      rolls += fmt.Sprintf("%v", genRandRange(1, 6))
    }
    if q == (size -1) { token += fmt.Sprintf("%s", dl[rolls]) } else { token += fmt.Sprintf("%s ", dl[rolls]) }
  }
 return token
}

func main() {
  if len(os.Args) >= 2 {
    if size, err := strconv.Atoi(os.Args[1]); err != nil { fmt.Printf("ERROR: Invalid length parameter\n") } else {
      q := genDiceToken(size)
      fmt.Printf("%s\n", q)
    }
  } else { fmt.Printf("NAME:\n  pwgen - Diceware password generator.\n\nUSAGE:\n  %s [Password length]\n", os.Args[0]) }
}
