package attributes

type Dexterity struct {
	Attribute
	InitMod					int
	AtkMod					int
	DefMod					int
	DodgeSavingThrowMod		int
	FeatOfAgility			int
}

func InitModDexterityDict() func(float32) int {
	innerMap := map[float32] int {
		3.0: 7,
		3.5: 7,
		4.0: 6,
		4.5: 6,
		5.0: 6,
		5.5: 5,
		6.0: 5,
		6.5: 5,
		7.0: 4,
		7.5: 4,
		8.0: 4,
		8.5: 3,
		9.0: 3,
		9.5: 3,
		10.0: 2,
		10.5: 2,
		11.0: 2,
		11.5: 1,
		12.0: 1,
		12.5: 1,
		13.0: 0,
		13.5: 0,
		14.0: 0,
		14.5: -1,
		15.0: -1,
		15.5: -1,
		16.0: -2,
		16.5: -2,
		17.0: -2,
		17.5: -3,
		18.0: -3,
		18.5: -3,
		19.0: -4,
		19.5: -4,
		20.0: -4,
		20.5: -5,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

func GetDexInitMod(val int, frac int) int {
	key := AttributeLookupKeyConstructor(val, frac)
	return InitModDexterityDict()(key)
}

func DexAtkModDict() func(float32) int {
	innerMap := map[float32] int {
		3.0: -4,
		3.5: -3,
		4.0: -3,
		4.5: -3,
		5.0: -3,
		5.5: -2,
		6.0: -2,
		6.5: -2,
		7.0: -2,
		7.5: -1,
		8.0: -1,
		8.5: -1,
		9.0: 0,
		9.5: 0,
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
		15.5: 2,
		16.0: 3,
		16.5: 3,
		17.0: 3,
		17.5: 3,
		18.0: 4,
		18.5: 4,
		19.0: 4,
		19.5: 4,
		20.0: 5,
		20.5: 5,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

func GetDexAtkMod(val int, frac int) int {
	key := AttributeLookupKeyConstructor(val, frac)
	return DexAtkModDict()(key)
}

func DexDefModDict() func(float32) int {
	innerMap := map[float32] int {
		3.0: -5,
		3.5: -5,
		4.0: -4,
		4.5: -4,
		5.0: -4,
		5.5: -3,
		6.0: -3,
		6.5: -3,
		7.0: -2,
		7.5: -2,
		8.0: -2,
		8.5: -1,
		9.0: -1,
		9.5: -1,
		10.0: 0,
		10.5: 0,
		11.0: 0,
		11.5: 1,
		12.0: 1,
		12.5: 1,
		13.0: 2,
		13.5: 2,
		14.0: 2,
		14.5: 3,
		15.0: 3,
		15.5: 3,
		16.0: 4,
		16.5: 4,
		17.0: 4,
		17.5: 5,
		18.0: 5,
		18.5: 5,
		19.0: 6,
		19.5: 6,
		20.0: 6,
		20.5: 7,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

func GetDexDefMod(val int, frac int) int {
	key := AttributeLookupKeyConstructor(val, frac)
	return DexDefModDict()(key)
}

func DexDodgeSavingThrowModDict() func(float32) int {
	innerMap := map[float32] int {
		3.0: -3,
		3.5: -3,
		4.0: -3,
		4.5: -3,
		5.0: -2,
		5.5: -2,
		6.0: -2,
		6.5: -2,
		7.0: -1,
		7.5: -1,
		8.0: -1,
		8.5: -1,
		9.0: 0,
		9.5: 0,
		10.0: 0,
		10.5: 0,
		11.0: 0,
		11.5: 0,
		12.0: 0,
		12.5: 0,
		13.0: 1,
		13.5: 1,
		14.0: 1,
		14.5: 1,
		15.0: 2,
		15.5: 2,
		16.0: 2,
		16.5: 2,
		17.0: 2,
		17.5: 2,
		18.0: 3,
		18.5: 3,
		19.0: 3,
		19.5: 3,
		20.0: 3,
		20.5: 3,
	}

	return func(key float32) int {
		return innerMap[key]
	}
}

func GetDexDodgeSavingThrowMod(val int, frac int) int {
	key := AttributeLookupKeyConstructor(val, frac)
	return DexDodgeSavingThrowModDict()(key)
}

func DexFeatOfAgilityDict() func(float32) int {
	innerMap := map[float32] int {
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

func GetDexFeatOfAgility(val int, frac int) int {
	key := AttributeLookupKeyConstructor(val, frac)
	return DexFeatOfAgilityDict()(key)
}
