package model

var InitMaterialsVertical = []Material{
	{"01.ЖРС",
		"ЖРС, концентрат, 62% Fe",
		"Сырьевые материалы",
		"ferroalloynet.com",
		"Циндао (Китай)",
		"CNF",
		"$/т",
		"A",
		[]Property{{"Средняя цена", "E", 3, "decimal"}},
	},

	{"01.ЖРС",
		"ЖРС, концентрат, 65% Fe",
		"Сырьевые материалы",
		"ferroalloynet.com",
		"Циндао (Китай)",
		"CNF",
		"$/т",
		"A",
		[]Property{{"Средняя цена", "L", 3, "decimal"}},
	},

	{"02.Лом",
		"Лом, 3A",
		"Сырьевые материалы",
		"metallplace.ru",
		"Урал (Россия)",
		"CPT",
		"₽/т",
		"A",
		[]Property{
			{"Мин цена", "D", 3, "decimal"},
			{"Макс цена", "E", 3, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"02.Лом",
		"Лом, HMS 1&2 (80:20)",
		"Сырьевые материалы",
		"steelmint.com",
		"Искендерун (Турция)",
		"CNF",
		"$/т",
		"A",
		[]Property{
			{"Средняя цена", "L", 3, "decimal"},
		},
	},

	{"03.Чугун",
		"Чугун, чушковый",
		"Сырьевые материалы",
		"steelmint.com",
		"Черное море (Россия)",
		"FOB",
		"$/т",
		"B",
		[]Property{
			{"Мин цена", "D", 159, "decimal"},
			{"Макс цена", "E", 159, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"04.Уголь",
		"Кокс. уголь, австралийский",
		"Сырьевые материалы",
		"mysteel.net",
		"Циндао (Китай)",
		"CFR",
		"$/т",
		"A",
		[]Property{
			{"Средняя цена", "E", 3, "decimal"},
		},
	},

	{"04.Уголь",
		"Кокс. уголь, российский",
		"Сырьевые материалы",
		"mysteel.net",
		"Циндао (Китай)",
		"CFR",
		"$/т",
		"A",
		[]Property{
			{"Средняя цена", "L", 303, "decimal"},
		},
	},

	{"05.Мет. кокс",
		"Кокс, 25-90 мм; 64% CSR",
		"Сырьевые материалы",
		"steelmint.com",
		"Таншань (Китай)",
		"FOB",
		"$/т",
		"B",
		[]Property{
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Заготовка, 130*130 мм; Ст3",
		"Сталь",
		"steelmint.com",
		"Черное море (Россия)",
		"FOB",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "D", 159, "decimal"},
			{"Макс цена", "E", 159, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Арматура, 12-25 мм; Ст3",
		"Сталь",
		"steelmint.com",
		"Черное море (Россия)",
		"FOB",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "J", 159, "decimal"},
			{"Макс цена", "K", 159, "decimal"},
			{"Средняя цена", "L", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Сляб, 150*250 мм; Ст3",
		"Сталь",
		"steelmint.com",
		"Черное море (Россия)",
		"FOB",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "P", 203, "decimal"},
			{"Макс цена", "Q", 203, "decimal"},
			{"Средняя цена", "R", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон, г/к 3 мм; SAE 1006",
		"Сталь",
		"steelmint.com",
		"Черное море (Россия)",
		"FOB",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "V", 159, "decimal"},
			{"Макс цена", "W", 159, "decimal"},
			{"Средняя цена", "X", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон, х/к 0,9 мм; SPCC",
		"Сталь",
		"steelmint.com",
		"Черное море (Россия)",
		"FOB",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "AB", 203, "decimal"},
			{"Макс цена", "AC", 203, "decimal"},
			{"Средняя цена", "AD", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Арматура, А1 8-40 мм; Ст3",
		"Сталь",
		"metallplace.ru",
		"Россия",
		"EXW",
		"₽/т",
		"A",
		[]Property{
			{"Средняя цена", "AJ", 212, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон г/к, 1,5-20 мм; Ст3",
		"Сталь",
		"steelmint.com",
		"Россия",
		"EXW",
		"₽/т",
		"A",
		[]Property{
			{"Средняя цена", "AP", 212, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон х/к, 0,5-1,5 мм; Ст08",
		"Сталь",
		"steelmint.com",
		"Россия",
		"EXW",
		"₽/т",
		"A",
		[]Property{
			{"Средняя цена", "AV", 213, "decimal"},
		},
	},

	{"07.ФС (М)",
		"FeMn, HC; 76% Mn",
		"Ферросплавы и руды",
		"crugroup.com",
		"Роттердам (ЕС)",
		"DDP",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "D", 11, "decimal"},
			{"Макс цена", "E", 11, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"FeSi; 75% Si",
		"Ферросплавы и руды",
		"crugroup.com",
		"Роттердам (ЕС)",
		"DDP",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "K", 11, "decimal"},
			{"Макс цена", "L", 11, "decimal"},
			{"Средняя цена", "M", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"SiMn; 65% Mn, 17% Si",
		"Ферросплавы и руды",
		"crugroup.com",
		"Роттердам (ЕС)",
		"DDP",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "R", 11, "decimal"},
			{"Макс цена", "S", 11, "decimal"},
			{"Средняя цена", "T", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"FeCr, HC; 62-70% Cr",
		"Ферросплавы и руды",
		"crugroup.com",
		"Роттердам (ЕС)",
		"DDP",
		"¢/фунт",
		"A",
		[]Property{
			{"Мин цена", "Y", 11, "decimal"},
			{"Макс цена", "Z", 11, "decimal"},
			{"Средняя цена", "AA", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"FeCr, LC; 0,1% Cr",
		"Ферросплавы и руды",
		"crugroup.com",
		"Роттердам (ЕС)",
		"DDP",
		"¢/фунт",
		"A",
		[]Property{
			{"Мин цена", "AF", 11, "decimal"},
			{"Макс цена", "AG", 11, "decimal"},
			{"Средняя цена", "AH", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"Mn руда, кусковая; 36-38% Mn",
		"Ферросплавы и руды",
		"crugroup.com",
		"Таншань (Китай)",
		"CIF",
		"$/1% Mn смт",
		"A",
		[]Property{
			{"Мин цена", "AM", 11, "decimal"},
			{"Макс цена", "AN", 11, "decimal"},
			{"Средняя цена", "AO", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"Cr руда, кусоквая; 42% Cr",
		"Ферросплавы и руды",
		"crugroup.com",
		"Таншань (Китай)",
		"CIF",
		"$/т",
		"A",
		[]Property{
			{"Мин цена", "AT", 11, "decimal"},
			{"Макс цена", "AU", 11, "decimal"},
			{"Средняя цена", "AV", 3, "decimal"},
		},
	},

	{"08.ГЭ",
		"ГЭ, HP; 450 мм",
		"Электроды",
		"steelmint.com",
		"Шэньси (Китай)",
		"EXW",
		"$/т",
		"A",
		[]Property{
			{"Средняя цена", "F", 212, "decimal"},
		},
	},

	{"08.ГЭ",
		"ГЭ, UHP; 600 мм",
		"Электроды",
		"steelmint.com",
		"Шэньси (Китай)",
		"EXW",
		"$/т",
		"A",
		[]Property{
			{"Средняя цена", "L", 212, "decimal"},
		},
	},

	{"09.Запасы",
		"Запасы в портах, Mn руда",
		"Ферросплавы и руды",
		"ferroalloynet.com",
		"Китай",
		"",
		"млн т",
		"B",
		[]Property{
			{"Запас", "F", 3, "decimal"},
		},
	},

	{"09.Запасы",
		"Запасы в портах, Cr руда",
		"Ферросплавы и руды",
		"ferroalloynet.com",
		"Китай",
		"",
		"млн т",
		"B",
		[]Property{
			{"Запас", "L", 3, "decimal"},
		},
	},

	{"09.Запасы",
		"Запасы в портах, Fe руда",
		"Ферросплавы и руды",
		"steelmint.com",
		"Китай",
		"",
		"млн т",
		"B",
		[]Property{
			{"Запас", "R", 3, "decimal"},
		},
	},
}
