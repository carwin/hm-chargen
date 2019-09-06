package attributes

type Wisdom struct {
	Attribute
	InitMod					int
	BpBonus					int
	DefMod					int
	MentalSavingThrowMod	int
}

func InitModWisdomDict() func(int) int {
	innerMap := map[int] int {
		1: 7,
		2: 6,
		3: 5,
		4: 4,
		5: 4,
		6: 4,
		7: 3,
		8: 3,
		9: 3,
		10: 2,
		11: 2,
		12: 1,
		13: 1,
		14: 1,
		15: 0,
		16: 0,
		17: 0,
		18: -1,
		19: -1,
		20: -1,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetWisInitMod(val int) int {
	return InitModWisdomDict()(val)
}

func BpWisdomBonusDict() func(int) int {
	innerMap := map[int] int {
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
		10: 0,
		11: 1,
		12: 2,
		13: 3,
		14: 6,
		15: 10,
		16: 15,
		17: 21,
		18: 28,
		19: 36,
		20: 45,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetWisBpBonus(val int) int {
	return BpWisdomBonusDict()(val)
}

func DefModWisdomDict() func(int) int {
	innerMap := map[int] int {
		1: -4,
		2: -3,
		3: -3,
		4: -2,
		5: -2,
		6: -2,
		7: -1,
		8: -1,
		9: -1,
		10: 0,
		11: 0,
		12: 1,
		13: 1,
		14: 1,
		15: 2,
		16: 2,
		17: 2,
		18: 3,
		19: 3,
		20: 3,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetWisdomDefMod(val int) int {
	return DefModWisdomDict()(val)
}

func MentalSavingThrowWisdomModDict() func(int) int {
	innerMap := map[int] int {
		1: -4,
		2: -4,
		3: -3,
		4: -3,
		5: -2,
		6: -2,
		7: -1,
		8: -1,
		9: 0,
		10: 0,
		11: 0,
		12: 0,
		13: 1,
		14: 1,
		15: 2,
		16: 2,
		17: 2,
		18: 3,
		19: 3,
		20: 3,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetWisMentalSavingThrowModifier(val int) int {
	return MentalSavingThrowWisdomModDict()(val)
}
