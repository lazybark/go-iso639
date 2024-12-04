package iso639

var (
	Abkhazian = Language{
		Codes:       []Code{AB, ABK},
		EnglishName: "Abkhazian",
		Scope:       Individual,
		Regions:     []Region{Eurasia, Georgia, Caucasus, Europe},
		Scripts:     []Script{LatinScript, Cyrillic, GeorgianScript},
		Family:      NorthwestCaucasian,
		NativeNames: []string{"Аҧсуа", "Apsua", "აფსუა", "Abkhaz"}}
	Afar = Language{
		Codes:       []Code{AA, AAR},
		EnglishName: "Afar",
		Scope:       Individual,
		Regions:     []Region{Africa},
		Type:        Living,
		Scripts:     []Script{LatinScript},
		Family:      Afroasiatic,
		NativeNames: []string{"Qafar af"}}
	Afrikaans = Language{
		Codes:       []Code{AF, AFR},
		EnglishName: "Afrikaans",
		Scope:       Individual,
	}
	Akan = Language{
		Codes:       []Code{AK, AKA},
		EnglishName: "Akan",
		Scope:       Macro,
		NativeNames: []string{"Ákán"},
		Variants: []Variant{
			{Code: FAT, EnglishName: "Fanti"},
		}}
	Albanian = Language{
		Codes:       []Code{SQ, SQI, ALB},
		EnglishName: "Albanian",
		Scope:       Macro,
		NativeNames: []string{"Shqip"},
		Variants: []Variant{
			{Code: AAE, EnglishName: "Arbëreshë Albanian"},
			{Code: AAT, EnglishName: "Arvanitika Albanian"},
			{Code: ALN, EnglishName: "Gheg Albanian"},
			{Code: ALS, EnglishName: "Tosk Albanian"},
		}}
	Arabic = Language{
		Codes:       []Code{AR, ARA},
		EnglishName: "Arabic",
		Type:        Living,
		Scope:       Macro,
		Scripts:     []Script{ArabicScript, LatinScript},
		Family:      Afroasiatic,
		Regions:     []Region{Africa, Asia, Eurasia, MiddleEast},
		NativeNames: []string{"اَلْعَرَبِيَّةُ", "al-ʿarabiyyah"},
		Variants: []Variant{
			{Code: AAO, EnglishName: "Algerian Saharan Arabic"},
			{Code: ABH, EnglishName: "Tajiki Arabic"},
			{Code: ABV, EnglishName: "Baharna Arabic"},
			{Code: ACM, EnglishName: "Mesopotamian Arabic"},
			{Code: ACQ, EnglishName: "Ta'izzi-Adeni Arabic"},
			{Code: ACW, EnglishName: "Hijazi Arabic"},
			{Code: ACX, EnglishName: "Omani Arabic"},
			{Code: ACY, EnglishName: "Cypriot Arabic"},
			{Code: ADF, EnglishName: "Dhofari Arabic"},
			{Code: AEB, EnglishName: "Tunisian Arabic"},
			{Code: AEC, EnglishName: "Saidi Arabic"},
			{Code: AFB, EnglishName: "Gulf Arabic"},
			{Code: APC, EnglishName: "Levantine Arabic"},
			{Code: APD, EnglishName: "Sudanese Arabic"},
			{Code: ARB, EnglishName: "Standard Arabic"},
			{Code: ARQ, EnglishName: "Algerian Arabic"},
			{Code: ARS, EnglishName: "Najdi Arabic"},
			{Code: ARY, EnglishName: "Moroccan Arabic"},
			{Code: ARZ, EnglishName: "Egyptian Arabic"},
			{Code: AUZ, EnglishName: "Uzbeki Arabic"},
			{Code: AVL, EnglishName: "Eastern Egyptian Bedawi Arabic"},
			{Code: AYH, EnglishName: "Hadrami Arabic"},
			{Code: AYL, EnglishName: "Libyan Arabic"},
			{Code: AYN, EnglishName: "Sanaani Arabic"},
			{Code: AYP, EnglishName: "North Mesopotamian Arabic"},
			{Code: PGA, EnglishName: "Sudanese Creole Arabic"},
			{Code: SHU, EnglishName: "Chadian Arabic"},
			{Code: SSH, EnglishName: "Shihhi Arabic"},
		}}
	Aragonese = Language{
		Codes:       []Code{AN, ARG},
		EnglishName: "Aragonese",
		Scope:       Individual,
		NativeNames: []string{"Aragonés"}}
	Amharic = Language{
		Codes:       []Code{AM, AMH},
		EnglishName: "Amharic",
		Scope:       Individual,
		NativeNames: []string{"አማርኛ", "Amarəñña"}}
	Armenian = Language{
		Codes:       []Code{HY, HYE, ARM},
		EnglishName: "Armenian",
		Scope:       Individual,
		NativeNames: []string{"Հայերեն", "Hayeren"}}
	Assamese = Language{
		Codes:       []Code{AS, ASM},
		EnglishName: "Assamese",
		Scope:       Individual,
		NativeNames: []string{"অসমীয়া", "Ôxômiya"}}
	Avaric = Language{
		Codes:       []Code{AV, AVA},
		EnglishName: "Avaric",
		Scope:       Individual,
		NativeNames: []string{"Авар мацӏ", "اوار ماض", "Avar maz"}}
	Avestan = Language{
		Codes:       []Code{AE, AVE},
		EnglishName: "Avestan",
		Scope:       Individual,
		NativeNames: []string{"Upastawakaēna"}}
	Aymara = Language{
		Codes:       []Code{AY, AYM},
		EnglishName: "Aymara",
		Scope:       Macro,
		NativeNames: []string{"Aymara"},
		Variants: []Variant{
			{Code: AYC, EnglishName: "Southern Aymara"},
			{Code: AYR, EnglishName: "Central Aymara"}}}
	Azerbaijani = Language{
		Codes:       []Code{AZ, AZE},
		EnglishName: "Azerbaijani",
		Scope:       Macro,
		NativeNames: []string{"Azərbaycan dili", "آذربایجان دیلی", "Азәрбајҹан дили"},
		Variants: []Variant{
			{Code: AZJ, EnglishName: "North Azerbaijani"},
			{Code: AZB, EnglishName: "South Azerbaijani"}}}
	Bambara = Language{
		Codes:       []Code{BM, BAM},
		EnglishName: "Bambara",
		Scope:       Individual,
		NativeNames: []string{"بَمَنَنكَن", "Bamanankan"}}
	Bashkir = Language{
		Codes:       []Code{BA, BAK},
		EnglishName: "Bashkir",
		Scope:       Individual,
		NativeNames: []string{"Башҡорт теле", "Başqort tele"}}
	Basque = Language{
		Codes:       []Code{EU, EUS, BAQ},
		EnglishName: "Basque",
		Scope:       Individual,
		NativeNames: []string{"Euskara"}}
	Belarusian = Language{
		Codes:       []Code{BE, BEL},
		EnglishName: "Belarusian",
		Scope:       Individual,
		NativeNames: []string{"Беларуская мова", "Belaruskaâ mova"}}
	Bengali = Language{
		Codes:       []Code{BN, BEN},
		EnglishName: "Bengali",
		Scope:       Individual,
		NativeNames: []string{"বাংলা", "Bāŋlā"}}
	Bislama = Language{
		Codes:       []Code{BI, BIS},
		EnglishName: "Bislama",
		Scope:       Individual,
	}
	Bosnian = Language{
		Codes:       []Code{BS, BOS},
		EnglishName: "Bosnian",
		Scope:       Individual,
		NativeNames: []string{"Босански", "Bosanski"}}
	Breton = Language{
		Codes:       []Code{BR, BRE},
		EnglishName: "Breton",
		Scope:       Individual,
		NativeNames: []string{"Brezhoneg"}}
	Bulgarian = Language{
		Codes:       []Code{BG, BUL},
		EnglishName: "Bulgarian",
		Scope:       Individual,
		NativeNames: []string{"Български", "Bulgarski"}}
	Burmese = Language{
		Codes:       []Code{MY, MYA, BUR},
		EnglishName: "Burmese",
		Scope:       Individual,
		NativeNames: []string{"မြန်မာစာ", "Mrãmācā"}}
	Catalan = Language{
		Codes:       []Code{CA, CAT},
		EnglishName: "Catalan",
		Scope:       Individual,
		NativeNames: []string{"Valencian", "Català", "Valencià"}}
	Chamorro = Language{
		Codes:       []Code{CH, CHA},
		EnglishName: "Chamorro",
		Scope:       Individual,
		NativeNames: []string{"Finu' Chamoru"}}
	Chechen = Language{
		Codes:       []Code{CE, CHE},
		EnglishName: "Chechen",
		Scope:       Individual,
		NativeNames: []string{"Нохчийн мотт", "Noxçiyn mott", "Chechnyan", "Chechnian"}}
	Chichewa = Language{
		Codes:       []Code{NY, NYA},
		EnglishName: "Chichewa",
		Scope:       Individual,
		NativeNames: []string{"Chewa", "Nyanja", "Chinyanja"}}
	Chinese = Language{
		Codes:       []Code{ZH, ZHO, CHI},
		EnglishName: "Chinese",
		Type:        Living,
		Scope:       Macro,
		Scripts:     []Script{ChineseScript, LatinScript},
		Regions:     []Region{Asia, Eurasia, Oceania},
		NativeNames: []string{"中文", "Zhōngwén", "汉语", "漢語", "Hànyǔ"},
		Variants: []Variant{
			{Code: CDO, EnglishName: "Min Dong Chinese"},
			{Code: CJY, EnglishName: "Jinyu Chinese"},
			{Code: CMN, EnglishName: "Mandarin Chinese"},
			{Code: CNP, EnglishName: "Northern Ping Chinese"},
			{Code: CPX, EnglishName: "Pu-Xian Chinese"},
			{Code: CSP, EnglishName: "Southern Ping Chinese"},
			{Code: CZH, EnglishName: "Huizhou Chinese"},
			{Code: CZO, EnglishName: "Min Zhong Chinese"},
			{Code: GAN, EnglishName: "Gan Chinese"},
			{Code: HAK, EnglishName: "Hakka Chinese"},
			{Code: HNM, EnglishName: "Hainanese"},
			{Code: HSN, EnglishName: "Xiang Chinese"},
			{Code: LUH, EnglishName: "Leizhou Chinese"},
			{Code: LZH, EnglishName: "Literary Chinese"},
			{Code: MNP, EnglishName: "Min Bei Chinese"},
			{Code: NAN, EnglishName: "Min Nan Chinese"},
			{Code: SJC, EnglishName: "Shaojiang Chinese"},
			{Code: WUU, EnglishName: "Wu Chinese"},
			{Code: YUE, EnglishName: "Yue Chinese"}}}
	ChurchSlavic = Language{
		Codes:       []Code{CU, CHU},
		EnglishName: "Church Slavonic",
		Scope:       Individual,
		NativeNames: []string{"Славе́нскїй ѧ҆зы́къ"}}
	Chuvash = Language{
		Codes:       []Code{CV, CHV},
		EnglishName: "Chuvash",
		Scope:       Individual,
		NativeNames: []string{"Чăвашла", "Çăvaşla"}}
	Cornish = Language{
		Codes:       []Code{KW, COR},
		EnglishName: "Cornish",
		Scope:       Individual,
		NativeNames: []string{"Kernowek"}}
	Corsican = Language{
		Codes:       []Code{CO, COS},
		EnglishName: "Corsican",
		Scope:       Individual,
		NativeNames: []string{"Corsu"}}
	Cree = Language{
		Codes:       []Code{CR, CRE},
		EnglishName: "Cree",
		Scope:       Macro,
		NativeNames: []string{"ᓀᐦᐃᔭᐁᐧᐃᐧᐣ", "Nehiyawewin"},
		Variants: []Variant{
			{Code: CRJ, EnglishName: "Southern East Cree"},
			{Code: CRK, EnglishName: "Plains Cree"},
			{Code: CRL, EnglishName: "Northern East Cree"},
			{Code: CRM, EnglishName: "Moose Cree"},
			{Code: CSW, EnglishName: "Swampy Cree"},
			{Code: CWD, EnglishName: "Woods Cree"}}}
	Croatian = Language{
		Codes:       []Code{HR, HRV},
		EnglishName: "Croatian",
		Scope:       Individual,
		NativeNames: []string{"Hrvatski", "Crovatian"}}
	Czech = Language{
		Codes:       []Code{CS, CES, CZE},
		EnglishName: "Czech",
		Scope:       Individual,
		NativeNames: []string{"Čeština", "Czechian"}}
	Danish = Language{
		Codes:       []Code{DA, DAN},
		EnglishName: "Danish",
		Scope:       Individual,
		NativeNames: []string{"Dansk"}}
	Divehi = Language{
		Codes:       []Code{DV, DIV},
		EnglishName: "Divehi",
		Scope:       Individual,
		NativeNames: []string{"Dhivehi", "Maldivian", "ދިވެހި "}}
	Dutch = Language{
		Codes:       []Code{NL, NLD, DUT},
		EnglishName: "Dutch",
		Scope:       Individual,
		NativeNames: []string{"Flemish", "Nederlands"},
		Variants: []Variant{
			{Code: VLS, EnglishName: "West Flemish"}}}
	Dzongkha = Language{
		Codes:       []Code{DZ, DZO},
		EnglishName: "Dzongkha",
		Scope:       Individual,
		NativeNames: []string{"རྫོང་ཁ་", "Bhutanese"}}
	English = Language{
		Codes:       []Code{EN, ENG},
		EnglishName: "English",
		Type:        Living,
		Scope:       Individual,
		Scripts:     []Script{LatinScript, AngloSaxonRunes},
		Family:      IndoEuropean,
		Regions:     []Region{Europe, Oceania, NorthAmerica, SouthAmerica, Africa, Asia, Eurasia, India},
	}
	Esperanto = Language{
		Codes:       []Code{EO, EPO},
		EnglishName: "Esperanto",
		Scope:       Individual,
	}
	Estonian = Language{
		Codes:       []Code{ET, EST},
		EnglishName: "Estonian",
		Scope:       Macro,
		NativeNames: []string{"Eesti keel"},
		Variants: []Variant{
			{Code: VRO, EnglishName: "Võro"},
			{Code: EKK, EnglishName: "Standard Estonian"}}}
	Ewe = Language{
		Codes:       []Code{EE, EWE},
		EnglishName: "Ewe",
		Scope:       Individual,
		NativeNames: []string{"Èʋegbe"}}
	Faroese = Language{
		Codes:       []Code{FO, FAO},
		EnglishName: "Faroese",
		Scope:       Individual,
		NativeNames: []string{"Føroyskt"}}
	Fijian = Language{
		Codes:       []Code{FJ, FIJ},
		EnglishName: "Fijian",
		Scope:       Individual,
		NativeNames: []string{"Na Vosa Vakaviti"}}
	Finnish = Language{
		Codes:       []Code{FI, FIN},
		EnglishName: "Finnish",
		Scope:       Individual,
		NativeNames: []string{"Suomi"}}
	French = Language{
		Codes:       []Code{FR, FRA, FRE},
		EnglishName: "French",
		Scope:       Individual,
		NativeNames: []string{"Français"}}
	WesternFrisian = Language{
		Codes:       []Code{FY, FRY},
		EnglishName: "Western Frisian",
		Scope:       Individual,
		NativeNames: []string{"Frysk", "West Frisian", "Frisian", "Fries"}}
	Fulah = Language{
		Codes:       []Code{FF, FUL},
		EnglishName: "Fulah",
		Scope:       Macro,
		NativeNames: []string{"𞤊𞤵𞤤𞤬𞤵𞤤𞤣𞤫", "ࢻُلْࢻُلْدٜ", "Fulfulde", "𞤆𞤵𞤤𞤢𞥄𞤪", "ݒُلَارْ", "Pulaar"},
		Variants: []Variant{
			{Code: FFM, EnglishName: "Maasina Fulfulde"},
			{Code: FUB, EnglishName: "Adamawa Fulfulde"},
			{Code: FUC, EnglishName: "Pulaar"},
			{Code: FUE, EnglishName: "Borgu Fulfulde"},
			{Code: FUF, EnglishName: "Pular"},
			{Code: FUH, EnglishName: "Western Niger Fulfulde"},
			{Code: FUI, EnglishName: "Bagirmi Fulfulde"},
			{Code: FUQ, EnglishName: "Central-Eastern Niger Fulfulde"},
			{Code: FUV, EnglishName: "Nigerian Fulfulde"}}}
	Gaelic = Language{
		Codes:       []Code{GD, GLA},
		EnglishName: "Gaelic",
		Scope:       Individual,
		NativeNames: []string{"Scottish Gaelic", "Gàidhlig"}}
	Galician = Language{
		Codes:       []Code{GL, GLG},
		EnglishName: "Galician",
		Scope:       Individual,
		NativeNames: []string{"Galego"}}
	Ganda = Language{
		Codes:       []Code{LG, LUG},
		EnglishName: "Ganda",
		Scope:       Individual,
		NativeNames: []string{"Luganda"}}
	Georgian = Language{
		Codes:       []Code{KA, KAT},
		EnglishName: "Georgian",
		Scope:       Individual,
		NativeNames: []string{"ქართული", "Kharthuli"}}
	German = Language{
		Codes:       []Code{DE, DEU},
		EnglishName: "German",
		Scope:       Individual,
		NativeNames: []string{"Deutsch"}}
	Greek = Language{
		Codes:       []Code{EL, ELL},
		EnglishName: "Greek",
		Scope:       Individual,
		NativeNames: []string{"Νέα Ελληνικά", "Néa Ellêniká"}}
	Greenlandic = Language{
		Codes:       []Code{KL, KAL},
		EnglishName: "Greenlandic",
		Scope:       Individual,
		NativeNames: []string{"Kalaallisut"}}
	Guarani = Language{
		Codes:       []Code{GN, GRN},
		EnglishName: "Guarani",
		Scope:       Macro,
		NativeNames: []string{"Avañe'ẽ"},
		Variants: []Variant{
			{Code: GNW, EnglishName: "Western Bolivian Guaraní"},
			{Code: GUG, EnglishName: "Paraguayan Guaraní"},
			{Code: GUI, EnglishName: "Eastern Bolivian Guaraní"},
			{Code: GUN, EnglishName: "Mbyá Guaraní"},
			{Code: NHD, EnglishName: "Chiripá"}}}
	Gujarati = Language{
		Codes:       []Code{GU, GUJ},
		EnglishName: "Gujarati",
		Scope:       Individual,
		NativeNames: []string{"ગુજરાતી", "Gujarātī"}}
	Haitian = Language{
		Codes:       []Code{HT, HAT},
		EnglishName: "Haitian",
		Scope:       Individual,
		NativeNames: []string{"Haitian Creole", "Kreyòl ayisyen"}}
	Hausa = Language{
		Codes:       []Code{HA, HAU},
		EnglishName: "Hausa",
		Scope:       Individual,
		NativeNames: []string{"هَرْشٜن هَوْس", "halshen Hausa"}}
	Hebrew = Language{
		Codes:       []Code{HE, HEB},
		EnglishName: "Hebrew",
		Scope:       Individual,
		NativeNames: []string{"עברית", "Ivrit"}}
	Herero = Language{
		Codes:       []Code{HZ, HER},
		EnglishName: "Herero",
		Scope:       Individual,
		NativeNames: []string{"Otjiherero"}}
	Hindi = Language{
		Codes:       []Code{HI, HIN},
		EnglishName: "Hindi",
		Scope:       Individual,
		NativeNames: []string{"हिन्दी", "Hindī"}}
	HiriMotu = Language{
		Codes:       []Code{HO, HMO},
		EnglishName: "Hiri Motu",
		Scope:       Individual,
		NativeNames: []string{"Police Motu"}}
	Hungarian = Language{
		Codes:       []Code{HU, HUN},
		EnglishName: "Hungarian",
		Scope:       Individual,
		NativeNames: []string{"Magyar nyelv"}}
	Icelandic = Language{
		Codes:       []Code{IS, ISL},
		EnglishName: "Icelandic",
		Scope:       Individual,
		NativeNames: []string{"Íslenska"}}
	Ido = Language{
		Codes:       []Code{IO, IDO},
		EnglishName: "Ido",
		Scope:       Individual,
	}
	Igbo = Language{
		Codes:       []Code{IG, IBO},
		EnglishName: "Igbo",
		Scope:       Individual,
		NativeNames: []string{"ásụ̀sụ́ Ìgbò"}}
	Indonesian = Language{
		Codes:       []Code{ID, IND},
		EnglishName: "Indonesian",
		Scope:       Individual,
		NativeNames: []string{"bahasa Indonesia"}}
	Interlingua = Language{
		Codes:       []Code{IA, INA},
		EnglishName: "Interlingua",
		Scope:       Individual,
	}
	Interlingue = Language{
		Codes:       []Code{IE, ILE},
		EnglishName: "Interlingue",
		Scope:       Individual,
		NativeNames: []string{"Occidental"}}
	Inuktitut = Language{
		Codes:       []Code{IU, IKU},
		EnglishName: "Inuktitut",
		Scope:       Macro,
		NativeNames: []string{"ᐃᓄᒃᑎᑐᑦ"},
		Variants: []Variant{
			{Code: IKE, EnglishName: "Eastern Canadian Inuktitut"},
			{Code: IKT, EnglishName: "Inuinnaqtun"}}}
	Inupiaq = Language{
		Codes:       []Code{IK, IPK},
		EnglishName: "Inupiaq",
		Scope:       Macro,
		NativeNames: []string{"Iñupiaq"},
		Variants: []Variant{
			{Code: ESI, EnglishName: "North Alaskan Inupiatun"},
			{Code: ESK, EnglishName: "Northwest Alaska Inupiatun"}}}
	Irish = Language{
		Codes:       []Code{GA, GLE},
		EnglishName: "Irish",
		Scope:       Individual,
		NativeNames: []string{"Gaeilge"}}
	Italian = Language{
		Codes:       []Code{IT, ITA},
		EnglishName: "Italian",
		Scope:       Individual,
		NativeNames: []string{"Italiano"}}
	Japanese = Language{
		Codes:       []Code{JA, JPN},
		EnglishName: "Japanese",
		Scope:       Individual,
		NativeNames: []string{"日本語", "Nihongo"}}
	Javanese = Language{
		Codes:       []Code{JV, JAV},
		EnglishName: "Javanese",
		Scope:       Individual,
		NativeNames: []string{"ꦧꦱꦗꦮ", "basa Jawa"}}
	Kannada = Language{
		Codes:       []Code{KN, KAN},
		EnglishName: "Kannada",
		Scope:       Individual,
		NativeNames: []string{"ಕನ್ನಡ", "Kannaḍa"}}
	Kanuri = Language{
		Codes:       []Code{KR, KAU},
		EnglishName: "Kanuri",
		Scope:       Macro,
		NativeNames: []string{"كَنُرِيِه", "Kànùrí"},
		Variants: []Variant{
			{Code: KBY, EnglishName: "Manga Kanuri"},
			{Code: KNC, EnglishName: "Central Kanuri"},
			{Code: KRT, EnglishName: "Tumari Kanuri"}}}
	Kashmiri = Language{
		Codes:       []Code{KS, KAS},
		EnglishName: "Kashmiri",
		Scope:       Individual,
		NativeNames: []string{"कॉशुर", "كأشُر", "Kosher"}}
	Kazakh = Language{
		Codes:       []Code{KK, KAZ},
		EnglishName: "Kazakh",
		Scope:       Individual,
		NativeNames: []string{"Қазақша", "Qazaqşa"}}
	CentralKhmer = Language{
		Codes:       []Code{KM, KHM},
		EnglishName: "Central Khmer",
		NativeNames: []string{"ខេមរភាសា", "Khémôrôphéasa", "Khmer", "Cambodian"}}
	Kikuyu = Language{
		Codes:       []Code{KI, KIK},
		EnglishName: "Kikuyu",
		Scope:       Individual,
		NativeNames: []string{"Gikuyu", "Gĩgĩkũyũ"}}
	Kinyarwanda = Language{
		Codes:       []Code{RW, KIN},
		EnglishName: "Kinyarwanda",
		Scope:       Individual,
		NativeNames: []string{"Ikinyarwanda"}}
	Kirghiz = Language{
		Codes:       []Code{KY, KIR},
		EnglishName: "Kirghiz",
		Scope:       Individual,
		NativeNames: []string{"Kyrgyz", "Кыргызча", "Kırgızça"}}
	Komi = Language{
		Codes:       []Code{KV, KOM},
		EnglishName: "Komi",
		Scope:       Macro,
		NativeNames: []string{"Коми кыв", "Zyran", "Zyrian"},
		Variants: []Variant{
			{Code: KOI, EnglishName: "Komi-Permyak"},
			{Code: KPV, EnglishName: "Komi-Zyrian"}}}
	Kongo = Language{
		Codes:       []Code{KG, KON},
		EnglishName: "Kongo",
		Scope:       Macro,
		NativeNames: []string{"Kikongo"},
		Variants: []Variant{
			{Code: KNG, EnglishName: "Koongo"},
			{Code: KWY, EnglishName: "San Salvador Kongo"},
			{Code: LDI, EnglishName: "Laari"}}}
	Korean = Language{
		Codes:       []Code{KO, KOR},
		EnglishName: "Korean",
		Scope:       Individual,
		NativeNames: []string{"한국어", "조선말", "Hangugeo", "Chosŏnmal"}}
	Kuanyama = Language{
		Codes:       []Code{KJ, KUA},
		EnglishName: "Kuanyama",
		Scope:       Individual,
		NativeNames: []string{"Kwanyama"}}
	Kurdish = Language{
		Codes:       []Code{KU, KUR},
		EnglishName: "Kurdish",
		Type:        Living,
		Scope:       Macro,
		Family:      IndoEuropean,
		Scripts:     []Script{LatinScript, ArabicScript, Cyrillic},
		Regions:     []Region{Asia, Europe, Eurasia, Caucasus},
		NativeNames: []string{"Kurdî", "کوردی"},
		Variants: []Variant{
			{Code: CKB, EnglishName: "Central Kurdish"},
			{Code: KMR, EnglishName: "Northern Kurdish"},
			{Code: SDH, EnglishName: "Southern Kurdish"}}}
	Lao = Language{
		Codes:       []Code{LO, LAO},
		EnglishName: "Lao",
		Scope:       Individual,
		NativeNames: []string{"ພາສາລາວ", "phasa Lao"}}
	Latin = Language{
		Codes:       []Code{LA, LAT},
		EnglishName: "Latin",
		Scope:       Individual,
		NativeNames: []string{"Latinum"}}
	Latvian = Language{
		Codes:       []Code{LV, LAV},
		EnglishName: "Latvian",
		Scope:       Macro,
		NativeNames: []string{"Latviski"},
		Variants: []Variant{
			{Code: LTG, EnglishName: "Latgalian"},
			{Code: LVS, EnglishName: "Standard Latvian"}}}
	Limburgan = Language{
		Codes:       []Code{LI, LIM},
		EnglishName: "Limburgan",
		Scope:       Individual,
		NativeNames: []string{"Limburger", "Limburgish", "Lèmburgs"}}
	Lingala = Language{
		Codes:       []Code{LN, LIN},
		EnglishName: "Lingala",
		Scope:       Individual,
		NativeNames: []string{"Lingála"}}
	Lithuanian = Language{
		Codes:       []Code{LT, LIT},
		EnglishName: "Lithuanian",
		Scope:       Individual,
		NativeNames: []string{"Lietuviškai"}}
	LubaKatanga = Language{
		Codes:       []Code{LU, LUB},
		EnglishName: "Luba-Katanga",
		Scope:       Individual,
		NativeNames: []string{"Kiluba", "Luba-Shaba"}}
	Luxembourgish = Language{
		Codes:       []Code{LB, LTZ},
		EnglishName: "Luxembourgish",
		Scope:       Individual,
		NativeNames: []string{"Letzeburgesch", "Lëtzebuergesch", "Luxembourgian"}}
	Macedonian = Language{
		Codes:       []Code{MK, MKD, MAC},
		EnglishName: "Macedonian",
		Scope:       Individual,
		NativeNames: []string{"Македонски", "Makedonski"}}
	Malagasy = Language{
		Codes:       []Code{MG, MLG},
		EnglishName: "Malagasy",
		Scope:       Macro,
		NativeNames: []string{"مَلَغَسِ"},
		Variants: []Variant{
			{Code: BHR, EnglishName: "Bara Malagasy"},
			{Code: BMM, EnglishName: "Northern Betsimisaraka Malagasy"},
			{Code: BZC, EnglishName: "Southern Betsimisaraka Malagasy"},
			{Code: MSH, EnglishName: "Masikoro Malagasy"},
			{Code: PLT, EnglishName: "Plateau Malagasy"},
			{Code: SKG, EnglishName: "Sakalava Malagasy"},
			{Code: TDX, EnglishName: "Tandroy-Mahafaly Malagasy"},
			{Code: TKG, EnglishName: "Tesaka Malagasy"},
			{Code: TXY, EnglishName: "Tanosy Malagasy"},
			{Code: XMV, EnglishName: "Antankarana Malagasy"},
			{Code: XMW, EnglishName: "Tsimihety Malagasy"}}}
	Malay = Language{
		Codes:       []Code{MS, MSA},
		EnglishName: "Malay",
		Scope:       Macro,
		NativeNames: []string{"بهاس ملايو", "bahasa Melayu"},
		Variants: []Variant{
			{Code: BJN, EnglishName: "Banjar"},
			{Code: BTJ, EnglishName: "Bacanese Malay"},
			{Code: BVE, EnglishName: "Berau Malay"},
			{Code: BVU, EnglishName: "Bukit Malay"},
			{Code: COA, EnglishName: "Cocos Islands Malay"},
			{Code: DUP, EnglishName: "Duano"},
			{Code: HJI, EnglishName: "Haji"},
			{Code: JAK, EnglishName: "Jakun"},
			{Code: JAX, EnglishName: "Jambi Malay"},
			{Code: KVB, EnglishName: "Kubu"},
			{Code: KVR, EnglishName: "Kerinci"},
			{Code: KXD, EnglishName: "Brunei"},
			{Code: LCE, EnglishName: "Loncong"},
			{Code: LCF, EnglishName: "Lubu"},
			{Code: LIW, EnglishName: "Col"},
			{Code: MAX, EnglishName: "North Moluccan Malay"},
			{Code: MEO, EnglishName: "Kedah Malay"},
			{Code: MFA, EnglishName: "Pattani Malay"},
			{Code: MFB, EnglishName: "Bangka"},
			{Code: MIN, EnglishName: "Minangkabau"},
			{Code: MQG, EnglishName: "Kota Bangun Kutai Malay"},
			{Code: MSI, EnglishName: "Sabah Malay"},
			{Code: MUI, EnglishName: "Musi"},
			{Code: ORN, EnglishName: "Orang Kanaq"},
			{Code: ORS, EnglishName: "Orang Seletar"},
			{Code: PEL, EnglishName: "Pekal"},
			{Code: PSE, EnglishName: "Central Malay"},
			{Code: TMW, EnglishName: "Temuan"},
			{Code: URK, EnglishName: "Urak Lawoi'"},
			{Code: VKK, EnglishName: "Kaur"},
			{Code: VKT, EnglishName: "Tenggarong Kutai Malay"},
			{Code: XMM, EnglishName: "Manado Malay"},
			{Code: ZLM, EnglishName: "Malay (individual language)"},
			{Code: ZMI, EnglishName: "Negeri Sembilan Malay"},
			{Code: ZSM, EnglishName: "Standard Malay"}}}
	Malayalam = Language{
		Codes:       []Code{ML, MAL},
		EnglishName: "Malayalam",
		Scope:       Individual,
		NativeNames: []string{"മലയാളം", "Malayāļã"}}
	Maltese = Language{
		Codes:       []Code{MT, MLT},
		EnglishName: "Maltese",
		Scope:       Individual,
		NativeNames: []string{"Malti"}}
	Manx = Language{
		Codes:       []Code{GV, GLV},
		EnglishName: "Manx",
		Scope:       Individual,
		NativeNames: []string{"Gaelg", "Gailck"}}
	Maori = Language{
		Codes:       []Code{MI, MRI, MAO},
		EnglishName: "Maori",
		Scope:       Individual,
		NativeNames: []string{"reo Māori"}}
	Marathi = Language{
		Codes:       []Code{MR, MAR},
		EnglishName: "Marathi",
		Scope:       Individual,
		NativeNames: []string{"मराठी", "Marāṭhī", "Maharashtran"}}
	Marshallese = Language{
		Codes:       []Code{MH, MAH},
		EnglishName: "Marshallese",
		Scope:       Individual,
		NativeNames: []string{"kajin M̧ajel‌̧", "Ebon"}}
	Mongolian = Language{
		Codes:       []Code{MN, MON},
		EnglishName: "Mongolian",
		Type:        Living,
		Scope:       Macro,
		Family:      Mongolic,
		Regions:     []Region{Asia},
		Scripts:     []Script{MongolianScript, Cyrillic},
		NativeNames: []string{"ᠮᠣᠩᠭᠣᠯ ᠬᠡᠯᠡ", "Монгол хэл", "Mongol xel"},
		Variants: []Variant{
			{Code: KHK, EnglishName: "Halh Mongolian"},
			{Code: MVF, EnglishName: "Peripheral Mongolian"}}}
	Nauru = Language{
		Codes:       []Code{NA, NAU},
		EnglishName: "Nauru",
		Scope:       Individual,
		NativeNames: []string{"dorerin Naoe", "Nauruan"}}
	Navajo = Language{
		Codes:       []Code{NV, NAV},
		EnglishName: "Navajo",
		Scope:       Individual,
		NativeNames: []string{"Navaho", "Diné bizaad", "Naabeehó bizaad"}}
	NorthNdebele = Language{
		Codes:       []Code{ND, NDE},
		EnglishName: "North Ndebele",
		Scope:       Individual,
		NativeNames: []string{"isiNdebele"}}
	SouthNdebele = Language{
		Codes:       []Code{NR, NBL},
		EnglishName: "South Ndebele",
		Scope:       Individual,
		NativeNames: []string{"isiNdebele"}}
	Ndonga = Language{
		Codes:       []Code{NG, NDO},
		EnglishName: "Ndonga",
		Scope:       Individual,
		NativeNames: []string{"Oshiwambo"}}
	Nepali = Language{
		Codes:       []Code{NE, NEP},
		EnglishName: "Nepali",
		Scope:       Macro,
		NativeNames: []string{"नेपाली", "Nepālī"},
		Variants: []Variant{
			{Code: DTY, EnglishName: "Dotyali"},
			{Code: NPI, EnglishName: "Nepali (individual language)"}}}
	Norwegian = Language{
		Codes:       []Code{NO, NOR},
		EnglishName: "Norwegian",
		Type:        Living,
		Scope:       Macro,
		Regions:     []Region{Europe, Eurasia},
		Family:      IndoEuropean,
		Scripts:     []Script{LatinScript},
		NativeNames: []string{"Norsk"},
		Variants: []Variant{
			{Code: NNO, EnglishName: "Norwegian Nynorsk"},
			{Code: NOB, EnglishName: "Norwegian Bokmål"}}}
	Occitan = Language{
		Codes:       []Code{OC, OCI},
		EnglishName: "Occitan",
		Scope:       Individual,
	}
	Ojibwa = Language{
		Codes:       []Code{OJ, OJI},
		EnglishName: "Ojibwa",
		Scope:       Macro,
		NativeNames: []string{"ᐊᓂᔑᓈᐯᒧᐎᓐ", "Anishinaabemowin"},
		Variants: []Variant{
			{Code: CIW, EnglishName: "Chippewa"},
			{Code: OJB, EnglishName: "Northwestern Ojibwa"},
			{Code: OJC, EnglishName: "Central Ojibwa"},
			{Code: OJG, EnglishName: "Eastern Ojibwa"},
			{Code: OJS, EnglishName: "Severn Ojibwa"},
			{Code: OJW, EnglishName: "Western Ojibwa"},
			{Code: OTW, EnglishName: "Ottawa"}}}
	Oriya = Language{
		Codes:       []Code{OR, ORI},
		EnglishName: "Oriya",
		Scope:       Macro,
		NativeNames: []string{"ଓଡ଼ିଆ", "Odia"},
		Variants: []Variant{
			{Code: ORY, EnglishName: "Odia"},
			{Code: SPV, EnglishName: "Sambalpuri"}}}
	Oromo = Language{
		Codes:       []Code{OM, ORM},
		EnglishName: "Oromo",
		Scope:       Macro,
		NativeNames: []string{"afaan Oromoo"},
		Variants: []Variant{
			{Code: GAX, EnglishName: "Borana-Arsi-Guji Oromo"},
			{Code: GAZ, EnglishName: "West Central Oromo"},
			{Code: HAE, EnglishName: "Eastern Oromo"},
			{Code: ORC, EnglishName: "Orma"}}}
	Ossetian = Language{
		Codes:       []Code{OS, OSS},
		EnglishName: "Ossetian",
		Scope:       Individual,
		NativeNames: []string{"Ossetic", "дигорон Ӕвзаг", "digoron Ævzag"}}
	Pali = Language{
		Codes:       []Code{PI, PLI},
		EnglishName: "Pali",
		Scope:       Individual,
	}
	Pashto = Language{
		Codes:       []Code{PS, PUS},
		EnglishName: "Pashto",
		Type:        Living,
		Scope:       Macro,
		Family:      IndoEuropean,
		Regions:     []Region{MiddleEast, CentralAsia, Eurasia, Asia},
		Scripts:     []Script{PashtoScript},
		NativeNames: []string{"Pushto", "پښتو", "Pax̌tow"},
		Variants: []Variant{
			{Code: PBT, EnglishName: "Southern Pashto"},
			{Code: PBU, EnglishName: "Northern Pashto"},
			{Code: PST, EnglishName: "Central Pashto"}}}
	Persian = Language{
		Codes:       []Code{FA, FAS, PER},
		EnglishName: "Persian",
		Type:        Living,
		Scope:       Macro,
		Family:      IndoEuropean,
		Regions:     []Region{MiddleEast, CentralAsia, Eurasia},
		Scripts:     []Script{PersianScript},
		NativeNames: []string{"فارسی", "Fārsiy"},
		Variants: []Variant{
			{Code: PES, EnglishName: "Iranian Persian"},
			{Code: PRS, EnglishName: "Dari"}}}
	Polish = Language{
		Codes:       []Code{PL, POL},
		EnglishName: "Polish",
		Scope:       Individual,
		NativeNames: []string{"Polski"}}
	Portuguese = Language{
		Codes:       []Code{PT, POR},
		EnglishName: "Portuguese",
		Scope:       Individual,
		NativeNames: []string{"Português"}}
	Punjabi = Language{
		Codes:       []Code{PA, PAN},
		EnglishName: "Punjabi",
		Scope:       Individual,
		NativeNames: []string{"ਪੰਜਾਬੀ", "Panjabi", "پنجابی", "Pãjābī"}}
	Quechua = Language{
		Codes:       []Code{QU, QUE},
		EnglishName: "Quechua",
		Scope:       Macro,
		NativeNames: []string{"Runa simi", "kichwa simi", "Nuna shimi"},
		Variants: []Variant{
			{Code: QUB, EnglishName: "Huallaga Huánuco Quechua"},
			{Code: QUD, EnglishName: "Calderón Highland Quichua"},
			{Code: QUF, EnglishName: "Lambayeque Quechua"},
			{Code: QUG, EnglishName: "Chimborazo Highland Quichua"},
			{Code: QUH, EnglishName: "South Bolivian Quechua"},
			{Code: QUK, EnglishName: "Chachapoyas Quechua"},
			{Code: QUL, EnglishName: "North Bolivian Quechua"},
			{Code: QUP, EnglishName: "Southern Pastaza Quechua"},
			{Code: QUR, EnglishName: "Yanahuanca Pasco Quechua"},
			{Code: QUS, EnglishName: "Santiago del Estero Quichua"},
			{Code: QUW, EnglishName: "Tena Lowland Quichua"},
			{Code: QUX, EnglishName: "Yauyos Quechua"},
			{Code: QUY, EnglishName: "Ayacucho Quechua"},
			{Code: QUZ, EnglishName: "Cusco Quechua"},
			{Code: QVA, EnglishName: "Ambo-Pasco Quechua"},
			{Code: QVC, EnglishName: "Cajamarca Quechua"},
			{Code: QVE, EnglishName: "Eastern Apurímac Quechua"},
			{Code: QVH, EnglishName: "Huamalíes-Dos de Mayo Huánuco Quechua"},
			{Code: QVI, EnglishName: "Imbabura Highland Quichua"},
			{Code: QVJ, EnglishName: "Loja Highland Quichua"},
			{Code: QVL, EnglishName: "Cajatambo North Lima Quechua"},
			{Code: QVM, EnglishName: "Margos-Yarowilca-Lauricocha Quechua"},
			{Code: QVN, EnglishName: "North Junín Quechua"},
			{Code: QVO, EnglishName: "Napo Lowland Quechua"},
			{Code: QVP, EnglishName: "Pacaraos Quechua"},
			{Code: QVS, EnglishName: "San Martín Quechua"},
			{Code: QVW, EnglishName: "Huaylla Wanca Quechua"},
			{Code: QVZ, EnglishName: "Northern Pastaza Quichua"},
			{Code: QWA, EnglishName: "Corongo Ancash Quechua"},
			{Code: QWC, EnglishName: "Classical Quechua"},
			{Code: QWH, EnglishName: "Huaylas Ancash Quechua"},
			{Code: QWS, EnglishName: "Sihuas Ancash Quechua"},
			{Code: QXA, EnglishName: "Chiquián Ancash Quechua"},
			{Code: QXC, EnglishName: "Chincha Quechua"},
			{Code: QXH, EnglishName: "Panao Huánuco Quechua"},
			{Code: QXL, EnglishName: "Salasaca Highland Quichua"},
			{Code: QXN, EnglishName: "Northern Conchucos Ancash Quechua"},
			{Code: QXO, EnglishName: "Southern Conchucos Ancash Quechua"},
			{Code: QXP, EnglishName: "Puno Quechua"},
			{Code: QXR, EnglishName: "Cañar Highland Quichua"},
			{Code: QXT, EnglishName: "Santa Ana de Tusi Pasco Quechua"},
			{Code: QXU, EnglishName: "Arequipa-La Unión Quechua"},
			{Code: QXW, EnglishName: "Jauja Wanca Quechua"}}}
	Romanian = Language{
		Codes:       []Code{RO, RON, RUM},
		EnglishName: "Romanian",
		Scope:       Individual,
		NativeNames: []string{"Românã", "Moldavian", " Moldovan"},
		Variants: []Variant{
			{Code: MO, EnglishName: "Moldavian"},
			{Code: MOL, EnglishName: "Moldavian"}}}
	Romansh = Language{
		Codes:       []Code{RM, ROH},
		EnglishName: "Romansh",
		Scope:       Individual,
		NativeNames: []string{"Rumantsch", "Rumàntsch", "Romauntsch", "Romontsch"}}
	Rundi = Language{
		Codes:       []Code{RN, RUN},
		EnglishName: "Rundi",
		Scope:       Individual,
		NativeNames: []string{"Ikirundi"}}
	Russian = Language{
		Codes:       []Code{RU, RUS},
		EnglishName: "Russian",
		Scope:       Individual,
		NativeNames: []string{"Русский язык", "Russkiĭ âzyk"}}
	NorthernSami = Language{
		Codes:       []Code{SE, SME},
		EnglishName: "Northern Sami",
		Scope:       Individual,
		NativeNames: []string{"Davvisámegiella"}}
	Samoan = Language{
		Codes:       []Code{SM, SMO},
		EnglishName: "Samoan",
		Scope:       Individual,
		NativeNames: []string{"gagana Sāmoa", "Sāmoa"}}
	Sango = Language{
		Codes:       []Code{SG, SAG},
		EnglishName: "Sango",
		Scope:       Individual,
		NativeNames: []string{"yângâ tî Sängö"}}
	Sanskrit = Language{
		Codes:       []Code{SA, SAN},
		EnglishName: "Sanskrit",
		Type:        Ancient,
		Scope:       Macro,
		Family:      IndoEuropean,
		Scripts:     []Script{Devanagari, Brahmi},
		Regions:     []Region{Asia, India, Eurasia},
		NativeNames: []string{"संस्कृतम्", "Saṃskṛtam"},
		Variants: []Variant{
			{Code: CLS, EnglishName: "Classical Sanskrit"},
			{Code: VSN, EnglishName: "Vedic Sanskrit"}}}
	Sardinian = Language{
		Codes:       []Code{SC, SRD},
		EnglishName: "Sardinian",
		Scope:       Macro,
		NativeNames: []string{"Sardu"},
		Variants: []Variant{
			{Code: SDC, EnglishName: "Sassarese Sardinian"},
			{Code: SDN, EnglishName: "Gallurese Sardinian"},
			{Code: SRC, EnglishName: "Logudorese Sardinian"},
			{Code: SRO, EnglishName: "Campidanese Sardinian"}}}
	Serbian = Language{
		Codes:       []Code{SR, SRP},
		EnglishName: "Serbian",
		Scope:       Individual,
		NativeNames: []string{"Српски", "Srpski"}}
	Shona = Language{
		Codes:       []Code{SN, SNA},
		EnglishName: "Shona",
		Scope:       Individual,
		NativeNames: []string{"chiShona"}}
	Sindhi = Language{
		Codes:       []Code{SD, SND},
		EnglishName: "Sindhi",
		Scope:       Individual,
		NativeNames: []string{"سنڌي", "सिन्धी", "Sindhī"}}
	Sinhala = Language{
		Codes:       []Code{SI, SIN},
		EnglishName: "Sinhala",
		Scope:       Individual,
		NativeNames: []string{"Sinhalese", "සිංහල", "Sinhala"}}
	Slovak = Language{
		Codes:       []Code{SK, SLK, SLO},
		EnglishName: "Slovak",
		Scope:       Individual,
		NativeNames: []string{"Slovenčina"}}
	Slovenian = Language{
		Codes:       []Code{SL, SLV},
		EnglishName: "Slovenian",
		Scope:       Individual,
		NativeNames: []string{"Slovenščina"}}
	Somali = Language{
		Codes:       []Code{SO, SOM},
		EnglishName: "Somali",
		Scope:       Individual,
		NativeNames: []string{"Soomaali", "𐒈𐒝𐒑𐒛𐒐𐒘", "سٝومالِ"}}
	SouthernSotho = Language{
		Codes:       []Code{ST, SOT},
		EnglishName: "Southern Sotho",
		Scope:       Individual,
		NativeNames: []string{"Sesotho"}}
	Spanish = Language{
		Codes:       []Code{ES, SPA},
		EnglishName: "Spanish",
		Scope:       Individual,
		NativeNames: []string{"Castilian", "Español", "Castellano"}}
	Sundanese = Language{
		Codes:       []Code{SU, SUN},
		EnglishName: "Sundanese",
		Scope:       Individual,
		NativeNames: []string{"basa Sunda", "ᮘᮞ ᮞᮥᮔ᮪ᮓ", "بَاسَا سُوْندَا"}}
	Swahili = Language{
		Codes:       []Code{SW, SWA},
		EnglishName: "Swahili",
		Type:        Living,
		Scope:       Macro,
		Scripts:     []Script{LatinScript, ArabicScript},
		Family:      NigerCongo,
		Regions:     []Region{Africa},
		NativeNames: []string{"Kiswahili", "كِسوَحِيلِ"},
		Variants: []Variant{
			{Code: SWC, EnglishName: "Congo Swahili"},
			{Code: SWH, EnglishName: "Swahili (individual language)"}}}
	Swati = Language{
		Codes:       []Code{SS, SSW},
		EnglishName: "Swati",
		Scope:       Individual,
		NativeNames: []string{"SiSwati"}}
	Swedish = Language{
		Codes:       []Code{SV, SWE},
		EnglishName: "Swedish",
		Scope:       Individual,
		NativeNames: []string{"Svenska"}}
	Tagalog = Language{
		Codes:       []Code{TL, TGL},
		EnglishName: "Tagalog",
		Scope:       Individual,
		NativeNames: []string{"Wikang Tagalog"}}
	Tahitian = Language{
		Codes:       []Code{TY, TAH},
		EnglishName: "Tahitian",
		Scope:       Individual,
		NativeNames: []string{"reo Tahiti", "Reo Mā`ohi"}}
	Tajik = Language{
		Codes:       []Code{TG, TGK},
		EnglishName: "Tajik",
		Scope:       Individual,
		NativeNames: []string{"Тоҷикӣ", "Tojikī"}}
	Tamil = Language{
		Codes:       []Code{TA, TAM},
		EnglishName: "Tamil",
		Scope:       Individual,
		NativeNames: []string{"தமிழ்", "Tamiḻ"}}
	Tatar = Language{
		Codes:       []Code{TT, TAT},
		EnglishName: "Tatar",
		Scope:       Individual,
		NativeNames: []string{"Татар теле", "Tatar tele", "تاتار تئلئ"}}
	Telugu = Language{
		Codes:       []Code{TE, TEL},
		EnglishName: "Telugu",
		Scope:       Individual,
		NativeNames: []string{"తెలుగు"}}
	Thai = Language{
		Codes:       []Code{TH, THA},
		EnglishName: "Thai",
		Scope:       Individual,
		NativeNames: []string{"ภาษาไทย", "Phasa Thai"}}
	Tibetan = Language{
		Codes:       []Code{BO, BOD, TIB},
		EnglishName: "Tibetan",
		Scope:       Individual,
		NativeNames: []string{"བོད་སྐད་", "Bodskad", "ལྷ་སའི་སྐད་", "Lhas'iskad"}}
	Tigrinya = Language{
		Codes:       []Code{TI, TIR},
		EnglishName: "Tigrinya",
		Scope:       Individual,
		NativeNames: []string{"ትግርኛ", "Təgrəñña"}}
	Tonga = Language{
		Codes:       []Code{TO, TON},
		EnglishName: "Tonga",
		Scope:       Individual,
		NativeNames: []string{"Tonga Islands", "lea faka-Tonga"}}
	Tsonga = Language{
		Codes:       []Code{TS, TSO},
		EnglishName: "Tsonga",
		Scope:       Individual,
	}
	Tswana = Language{
		Codes:       []Code{TN, TSN},
		EnglishName: "Tswana",
		Scope:       Individual,
		NativeNames: []string{"Setswana", "Sechuana"}}
	Turkish = Language{
		Codes:       []Code{TR, TUR},
		EnglishName: "Turkish",
		Scope:       Individual,
		NativeNames: []string{"Türkçe"}}
	Turkmen = Language{
		Codes:       []Code{TK, TUK},
		EnglishName: "Turkmen",
		Scope:       Individual,
		NativeNames: []string{"Türkmençe", "Түркменче", "تۆرکمنچه"}}
	Twi = Language{
		Codes:       []Code{TW, TWI},
		EnglishName: "Twi",
		Scope:       Individual,
	}
	Uighur = Language{
		Codes:       []Code{UG, UIG},
		EnglishName: "Uighur",
		Scope:       Individual,
		NativeNames: []string{"Uyghur", "ئۇيغۇر تىلى", "Уйғур тили", "Uyƣur tili"}}
	Ukrainian = Language{
		Codes:       []Code{UK, UKR},
		EnglishName: "Ukrainian",
		Scope:       Individual,
		NativeNames: []string{"Українська", "Ukraїnska"}}
	Urdu = Language{
		Codes:       []Code{UR, URD},
		EnglishName: "Urdu",
		Scope:       Individual,
		NativeNames: []string{"اُردُو", "Urduw"}}
	Uzbek = Language{
		Codes:       []Code{UZ, UZB},
		EnglishName: "Uzbek",
		Type:        Living,
		Scope:       Macro,
		Scripts:     []Script{Cyrillic, LatinScript, ArabicScript},
		Family:      Turkic,
		Regions:     []Region{Asia, CentralAsia},
		NativeNames: []string{"Ózbekça", "ўзбекча", "ئوزبېچه"},
		Variants: []Variant{
			{Code: UZN, EnglishName: "Northern Uzbek"},
			{Code: UZS, EnglishName: "Southern Uzbek"}}}
	Venda = Language{
		Codes:       []Code{VE, VEN},
		EnglishName: "Venda",
		Scope:       Individual,
		NativeNames: []string{"Tshivenḓa"}}
	Vietnamese = Language{
		Codes:       []Code{VI, VIE},
		EnglishName: "Vietnamese",
		Scope:       Individual,
		NativeNames: []string{"Tiếng Việt"}}
	Volapük = Language{
		Codes:       []Code{VO, VOL},
		EnglishName: "Volapük",
		Scope:       Individual,
	}
	Walloon = Language{
		Codes:       []Code{WA, WLN},
		EnglishName: "Walloon",
		Scope:       Individual,
		NativeNames: []string{"Walon"}}
	Welsh = Language{
		Codes:       []Code{CY, CYM, WEL},
		EnglishName: "Welsh",
		Scope:       Individual,
		NativeNames: []string{"Cymraeg"}}
	Wolof = Language{
		Codes:       []Code{WO, WOL},
		EnglishName: "Wolof",
		Scope:       Individual,
		NativeNames: []string{"وࣷلࣷفْ"}}
	Xhosa = Language{
		Codes:       []Code{XH, XHO},
		EnglishName: "Xhosa",
		Scope:       Individual,
		NativeNames: []string{"isiXhosa"}}
	SichuanYi = Language{
		Codes:       []Code{II, III},
		EnglishName: "Sichuan Yi",
		Scope:       Individual,
		NativeNames: []string{"Nuosu", "ꆈꌠꉙ", "Nuosuhxop"}}
	Yoruba = Language{
		Codes:       []Code{YO, YOR},
		EnglishName: "Yoruba",
		Scope:       Individual,
		NativeNames: []string{"èdè Yorùbá"}}
	Yiddish = Language{
		Codes:       []Code{YI, YID},
		EnglishName: "Yiddish",
		Scope:       Macro,
		Family:      IndoEuropean,
		Scripts:     []Script{HebrewScript, LatinScript},
		Regions:     []Region{Eurasia, Asia},
		NativeNames: []string{"ייִדיש", "Yidiš"},
		Variants: []Variant{
			{Code: YDD, EnglishName: "Eastern Yiddish"},
			{Code: YIH, EnglishName: "Western Yiddish"}}}
	Zhuang = Language{
		Codes:       []Code{ZA, ZHA},
		EnglishName: "Zhuang",
		Scope:       Macro,
		NativeNames: []string{"Chuang", "話僮", "Vahcuengh"},
		Variants: []Variant{
			{Code: ZCH, EnglishName: "Central Hongshuihe Zhuang"},
			{Code: ZEH, EnglishName: "Eastern Hongshuihe Zhuang"},
			{Code: ZGB, EnglishName: "Guibei Zhuang"},
			{Code: ZGM, EnglishName: "Minz Zhuang"},
			{Code: ZGN, EnglishName: "Guibian Zhuang"},
			{Code: ZHD, EnglishName: "Dai Zhuang"},
			{Code: ZHN, EnglishName: "Nong Zhuang"},
			{Code: ZLJ, EnglishName: "Liujiang Zhuang"},
			{Code: ZLN, EnglishName: "Lianshan Zhuang"},
			{Code: ZLQ, EnglishName: "Liuqian Zhuang"},
			{Code: ZQE, EnglishName: "Qiubei Zhuang"},
			{Code: ZYB, EnglishName: "Yongbei Zhuang"},
			{Code: ZYG, EnglishName: "Yang Zhuang"},
			{Code: ZYJ, EnglishName: "Youjiang Zhuang"},
			{Code: ZYN, EnglishName: "Yongnan Zhuang"},
			{Code: ZZJ, EnglishName: "Zuojiang Zhuang"},
		}}
	Zulu = Language{
		Codes:       []Code{ZU, ZUL},
		EnglishName: "Zulu",
		Scope:       Individual,
		NativeNames: []string{"isiZulu"}}
	AncientGreek = Language{
		Codes:       []Code{GRC},
		EnglishName: "Ancient Greek",
		Scope:       Individual,
		NativeNames: []string{"Ἀρχαία ἑλληνικὴ", "Archaía ellēnikḗ"}}
	Zapotec = Language{
		Codes:       []Code{ZAP},
		EnglishName: "Zapotec",
		Scope:       Individual,
		NativeNames: []string{"Didxsaj"}}
	Blissymbols = Language{
		Codes:       []Code{ZBL},
		EnglishName: "Blissymbols",
		Scope:       Individual,
	}
	Sumerian = Language{
		Codes:       []Code{SUX},
		EnglishName: "Sumerian",
		Type:        Ancient,
		Regions:     []Region{MiddleEast},
		Family:      LanguageIsolate,
		Scripts:     []Script{Cuneiform},
	}
)
