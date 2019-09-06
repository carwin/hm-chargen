package attributes

type Charisma struct {
	Attribute
	BpBonus				int
	StartingHonorMod	int
	TurningMod			int
	MoraleMod			int
	MaxProtoges			int
}

func ChaBpBonusDict() func(int) int {
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
		11: 0,
		12: 1,
		13: 3,
		14: 6,
		15: 10,
		16: 15,
		17: 21,
		18: 28,
		19: 36,
		20: 45,
		21: 55,
		22: 66,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetChaBpBonus(val int) int {
	return ChaBpBonusDict()(val)
}

func ChaStartingHonorModDict() func(int) int {
	innerMap := map[int] int {
		1: -6,
		2: -5,
		3: -4,
		4: -3,
		5: -3,
		6: -2,
		7: -2,
		8: -1,
		9: -1,
		10: 0,
		11: 0,
		12: 1,
		13: 1,
		14: 2,
		15: 2,
		16: 3,
		17: 3,
		18: 4,
		19: 4,
		20: 5,
		21: 5,
		22: 6,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetChaStartingHonorMod(val int) int {
	return ChaStartingHonorModDict()(val)
}

func ChaTurningModDict() func(int) int {
	innerMap := map[int] int {
		1: -9,
		2: -8,
		3: -7,
		4: -6,
		5: -5,
		6: -4,
		7: -3,
		8: -2,
		9: -1,
		10: 0,
		11: 1,
		12: 2,
		13: 3,
		14: 4,
		15: 5,
		16: 6,
		17: 7,
		18: 8,
		19: 9,
		20: 10,
		21: 11,
		22: 12,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetChaTurningMod(val int) int {
	return ChaTurningModDict()(val)
}

func ChaMoraleModDict() func(int) int {
	innerMap := map[int] int {
		1: -5,
		2: -4,
		3: -4,
		4: -3,
		5: -3,
		6: -2,
		7: -2,
		8: -1,
		9: -1,
		10: 0,
		11: 1,
		12: 1,
		13: 2,
		14: 2,
		15: 3,
		16: 3,
		17: 4,
		18: 4,
		19: 5,
		20: 5,
		21: 6,
		22: 6,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetChaMoraleMod(val int) int {
	return ChaMoraleModDict()(val)
}

func ChaMaxProtogesDict() func(int) int {
	innerMap := map[int] int {
		1: 0,
		2: 1,
		3: 1,
		4: 1,
		5: 1,
		6: 2,
		7: 2,
		8: 2,
		9: 2,
		10: 3,
		11: 3,
		12: 3,
		13: 3,
		14: 3,
		15: 3,
		16: 4,
		17: 4,
		18: 4,
		19: 4,
		20: 4,
		21: 5,
		22: 5,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetChaMaxProtoges(val int) int {
	return ChaMaxProtogesDict()(val)
}