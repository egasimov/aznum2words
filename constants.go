package aznum2words

const (
	MaxNumberDigitCount int = 66
	MaxNumberScaleCount int = 15
)

const (

	//single digits numbers
	ZeroAsString  string = "sıfır"  // 0
	OneAsString   string = "bir"    // 1
	TwoAsString   string = "iki"    // 2
	ThreeAsString string = "üç"     // 3
	FourAsString  string = "dörd"   // 4
	FiveAsString  string = "beş"    // 5
	SixAsString   string = "altı"   // 6
	SevenAsString string = "yeddi"  // 7
	EightAsString string = "səkkiz" // 8
	NineAsString  string = "doqquz" // 9

	//two digits numbers
	TenAsString     string = "on"     // 10
	TwentyAsString  string = "iyirmi" // 20
	ThirtyAsString  string = "otuz"   // 30
	FortyAsString   string = "qırx"   // 40
	FiftyAsString   string = "əlli"   // 50
	SixtyAsString   string = "altmış" // 60
	SeventyAsString string = "yetmiş" // 70
	EightyAsString  string = "səksən" // 80
	NinetyAsString  string = "doxsan" // 90

	//three digits numbers
	HundredAsString string = "yüz" // 10^2

	//others
	ThousandAsString        string = "min"            // 10^3
	MillionAsString         string = "milyon"         // 10^6
	BillionAsString         string = "milyard"        // 10^9
	TrillionAsString        string = "trilyon"        // 10^12
	QuadrillionAsString     string = "kvadrilyon"     // 10^15
	QuintillionAsString     string = "kvintilyon"     // 10^18
	SextillionAsString      string = "sekstilyon"     // 10^21
	SeptillionAsString      string = "septilyon"      // 10^24
	OctillionAsString       string = "oktilyon"       // 10^27
	NonillionAsString       string = "nonilyon"       // 10^30
	DecillionAsString       string = "desilyon"       //10^33
	UndecillionAsString     string = "undesilyon"     //10^36
	DodecillionAsString     string = "dodesilyon"     //10^39
	TredecillionAsString    string = "tredesilyon"    //10^42
	CathodecillionAsString  string = "katordesilyon"  //10^45
	KendecillionAsString    string = "kendesilyon"    //10^48
	SexdecillionAsString    string = "seksdesilyon"   //10^51
	SeptendecillionAsString string = "septendesilyon" //10^54
	OctodecillionAsString   string = "oktodesilyon"   //10^57
	NovedesillionAsString   string = "novemdesilyon"  //10^60
	VigintillionAsString    string = "vigintilyon"    //10^63
)

const (
	NegativeAsWord  string = "mənfi" // -
	SeparatorAsWord string = "tam"
)
