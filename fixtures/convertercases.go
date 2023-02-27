package fixtures

type TestCaseConvertDigit2Word struct {
	Description   string
	Given         string
	Expected      string
	IsErrExpected bool
}

type TestCaseConvertIntPart struct {
	Description string
	Given       int
	Expected    string
}

func ConvertIntPartsForSingleDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: Single digit number"

	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       0,
			Expected:    "sıfır",
		},
		{
			Description: Description,
			Given:       1,
			Expected:    "bir",
		},
		{
			Description: Description,
			Given:       2,
			Expected:    "iki",
		},
		{
			Description: Description,
			Given:       3,
			Expected:    "üç",
		},
		{
			Description: Description,
			Given:       4,
			Expected:    "dörd",
		}, {
			Description: Description,
			Given:       5,
			Expected:    "beş",
		},
		{
			Description: Description,
			Given:       6,
			Expected:    "altı",
		},
		{
			Description: Description,
			Given:       7,
			Expected:    "yeddi",
		}, {
			Description: Description,
			Given:       8,
			Expected:    "səkkiz",
		},
		{
			Description: Description,
			Given:       9,
			Expected:    "doqquz",
		},
	}
}

func ConvertIntPartsForTwoDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: Two digits number"
	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       10,
			Expected:    "on",
		},
		{
			Description: Description,
			Given:       11,
			Expected:    "on bir",
		},
		{
			Description: Description,
			Given:       19,
			Expected:    "on doqquz",
		},
		{
			Description: Description,
			Given:       27,
			Expected:    "iyirmi yeddi",
		},
		{
			Description: Description,
			Given:       32,
			Expected:    "otuz iki",
		},
		{
			Description: Description,
			Given:       48,
			Expected:    "qırx səkkiz",
		},
		{
			Description: Description,
			Given:       54,
			Expected:    "əlli dörd",
		},
		{
			Description: Description,
			Given:       67,
			Expected:    "altmış yeddi",
		},
		{
			Description: Description,
			Given:       70,
			Expected:    "yetmiş",
		},
		{
			Description: Description,
			Given:       82,
			Expected:    "səksən iki",
		},
		{
			Description: Description,
			Given:       96,
			Expected:    "doxsan altı",
		}}
}

func ConvertIntPartsForThreeDigitsNumbers() []TestCaseConvertIntPart {
	return []TestCaseConvertIntPart{
		{
			Description: "Test Case: Single three digits number",
			Given:       100,
			Expected:    "bir yüz",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       111,
			Expected:    "bir yüz on bir",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       119,
			Expected:    "bir yüz on doqquz",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       271,
			Expected:    "iki yüz yetmiş bir",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       321,
			Expected:    "üç yüz iyirmi bir",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       485,
			Expected:    "dörd yüz səksən beş",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       541,
			Expected:    "beş yüz qırx bir",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       594,
			Expected:    "beş yüz doxsan dörd",
		},
		{
			Description: "Test Case: Single three digits number",
			Given:       670,
			Expected:    "altı yüz yetmiş",
		},
		{
			Description: "Test Case: Single two digits number",
			Given:       701,
			Expected:    "yeddi yüz bir",
		},
		{
			Description: "Test Case: Single two digits number",
			Given:       820,
			Expected:    "səkkiz yüz iyirmi",
		},
		{
			Description: "Test Case: Single two digits number",
			Given:       906,
			Expected:    "doqquz yüz altı",
		}}
}

func ConvertIntPartsForFourDigitsNumbers() []TestCaseConvertIntPart {
	return []TestCaseConvertIntPart{
		{
			Description: "Test Case: four digits number",
			Given:       1000,
			Expected:    "bir min",
		},
		{
			Description: "Test Case: four digits number",
			Given:       1111,
			Expected:    "bir min bir yüz on bir",
		},
		{
			Description: "Test Case: four digits number",
			Given:       1019,
			Expected:    "bir min on doqquz",
		},
		{
			Description: "Test Case: four digits number",
			Given:       2701,
			Expected:    "iki min yeddi yüz bir",
		},
		{
			Description: "Test Case: four digits number",
			Given:       3210,
			Expected:    "üç min iki yüz on",
		},
		{
			Description: "Test Case: three digits number",
			Given:       4850,
			Expected:    "dörd min səkkiz yüz əlli",
		},
		{
			Description: "Test Case: three digits number",
			Given:       5412,
			Expected:    "beş min dörd yüz on iki",
		},
		{
			Description: "Test Case: three digits number",
			Given:       9594,
			Expected:    "doqquz min beş yüz doxsan dörd",
		},
		{
			Description: "Test Case: three digits number",
			Given:       8670,
			Expected:    "səkkiz min altı yüz yetmiş",
		},
		{
			Description: "Test Case: three digits number",
			Given:       7001,
			Expected:    "yeddi min bir",
		},
		{
			Description: "Test Case: three digits number",
			Given:       8200,
			Expected:    "səkkiz min iki yüz",
		},
		{
			Description: "Test Case: three digits number",
			Given:       9526,
			Expected:    "doqquz min beş yüz iyirmi altı",
		}}
}

func ConvertIntPartsForFiveDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: five digits number"
	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       10000,
			Expected:    "on min",
		},
		{
			Description: Description,
			Given:       11111,
			Expected:    "on bir min bir yüz on bir",
		},
		{
			Description: Description,
			Given:       10192,
			Expected:    "on min bir yüz doxsan iki",
		},
		{
			Description: Description,
			Given:       27021,
			Expected:    "iyirmi yeddi min iyirmi bir",
		},
		{
			Description: Description,
			Given:       13210,
			Expected:    "on üç min iki yüz on",
		},
		{
			Description: Description,
			Given:       48550,
			Expected:    "qırx səkkiz min beş yüz əlli",
		},
		{
			Description: Description,
			Given:       95412,
			Expected:    "doxsan beş min dörd yüz on iki",
		},
		{
			Description: Description,
			Given:       79594,
			Expected:    "yetmiş doqquz min beş yüz doxsan dörd",
		},
		{
			Description: Description,
			Given:       86070,
			Expected:    "səksən altı min yetmiş",
		},
		{
			Description: Description,
			Given:       70010,
			Expected:    "yetmiş min on",
		},
		{
			Description: Description,
			Given:       81200,
			Expected:    "səksən bir min iki yüz",
		},
		{
			Description: Description,
			Given:       10001,
			Expected:    "on min bir",
		}}
}

func ConvertIntPartsForSixDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: six digits number"
	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       100000,
			Expected:    "bir yüz min",
		},
		{
			Description: Description,
			Given:       111111,
			Expected:    "bir yüz on bir min bir yüz on bir",
		},
		{
			Description: Description,
			Given:       110192,
			Expected:    "bir yüz on min bir yüz doxsan iki",
		},
		{
			Description: Description,
			Given:       127021,
			Expected:    "bir yüz iyirmi yeddi min iyirmi bir",
		},
		{
			Description: Description,
			Given:       113210,
			Expected:    "bir yüz on üç min iki yüz on",
		},
		{
			Description: Description,
			Given:       148550,
			Expected:    "bir yüz qırx səkkiz min beş yüz əlli",
		},
		{
			Description: Description,
			Given:       195412,
			Expected:    "bir yüz doxsan beş min dörd yüz on iki",
		},
		{
			Description: Description,
			Given:       179594,
			Expected:    "bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		{
			Description: Description,
			Given:       186070,
			Expected:    "bir yüz səksən altı min yetmiş",
		},
		{
			Description: Description,
			Given:       170010,
			Expected:    "bir yüz yetmiş min on",
		},
		{
			Description: Description,
			Given:       181200,
			Expected:    "bir yüz səksən bir min iki yüz",
		},
		{
			Description: Description,
			Given:       110001,
			Expected:    "bir yüz on min bir",
		}}
}

func ConvertIntPartsForSevenDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: seven digits number"
	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       1000000,
			Expected:    "bir milyon",
		},
		{
			Description: Description,
			Given:       1111111,
			Expected:    "bir milyon bir yüz on bir min bir yüz on bir",
		},
		{
			Description: Description,
			Given:       1110192,
			Expected:    "bir milyon bir yüz on min bir yüz doxsan iki",
		},
		{
			Description: Description,
			Given:       1127021,
			Expected:    "bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		{
			Description: Description,
			Given:       1113210,
			Expected:    "bir milyon bir yüz on üç min iki yüz on",
		},
		{
			Description: Description,
			Given:       1148550,
			Expected:    "bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		{
			Description: Description,
			Given:       1195412,
			Expected:    "bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		{
			Description: Description,
			Given:       1179594,
			Expected:    "bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		{
			Description: Description,
			Given:       1186070,
			Expected:    "bir milyon bir yüz səksən altı min yetmiş",
		},
		{
			Description: Description,
			Given:       1170010,
			Expected:    "bir milyon bir yüz yetmiş min on",
		},
		{
			Description: Description,
			Given:       1181200,
			Expected:    "bir milyon bir yüz səksən bir min iki yüz",
		},
		{
			Description: Description,
			Given:       1110001,
			Expected:    "bir milyon bir yüz on min bir",
		}}
}

func ConvertIntPartsForEightDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: eight digits number"
	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       11000000,
			Expected:    "on bir milyon",
		},
		{
			Description: Description,
			Given:       11111111,
			Expected:    "on bir milyon bir yüz on bir min bir yüz on bir",
		},
		{
			Description: Description,
			Given:       11110192,
			Expected:    "on bir milyon bir yüz on min bir yüz doxsan iki",
		},
		{
			Description: Description,
			Given:       11127021,
			Expected:    "on bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		{
			Description: Description,
			Given:       11113210,
			Expected:    "on bir milyon bir yüz on üç min iki yüz on",
		},
		{
			Description: Description,
			Given:       21148550,
			Expected:    "iyirmi bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		{
			Description: Description,
			Given:       31195412,
			Expected:    "otuz bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		{
			Description: Description,
			Given:       81179594,
			Expected:    "səksən bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		{
			Description: Description,
			Given:       41186070,
			Expected:    "qırx bir milyon bir yüz səksən altı min yetmiş",
		},
		{
			Description: Description,
			Given:       91170010,
			Expected:    "doxsan bir milyon bir yüz yetmiş min on",
		},
		{
			Description: Description,
			Given:       71181200,
			Expected:    "yetmiş bir milyon bir yüz səksən bir min iki yüz",
		},
		{
			Description: Description,
			Given:       61110001,
			Expected:    "altmış bir milyon bir yüz on min bir",
		}}
}

func ConvertIntPartsForNineDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: eight digits number"
	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       211000000,
			Expected:    "iki yüz on bir milyon",
		},
		{
			Description: Description,
			Given:       211111111,
			Expected:    "iki yüz on bir milyon bir yüz on bir min bir yüz on bir",
		},
		{
			Description: Description,
			Given:       411110192,
			Expected:    "dörd yüz on bir milyon bir yüz on min bir yüz doxsan iki",
		},
		{
			Description: Description,
			Given:       111127021,
			Expected:    "bir yüz on bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		{
			Description: Description,
			Given:       611113210,
			Expected:    "altı yüz on bir milyon bir yüz on üç min iki yüz on",
		},
		{
			Description: Description,
			Given:       721148550,
			Expected:    "yeddi yüz iyirmi bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		{
			Description: Description,
			Given:       831195412,
			Expected:    "səkkiz yüz otuz bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		{
			Description: Description,
			Given:       881179594,
			Expected:    "səkkiz yüz səksən bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		{
			Description: Description,
			Given:       141186070,
			Expected:    "bir yüz qırx bir milyon bir yüz səksən altı min yetmiş",
		},
		{
			Description: Description,
			Given:       191170010,
			Expected:    "bir yüz doxsan bir milyon bir yüz yetmiş min on",
		},
		{
			Description: Description,
			Given:       171181200,
			Expected:    "bir yüz yetmiş bir milyon bir yüz səksən bir min iki yüz",
		},
		{
			Description: Description,
			Given:       961110001,
			Expected:    "doqquz yüz altmış bir milyon bir yüz on min bir",
		}}
}

func ConvertIntPartsForTenDigitsNumbers() []TestCaseConvertIntPart {
	Description := "Test Case: eight digits number"
	return []TestCaseConvertIntPart{
		{
			Description: Description,
			Given:       1211000000,
			Expected:    "bir milyard iki yüz on bir milyon",
		},
		{
			Description: Description,
			Given:       3211111111,
			Expected:    "üç milyard iki yüz on bir milyon bir yüz on bir min bir yüz on bir",
		},
		{
			Description: Description,
			Given:       7411110192,
			Expected:    "yeddi milyard dörd yüz on bir milyon bir yüz on min bir yüz doxsan iki",
		},
		{
			Description: Description,
			Given:       1111127021,
			Expected:    "bir milyard bir yüz on bir milyon bir yüz iyirmi yeddi min iyirmi bir",
		},
		{
			Description: Description,
			Given:       5611113210,
			Expected:    "beş milyard altı yüz on bir milyon bir yüz on üç min iki yüz on",
		},
		{
			Description: Description,
			Given:       9721148550,
			Expected:    "doqquz milyard yeddi yüz iyirmi bir milyon bir yüz qırx səkkiz min beş yüz əlli",
		},
		{
			Description: Description,
			Given:       8831195412,
			Expected:    "səkkiz milyard səkkiz yüz otuz bir milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		{
			Description: Description,
			Given:       5881179594,
			Expected:    "beş milyard səkkiz yüz səksən bir milyon bir yüz yetmiş doqquz min beş yüz doxsan dörd",
		},
		{
			Description: Description,
			Given:       4141186070,
			Expected:    "dörd milyard bir yüz qırx bir milyon bir yüz səksən altı min yetmiş",
		},
		{
			Description: Description,
			Given:       6091170010,
			Expected:    "altı milyard doxsan bir milyon bir yüz yetmiş min on",
		},
		{
			Description: Description,
			Given:       7171181200,
			Expected:    "yeddi milyard bir yüz yetmiş bir milyon bir yüz səksən bir min iki yüz",
		},
		{
			Description: Description,
			Given:       3961110001,
			Expected:    "üç milyard doqquz yüz altmış bir milyon bir yüz on min bir",
		}}
}

