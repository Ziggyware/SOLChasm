package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type AbilityTypeEnum int

const (
	Damage AbilityTypeEnum = iota
	Heal
)

func (v AbilityTypeEnum) String() string {
	return [...]string{"Damage", "Heal"}[v]
}

type RarityTypeEnum int

const (
	Common RarityTypeEnum = iota
	Uncommon
	Rare
	Epic
	Legendary
)

func (v RarityTypeEnum) String() string {
	return [...]string{"Common", "Uncommon", "Rare", "Epic", "Legendary"}[v]
}

type Ability struct {
	Name       string
	Rarity     RarityTypeEnum
	FlavorText string
	Type       AbilityTypeEnum
	Desc       string
}

type CharacterCard struct {
	Name       string
	Rarity     RarityTypeEnum
	FlavorText string
	Ability1   Ability
	Ability2   Ability
	Ability3   Ability
}

func loadCards() []CharacterCard {
	return []CharacterCard{
		{
			Name:       "Test1",
			Rarity:     Common,
			FlavorText: "test flavor",
			Ability1: Ability{
				Name:       "Fireball",
				Rarity:     Common,
				FlavorText: "fire from hell",
				Type:       Damage,
				Desc:       "Deal 1 to 10 damage",
			},
		},
	}
}

func LoadPage(page string) string {
	bt, _ := os.ReadFile(page)
	txt := string(bt)

	bt, _ = os.ReadFile("./cards/characterCardTemplate.html")
	cardTxt := string(bt)

	card1 := loadCards()[0]
	ret := strings.Replace(cardTxt, "{{CardName}}", card1.Name, -1)
	ret = strings.Replace(ret, "{{Rarity}}", fmt.Sprint(card1.Rarity), -1)
	ret = strings.Replace(ret, "{{FlavorText}}", card1.FlavorText, -1)

	txt = strings.Replace(txt, "{{card}}", ret, -1)
	return ret
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(LoadPage("index.html")))
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
