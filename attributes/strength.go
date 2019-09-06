package attributes

import "fmt"

type Strength struct {
	Attribute
	DmgMod		int
	FeatofStr	int
	Lift		int
	CarryNoEnc	int
	CarryLtEnc	int
	CarryMdEnc	int
	CarryHvEnc	int
	Drag		int
}

// Damage Modifier
// ----------------------------------------------------------------------------

// A closure function that houses the data used for Damage Modifiers based on
// Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func DamageModDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: -7,
		1.5: -6,
		2.0: -6,
		2.5: -5,
		3.0: -5,
		3.5: -4,
		4.0: -4,
		4.5: -4,
		5.0: -3,
		5.5: -3,
		6.0: -3,
		6.5: -2,
		7.0: -2,
		7.5: -2,
		8.0: -1,
		8.5: -1,
		9.0: -1,
		9.5: -1,
		10.0: 0,
		10.5: 0,
		11.0: 0,
		11.5: 0,
		12.0: 1,
		12.5: 1,
		13.0: 1,
		13.5: 1,
		14.0: 2,
		14.5: 2,
		15.0: 2,
		15.5: 3,
		16.0: 3,
		16.5: 3,
		17.0: 4,
		17.5: 4,
		18.0: 4,
		18.5: 5,
		19.0: 5,
		19.5: 6,
		20.0: 6,
		20.5: 7,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

// The getter function for the Damage Modifier based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the Damage Modifier value.
func GetStrDamageMod(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return DamageModDict()(attrKey)
}

// Feat of Strength
// ----------------------------------------------------------------------------

// The getter function for the Feat of Strength modifier based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the Feat of Strength modifier value.
func GetFeatOfStrength(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return FeatOfStrengthDict()(attrKey)
}

// A closure function that houses the data used for Feat of Strength based on
// Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func FeatOfStrengthDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: -14,
		1.5: -13,
		2.0: -12,
		2.5: -11,
		3.0: -10,
		3.5: -9,
		4.0: -9,
		4.5: -8,
		5.0: -7,
		5.5: -7,
		6.0: -6,
		6.5: -5,
		7.0: -5,
		7.5: -4,
		8.0: -3,
		8.5: -3,
		9.0: -2,
		9.5: -1,
		10.0: 0,
		10.5: 0,
		11.0: 0,
		11.5: 0,
		12.0: 1,
		12.5: 2,
		13.0: 3,
		13.5: 4,
		14.0: 5,
		14.5: 6,
		15.0: 7,
		15.5: 8,
		16.0: 9,
		16.5: 10,
		17.0: 11,
		17.5: 12,
		18.0: 13,
		18.5: 14,
		19.0: 15,
		19.5: 16,
		20.0: 17,
		20.5: 18,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

// Lift
// ----------------------------------------------------------------------------
// The getter function for the Lift modifier based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the maximum Lift value.
func GetStrLift(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return LiftStrengthDict()(attrKey)
}

// A closure function that houses the data used for maximum Lift based on
// Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func LiftStrengthDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: 32,
		1.5: 42,
		2.0: 52,
		2.5: 58,
		3.0: 64,
		3.5: 76,
		4.0: 88,
		4.5: 99,
		5.0: 110,
		5.5: 120,
		6.0: 130,
		6.5: 140,
		7.0: 149,
		7.5: 157,
		8.0: 166,
		8.5: 173,
		9.0: 181,
		9.5: 187,
		10.0: 194,
		10.5: 200,
		11.0: 205,
		11.5: 210,
		12.0: 215,
		12.5: 220,
		13.0: 225,
		13.5: 230,
		14.0: 235,
		14.5: 240,
		15.0: 245,
		15.5: 267,
		16.0: 291,
		16.5: 318,
		17.0: 347,
		17.5: 380,
		18.0: 417,
		18.5: 458,
		19.0: 504,
		19.5: 554,
		20.0: 612,
		20.5: 675,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

// Carry Capacity (no encumbrance)
// ----------------------------------------------------------------------------
// The getter function for the unencumbered Carry Capacity based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the unencumbered Carry Capacity value.
func GetCarryNoEnc(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return GetCarryNoEncStrengthDict()(attrKey)
}

