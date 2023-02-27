package fixtures

type TestCaseSpellNumber struct {
	Description string
	Given       string
	Expected    string
}

func SpellNumberForNegativeIntegerNumbers() []TestCaseSpellNumber {
	testCaseDescription := "Test Case: negative integer number"

	return []TestCaseSpellNumber{
		{
			Description: testCaseDescription,
			Given:       "-10000",
			Expected:    "mənfi on min",
		},
		{
			Description: testCaseDescription,
			Given:       "-11111",
			Expected:    "mənfi on bir min bir yüz on bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-10192",
			Expected:    "mənfi on min bir yüz doxsan iki",
		},
		{
			Description: testCaseDescription,
			Given:       "-27021",
			Expected:    "mənfi iyirmi yeddi min iyirmi bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-13210",
			Expected:    "mənfi on üç min iki yüz on",
		},
		{
			Description: testCaseDescription,
			Given:       "-48550",
			Expected:    "mənfi qırx səkkiz min beş yüz əlli",
		},
		{
			Description: testCaseDescription,
			Given:       "-95412",
			Expected:    "mənfi doxsan beş min dörd yüz on iki",
		},
		{
			Description: testCaseDescription,
			Given:       "-79594",
			Expected:    "mənfi yetmiş doqquz min beş yüz doxsan dörd",
		},
		{
			Description: testCaseDescription,
			Given:       "-86070",
			Expected:    "mənfi səksən altı min yetmiş",
		},
		{
			Description: testCaseDescription,
			Given:       "-700",
			Expected:    "mənfi yeddi yüz",
		},
		{
			Description: testCaseDescription,
			Given:       "-81",
			Expected:    "mənfi səksən bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-10000",
			Expected:    "mənfi on min",
		}}
}

func SpellNumberForPositiveIntegerNumbers() []TestCaseSpellNumber {
	testCaseDescription := "Test Case: positive integer number"

	return []TestCaseSpellNumber{
		{
			Description: testCaseDescription,
			Given:       "493882371553121860890561055192142938414552660618128252927700430053",
			Expected:    "dörd yüz doxsan üç vigintilyon səkkiz yüz səksən iki novemdesilyon üç yüz yetmiş bir oktodesilyon beş yüz əlli üç septendesilyon bir yüz iyirmi bir seksdesilyon səkkiz yüz altmış kendesilyon səkkiz yüz doxsan katordesilyon beş yüz altmış bir tredesilyon əlli beş dodesilyon bir yüz doxsan iki undesilyon bir yüz qırx iki desilyon doqquz yüz otuz səkkiz nonilyon dörd yüz on dörd oktilyon beş yüz əlli iki septilyon altı yüz altmış sekstilyon altı yüz on səkkiz kvintilyon bir yüz iyirmi səkkiz kvadrilyon iki yüz əlli iki trilyon doqquz yüz iyirmi yeddi milyard yeddi yüz milyon dörd yüz otuz min əlli üç",
		},
	}
}

func SpellNumberForNegativeFloatingPointNumbers() []TestCaseSpellNumber {
	testCaseDescription := "Test Case: negative floating point number"

	return []TestCaseSpellNumber{
		{
			Description: testCaseDescription,
			Given:       "-12.1",
			Expected:    "mənfi on iki tam onda bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-111.11",
			Expected:    "mənfi bir yüz on bir tam yüzdə on bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-10.192",
			Expected:    "mənfi on tam bir mində bir yüz doxsan iki",
		},
		{
			Description: testCaseDescription,
			Given:       "-2.7021",
			Expected:    "mənfi iki tam on mində yeddi min iyirmi bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.13211",
			Expected:    "mənfi sıfır tam yüz mində on üç min iki yüz on bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.248551",
			Expected:    "mənfi sıfır tam bir milyonda iki yüz qırx səkkiz min beş yüz əlli bir",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.6195412",
			Expected:    "mənfi sıfır tam on milyonda altı milyon bir yüz doxsan beş min dörd yüz on iki",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.79354594",
			Expected:    "mənfi sıfır tam yüz milyonda yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.179354594",
			Expected:    "mənfi sıfır tam bir milyardda bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.1179354594",
			Expected:    "mənfi sıfır tam on milyardda bir milyard bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.41179354594",
			Expected:    "mənfi sıfır tam yüz milyardda qırx bir milyard bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		},
		{
			Description: testCaseDescription,
			Given:       "-0.541179354594",
			Expected:    "mənfi sıfır tam bir trilyonda beş yüz qırx bir milyard bir yüz yetmiş doqquz milyon üç yüz əlli dörd min beş yüz doxsan dörd",
		}}
}
