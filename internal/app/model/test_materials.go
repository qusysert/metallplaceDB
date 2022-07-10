package model

var TestMaterials = []Material{
	{"sheet",
		"Material 1",
		"Website 1",
		"Global",
		"usd/t",
		"A",
		[]Property{{"min_price", "C", 2, "decimal"},
			{"max_price", "D", 2, "decimal"},
			{"med_price", "E", 2, "decimal"},
		},
	},

	{"sheet",
		"Material 2",
		"Website 2",
		"Global",
		"usd/t",
		"A",
		[]Property{{"min_price", "I", 2, "decimal"},
			{"max_price", "J", 2, "decimal"},
			{"med_price", "K", 2, "decimal"},
		},
	},
}
