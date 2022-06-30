package model

var InitMaterials = []Material{
	{"fe",
		"Iron ore (62% Fe)",
		"ferroalloynet.com",
		"Worldwide",
		"usd/t",
		"A",
		[]Property{{"price", "C", 3, "decimal"}},
	},

	{"fe",
		"Iron ore (65% Fe)",
		"ferroalloynet.com",
		"Worldwide",
		"usd/t",
		"A",
		[]Property{{"price", "H", 3, "decimal"}},
	},

	{"fe",
		"Iron ore (62% Fe)",
		"mysteel.net",
		"China",
		"usd/t",
		"A",
		[]Property{{"price", "M", 741, "decimal"}},
	},

	{"scrap",
		"Scrap 3A",
		"metalexpert.com",
		"Russia",
		"rub/t",
		"A",
		[]Property{{"min_price", "D", 3, "decimal"},
			{"max_price", "E", 3, "decimal"},
			{"med_price", "F", 3, "decimal"},
		},
	},

	{"scrap",
		"Scrap HMS (80:20)",
		"Not specified",
		"Turkey",
		"usd/t",
		"A",
		[]Property{{"min_price", "J", 209, "decimal"},
			{"max_price", "K", 209, "decimal"},
			{"med_price", "L", 3, "decimal"},
		},
	},
}
