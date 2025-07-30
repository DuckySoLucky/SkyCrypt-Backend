package constants

var MINION_TYPES = []string{"farming", "mining", "combat", "foraging", "fishing"}

var MINION_SLOTS = map[int]int{
	0:   5,
	5:   6,
	15:  7,
	30:  8,
	50:  9,
	75:  10,
	100: 11,
	125: 12,
	150: 13,
	175: 14,
	200: 15,
	225: 16,
	250: 17,
	275: 18,
	300: 19,
	350: 20,
	400: 21,
	450: 22,
	500: 23,
	550: 24,
	600: 25,
	650: 26,
	700: 27,
}

type minionInfo struct {
	Texture string
	MaxTier int
	Name    string
}

var MINIONS = map[string]map[string]minionInfo{
	"farming": {
		"COCOA":        {Texture: "http://localhost:8080/api/head/acb680e96f6177cd8ffaf27e9625d8b544d720afc50738801818d0e745c0e5f7", MaxTier: 12},
		"PUMPKIN":      {Texture: "http://localhost:8080/api/head/f3fb663e843a7da787e290f23c8af2f97f7b6f572fa59a0d4d02186db6eaabb7", MaxTier: 12},
		"CHICKEN":      {Texture: "http://localhost:8080/api/head/a04b7da13b0a97839846aa5648f5ac6736ba0ca9fbf38cd366916e417153fd7f", MaxTier: 12},
		"MUSHROOM":     {Texture: "http://localhost:8080/api/head/4a3b58341d196a9841ef1526b367209cbc9f96767c24f5f587cf413d42b74a93", MaxTier: 12},
		"CACTUS":       {Texture: "http://localhost:8080/api/head/ef93ec6e67a6cd272c9a9684b67df62584cb084a265eee3cde141d20e70d7d72", MaxTier: 12},
		"PIG":          {Texture: "http://localhost:8080/api/head/a9bb5f0c56408c73cfa412345c8fc51f75b6c7311ae60e7099c4781c48760562", MaxTier: 12},
		"WHEAT":        {Texture: "http://localhost:8080/api/head/bbc571c5527336352e2fee2b40a9edfa2e809f64230779aa01253c6aa535881b", MaxTier: 12},
		"COW":          {Texture: "http://localhost:8080/api/head/c2fd8976e1b64aebfd38afbe62aa1429914253df3417ace1f589e5cf45fbd717", MaxTier: 12},
		"RABBIT":       {Texture: "http://localhost:8080/api/head/ef59c052d339bb6305cad370fd8c52f58269a957dfaf433a255597d95e68a373", MaxTier: 12},
		"SUGAR_CANE":   {Texture: "http://localhost:8080/api/head/2fced0e80f0d7a5d1f45a1a7217e6a99ea9720156c63f6efc84916d4837fabde", MaxTier: 12},
		"MELON":        {Texture: "http://localhost:8080/api/head/95d54539ac8d3fba9696c91f4dcc7f15c320ab86029d5c92f12359abd4df811e", MaxTier: 12},
		"NETHER_WARTS": {Texture: "http://localhost:8080/api/head/71a4620bb3459c1c2fa74b210b1c07b4a02254351f75173e643a0e009a63f558", MaxTier: 12},
		"CARROT":       {Texture: "http://localhost:8080/api/head/4baea990b45d330998cb0c1f8515c27b24f93bff1df0db056e647f8200d03b9d", MaxTier: 12},
		"POTATO":       {Texture: "http://localhost:8080/api/head/7dda35a044cb0374b516015d991a0f65bf7d0fb6566e350496642cf2059ff1d9", MaxTier: 12},
		"SHEEP":        {Texture: "http://localhost:8080/api/head/fd15d4b8bce708f77f963f1b4e87b1b969fef1766a3e9b67b249c59d5e80e8c5", MaxTier: 12},
	},
	"mining": {
		"HARD_STONE":  {Texture: "http://localhost:8080/api/head/1e8bab9493708beda34255606d5883b8762746bcbe6c94e8ca78a77a408c8ba8", MaxTier: 12},
		"RED_SAND":    {Texture: "http://localhost:8080/api/head/9d24991435e4e7fb1a9ad23db75c80aec300d003ec0c5963e0ed658634027889", MaxTier: 12},
		"MYCELIUM":    {Texture: "http://localhost:8080/api/head/fc8ebad72b77df3990e07bc869a99a8f8962d3c19c76e39d99553cae4131cc8", MaxTier: 12},
		"COBBLESTONE": {Texture: "http://localhost:8080/api/head/2f93289a82bd2a06cbbe61b733cfdc1f1bd93c4340f7a90abd9bdda774109071", MaxTier: 12},
		"OBSIDIAN":    {Texture: "http://localhost:8080/api/head/320c29ab966637cb9aecc34ee76d5a0130461e0c4fdb08cdaf80939fa1209102", MaxTier: 12},
		"GLOWSTONE":   {Texture: "http://localhost:8080/api/head/20f4d7c26b0310990a7d3a3b45948b95dd4ab407a16a4b6d3b7cb4fba031aeed", MaxTier: 12},
		"GRAVEL":      {Texture: "http://localhost:8080/api/head/7458507ed31cf9a38986ac8795173c609637f03da653f30483a721d3fbe602d"},
		"SAND":        {Texture: "http://localhost:8080/api/head/81f8e2ad021eefd1217e650e848b57622144d2bf8a39fbd50dab937a7eac10de"},
		"ICE":         {Texture: "http://localhost:8080/api/head/e500064321b12972f8e5750793ec1c823da4627535e9d12feaee78394b86dabe", MaxTier: 12},
		"SNOW":        {Texture: "http://localhost:8080/api/head/f6d180684c3521c9fc89478ba4405ae9ce497da8124fa0da5a0126431c4b78c3", MaxTier: 12},
		"COAL":        {Texture: "http://localhost:8080/api/head/425b8d2ea965c780652d29c26b1572686fd74f6fe6403b5a3800959feb2ad935", MaxTier: 12},
		"IRON":        {Texture: "http://localhost:8080/api/head/af435022cb3809a68db0fccfa8993fc1954dc697a7181494905b03fdda035e4a", MaxTier: 12},
		"GOLD":        {Texture: "http://localhost:8080/api/head/f6da04ed8c810be29bba53c62e712d65cfb25238117b94d7e85a4615775bf14f", MaxTier: 12},
		"DIAMOND":     {Texture: "http://localhost:8080/api/head/2354bbe604dfe58bf92e7729730d0c8e37844e831ee3816d7e8427c27a1824a2", MaxTier: 12},
		"LAPIS":       {Texture: "http://localhost:8080/api/head/64fd97b9346c1208c1db3957530cdfc5789e3e65943786b0071cf2b2904a6b5c", MaxTier: 12},
		"REDSTONE":    {Texture: "http://localhost:8080/api/head/1edefcf1a89d687a0a4ecf1589977af1e520fc673c48a0434be426612e8faa67", MaxTier: 12},
		"EMERALD":     {Texture: "http://localhost:8080/api/head/9bf57f3401b130c6b53808f2b1e119cc7b984622dac7077bbd53454e1f65bbf0", MaxTier: 12},
		"MITHRIL":     {Texture: "http://localhost:8080/api/head/c62fa670ff8599b32ab344195ba15f3ef64c3a8aa8a37821c08375950cb74cd0", MaxTier: 12},
		"QUARTZ":      {Texture: "http://localhost:8080/api/head/d270093be62dfd3019f908043db570b5dfd366fd5345fccf9da340e75c701a60", MaxTier: 12},
		"ENDER_STONE": {Name: "End Stone", Texture: "http://localhost:8080/api/head/7994be3dcfbb4ed0a5a7495b7335af1a3ced0b5888b5007286a790767c3b57e6"},
	},
	"combat": {
		"ZOMBIE":     {Texture: "http://localhost:8080/api/head/196063a884d3901c41f35b69a8c9f401c61ac9f6330f964f80c35352c3e8bfb0"},
		"REVENANT":   {Texture: "http://localhost:8080/api/head/a3dce8555923558d8d74c2a2b261b2b2d630559db54ef97ed3f9c30e9a20aba", MaxTier: 12},
		"SKELETON":   {Texture: "http://localhost:8080/api/head/2fe009c5cfa44c05c88e5df070ae2533bd682a728e0b33bfc93fd92a6e5f3f64"},
		"CREEPER":    {Texture: "http://localhost:8080/api/head/54a92c2f8c1b3774e80492200d0b2218d7b019314a73c9cb5b9f04cfcacec471"},
		"SPIDER":     {Texture: "http://localhost:8080/api/head/e77c4c284e10dea038f004d7eb43ac493de69f348d46b5c1f8ef8154ec2afdd0"},
		"TARANTULA":  {Texture: "http://localhost:8080/api/head/97e86007064c9ce26eb4bad8ac9aa30aac309e70a9e0b615936318dea40a721"},
		"CAVESPIDER": {Name: "Cave Spider", Texture: "http://localhost:8080/api/head/5d815df973bcd01ee8dfdb3bd74f0b7cb8fef2a70559e4faa5905127bbb4a435"},
		"BLAZE":      {Texture: "http://localhost:8080/api/head/3208fbd64e97c6e00853d36b3a201e4803cae43dcbd6936a3cece050912e1f20", MaxTier: 12},
		"MAGMA_CUBE": {Texture: "http://localhost:8080/api/head/18c9a7a24da7e3182e4f62fa62762e21e1680962197c7424144ae1d2c42174f7", MaxTier: 12},
		"ENDERMAN":   {Texture: "http://localhost:8080/api/head/e460d20ba1e9cd1d4cfd6d5fb0179ff41597ac6d2461bd7ccdb58b20291ec46e"},
		"GHAST":      {Texture: "http://localhost:8080/api/head/2478547d122ec83a818b46f3b13c5230429559e40c7d144d4ec225f92c1494b3", MaxTier: 12},
		"SLIME":      {Texture: "http://localhost:8080/api/head/c95eced85db62c922724efca804ea0060c4a87fcdedf2fd5c4f9ac1130a6eb26"},
		"VOIDLING":   {Texture: "http://localhost:8080/api/head/3a851ed2ce5c2c0523af772d206d9555e2e1383ec87946e6ff4c51186e29ef7f"},
		"INFERNO":    {Texture: "http://localhost:8080/api/head/665c54366f88fb3280b1c3fc500ce2b799c8dd327ab6d41c9bc959488f5cfd92"},
		"VAMPIRE":    {Texture: "http://localhost:8080/api/head/5b0c2db42e90f83fae6551c96e83669211a77c2c155c54d1523af3079f9565ed"},
	},
	"foraging": {
		"OAK":      {Texture: "http://localhost:8080/api/head/57e4a30f361204ea9cded3fbff850160731a0081cc452cfe26aed48e97f6364b"},
		"BIRCH":    {Texture: "http://localhost:8080/api/head/eb74109dbb88178afb7a9874afc682904cedb3df75978a51f7beeb28f924251"},
		"SPRUCE":   {Texture: "http://localhost:8080/api/head/7ba04bfe516955fd43932dcb33bd5eac20b38a231d9fa8415b3fb301f60f7363"},
		"DARK_OAK": {Texture: "http://localhost:8080/api/head/5ecdc8d6b2b7e081ed9c36609052c91879b89730b9953adbc987e25bf16c5581"},
		"ACACIA":   {Texture: "http://localhost:8080/api/head/42183eaf5b133b838db13d145247e389ab4b4f33c67846363792dc3d82b524c0"},
		"JUNGLE":   {Texture: "http://localhost:8080/api/head/2fe73d981690c1be346a16331819c4e8800859fcdc3e5153718c6ad45861924c"},
		"FLOWER":   {Texture: "http://localhost:8080/api/head/baa7c59b2f792d8d091aecacf47a19f8ab93f3fd3c48f6930b1c2baeb09e0f9b", MaxTier: 12},
	},
	"fishing": {
		"FISHING": {Texture: "http://localhost:8080/api/head/53ea0fd89524db3d7a3544904933830b4fc8899ef60c113d948bb3c4fe7aabb1", MaxTier: 12},
		"CLAY":    {Texture: "http://localhost:8080/api/head/af9b312c8f53da289060e6452855072e07971458abbf338ddec351e16c171ff8", MaxTier: 12},
	},
}

var MINION_CATEGORY_ICONS = map[string]string{
	"farming":  "http://localhost:8080/api/item/GOLD_HOE",
	"mining":   "http://localhost:8080/api/item/STONE_PICKAXE",
	"combat":   "http://localhost:8080/api/item/STONE_SWORD",
	"foraging": "http://localhost:8080/api/item/SAPLING:3",
	"fishing":  "http://localhost:8080/api/item/FISHING_ROD",
}

const MinionsMaxSlots = 26
