package constants

var KUUDRA_COMPLETIONS_MULTIPLIER = map[string]int{
	"none":     1,
	"hot":      2,
	"burning":  3,
	"fiery":    4,
	"infernal": 5,
}

type kuudraTier struct {
	Name    string
	Texture string
}

var KUUDRA_TIERS = map[string]kuudraTier{
	"none": {
		Name:    "Basic",
		Texture: "/api/head/bfd3e71838c0e76f890213120b4ce7449577736604338a8d28b4c86db2547e71",
	},
	"hot": {
		Name:    "Hot",
		Texture: "/api/head/c0259e8964c3deb95b1233bb2dc82c986177e63ae36c11265cb385180bb91cc0",
	},
	"burning": {
		Name:    "Burning",
		Texture: "/api/head/330f6f6e63b245f839e3ccdce5a5f22056201d0274411dfe5d94bbe449c4ece",
	},
	"fiery": {
		Name:    "Fiery",
		Texture: "/api/head/bd854393bbf9444542502582d4b5a23cc73896506e2fc739d545bc35bc7b1c06",
	},
	"infernal": {
		Name:    "Infernal",
		Texture: "/api/head/82ee25414aa7efb4a2b4901c6e33e5eaa705a6ab212ebebfd6a4de984125c7a0",
	},
}

type dojoChallenge struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
}

var DOJO = map[string]dojoChallenge{
	"mob_kb": {
		Name:    "Force",
		Texture: "/api/item/STICK",
	},
	"wall_jump": {
		Name:    "Stamina",
		Texture: "/api/item/RABBIT_FOOT",
	},
	"archer": {
		Name:    "Mastery",
		Texture: "/api/item/BOW",
	},
	"sword_swap": {
		Name:    "Discipline",
		Texture: "/api/item/DIAMOND_SWORD",
	},
	"snake": {
		Name:    "Swiftness",
		Texture: "/api/item/LEAD",
	},
	"lock_head": {
		Name:    "Control",
		Texture: "/api/item/ENDER_EYE",
	},
	"fireball": {
		Name:    "Tenacity",
		Texture: "/api/item/FIRE_CHARGE",
	},
}
