package constants

const MAX_PEAK_OF_THE_MOUNTAIN_LEVEL = 10

type Reward struct {
	TokenOfTheMountain                int
	SkyblockExperience                int
	AccessToForge                     int
	NewForgeableItems                 int
	ForgeSlot                         int
	AccessCrystalHollows              int
	EmissaryBraumCrystalHollows       int
	PickaxeAbilityLevel               int
	CommissionSlot                    int
	MithrilPowderWhenMiningMithril    int
	GemstonePowderWhenMiningGemstones int
	GlacitePowderWhenMiningGlacite    int
	ChanceForGlaciteMineshaftToSpawn  string
}

var HOTM_REWARDS = map[int]Reward{
	1:  {TokenOfTheMountain: 1, SkyblockExperience: 35},
	2:  {TokenOfTheMountain: 2, SkyblockExperience: 45, AccessToForge: 0, NewForgeableItems: 0},
	3:  {TokenOfTheMountain: 2, SkyblockExperience: 60, ForgeSlot: 1, NewForgeableItems: 0, AccessCrystalHollows: 0, EmissaryBraumCrystalHollows: 0},
	4:  {TokenOfTheMountain: 2, SkyblockExperience: 75, ForgeSlot: 1, NewForgeableItems: 0},
	5:  {TokenOfTheMountain: 2, SkyblockExperience: 90, NewForgeableItems: 0},
	6:  {TokenOfTheMountain: 2, SkyblockExperience: 100, NewForgeableItems: 0},
	7:  {TokenOfTheMountain: 3, SkyblockExperience: 130, NewForgeableItems: 0},
	8:  {TokenOfTheMountain: 2, SkyblockExperience: 180, NewForgeableItems: 0},
	9:  {TokenOfTheMountain: 2, SkyblockExperience: 210, NewForgeableItems: 0},
	10: {TokenOfTheMountain: 2, SkyblockExperience: 240, NewForgeableItems: 0},
}

var POTM_REWARDS = map[int]Reward{
	1:  {TokenOfTheMountain: 1, SkyblockExperience: 25},
	2:  {PickaxeAbilityLevel: 1, SkyblockExperience: 35},
	3:  {CommissionSlot: 1, SkyblockExperience: 50},
	4:  {MithrilPowderWhenMiningMithril: 1, SkyblockExperience: 65},
	5:  {TokenOfTheMountain: 1, SkyblockExperience: 75},
	6:  {GemstonePowderWhenMiningGemstones: 2, SkyblockExperience: 100},
	7:  {TokenOfTheMountain: 1, SkyblockExperience: 125},
	8:  {GlacitePowderWhenMiningGlacite: 3, SkyblockExperience: 150},
	9:  {ChanceForGlaciteMineshaftToSpawn: "10%", SkyblockExperience: 175},
	10: {TokenOfTheMountain: 2, SkyblockExperience: 200},
}

var PRECURSOR_PARTS = map[string]string{
	"ELECTRON_TRANSMITTER": "Electron Transmitter",
	"FTX_3070":             "FTX 3070",
	"ROBOTRON_REFLECTOR":   "Robotron Reflector",
	"SUPERLITE_MOTOR":      "Superlite Motor",
	"CONTROL_SWITCH":       "Control Switch",
	"SYNTHETIC_HEART":      "Synthetic Heart",
}

var GEMSTONE_CRYSTALS = []string{
	"jade", "amber", "amethyst", "sapphire", "topaz", "jasper", "ruby", "opal", "aquamarine", "peridot", "citrine", "onyx",
}

var FOSSILS = []string{
	"CLAW", "SPINE", "CLUBBED", "UGLY", "HELIX", "FOOTPRINT", "WEBBED", "TUSK",
}

var CORPSES = map[string]string{
	"lapis":    "/api/item/LAPIS_ARMOR_HELMET",
	"umber":    "/api/item/ARMOR_OF_YOG_HELMET",
	"tungsten": "/api/item/MINERAL_HELMET",
	"vanguard": "/api/item/VANGUARD_HELMET",
}

type forgeItem struct {
	Name     string
	Duration int64
}

