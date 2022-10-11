package model

var InitMaterialsVertical = []Material{
	{"01.ЖРС",
		"ЖРС (62%), CNF Китай",
		"ferroalloynet.com",
		"Глобальный",
		"USD/т",
		"A",
		[]Property{{"Средняя цена", "E", 3, "decimal"}},
	},

	{"01.ЖРС",
		"ЖРС (65%), Platts",
		"ferroalloynet.com",
		"Глобальный",
		"USD/t",
		"A",
		[]Property{{"Средняя цена", "L", 3, "decimal"}},
	},

	{"02.Лом",
		"Лом 3А, CPT Урал",
		"-",
		"-",
		"руб/т",
		"A",
		[]Property{
			{"Мин цена", "D", 3, "decimal"},
			{"Макс цена", "E", 3, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"02.Лом",
		"Лом HMS (80:20), CNF Турция",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Средняя цена", "L", 3, "decimal"},
		},
	},

	{"03.Чугун",
		"Чугун, FOB Черное море",
		"steelmint.com",
		"-",
		"USD/т",
		"B",
		[]Property{
			{"Мин цена", "D", 159, "decimal"},
			{"Макс цена", "E", 159, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"04.Уголь",
		"Кокс. уголь, CFR Китай",
		"mysteel.net",
		"Австралия",
		"USD/т",
		"A",
		[]Property{
			{"Средняя цена", "E", 3, "decimal"},
		},
	},

	{"04.Уголь",
		"Кокс. уголь, CFR Китай",
		"mysteel.net",
		"Россия",
		"USD/т",
		"A",
		[]Property{
			{"Средняя цена", "L", 303, "decimal"},
		},
	},

	{"05.Мет. кокс",
		"Металлургический кокс, FOB Китай",
		"steelmint.com",
		"Глобальный",
		"USD/т",
		"B",
		[]Property{
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Заготовка, FOB Черное море",
		"steelmint.com",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "D", 159, "decimal"},
			{"Макс цена", "E", 159, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Арматура, FOB Черное море",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "J", 159, "decimal"},
			{"Макс цена", "K", 159, "decimal"},
			{"Средняя цена", "L", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Сляб, FOB Черное море",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "P", 203, "decimal"},
			{"Макс цена", "Q", 203, "decimal"},
			{"Средняя цена", "R", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон г/к, FOB Черное море",
		"steelmint.com",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "V", 159, "decimal"},
			{"Макс цена", "W", 159, "decimal"},
			{"Средняя цена", "X", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон х/к, FOB Черное море",
		"steelmint.com",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "AB", 203, "decimal"},
			{"Макс цена", "AC", 203, "decimal"},
			{"Средняя цена", "AD", 3, "decimal"},
		},
	},

	{"06.Сталь",
		"Арматура А1, EXW Россия",
		"-",
		"-",
		"руб/т",
		"A",
		[]Property{
			{"Средняя цена", "AJ", 212, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон г/к, EXW Россия",
		"steelmint.com",
		"-",
		"руб/т",
		"A",
		[]Property{
			{"Средняя цена", "AP", 212, "decimal"},
		},
	},

	{"06.Сталь",
		"Рулон х/к, EXW Россия",
		"steelmint.com",
		"-",
		"руб/т",
		"A",
		[]Property{
			{"Средняя цена", "AV", 213, "decimal"},
		},
	},

	{"07.ФС (М)",
		"FeMn76, DDP Европа",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "D", 11, "decimal"},
			{"Макс цена", "E", 11, "decimal"},
			{"Средняя цена", "F", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"FeSi75, DDP Европа",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "K", 11, "decimal"},
			{"Макс цена", "L", 11, "decimal"},
			{"Средняя цена", "M", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"SiMn65, DDP Европа",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "R", 11, "decimal"},
			{"Макс цена", "S", 11, "decimal"},
			{"Средняя цена", "T", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"HC FeCr, DDP Европа",
		"-",
		"-",
		"¢/фунт",
		"A",
		[]Property{
			{"Мин цена", "Y", 11, "decimal"},
			{"Макс цена", "Z", 11, "decimal"},
			{"Средняя цена", "AA", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"LC FeCr, DDP Европа",
		"-",
		"-",
		"¢/фунт",
		"A",
		[]Property{
			{"Мин цена", "AF", 11, "decimal"},
			{"Макс цена", "AG", 11, "decimal"},
			{"Средняя цена", "AH", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"Mn руда (36-38), CIF Китай",
		"-",
		"-",
		"USD/1%",
		"A",
		[]Property{
			{"Мин цена", "AM", 11, "decimal"},
			{"Макс цена", "AN", 11, "decimal"},
			{"Средняя цена", "AO", 3, "decimal"},
		},
	},

	{"07.ФС (М)",
		"Cr руда, CIF Китай",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Мин цена", "AT", 11, "decimal"},
			{"Макс цена", "AU", 11, "decimal"},
			{"Средняя цена", "AV", 3, "decimal"},
		},
	},

	{"08.ГЭ",
		"ГЭ 450 мм, EXW Китай",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Средняя цена", "F", 212, "decimal"},
		},
	},

	{"08.ГЭ",
		"ГЭ 600 мм, EXW Китай",
		"-",
		"-",
		"USD/т",
		"A",
		[]Property{
			{"Средняя цена", "L", 212, "decimal"},
		},
	},

	{"09.Запасы",
		"Запасы Mn руд в портах Китая",
		"-",
		"-",
		"млн тонн",
		"B",
		[]Property{
			{"Запас", "F", 3, "decimal"},
		},
	},

	{"09.Запасы",
		"Запасы Cr руд в портах Китая",
		"-",
		"-",
		"млн тонн",
		"B",
		[]Property{
			{"Запас", "L", 3, "decimal"},
		},
	},

	{"09.Запасы",
		"Запасы железной руды в портах Китая",
		"-",
		"-",
		"млн тонн",
		"B",
		[]Property{
			{"Запас", "R", 3, "decimal"},
		},
	},
}
