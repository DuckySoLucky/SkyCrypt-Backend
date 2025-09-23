package constants

type floor struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
	ID      string `json:"id"`
}

type dungeons struct {
	Floors map[string][]floor `json:"floors"`
}

var DUNGEONS = dungeons{
	Floors: map[string][]floor{
		"catacombs": {
			{
				Name:    "Entrance",
				Texture: "/api/head/35c3024f4d9d12ddf5959b6aea3c810f5ee85176aab1b2e7f462aa1c194c342b",
				ID:      "0",
			},
			{
				Name:    "Floor 1",
				Texture: "/api/head/726f384acdfbb7218e96efac630e9ae1a14fd2f820ab660cc68322a59b165a12",
				ID:      "1",
			},
			{
				Name:    "Floor 2",
				Texture: "/api/head/ebaf2ae74553a64587840d6e70fb27d2c0ae2f8bdfacbe56654c8db4001cdc98",
				ID:      "2",
			},
			{
				Name:    "Floor 3",
				Texture: "/api/head/5a2f67500a65f3ce79d34ec150de93df8f60ebe52e248f5e1cdb69b0726256f7",
				ID:      "3",
			},
			{
				Name:    "Floor 4",
				Texture: "/api/head/5720917cda0567442617f2721e88be9d2ffbb0b26a3f4c2fe21655814d4f4476",
				ID:      "4",
			},
			{
				Name:    "Floor 5",
				Texture: "/api/head/5720917cda0567442617f2721e88be9d2ffbb0b26a3f4c2fe21655814d4f4476",
				ID:      "5",
			},
			{
				Name:    "Floor 6",
				Texture: "/api/head/3ce69d2ddcc81c9fc2e9948c92003eb0f7ebf0e7e952e801b7f2069dcee76d85",
				ID:      "6",
			},
			{
				Name:    "Floor 7",
				Texture: "/api/head/76965e3fd619de6b0a7ce1673072520a9360378e1cb8c19d4baf0c86769d3764",
				ID:      "7",
			},
		},
		"master_catacombs": {
			{
				Name:    "Floor 1",
				Texture: "/api/head/1eb5b21af330af122b268b7aa390733bd1b699b4d0923233ecd24f81e08b9bce",
				ID:      "1",
			},
			{
				Name:    "Floor 2",
				Texture: "/api/head/32292e4e0fa62667256ef8da0f01982a996499f4d5d894bd058c3e6f3d2fb2d9",
				ID:      "2",
			},
			{
				Name:    "Floor 3",
				Texture: "/api/head/c969f6b148648aa8d027228a52fb5a3ca1ee84dc76e47851f14db029a730a8a3",
				ID:      "3",
			},
			{
				Name:    "Floor 4",
				Texture: "/api/head/d7b69021f9c09647dfd9b34df3deaff70cfc740f6a26f612dd47503fc34c97f0",
				ID:      "4",
			},
			{
				Name:    "Floor 5",
				Texture: "/api/head/d65cbce40e60e7a59a87fa8f4ecb6ccfc1514338c262614bf33739a6263f5405",
				ID:      "5",
			},
			{
				Name:    "Floor 6",
				Texture: "/api/head/d65cbce40e60e7a59a87fa8f4ecb6ccfc1514338c262614bf33739a6263f5405",
				ID:      "6",
			},
			{
				Name:    "Floor 7",
				Texture: "/api/head/d65cbce40e60e7a59a87fa8f4ecb6ccfc1514338c262614bf33739a6263f5405",
				ID:      "7",
			},
		},
	},
}
