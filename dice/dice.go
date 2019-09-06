package dice

import (
  "time"
  "math/rand"
)

func RollDice(count int, sides int) int {
  min := 1
  max := sides
  seed := rand.NewSource(time.Now().UnixNano())
  var total int

  for i := 0; i < count; i++ {
    total += rand.New(seed).Intn(max - min + 1) + min
  }

  // This should return an int
  return total
}
