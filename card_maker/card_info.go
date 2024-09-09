package card_maker

import (
	"encoding/json"
	"log"
)

type CardInfo struct {
	Number string `json:"number"`

	Type          string   `json:"type"`
	Name          string   `json:"name"`
	Category      string   `json:"category"`
	Tag           string   `json:"tag"`
	Description   string   `json:"description"`
	Quote         string   `json:"quote"`
	ElementsCost  Elements `json:"elements_cost"`
	ElementsGain  Elements `json:"elements_gain"`
	VersionNumber string   `json:"version_number"`
	VersionName   string   `json:"version_name"`

	// unit cards
	Attack int `json:"attack"`
	Life   int `json:"life"`

	// ability
	Duration        int      `json:"duration"`
	Power           int      `json:"power"`
	ElementsExpense Elements `json:"elements_expense"`

	// spawns
	Spawns []string `json:"spawns"`

	OutputPath string `json:"output_path"`
}

// make it comparable
func (cardInfo CardInfo) Equals(other CardInfo) bool {
	cond1 := cardInfo.Number == other.Number &&
		cardInfo.Type == other.Type &&
		cardInfo.Name == other.Name &&
		cardInfo.Category == other.Category &&
		cardInfo.Tag == other.Tag &&
		cardInfo.Description == other.Description &&
		cardInfo.Quote == other.Quote &&
		cardInfo.ElementsCost.Equals(other.ElementsCost) &&
		cardInfo.ElementsGain.Equals(other.ElementsGain) &&
		cardInfo.VersionNumber == other.VersionNumber &&
		cardInfo.VersionName == other.VersionName &&
		cardInfo.Attack == other.Attack &&
		cardInfo.Life == other.Life &&
		cardInfo.Duration == other.Duration &&
		cardInfo.Power == other.Power &&
		cardInfo.ElementsExpense.Equals(other.ElementsExpense) &&
		cardInfo.OutputPath == other.OutputPath
	if !cond1 {
		return false
	}

	// compare spawns
	if len(cardInfo.Spawns) != len(other.Spawns) {
		return false
	}
	for i, spawn := range cardInfo.Spawns {
		if spawn != other.Spawns[i] {
			return false
		}
	}
	return cond1
}

func (cardInfo CardInfo) String() string {
	jsonByte, err := json.MarshalIndent(cardInfo, "", "  ")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		return string(jsonByte)
	}
}
