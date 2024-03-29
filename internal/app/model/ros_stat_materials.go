package model

type MaterialCoordinates struct {
	Material string `json:"material"`
	Sheet    string `json:"sheet"`
	Row      int    `json:"row"`
}

var RosStatCoordinates = []MaterialCoordinates{
	{Sheet: "Н-П_24", Row: 56, Material: "Ферромолибден"},
	{Sheet: "Н-П_24", Row: 64, Material: "Феррованадий"},
	{Sheet: "Н-П_24", Row: 49, Material: "Ферросилиций"},
	{Sheet: "Н-П_24", Row: 91, Material: "Ферромарганец"},
	{Sheet: "Н-П_24", Row: 85, Material: "Ферросиликомарганец"},
	{Sheet: "Н-П_24", Row: 75, Material: "Феррохром"},
	{Sheet: "Н-П_24", Row: 81, Material: "Ферросиликохром"},

	{Sheet: "Н-П_07", Row: 11, Material: "Руда железная сырая"},
	{Sheet: "Н-П_07", Row: 19, Material: "Руда железная товарная необогащенная"},
	{Sheet: "Н-П_07", Row: 26, Material: "Концентрат железорудный"},
	{Sheet: "Н-П_07", Row: 56, Material: "Агломерат железорудный"},
	{Sheet: "Н-П_07", Row: 64, Material: "Окатыши железорудные (окисленные)"},

	{Sheet: "Н-П_24", Row: 100, Material: "Продукты прямого восстановления железной руды и прочее губчатое железо в кусках, окатышах или аналогичных формах; железо с минимальным содержанием основного элемента 99,94 % в кусках, окатышах или аналогичных формах"},

	{Sheet: "Н-П_05", Row: 25, Material: "Антрацит"},
	{Sheet: "Н-П_05", Row: 57, Material: "Антрацит обогащенный"},

	{Sheet: "Н-П_05", Row: 30, Material: "Уголь коксующийся"},
	{Sheet: "Н-П_05", Row: 62, Material: "Уголь коксующийся обогащенный"},
	{Sheet: "Н-П_05", Row: 36, Material: "Уголь, за исключением антрацита, угля коксующегося и угля бурого"},
	{Sheet: "Н-П_05", Row: 68, Material: "Уголь обогащенный, за исключением антрацита, угля коксующегося и угля бурого (лигнита)"},
	{Sheet: "Н-П_05", Row: 74, Material: "Уголь бурый рядовой (лигнит)"},

	{Sheet: "Н-П_24", Row: 25, Material: "Чугун передельный в чушках, болванках или в прочих первичных формах"},
	{Sheet: "Н-П_24", Row: 34, Material: "Чугун литейный"},
}
