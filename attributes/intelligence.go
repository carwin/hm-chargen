package attributes

type Intelligence struct {
	Attribute
	AtkMod		int
	BpBonus		int
}

func AtkModIntelligenceDict() func(int) int {
	innerMap := map[int] int {
		1: -5,
		2: -4,
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

func GetIntDmgMod(val int) int {
	return AtkModIntelligenceDict()(val)
}

func BpIntelligenceBonusDict() func(int) int {
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

func GetIntBpBonus(val int) int {
	return BpIntelligenceBonusDict()(val)
}