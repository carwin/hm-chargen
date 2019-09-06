package attributes

type Looks struct {
	Attribute
	CharismaMod			int
	StartingHonorMod	int
	StartingFameMod		int
}

func LksCharismaModDict() func(int) int {
	innerMap := map[int] int {
		1: -6,
		2: -5,
		3: -5,
		4: -4,
		5: -3,
		6: -2,
		7: -2,
		8: -1,
		9: -1,
		10: 0,
		11: 0,
		12: 0,
		13: 1,
		14: 1,
		15: 2,
		16: 2,
		17: 3,
		18: 4,
		19: 5,
		20: 6,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetLksCharismaMod(val int) int {
	return LksCharismaModDict()(val)
}

func LksStartingHonorModDict() func(int) int {
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
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetLksStartingHonorMod(val int) int {
	return LksStartingHonorModDict()(val)
}

func LksStartingFameModDict() func(int) int {
	innerMap := map[int] int {
		1: 9,
		2: 6,
		3: 4,
		4: 2,
		5: 1,
		6: 1,
		7: 0,
		8: 0,
		9: 0,
		10: 0,
		11: 0,
		12: 0,
		13: 0,
		14: 1,
		15: 2,
		16: 3,
		17: 5,
		18: 7,
		19: 8,
		20: 9,
	}

	return func(key int) int {
		return innerMap[key]
	}
}

func GetLksStartingFameMod(val int) int {
	return LksStartingFameModDict()(val)
}