var FORGE = map[string]forgeItem{
	"REFINED_DIAMOND":                {Name: "Refined Diamond", Duration: 28800000},
	"REFINED_MITHRIL":                {Name: "Refined Mithril", Duration: 21600000},
	"REFINED_TITANIUM":               {Name: "Refined Titanium", Duration: 43200000},
	"REFINED_TUNGSTEN":               {Name: "Refined Tungsten", Duration: 3600000},
	"REFINED_UMBER":                  {Name: "Refined Umber", Duration: 3600000},
	"MITHRIL_NECKLACE":               {Name: "Mithril Necklace", Duration: 3600000},
	"MITHRIL_CLOAK":                  {Name: "Mithril Cloak", Duration: 3600000},
	"MITHRIL_BELT":                   {Name: "Mithril Belt", Duration: 3600000},
	"MITHRIL_GAUNTLET":               {Name: "Mithril Gauntlet", Duration: 3600000},
	"TITANIUM_NECKLACE":              {Name: "Titanium Necklace", Duration: 16200000},
	"TITANIUM_CLOAK":                 {Name: "Titanium Cloak", Duration: 16200000},
	"TITANIUM_BELT":                  {Name: "Titanium Belt", Duration: 16200000},
	"TITANIUM_GAUNTLET":              {Name: "Titanium Gauntlet", Duration: 16200000},
	"TITANIUM_TALISMAN":              {Name: "Titanium Talisman", Duration: 50400000},
	"TITANIUM_RING":                  {Name: "Titanium Ring", Duration: 72000000},
	"TITANIUM_ARTIFACT":              {Name: "Titanium Artifact", Duration: 129600000},
	"TITANIUM_RELIC":                 {Name: "Titanium Relic", Duration: 259200000},
	"DIVAN_POWDER_COATING":           {Name: "Divan Powder Coating", Duration: 129600000},
	"DIVAN_HELMET":                   {Name: "Helmet Of Divan", Duration: 86400000},
	"DIVAN_CHESTPLATE":               {Name: "Chestplate Of Divan", Duration: 86400000},
	"DIVAN_LEGGINGS":                 {Name: "Leggings Of Divan", Duration: 86400000},
	"DIVAN_BOOTS":                    {Name: "Boots Of Divan", Duration: 86400000},
	"AMBER_NECKLACE":                 {Name: "Amber Necklace", Duration: 86400000},
	"SAPPHIRE_CLOAK":                 {Name: "Sapphire Cloak", Duration: 86400000},
	"JADE_BELT":                      {Name: "Jade Belt", Duration: 86400000},
	"AMETHYST_GAUNTLET":              {Name: "Amethyst Gauntlet", Duration: 86400000},
	"GEMSTONE_CHAMBER":               {Name: "Gemstone Chamber", Duration: 14400000},
	"DWARVEN_HANDWARMERS":            {Name: "Dwarven Handwarmers", Duration: 14400000},
	"DWARVEN_METAL":                  {Name: "Dwarven Metal Talisman", Duration: 86400000},
	"DIVAN_PENDANT":                  {Name: "Pendant of Divan", Duration: 604800000},
	"POWER_RELIC":                    {Name: "Relic of Power", Duration: 28800000},
	"PERFECT_AMBER_GEM":              {Name: "Perfect Amber Gemstone", Duration: 72000000},
	"PERFECT_AMETHYST_GEM":           {Name: "Perfect Amethyst Gemstone", Duration: 72000000},
	"PERFECT_JADE_GEM":               {Name: "Perfect Jade Gemstone", Duration: 72000000},
	"PERFECT_JASPER_GEM":             {Name: "Perfect Jasper Gemstone", Duration: 72000000},
	"PERFECT_OPAL_GEM":               {Name: "Perfect Opal Gemstone", Duration: 72000000},
	"PERFECT_RUBY_GEM":               {Name: "Perfect Ruby Gemstone", Duration: 72000000},
	"PERFECT_SAPPHIRE_GEM":           {Name: "Perfect Sapphire Gemstone", Duration: 72000000},
	"PERFECT_TOPAZ_GEM":              {Name: "Perfect Topaz Gemstone", Duration: 72000000},
	"PERFECT_AQUAMARINE_GEM":         {Name: "Perfect Aquamarine Gem", Duration: 72000000},
	"PERFECT_CITRINE_GEM":            {Name: "Perfect Citrine Gem", Duration: 72000000},
	"PERFECT_ONYX_GEM":               {Name: "Perfect Onyx Gem", Duration: 72000000},
	"PERFECT_PERIDOT_GEM":            {Name: "Perfect Peridot Gem", Duration: 72000000},
	"BEJEWELED_HANDLE":               {Name: "Bejeweled Handle", Duration: 30000},
	"DRILL_ENGINE":                   {Name: "Drill Motor", Duration: 108000000},
	"FUEL_TANK":                      {Name: "Fuel Canister", Duration: 36000000},
	"GEMSTONE_MIXTURE":               {Name: "Gemstone Mixture", Duration: 14400000},
	"GLACITE_AMALGAMATION":           {Name: "Glacite Amalgamation", Duration: 14400000},
	"GOLDEN_PLATE":                   {Name: "Golden Plate", Duration: 21600000},
	"MITHRIL_PLATE":                  {Name: "Mithril Plate", Duration: 64800000},
	"TUNGSTEN_PLATE":                 {Name: "Tungsten Plate", Duration: 10800000},
	"UMBER_PLATE":                    {Name: "Umber Plate", Duration: 10800000},
	"PERFECT_PLATE":                  {Name: "Perfect Plate", Duration: 1800000},
	"DIAMONITE":                      {Name: "Diamonite", Duration: 21600000},
	"POCKET_ICEBERG":                 {Name: "Pocket Iceberg", Duration: 21600000},
	"PETRIFIED_STARFALL":             {Name: "Petrified Starfall", Duration: 21600000},
	"PURE_MITHRIL":                   {Name: "Pure Mithril", Duration: 21600000},
	"ROCK_GEMSTONE":                  {Name: "Dwarven Geode", Duration: 21600000},
	"TITANIUM_TESSERACT":             {Name: "Titanium Tesseract", Duration: 21600000},
	"GLEAMING_CRYSTAL":               {Name: "Gleaming Crystal", Duration: 21600000},
	"HOT_STUFF":                      {Name: "Scorched Topaz", Duration: 21600000},
	"AMBER_MATERIAL":                 {Name: "Amber Material", Duration: 21600000},
	"FRIGID_HUSK":                    {Name: "Frigid Husk", Duration: 21600000},
	"BEJEWELED_COLLAR":               {Name: "Bejeweled Collar", Duration: 7200000},
	"MOLE":                           {Name: "[Lvl 1] Mole", Duration: 259200000},
	"AMMONITE":                       {Name: "[Lvl 1] Ammonite", Duration: 259200000},
	"PENGUIN":                        {Name: "[Lvl 1] Penguin", Duration: 604800000},
	"TYRANNOSAURUS":                  {Name: "[Lvl 1] T-Rex", Duration: 604800000},
	"SPINOSAURUS":                    {Name: "[Lvl 1] Spinosaurus", Duration: 604800000},
	"GOBLIN":                         {Name: "[Lvl 1] Goblin", Duration: 604800000},
	"ANKYLOSAURUS":                   {Name: "[Lvl 1] Ankylosaurus", Duration: 604800000},
	"MAMMOTH":                        {Name: "[Lvl 1] Mammoth", Duration: 604800000},
	"MITHRIL_DRILL_1":                {Name: "Mithril Drill SX-R226", Duration: 14400000},
	"MITHRIL_DRILL_2":                {Name: "Mithril Drill SX-R326", Duration: 30000},
	"GEMSTONE_DRILL_1":               {Name: "Ruby Drill TX-15", Duration: 14400000},
	"GEMSTONE_DRILL_2":               {Name: "Gemstone Drill LT-522", Duration: 30000},
	"GEMSTONE_DRILL_3":               {Name: "Topaz Drill KGR-12", Duration: 30000},
	"GEMSTONE_DRILL_4":               {Name: "Jasper Drill X", Duration: 30000},
	"POLISHED_TOPAZ_ROD":             {Name: "Polished Topaz Rod", Duration: 43200000},
	"TITANIUM_DRILL_1":               {Name: "Titanium Drill DR-X355", Duration: 14400000},
	"TITANIUM_DRILL_2":               {Name: "Titanium Drill DR-X455", Duration: 30000},
	"TITANIUM_DRILL_3":               {Name: "Titanium Drill DR-X555", Duration: 30000},
	"TITANIUM_DRILL_4":               {Name: "Titanium Drill DR-X655", Duration: 30000},
	"CHISEL":                         {Name: "Chisel", Duration: 14400000},
	"REINFORCED_CHISEL":              {Name: "Reinforced Chisel", Duration: 30000},
	"GLACITE_CHISEL":                 {Name: "Glacite-Plated Chisel", Duration: 30000},
	"PERFECT_CHISEL":                 {Name: "Perfect Chisel", Duration: 30000},
	"DIVAN_DRILL":                    {Name: "Divan's Drill", Duration: 30000},
	"STARFALL_SEASONING":             {Name: "Starfall Seasoning", Duration: 64800000},
	"GOBLIN_OMELETTE":                {Name: "Goblin Omelette", Duration: 64800000},
	"GOBLIN_OMELETTE_BLUE_CHEESE":    {Name: "Blue Cheese Goblin Omelette", Duration: 64800000},
	"GOBLIN_OMELETTE_PESTO":          {Name: "Pesto Goblin Omelette", Duration: 64800000},
	"GOBLIN_OMELETTE_SPICY":          {Name: "Spicy Goblin Omelette", Duration: 64800000},
	"GOBLIN_OMELETTE_SUNNY_SIDE":     {Name: "Sunny Side Goblin Omelette", Duration: 64800000},
	"TUNGSTEN_KEYCHAIN":              {Name: "Tungsten Regulator", Duration: 64800000},
	"MITHRIL_DRILL_ENGINE":           {Name: "Mithril-Plated Drill Engine", Duration: 86400000},
	"TITANIUM_DRILL_ENGINE":          {Name: "Titanium-Plated Drill Engine", Duration: 30000},
	"RUBY_POLISHED_DRILL_ENGINE":     {Name: "Ruby-polished Drill Engine", Duration: 30000},
	"SAPPHIRE_POLISHED_DRILL_ENGINE": {Name: "Sapphire-polished Drill Engine", Duration: 30000},
	"AMBER_POLISHED_DRILL_ENGINE":    {Name: "Amber-polished Drill Engine", Duration: 30000},
	"MITHRIL_FUEL_TANK":              {Name: "Mithril-Infused Fuel Tank", Duration: 86400000},
	"TITANIUM_FUEL_TANK":             {Name: "Titanium-Infused Fuel Tank", Duration: 30000},
	"GEMSTONE_FUEL_TANK":             {Name: "Gemstone Fuel Tank", Duration: 30000},
	"PERFECTLY_CUT_FUEL_TANK":        {Name: "Perfectly-Cut Fuel Tank", Duration: 30000},
	"BEACON_2":                       {Name: "Beacon II", Duration: 72000000},
	"BEACON_3":                       {Name: "Beacon III", Duration: 108000000},
	"BEACON_4":                       {Name: "Beacon IV", Duration: 144000000},
	"BEACON_5":                       {Name: "Beacon V", Duration: 180000000},
	"FORGE_TRAVEL_SCROLL":            {Name: "Travel Scroll to the Dwarven Forge", Duration: 18000000},
	"BASE_CAMP_TRAVEL_SCROLL":        {Name: "Travel Scroll to the Dwarven Base Camp", Duration: 36000000},
	"POWER_CRYSTAL":                  {Name: "Power Crystal", Duration: 7200000},
	"SECRET_RAILROAD_PASS":           {Name: "Secret Railroad Pass", Duration: 30000},
	"TUNGSTEN_KEY":                   {Name: "Tungsten Key", Duration: 1800000},
	"UMBER_KEY":                      {Name: "Umber Key", Duration: 1800000},
	"SKELETON_KEY":                   {Name: "Skeleton Key", Duration: 1800000},
	"PORTABLE_CAMPFIRE":              {Name: "Portable Campfire", Duration: 1800000},
}
