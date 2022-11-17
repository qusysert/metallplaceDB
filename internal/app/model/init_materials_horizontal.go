package model

var InitMaterialsHorizontal = []MaterialHorizontal{
	// ЛОМ ЧЕРНЫХ МЕТАЛЛОВ
	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (80:20), CNF, (недельный)",
		"-",
		"Турция, Искендерун",
		"USD/т",
		"1",
		[]Property{
			{"Мин цена", "F", 3, "decimal"},
			{"Макс цена", "G", 3, "decimal"},
			{"Средняя цена", "H", 3, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (80:20), FOB, (недельный)",
		"-",
		"Европа, Роттердам",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 4, "decimal"},
			{"Макс цена", "G", 4, "decimal"},
			{"Средняя цена", "H", 4, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (80:20), CNF, (недельный)",
		"-",
		"Бангладеш, Читтагонг",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 5, "decimal"},
			{"Макс цена", "G", 5, "decimal"},
			{"Средняя цена", "H", 5, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, шредированный, CNF, (недельный)",
		"-",
		"Бангладеш, Читтагонг",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 6, "decimal"},
			{"Макс цена", "G", 6, "decimal"},
			{"Средняя цена", "H", 6, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, H2, CNF, (недельный)",
		"-",
		"Бангладеш, Читтагонг",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "FZ", 7, "decimal"},
			{"Макс цена", "GA", 7, "decimal"},
			{"Средняя цена", "GB", 7, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (80:20), CNF, (недельный)",
		"-",
		"Вьетнам, Хайфон",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "FZ", 8, "decimal"},
			{"Макс цена", "GA", 8, "decimal"},
			{"Средняя цена", "GB", 8, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, H2, CNF, (недельный)",
		"-",
		"Вьетнам, Хайфон",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "FZ", 9, "decimal"},
			{"Макс цена", "GA", 9, "decimal"},
			{"Средняя цена", "GB", 9, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (80:20), CNF, (недельный)",
		"-",
		"Индия, Нава-Шева",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "BV", 10, "decimal"},
			{"Макс цена", "BW", 10, "decimal"},
			{"Средняя цена", "BX", 10, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, шредированный, CNF, (недельный)",
		"-",
		"Индия, Нава-Шева",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "BV", 11, "decimal"},
			{"Макс цена", "BW", 11, "decimal"},
			{"Средняя цена", "BX", 11, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, шредированный, CNF, (недельный)",
		"-",
		"Пакистан, Касим",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 12, "decimal"},
			{"Макс цена", "G", 12, "decimal"},
			{"Средняя цена", "H", 12, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (80:20), CNF, (недельный)",
		"-",
		"Тайвань, Тайчжун",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 13, "decimal"},
			{"Макс цена", "G", 13, "decimal"},
			{"Средняя цена", "H", 13, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (70:30), CNF, (недельный)",
		"-",
		"Таиланд, Лаем-Чабанг",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "TU", 14, "decimal"},
			{"Макс цена", "TV", 14, "decimal"},
			{"Средняя цена", "TW", 14, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, HMS 1&2 (80:20), FOB, (недельный)",
		"-",
		"США, Восточное побережье",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "WM", 15, "decimal"},
			{"Макс цена", "WN", 15, "decimal"},
			{"Средняя цена", "WO", 15, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, шредированный, FOB, (недельный)",
		"-",
		"США, Восточное побережье",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "WM", 16, "decimal"},
			{"Макс цена", "WN", 16, "decimal"},
			{"Средняя цена", "WO", 16, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Лом, H2, FOB, (недельный)",
		"-",
		"Япония, Токийский залив",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "WM", 17, "decimal"},
			{"Макс цена", "WN", 17, "decimal"},
			{"Средняя цена", "WO", 17, "decimal"},
		},
	},

	// ЗАГОТОВКА
	{"10.Сводка (неделя)",
		"Заготовка, 150*150 мм, CNF, (недельная)",
		"-",
		"Китай, Цзянъинь",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "OA", 19, "decimal"},
			{"Макс цена", "OB", 19, "decimal"},
			{"Средняя цена", "OC", 19, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Заготовка, 150*150 мм, FOB, (недельная)",
		"-",
		"Индия, Восточное побережье",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 20, "decimal"},
			{"Макс цена", "G", 20, "decimal"},
			{"Средняя цена", "H", 20, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Заготовка, 95*95 мм, FOB, (недельная)",
		"-",
		"Индия, Восточное побережье",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 21, "decimal"},
			{"Макс цена", "G", 21, "decimal"},
			{"Средняя цена", "H", 21, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Заготовка, 130*130 мм, FOB, (недельная)",
		"-",
		"Иран, Бендер Аббас",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 22, "decimal"},
			{"Макс цена", "G", 22, "decimal"},
			{"Средняя цена", "H", 22, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Заготовка, 120*120 мм, CNF, (недельная)",
		"-",
		"Турция, Искендерун",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 23, "decimal"},
			{"Макс цена", "G", 23, "decimal"},
			{"Средняя цена", "H", 23, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Заготовка, 150*150 мм, CNF, (недельная)",
		"-",
		"Филиппины, Манила",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "BV", 24, "decimal"},
			{"Макс цена", "BW", 24, "decimal"},
			{"Средняя цена", "BX", 24, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Заготовка, 100*100 мм, CNF, (недельная)",
		"-",
		"Непал, Раксол",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "FZ", 25, "decimal"},
			{"Макс цена", "GA", 25, "decimal"},
			{"Средняя цена", "GB", 25, "decimal"},
		},
	},

	// СОРТОВОЙ ПРОКАТ
	{"10.Сводка (неделя)",
		"Сортовой прокат, Арматура 10-40 мм, CNF, (недельная)",
		"-",
		"Китай, Гонконг",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "TU", 27, "decimal"},
			{"Макс цена", "TV", 27, "decimal"},
			{"Средняя цена", "TW", 27, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Сортовой прокат, Арматура 12-32 мм, FOB, (недельная)",
		"-",
		"Турция, Искендерун",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "TU", 28, "decimal"},
			{"Макс цена", "TV", 28, "decimal"},
			{"Средняя цена", "TW", 28, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Сортовой прокат, Катанка 5,5-8 мм, CNF, (недельная)",
		"-",
		"Непал, Раксол",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "TU", 29, "decimal"},
			{"Макс цена", "TV", 29, "decimal"},
			{"Средняя цена", "TW", 29, "decimal"},
		},
	},

	// РУЛОННАЯ СТАЛЬ
	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 2 мм, CNF, (недельная)",
		"-",
		"ОАЭ, Абу-Даби",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 31, "decimal"},
			{"Макс цена", "G", 31, "decimal"},
			{"Средняя цена", "H", 31, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 1,5-3 мм, CNF, (недельная)",
		"-",
		"Турция, Искендерун",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 32, "decimal"},
			{"Макс цена", "G", 32, "decimal"},
			{"Средняя цена", "H", 32, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 3-4 мм, CNF, (недельная)",
		"-",
		"Европа, Западная Европа",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 33, "decimal"},
			{"Макс цена", "G", 33, "decimal"},
			{"Средняя цена", "H", 33, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 2-12 мм, CNF, (недельная)",
		"-",
		"Индия, Западное побережье",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 34, "decimal"},
			{"Макс цена", "G", 34, "decimal"},
			{"Средняя цена", "H", 34, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 2-20 мм, CNF, (недельная)",
		"-",
		"Вьетнам, Хошимин",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 35, "decimal"},
			{"Макс цена", "G", 35, "decimal"},
			{"Средняя цена", "H", 35, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 2 мм, CNF, (недельная)",
		"-",
		"Бангладеш, Читтагонг",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 36, "decimal"},
			{"Макс цена", "G", 36, "decimal"},
			{"Средняя цена", "H", 36, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 3-12 мм, FOB, (недельная)",
		"-",
		"Китай, Ричжао",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 37, "decimal"},
			{"Макс цена", "G", 37, "decimal"},
			{"Средняя цена", "H", 37, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 1,6-2,5 мм, FOB, (недельная)",
		"-",
		"Южная Корея, Сеул",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 38, "decimal"},
			{"Макс цена", "G", 38, "decimal"},
			{"Средняя цена", "H", 38, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 2 мм, FOB, (недельная)",
		"-",
		"Япония, Токио",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 39, "decimal"},
			{"Макс цена", "G", 39, "decimal"},
			{"Средняя цена", "H", 39, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 2 мм, CNF, (недельная)",
		"-",
		"Пакистан, Касим",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "WM", 40, "decimal"},
			{"Макс цена", "WN", 40, "decimal"},
			{"Средняя цена", "WO", 40, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, г/к рулон 2 мм, CNF, (недельная)",
		"-",
		"Непал, Раксол",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "WM", 41, "decimal"},
			{"Макс цена", "WN", 41, "decimal"},
			{"Средняя цена", "WO", 41, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, х/к рулон 0,9 мм, FOB, (недельная)",
		"-",
		"Китай, Ричжао",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 42, "decimal"},
			{"Макс цена", "G", 42, "decimal"},
			{"Средняя цена", "H", 42, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Рулонная сталь, х/к рулон 0,9 мм, CNF, (недельная)",
		"-",
		"Индия, Восточное побережье",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "F", 43, "decimal"},
			{"Макс цена", "G", 43, "decimal"},
			{"Средняя цена", "H", 43, "decimal"},
		},
	},

	// ЧУГУН
	{"10.Сводка (неделя)",
		"Чугун, чушковый, FOB, (недельная)",
		"-",
		"Индия, Восточное побережье",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "HM", 45, "decimal"},
			{"Макс цена", "HN", 45, "decimal"},
			{"Средняя цена", "HO", 45, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Чугун, чушковый, FOB, (недельная)",
		"-",
		"Бразилия, Понта-Да-Мадейра",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "HM", 46, "decimal"},
			{"Макс цена", "HN", 46, "decimal"},
			{"Средняя цена", "HO", 46, "decimal"},
		},
	},

	{"10.Сводка (неделя)",
		"Чугун, чушковый, FOB, (недельная)",
		"-",
		"Италия, Маргера",
		"USD/т",
		"1",
		[]Property{{"Мин цена", "HM", 47, "decimal"},
			{"Макс цена", "HN", 47, "decimal"},
			{"Средняя цена", "HO", 47, "decimal"},
		},
	},
}
