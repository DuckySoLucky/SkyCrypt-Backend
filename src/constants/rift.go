package constants

type riftEye struct {
	Name    string
	Id      string
	Texture string
}

var RIFT_EYES = []riftEye{
	{Name: "The Intruder", Id: "dreadfarm", Texture: "/api/head/17db1923d03c4ef4e9f6e872c5a6ad2578b1aff2b281fbc3ffa7466c825fb9"},
	{Name: "The Gill-Man", Id: "wizard_tower", Texture: "/api/head/17db1923d03c4ef4e9f6e872c5a6ad2578b1aff2b281fbc3ffa7466c825fb9"},
	{Name: "The Baba Yaga", Id: "plaza", Texture: "/api/head/17db1923d03c4ef4e9f6e872c5a6ad2578b1aff2b281fbc3ffa7466c825fb9"},
	{Name: "The Bankster", Id: "fisherman_hut", Texture: "/api/head/17db1923d03c4ef4e9f6e872c5a6ad2578b1aff2b281fbc3ffa7466c825fb9"},
	{Name: "The Gooey", Id: "colosseum", Texture: "/api/head/17db1923d03c4ef4e9f6e872c5a6ad2578b1aff2b281fbc3ffa7466c825fb9"},
	{Name: "The Prince", Id: "castle", Texture: "/api/head/17db1923d03c4ef4e9f6e872c5a6ad2578b1aff2b281fbc3ffa7466c825fb9"},
	{Name: "The 7th Sin", Id: "mountaintop", Texture: "/api/head/17db1923d03c4ef4e9f6e872c5a6ad2578b1aff2b281fbc3ffa7466c825fb9"},
}

type riftTimecharm struct {
	Name    string
	ID      string
	Texture string
}

var RIFT_TIMECHARMS = []riftTimecharm{
	{Name: "Supreme Timecharm", ID: "wyldly_supreme", Texture: "/api/item/LEAVES:1"},
	{Name: "mrahcemiT esrevrorriM", ID: "mirrored", Texture: "/api/item/GLASS"},
	{Name: "Chicken N Egg Timecharm", ID: "chicken_n_egg", Texture: "/api/item/SOUL_SAND"},
	{Name: "SkyBlock Citizen Timecharm", ID: "citizen", Texture: "/api/item/JUKEBOX"},
	{Name: "Living Timecharm", ID: "lazy_living", Texture: "/api/item/LAPIS_ORE"},
	{Name: "Globulate Timecharm", ID: "slime", Texture: "/api/item/SLIME_BLOCK"},
	{Name: "Vampiric Timecharm", ID: "vampiric", Texture: "/api/item/REDSTONE_BLOCK"},
	{Name: "Celestial Timecharm", ID: "mountain", Texture: "/api/item/LAPIS_BLOCK"},
}

const RIFT_ENIGMA_SOULS = 52
const RIFT_MAX_GRUBBER_STACKS = 5