func ConvertOneDigitIntoWord() []TestCaseConvertDigit2Word {
	Description := "Test Case: convertOneDigitIntoWord"

	return []TestCaseConvertDigit2Word{
		{
			Description:   Description,
			Given:         "0",
			Expected:      "sıfır",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "1",
			Expected:      "bir",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "2",
			Expected:      "iki",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "3",
			Expected:      "üç",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "4",
			Expected:      "dörd",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "5",
			Expected:      "beş",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "6",
			Expected:      "altı",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "7",
			Expected:      "yeddi",
			IsErrExpected: false,
		},
		{
			Given:         "8",
			Expected:      "səkkiz",
			IsErrExpected: false,
		},
		{
			Given:         "9",
			Expected:      "doqquz",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "",
			Expected:      "sıfır",
			IsErrExpected: true,
		},
		{
			Description:   Description,
			Given:         "?",
			Expected:      "sıfır",
			IsErrExpected: true,
		},
	}
}

func ConvertTwoDigitsIntoWord() []TestCaseConvertDigit2Word {
	Description := "Test Case: convertTwoDigitsIntoWord"

	return []TestCaseConvertDigit2Word{
		{
			Description:   Description,
			Given:         "10",
			Expected:      "on",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "11",
			Expected:      "on bir",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "20",
			Expected:      "iyirmi",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "25",
			Expected:      "iyirmi beş",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "30",
			Expected:      "otuz",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "34",
			Expected:      "otuz dörd",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "40",
			Expected:      "qırx",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "50",
			Expected:      "əlli",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "60",
			Expected:      "altmış",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "62",
			Expected:      "altmış iki",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "70",
			Expected:      "yetmiş",
			IsErrExpected: false,
		},
		{
			Given:         "80",
			Expected:      "səksən",
			IsErrExpected: false,
		},
		{
			Given:         "90",
			Expected:      "doxsan",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "99",
			Expected:      "doxsan doqquz",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "?,",
			Expected:      "sıfır",
			IsErrExpected: true,
		},
		{
			Description:   Description,
			Given:         "?",
			Expected:      "sıfır",
			IsErrExpected: true,
		},
		{
			Description:   Description,
			Given:         "",
			Expected:      "sıfır",
			IsErrExpected: true,
		},
	}
}

func ConvertThreeDigitsIntoWord() []TestCaseConvertDigit2Word {
	Description := "Test Case: convertThreeDigitsIntoWord"

	return []TestCaseConvertDigit2Word{
		{
			Description:   Description,
			Given:         "100",
			Expected:      "bir yüz",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "101",
			Expected:      "bir yüz bir",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "210",
			Expected:      "iki yüz on",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "250",
			Expected:      "iki yüz əlli",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "311",
			Expected:      "üç yüz on bir",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "304",
			Expected:      "üç yüz dörd",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "400",
			Expected:      "dörd yüz",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "509",
			Expected:      "beş yüz doqquz",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "670",
			Expected:      "altı yüz yetmiş",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "620",
			Expected:      "altı yüz iyirmi",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "780",
			Expected:      "yeddi yüz səksən",
			IsErrExpected: false,
		},
		{
			Given:         "890",
			Expected:      "səkkiz yüz doxsan",
			IsErrExpected: false,
		},
		{
			Given:         "999",
			Expected:      "doqquz yüz doxsan doqquz",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "965",
			Expected:      "doqquz yüz altmış beş",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "0",
			Expected:      "",
			IsErrExpected: true,
		},
		{
			Description:   Description,
			Given:         "?",
			Expected:      "",
			IsErrExpected: true,
		},
		{
			Description:   Description,
			Given:         "34",
			Expected:      "",
			IsErrExpected: true,
		},
	}
}

func TripleToWord() []TestCaseConvertDigit2Word {
	Description := "Test Case: tripleToWord"

	return []TestCaseConvertDigit2Word{
		{
			Description:   Description,
			Given:         "",
			Expected:      "",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "1",
			Expected:      "bir",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "12",
			Expected:      "on iki",
			IsErrExpected: false,
		},
		{
			Description:   Description,
			Given:         "250",
			Expected:      "iki yüz əlli",
			IsErrExpected: false,
		},

		{
			Description:   Description,
			Given:         "?",
			Expected:      "",
			IsErrExpected: true,
		},
		{
			Description:   Description,
			Given:         "??",
			Expected:      "",
			IsErrExpected: true,
		},
		{
			Description:   Description,
			Given:         "???",
			Expected:      "",
			IsErrExpected: true,
		},
	}
}
