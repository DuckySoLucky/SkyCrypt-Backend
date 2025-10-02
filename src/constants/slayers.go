package constants

var SLAYER_XP = map[string]map[int]int{
	"zombie": {
		1: 5,
		2: 15,
		3: 200,
		4: 1000,
		5: 5000,
		6: 20000,
		7: 100000,
		8: 400000,
		9: 1000000,
	},
	"spider": {
		1: 5,
		2: 25,
		3: 200,
		4: 1000,
		5: 5000,
		6: 20000,
		7: 100000,
		8: 400000,
		9: 1000000,
	},
	"wolf": {
		1: 10,
		2: 30,
		3: 250,
		4: 1500,
		5: 5000,
		6: 20000,
		7: 100000,
		8: 400000,
		9: 1000000,
	},
	"enderman": {
		1: 10,
		2: 30,
		3: 250,
		4: 1500,
		5: 5000,
		6: 20000,
		7: 100000,
		8: 400000,
		9: 1000000,
	},
	"blaze": {
		1: 10,
		2: 30,
		3: 250,
		4: 1500,
		5: 5000,
		6: 20000,
		7: 100000,
		8: 400000,
		9: 1000000,
	},
	"vampire": {
		1: 20,
		2: 75,
		3: 240,
		4: 840,
		5: 2400,
	},
}

type slayerInfo struct {
	Name      string
	Head      string
	Levelling map[int]int
}

var SLAYER_INFO = map[string]slayerInfo{
	"zombie": {
		Name:      "Revenant Horror",
		Head:      "/api/head/1fc0184473fe882d2895ce7cbc8197bd40ff70bf10d3745de97b6c2a9c5fc78f",
		Levelling: SLAYER_XP["zombie"],
	},
	"spider": {
		Name:      "Tarantula Broodfather",
		Head:      "/api/head/9d7e3b19ac4f3dee9c5677c135333b9d35a7f568b63d1ef4ada4b068b5a25",
		Levelling: SLAYER_XP["spider"],
	},
	"wolf": {
		Name:      "Sven Packmaster",
		Head:      "/api/head/f83a2aa9d3734b919ac24c9659e5e0f86ecafbf64d4788cfa433bbec189e8",
		Levelling: SLAYER_XP["wolf"],
	},
	"enderman": {
		Name:      "Voidgloom Seraph",
		Head:      "/api/head/1b09a3752510e914b0bdc9096b392bb359f7a8e8a9566a02e7f66faff8d6f89e",
		Levelling: SLAYER_XP["enderman"],
	},
	"blaze": {
		Name:      "Inferno Demonlord",
		Head:      "/api/head/b20657e24b56e1b2f8fc219da1de788c0c24f36388b1a409d0cd2d8dba44aa3b",
		Levelling: SLAYER_XP["blaze"],
	},
	"vampire": {
		Name:      "Riftstalker Bloodfiend",
		Head:      "/api/head/5aa29ea961757dc3c90bfabf302c5abe9d308fb4a7d3864e5788ad2cc9160aa2",
		Levelling: SLAYER_XP["vampire"],
	},
}
