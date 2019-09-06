package attributes

type Attribute struct {
	Value       int
	Fractional  int
}

// A function that constructs the key for lookup.
// Takes in two integer arguments representing an ability score and its
// fractional value (13 - 57/100)
//  - Returns a float32 value for key lookup.
func AttributeLookupKeyConstructor(val int, frac int) float32 {
	key := float32(val)
	if frac >= 51 {
		key += .5
	}

	return key
}
