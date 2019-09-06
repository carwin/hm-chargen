package attributes

type Constitution struct {
	Attribute
	PhysicalSavingThrowMod	int
}

func ConPhysicalSavingThrowModDict() func(int) int {
	innerMap := map[int] int {
		1: -5,
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
		21: 4,
		22: 4,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetConPhysicalSavingThrowMod(val int) int {
	return ConPhysicalSavingThrowModDict()(val)
}