// A closure function that houses the data used for unencumbered Carry Capacity
// based on  Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func GetCarryNoEncStrengthDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: 3,
		1.5: 3,
		2.0: 4,
		2.5: 5,
		3.0: 5,
		3.5: 6,
		4.0: 6,
		4.5: 7,
		5.0: 7,
		5.5: 8,
		6.0: 8,
		6.5: 9,
		7.0: 9,
		7.5: 10,
		8.0: 10,
		8.5: 10,
		9.0: 11,
		9.5: 11,
		10.0: 11,
		10.5: 11,
		11.0: 12,
		11.5: 13,
		12.0: 14,
		12.5: 15,
		13.0: 17,
		13.5: 18,
		14.0: 19,
		14.5: 21,
		15.0: 23,
		15.5: 25,
		16.0: 27,
		16.5: 30,
		17.0: 32,
		17.5: 36,
		18.0: 39,
		18.5: 43,
		19.0: 47,
		19.5: 52,
		20.0: 58,
		20.5: 64,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}


// Carry Capacity (light encumbrance)
// ----------------------------------------------------------------------------
// The getter function for the lt. encumbered Carry Capacity based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the light encumbrance Carry Capacity value.
func GetCarryLtEnc(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return GetCarryLtEncStrengthDict()(attrKey)
}

// A closure function that houses the data used for lt. encumbered Carry Capacity
// based on  Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func GetCarryLtEncStrengthDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: 5,
		1.5: 6,
		2.0: 8,
		2.5: 9,
		3.0: 10,
		3.5: 11,
		4.0: 12,
		4.5: 13,
		5.0: 15,
		5.5: 16,
		6.0: 16,
		6.5: 17,
		7.0: 18,
		7.5: 19,
		8.0: 20,
		8.5: 20,
		9.0: 21,
		9.5: 22,
		10.0: 22,
		10.5: 23,
		11.0: 24,
		11.5: 26,
		12.0: 28,
		12.5: 31,
		13.0: 33,
		13.5: 36,
		14.0: 39,
		14.5: 42,
		15.0: 46,
		15.5: 50,
		16.0: 54,
		16.5: 59,
		17.0: 65,
		17.5: 71,
		18.0: 78,
		18.5: 86,
		19.0: 95,
		19.5: 105,
		20.0: 116,
		20.5: 128,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}


// Carry Capacity (medium encumbrance)
// ----------------------------------------------------------------------------
// The getter function for the md. encumbered Carry Capacity based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the medium encumbrance Carry Capacity value.
func GetCarryMdEnc(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return GetCarryMdEncStrengthDict()(attrKey)
}

// A closure function that houses the data used for md. encumbered Carry Capacity
// based on  Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func GetCarryMdEncStrengthDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: 10,
		1.5: 13,
		2.0: 16,
		2.5: 18,
		3.0: 20,
		3.5: 22,
		4.0: 24,
		4.5: 26,
		5.0: 29,
		5.5: 31,
		6.0: 32,
		6.5: 34,
		7.0: 36,
		7.5: 38,
		8.0: 39,
		8.5: 40,
		9.0: 42,
		9.5: 43,
		10.0: 44,
		10.5: 45,
		11.0: 48,
		11.5: 52,
		12.0: 56,
		12.5: 61,
		13.0: 66,
		13.5: 71,
		14.0: 77,
		14.5: 84,
		15.0: 91,
		15.5: 99,
		16.0: 108,
		16.5: 118,
		17.0: 129,
		17.5: 142,
		18.0: 156,
		18.5: 171,
		19.0: 189,
		19.5: 209,
		20.0: 231,
		20.5: 256,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}


// Carry Capacity (heavy encumbrance)
// ----------------------------------------------------------------------------
// The getter function for the md. encumbered Carry Capacity based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the heavy encumbrance Carry Capacity value.
func GetCarryHvEnc(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return GetCarryHvEncStrengthDict()(attrKey)
}

