package main

import (
	"fmt"
	"hm-chargen/attributes"
	"hm-chargen/dice"
)

type AbilityScores struct {
	Strength 		attributes.Strength
	Intelligence	attributes.Intelligence
	Wisdom 			attributes.Wisdom
	Dexterity		attributes.Dexterity
	Constitution	attributes.Constitution
	Looks 			attributes.Looks
	Charisma		attributes.Charisma
}

// Constructor function for ability scores.
func NewAbilityScores(str attributes.Strength, intel attributes.Intelligence, wis attributes.Wisdom, dex attributes.Dexterity, con attributes.Constitution, lks attributes.Looks, cha attributes.Charisma) *AbilityScores {
	as := AbilityScores{Strength: str, Intelligence: intel, Wisdom: wis, Dexterity: dex, Constitution: con, Looks: lks, Charisma: cha}

	// Return memory value I guess.
	return &as
}

// Step 1: Starting BP
// -----------------------------------------------------------------------------
var buildingPoints int = 40

// Step 2: Roll Ability Scores
// -----------------------------------------------------------------------------
func rollAbilityScores() AbilityScores {

	// Roll and prepare Strength
	str := new(attributes.Strength)
	str.Attribute.Value = dice.RollDice(3, 6)
	str.Attribute.Fractional = dice.RollDice(1, 99)
	str.DmgMod = attributes.GetStrDamageMod(str.Attribute.Value, str.Attribute.Fractional)
	str.FeatofStr = attributes.GetFeatOfStrength(str.Attribute.Value, str.Attribute.Fractional)
	str.Lift = attributes.GetStrLift(str.Attribute.Value, str.Attribute.Fractional)
	str.CarryNoEnc = attributes.GetCarryNoEnc(str.Attribute.Value, str.Attribute.Fractional)
	str.CarryLtEnc = attributes.GetCarryLtEnc(str.Attribute.Value, str.Attribute.Fractional)
	str.CarryMdEnc = attributes.GetCarryMdEnc(str.Attribute.Value, str.Attribute.Fractional)
	str.Drag = attributes.GetStrDrag(str.Attribute.Value, str.Attribute.Fractional)

	// Roll and prepare Intelligence
	intel := new(attributes.Intelligence)
	intel.Attribute.Value = dice.RollDice(3, 6)
	intel.Attribute.Fractional = dice.RollDice(1, 99)
	intel.AtkMod = attributes.GetIntDmgMod(intel.Attribute.Value)
	intel.BpBonus = attributes.GetIntBpBonus(intel.Attribute.Value)

	// Roll and prepare Wisdom
	wis := new(attributes.Wisdom)
	wis.Attribute.Value = dice.RollDice(3, 6)
	wis.Attribute.Fractional = dice.RollDice(1, 99)
	wis.InitMod = attributes.GetWisInitMod(wis.Attribute.Value)
	wis.BpBonus = attributes.GetWisBpBonus(wis.Attribute.Value)
	wis.DefMod = attributes.GetWisdomDefMod(wis.Attribute.Value)
	wis.MentalSavingThrowMod = attributes.GetWisMentalSavingThrowModifier(wis.Attribute.Value)

	// Roll and prepare Dexterity
	dex := new(attributes.Dexterity)
	dex.Attribute.Value = dice.RollDice(3, 6)
	dex.Attribute.Fractional = dice.RollDice(1, 99)
	dex.InitMod = attributes.GetDexInitMod(dex.Attribute.Value, dex.Attribute.Fractional)
	dex.AtkMod = attributes.GetDexAtkMod(dex.Attribute.Value, dex.Attribute.Fractional)
	dex.DefMod = attributes.GetDexDefMod(dex.Attribute.Value, dex.Attribute.Fractional)
	dex.DodgeSavingThrowMod = attributes.GetDexDodgeSavingThrowMod(dex.Attribute.Value, dex.Attribute.Fractional)
	dex.FeatOfAgility = attributes.GetDexFeatOfAgility(dex.Attribute.Value, dex.Attribute.Fractional)

	// Roll and prepare Constitution
	con := new(attributes.Constitution)
	con.Attribute.Value = dice.RollDice(3, 6)
	con.Attribute.Fractional = dice.RollDice(1, 99)
	con.PhysicalSavingThrowMod = attributes.GetConPhysicalSavingThrowMod(con.Attribute.Value)

	// Roll and prepare Looks
	lks := new(attributes.Looks)
	lks.Attribute.Value = dice.RollDice(3, 6)
	lks.Attribute.Fractional = dice.RollDice(1, 99)
	lks.CharismaMod = attributes.GetLksCharismaMod(lks.Attribute.Value)
	lks.StartingHonorMod = attributes.GetLksStartingHonorMod(lks.Attribute.Value)
	lks.StartingFameMod = attributes.GetLksStartingFameMod(lks.Attribute.Value)

	// Roll and prepare Charisma
	cha := new(attributes.Charisma)
	cha.Attribute.Value = dice.RollDice(3, 6)
	cha.Attribute.Fractional = dice.RollDice(1, 99)
	cha.BpBonus = attributes.GetChaBpBonus(cha.Attribute.Value)
	cha.StartingHonorMod = attributes.GetChaStartingHonorMod(cha.Attribute.Value)
	cha.TurningMod = attributes.GetChaTurningMod(cha.Attribute.Value)
	cha.MoraleMod = attributes.GetChaMoraleMod(cha.Attribute.Value)
	cha.MaxProtoges = attributes.GetChaMaxProtoges(cha.Attribute.Value)

	// Assign strength to ability scores
	//abilityScores := new(AbilityScores)
	//abilityScores.Strength = *str

	// Initialize abilityScores
	abilityScores := NewAbilityScores(*str, *intel, *wis, *dex, *con, *lks, *cha)

	// Add BP from Int. bonus.
	buildingPoints += abilityScores.Intelligence.BpBonus
	// Add BP from Wis. bonus.
	buildingPoints += abilityScores.Wisdom.BpBonus
	// Add BP from Cha. bonus.
	buildingPoints += abilityScores.Charisma.BpBonus

	// Return a pointer to the new ability scores
	return *abilityScores

}

func main() {
	// Roll Ability Scores
	abilityScores := rollAbilityScores()


	// Log out basic stats
	fmt.Printf("Strength: %d/%d\n", abilityScores.Strength.Attribute.Value, abilityScores.Strength.Attribute.Fractional)
	fmt.Printf("Intelligence: %d/%d\n", abilityScores.Intelligence.Attribute.Value, abilityScores.Intelligence.Attribute.Fractional)
	fmt.Printf("Wisdom: %d/%d\n", abilityScores.Wisdom.Attribute.Value, abilityScores.Wisdom.Attribute.Fractional)
	fmt.Printf("Dexterity: %d/%d\n", abilityScores.Dexterity.Attribute.Value, abilityScores.Dexterity.Attribute.Fractional)
	fmt.Printf("Constitution: %d/%d\n", abilityScores.Constitution.Attribute.Value, abilityScores.Constitution.Attribute.Fractional)
	fmt.Printf("Looks: %d/%d\n", abilityScores.Looks.Attribute.Value, abilityScores.Looks.Attribute.Fractional)
	fmt.Printf("Charisma: %d/%d\n", abilityScores.Charisma.Attribute.Value, abilityScores.Charisma.Attribute.Fractional)
	fmt.Printf("Building Points: %d\n", buildingPoints)

}