// A closure function that houses the data used for hv. encumbered Carry Capacity
// based on  Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func GetCarryHvEncStrengthDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: 15,
		1.5: 20,
		2.0: 24,
		2.5: 27,
		3.0: 30,
		3.5: 33,
		4.0: 36,
		4.5: 39,
		5.0: 44,
		5.5: 47,
		6.0: 48,
		6.5: 51,
		7.0: 54,
		7.5: 57,
		8.0: 59,
		8.5: 60,
		9.0: 63,
		9.5: 65,
		10.0: 66,
		10.5: 68,
		11.0: 72,
		11.5: 78,
		12.0: 84,
		12.5: 92,
		13.0: 99,
		13.5: 107,
		14.0: 116,
		14.5: 126,
		15.0: 137,
		15.5: 149,
		16.0: 162,
		16.5: 177,
		17.0: 194,
		17.5: 213,
		18.0: 234,
		18.5: 257,
		19.0: 284,
		19.5: 314,
		20.0: 347,
		20.5: 384,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}


// Drag
// ----------------------------------------------------------------------------
// The getter function for the Drag capacity based on Strength
// Takes the ability score and the fractional ability score as arguments and
// uses them to call AttributeLookupKeyConstructor() before doing the lookup in
// DamageModDict() using the constructed key.
// Returns the Drag capacity value.
func GetStrDrag(val int, frac int) int {
	attrKey := AttributeLookupKeyConstructor(val, frac)
	return GetDragStrengthDict()(attrKey)
}

// A closure function that houses the data used for Drag Capacity
// based on  Strength. Takes a float value, typically constructed with
// AttributeLookupKeyConstructor(), as an argument to do the lookup.
func GetDragStrengthDict() func(float32) int {
	// innerMap is captured in the closure returned below
	innerMap := map[float32]int {
		1.0: 80,
		1.5: 105,
		2.0: 130,
		2.5: 145,
		3.0: 160,
		3.5: 190,
		4.0: 220,
		4.5: 248,
		5.0: 275,
		5.5: 300,
		6.0: 325,
		6.5: 350,
		7.0: 373,
		7.5: 393,
		8.0: 415,
		8.5: 433,
		9.0: 453,
		9.5: 468,
		10.0: 485,
		10.5: 500,
		11.0: 513,
		11.5: 525,
		12.0: 538,
		12.5: 550,
		13.0: 563,
		13.5: 575,
		14.0: 588,
		14.5: 600,
		15.0: 613,
		15.5: 668,
		16.0: 728,
		16.5: 795,
		17.0: 868,
		17.5: 950,
		18.0: 1043,
		18.5: 1145,
		19.0: 1260,
		19.5: 1385,
		20.0: 1530,
		20.5: 1688,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

// I'm not sure what to do with encumbrance. I'll put it here for now.
type Encumbrance struct {
	DamageReduction 	int
	DefenseAdjustment 	int
	InitiativeModifier	int
	SpeedModifier 		int
	FatigueModifier 	int
	SprintSpeed 		float32
	RunSpeed 			float32
	JogSpeed 			float32
	WalkSpeed 			float32
	CrawlSpeed 			float32
}

func GetEncumbranceLevel(gearWeight int, strVal int, strFrac int) int {
	noEnc := GetCarryNoEnc(strVal, strFrac)
	ltEnc := GetCarryLtEnc(strVal, strFrac)
	mdEnc := GetCarryMdEnc(strVal, strFrac)
	hvEnc := GetCarryHvEnc(strVal, strFrac)

	fmt.Println(noEnc)
	fmt.Println(ltEnc)
	fmt.Println(mdEnc)
	fmt.Println(hvEnc)

	var encumbranceLevel int

	if gearWeight <= noEnc {
		encumbranceLevel = 0
	} else if gearWeight <= ltEnc {
		encumbranceLevel = 1
	} else if gearWeight <= mdEnc {
		encumbranceLevel = 2
	} else if gearWeight > mdEnc {
		encumbranceLevel = 3
	}

	return encumbranceLevel
}

//func GetEncumbranceDamageReduction(encLevel int) int {
//
//}