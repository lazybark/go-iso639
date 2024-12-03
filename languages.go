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
		Codes:       []Code{"br", "bre"},
		EnglishName: "Breton",
		Scope:       Individual,
		NativeNames: []string{"Brezhoneg"}}
	Bulgarian = Language{
		Codes:       []Code{"bg", "bul"},
		EnglishName: "Bulgarian",
		Scope:       Individual,
		NativeNames: []string{"Български", "Bulgarski"}}
	Burmese = Language{
		Codes:       []Code{"my", "mya", "bur"},
		EnglishName: "Burmese",
		Scope:       Individual,
		NativeNames: []string{"မြန်မာစာ", "Mrãmācā"}}
	Catalan = Language{
		Codes:       []Code{"ca", "cat"},
		EnglishName: "Catalan",
		Scope:       Individual,
		NativeNames: []string{"Valencian", "Català", "Valencià"}}
	Chamorro = Language{
		Codes:       []Code{"ch", "cha"},
		EnglishName: "Chamorro",
		Scope:       Individual,
		NativeNames: []string{"Finu' Chamoru"}}
	Chechen = Language{
		Codes:       []Code{"ce", "che"},
		EnglishName: "Chechen",
		Scope:       Individual,
		NativeNames: []string{"Нохчийн мотт", "Noxçiyn mott", "Chechnyan", "Chechnian"}}
	Chichewa = Language{
		Codes:       []Code{"ny", "nya"},
		EnglishName: "Chichewa",
		Scope:       Individual,
		NativeNames: []string{"Chewa", "Nyanja", "Chinyanja"}}
	Chinese = Language{
		Codes:       []Code{"zh", "zho", "chi"},
		EnglishName: "Chinese",
		Scope:       Macro,
		NativeNames: []string{"中文", "Zhōngwén", "汉语", "漢語", "Hànyǔ"},
		Variants: []Variant{
			{Code: "cdo", EnglishName: "Min Dong Chinese"},
			{Code: "cjy", EnglishName: "Jinyu Chinese"},
			{Code: "cmn", EnglishName: "Mandarin Chinese"},
			{Code: "cnp", EnglishName: "Northern Ping Chinese"},
			{Code: "cpx", EnglishName: "Pu-Xian Chinese"},
			{Code: "csp", EnglishName: "Southern Ping Chinese"},
			{Code: "czh", EnglishName: "Huizhou Chinese"},
			{Code: "czo", EnglishName: "Min Zhong Chinese"},
			{Code: "gan", EnglishName: "Gan Chinese"},
			{Code: "hak", EnglishName: "Hakka Chinese"},
			{Code: "hnm", EnglishName: "Hainanese"},
			{Code: "hsn", EnglishName: "Xiang Chinese"},
			{Code: "luh", EnglishName: "Leizhou Chinese"},
			{Code: "lzh", EnglishName: "Literary Chinese"},
			{Code: "mnp", EnglishName: "Min Bei Chinese"},
			{Code: "nan", EnglishName: "Min Nan Chinese"},
			{Code: "sjc", EnglishName: "Shaojiang Chinese"},
			{Code: "wuu", EnglishName: "Wu Chinese"},
			{Code: "yue", EnglishName: "Yue Chinese"}}}
	ChurchSlavic = Language{
		Codes:       []Code{"cu", "chu"},
		EnglishName: "Church Slavonic",
		Scope:       Individual,
		NativeNames: []string{"Славе́нскїй ѧ҆зы́къ"}}
	Chuvash = Language{
		Codes:       []Code{"cv", "chv"},
		EnglishName: "Chuvash",
		Scope:       Individual,
		NativeNames: []string{"Чăвашла", "Çăvaşla"}}
	Cornish = Language{
		Codes:       []Code{"kw", "cor"},
		EnglishName: "Cornish",
		Scope:       Individual,
		NativeNames: []string{"Kernowek"}}
	Corsican = Language{
		Codes:       []Code{"co", "cos"},
		EnglishName: "Corsican",
		Scope:       Individual,
		NativeNames: []string{"Corsu"}}
	Cree = Language{
		Codes:       []Code{"cr", "cre"},
		EnglishName: "Cree",
		Scope:       Macro,
		NativeNames: []string{"ᓀᐦᐃᔭᐁᐧᐃᐧᐣ", "Nehiyawewin"},
		Variants: []Variant{
			{Code: "crj", EnglishName: "Southern East Cree"},
			{Code: "crk", EnglishName: "Plains Cree"},
			{Code: "crl", EnglishName: "Northern East Cree"},
			{Code: "crm", EnglishName: "Moose Cree"},
			{Code: "csw", EnglishName: "Swampy Cree"},
			{Code: "cwd", EnglishName: "Woods Cree"}}}
	Croatian = Language{
		Codes:       []Code{"hr", "hrv"},
		EnglishName: "Croatian",
		Scope:       Individual,
		NativeNames: []string{"Hrvatski", "Crovatian"}}
	Czech = Language{
		Codes:       []Code{"cs", "ces", "cze"},
		EnglishName: "Czech",
		Scope:       Individual,
		NativeNames: []string{"Čeština", "Czechian"}}
	Danish = Language{
		Codes:       []Code{"da", "dan"},
		EnglishName: "Danish",
		Scope:       Individual,
		NativeNames: []string{"Dansk"}}
	Divehi = Language{
		Codes:       []Code{"dv", "div"},
		EnglishName: "Divehi",
		Scope:       Individual,
		NativeNames: []string{"Dhivehi", "Maldivian", "ދިވެހި "}}
	Dutch = Language{
		Codes:       []Code{"nl", "nld", "dut"},
		EnglishName: "Dutch",
		Scope:       Individual,
		NativeNames: []string{"Flemish", "Nederlands"},
		Variants: []Variant{
			{Code: "vls", EnglishName: "West Flemish"}}}
	Dzongkha = Language{
		Codes:       []Code{"dz", "dzo"},
		EnglishName: "Dzongkha",
		Scope:       Individual,
		NativeNames: []string{"རྫོང་ཁ་", "Bhutanese"}}
	English = Language{
		Codes:       []Code{"en", "eng"},
		EnglishName: "English",
		Type:        Living,
		Scope:       Individual,
		Scripts:     []Script{LatinScript, AngloSaxonRunes},
		Family:      IndoEuropean,
		Regions:     []Region{Europe, Oceania, NorthAmerica, SouthAmerica, Africa, Asia, Eurasia, India},
	}
	Esperanto = Language{
		Codes:       []Code{"eo", "epo"},
		EnglishName: "Esperanto",
		Scope:       Individual,
	}
	Estonian = Language{
		Codes:       []Code{"et", "est"},
		EnglishName: "Estonian",
		Scope:       Macro,
		NativeNames: []string{"Eesti keel"},
		Variants: []Variant{
			{Code: "vro", EnglishName: "Võro"},
			{Code: "ekk", EnglishName: "Standard Estonian"}}}
	Ewe = Language{
		Codes:       []Code{"ee", "ewe"},
		EnglishName: "Ewe",
		Scope:       Individual,
		NativeNames: []string{"Èʋegbe"}}
	Faroese = Language{
		Codes:       []Code{"fo", "fao"},
		EnglishName: "Faroese",
		Scope:       Individual,
		NativeNames: []string{"Føroyskt"}}
	Fijian = Language{
		Codes:       []Code{"fj", "fij"},
		EnglishName: "Fijian",
		Scope:       Individual,
		NativeNames: []string{"Na Vosa Vakaviti"}}
	Finnish = Language{
		Codes:       []Code{"fi", "fin"},
		EnglishName: "Finnish",
		Scope:       Individual,
		NativeNames: []string{"Suomi"}}
	French = Language{
		Codes:       []Code{"fr", "fra", "fre"},
		EnglishName: "French",
		Scope:       Individual,
		NativeNames: []string{"Français"}}
	WesternFrisian = Language{
		Codes:       []Code{"fy", "fry"},
		EnglishName: "Western Frisian",
		Scope:       Individual,
		NativeNames: []string{"Frysk", "West Frisian", "Frisian", "Fries"}}
	Fulah = Language{
		Codes:       []Code{"ff", "ful"},
		EnglishName: "Fulah",
		Scope:       Macro,
		NativeNames: []string{"𞤊𞤵𞤤𞤬𞤵𞤤𞤣𞤫", "ࢻُلْࢻُلْدٜ", "Fulfulde", "𞤆𞤵𞤤𞤢𞥄𞤪", "ݒُلَارْ", "Pulaar"},
		Variants: []Variant{
			{Code: "ffm", EnglishName: "Maasina Fulfulde"},
			{Code: "fub", EnglishName: "Adamawa Fulfulde"},
			{Code: "fuc", EnglishName: "Pulaar"},
			{Code: "fue", EnglishName: "Borgu Fulfulde"},
			{Code: "fuf", EnglishName: "Pular"},
			{Code: "fuh", EnglishName: "Western Niger Fulfulde"},
			{Code: "fui", EnglishName: "Bagirmi Fulfulde"},
			{Code: "fuq", EnglishName: "Central-Eastern Niger Fulfulde"},
			{Code: "fuv", EnglishName: "Nigerian Fulfulde"}}}
	Gaelic = Language{
		Codes:       []Code{"gd", "gla"},
		EnglishName: "Gaelic",
		Scope:       Individual,
		NativeNames: []string{"Scottish Gaelic", "Gàidhlig"}}
	Galician = Language{
		Codes:       []Code{"gl", "glg"},
		EnglishName: "Galician",
		Scope:       Individual,
		NativeNames: []string{"Galego"}}
	Ganda = Language{
		Codes:       []Code{"lg", "lug"},
		EnglishName: "Ganda",
		Scope:       Individual,
		NativeNames: []string{"Luganda"}}
	Georgian = Language{
		Codes:       []Code{"ka", "kat"},
		EnglishName: "Georgian",
		Scope:       Individual,
		NativeNames: []string{"ქართული", "Kharthuli"}}
	German = Language{
		Codes:       []Code{"de", "deu"},
		EnglishName: "German",
		Scope:       Individual,
		NativeNames: []string{"Deutsch"}}
	Greek = Language{
		Codes:       []Code{"el", "ell"},
		EnglishName: "Greek",
		Scope:       Individual,
		NativeNames: []string{"Νέα Ελληνικά", "Néa Ellêniká"}}
	Greenlandic = Language{
		Codes:       []Code{"kl", "kal"},
		EnglishName: "Greenlandic",
		Scope:       Individual,
		NativeNames: []string{"Kalaallisut"}}
	Guarani = Language{
		Codes:       []Code{"gn", "grn"},
		EnglishName: "Guarani",
		Scope:       Macro,
		NativeNames: []string{"Avañe'ẽ"},
		Variants: []Variant{
			{Code: "gnw", EnglishName: "Western Bolivian Guaraní"},
			{Code: "gug", EnglishName: "Paraguayan Guaraní"},
			{Code: "gui", EnglishName: "Eastern Bolivian Guaraní"},
			{Code: "gun", EnglishName: "Mbyá Guaraní"},
			{Code: "nhd", EnglishName: "Chiripá"}}}
	Gujarati = Language{
		Codes:       []Code{"gu", "guj"},
		EnglishName: "Gujarati",
		Scope:       Individual,
		NativeNames: []string{"ગુજરાતી", "Gujarātī"}}
	Haitian = Language{
		Codes:       []Code{"ht", "hat"},
		EnglishName: "Haitian",
		Scope:       Individual,
		NativeNames: []string{"Haitian Creole", "Kreyòl ayisyen"}}
	Hausa = Language{
		Codes:       []Code{"ha", "hau"},
		EnglishName: "Hausa",
		Scope:       Individual,
		NativeNames: []string{"هَرْشٜن هَوْس", "halshen Hausa"}}
	Hebrew = Language{
		Codes:       []Code{"he", "heb"},
		EnglishName: "Hebrew",
		Scope:       Individual,
		NativeNames: []string{"עברית", "Ivrit"}}
	Herero = Language{
		Codes:       []Code{"hz", "her"},
		EnglishName: "Herero",
		Scope:       Individual,
		NativeNames: []string{"Otjiherero"}}
	Hindi = Language{
		Codes:       []Code{"hi", "hin"},
		EnglishName: "Hindi",
		Scope:       Individual,
		NativeNames: []string{"हिन्दी", "Hindī"}}
	HiriMotu = Language{
		Codes:       []Code{"ho", "hmo"},
		EnglishName: "Hiri Motu",
		Scope:       Individual,
		NativeNames: []string{"Police Motu"}}
	Hungarian = Language{
		Codes:       []Code{"hu", "hun"},
		EnglishName: "Hungarian",
		Scope:       Individual,
		NativeNames: []string{"Magyar nyelv"}}
	Icelandic = Language{
		Codes:       []Code{"is", "isl"},
		EnglishName: "Icelandic",
		Scope:       Individual,
		NativeNames: []string{"Íslenska"}}
	Ido = Language{
		Codes:       []Code{"io", "ido"},
		EnglishName: "Ido",
		Scope:       Individual,
	}
	Igbo = Language{
		Codes:       []Code{"ig", "ibo"},
		EnglishName: "Igbo",
		Scope:       Individual,
		NativeNames: []string{"ásụ̀sụ́ Ìgbò"}}
	Indonesian = Language{
		Codes:       []Code{"id", "ind"},
		EnglishName: "Indonesian",
		Scope:       Individual,
		NativeNames: []string{"bahasa Indonesia"}}
	Interlingua = Language{
		Codes:       []Code{"ia", "ina"},
		EnglishName: "Interlingua",
		Scope:       Individual,
	}
	Interlingue = Language{
		Codes:       []Code{"ie", "ile"},
		EnglishName: "Interlingue",
		Scope:       Individual,
		NativeNames: []string{"Occidental"}}
	Inuktitut = Language{
		Codes:       []Code{"iu", "iku"},
		EnglishName: "Inuktitut",
		Scope:       Macro,
		NativeNames: []string{"ᐃᓄᒃᑎᑐᑦ"},
		Variants: []Variant{
			{Code: "ike", EnglishName: "Eastern Canadian Inuktitut"},
			{Code: "ikt", EnglishName: "Inuinnaqtun"}}}
	Inupiaq = Language{
		Codes:       []Code{"ik", "ipk"},
		EnglishName: "Inupiaq",
		Scope:       Macro,
		NativeNames: []string{"Iñupiaq"},
		Variants: []Variant{
			{Code: "esi", EnglishName: "North Alaskan Inupiatun"},
			{Code: "esk", EnglishName: "Northwest Alaska Inupiatun"}}}
	Irish = Language{
		Codes:       []Code{"ga", "gle"},
		EnglishName: "Irish",
		Scope:       Individual,
		NativeNames: []string{"Gaeilge"}}
	Italian = Language{
		Codes:       []Code{"it", "ita"},
		EnglishName: "Italian",
		Scope:       Individual,
		NativeNames: []string{"Italiano"}}
	Japanese = Language{
		Codes:       []Code{"ja", "jpn"},
		EnglishName: "Japanese",
		Scope:       Individual,
		NativeNames: []string{"日本語", "Nihongo"}}
	Javanese = Language{
		Codes:       []Code{"jv", "jav"},
		EnglishName: "Javanese",
		Scope:       Individual,
		NativeNames: []string{"ꦧꦱꦗꦮ", "basa Jawa"}}
	Kannada = Language{
		Codes:       []Code{"kn", "kan"},
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
		Codes:       []Code{"km", "khm"},
		EnglishName: "Central Khmer",
		NativeNames: []string{"ខេមរភាសា", "Khémôrôphéasa", "Khmer", "Cambodian"}}
	Kikuyu = Language{
		Codes:       []Code{"ki", "kik"},
		EnglishName: "Kikuyu",
		Scope:       Individual,
		NativeNames: []string{"Gikuyu", "Gĩgĩkũyũ"}}
	Kinyarwanda = Language{
		Codes:       []Code{"rw", "kin"},
		EnglishName: "Kinyarwanda",
		Scope:       Individual,
		NativeNames: []string{"Ikinyarwanda"}}
	Kirghiz = Language{
		Codes:       []Code{"ky", "kir"},
		EnglishName: "Kirghiz",
		Scope:       Individual,
		NativeNames: []string{"Kyrgyz", "Кыргызча", "Kırgızça"}}
	Komi = Language{
		Codes:       []Code{"kv", "kom"},
		EnglishName: "Komi",
		Scope:       Macro,
		NativeNames: []string{"Коми кыв", "Zyran", "Zyrian"},
		Variants: []Variant{
			{Code: "koi", EnglishName: "Komi-Permyak"},
			{Code: "kpv", EnglishName: "Komi-Zyrian"}}}
	Kongo = Language{
		Codes:       []Code{"kg", "kon"},
		EnglishName: "Kongo",
		Scope:       Macro,
		NativeNames: []string{"Kikongo"},
		Variants: []Variant{
			{Code: "kng", EnglishName: "Koongo"},
			{Code: "kwy", EnglishName: "San Salvador Kongo"},
			{Code: "ldi", EnglishName: "Laari"}}}
	Korean = Language{
		Codes:       []Code{"ko", "kor"},
		EnglishName: "Korean",
		Scope:       Individual,
		NativeNames: []string{"한국어", "조선말", "Hangugeo", "Chosŏnmal"}}
	Kuanyama = Language{
		Codes:       []Code{"kj", "kua"},
		EnglishName: "Kuanyama",
		Scope:       Individual,
		NativeNames: []string{"Kwanyama"}}
	Kurdish = Language{
		Codes:       []Code{"ku", "kur"},
		EnglishName: "Kurdish",
		Type:        Living,
		Scope:       Macro,
		Family:      IndoEuropean,
		Scripts:     []Script{LatinScript, ArabicScript, Cyrillic},
		Regions:     []Region{Asia, Europe, Eurasia, Caucasus},
		NativeNames: []string{"Kurdî", "کوردی"},
		Variants: []Variant{
			{Code: "ckb", EnglishName: "Central Kurdish"},
			{Code: "kmr", EnglishName: "Northern Kurdish"},
			{Code: "sdh", EnglishName: "Southern Kurdish"}}}
	Lao = Language{
		Codes:       []Code{"lo", "lao"},
		EnglishName: "Lao",
		Scope:       Individual,
		NativeNames: []string{"ພາສາລາວ", "phasa Lao"}}
	Latin = Language{
		Codes:       []Code{"la", "lat"},
		EnglishName: "Latin",
		Scope:       Individual,
		NativeNames: []string{"Latinum"}}
	Latvian = Language{
		Codes:       []Code{"lv", "lav"},
		EnglishName: "Latvian",
		Scope:       Macro,
		NativeNames: []string{"Latviski"},
		Variants: []Variant{
			{Code: "ltg", EnglishName: "Latgalian"},
			{Code: "lvs", EnglishName: "Standard Latvian"}}}
	Limburgan = Language{
		Codes:       []Code{"li", "lim"},
		EnglishName: "Limburgan",
		Scope:       Individual,
		NativeNames: []string{"Limburger", "Limburgish", "Lèmburgs"}}
	Lingala = Language{
		Codes:       []Code{"ln", "lin"},
		EnglishName: "Lingala",
		Scope:       Individual,
		NativeNames: []string{"Lingála"}}
	Lithuanian = Language{
		Codes:       []Code{"lt", "lit"},
		EnglishName: "Lithuanian",
		Scope:       Individual,
		NativeNames: []string{"Lietuviškai"}}
	LubaKatanga = Language{
		Codes:       []Code{"lu", "lub"},
		EnglishName: "Luba-Katanga",
		Scope:       Individual,
		NativeNames: []string{"Kiluba", "Luba-Shaba"}}
	Luxembourgish = Language{
		Codes:       []Code{"lb", "ltz"},
		EnglishName: "Luxembourgish",
		Scope:       Individual,
		NativeNames: []string{"Letzeburgesch", "Lëtzebuergesch", "Luxembourgian"}}
	Macedonian = Language{
		Codes:       []Code{"mk", "mkd", "mac"},
		EnglishName: "Macedonian",
		Scope:       Individual,
		NativeNames: []string{"Македонски", "Makedonski"}}
	Malagasy = Language{
		Codes:       []Code{"mg", "mlg"},
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
		Codes:       []Code{"ms", "msa"},
		EnglishName: "Malay",
		Scope:       Macro,
		NativeNames: []string{"بهاس ملايو", "bahasa Melayu"},
		Variants: []Variant{
			{Code: "bjn", EnglishName: "Banjar"},
			{Code: "btj", EnglishName: "Bacanese Malay"},
			{Code: "bve", EnglishName: "Berau Malay"},
			{Code: "bvu", EnglishName: "Bukit Malay"},
			{Code: "coa", EnglishName: "Cocos Islands Malay"},
			{Code: "dup", EnglishName: "Duano"},
			{Code: "hji", EnglishName: "Haji"},
			{Code: "jak", EnglishName: "Jakun"},
			{Code: "jax", EnglishName: "Jambi Malay"},
			{Code: "kvb", EnglishName: "Kubu"},
			{Code: "kvr", EnglishName: "Kerinci"},
			{Code: "kxd", EnglishName: "Brunei"},
			{Code: "lce", EnglishName: "Loncong"},
			{Code: "lcf", EnglishName: "Lubu"},
			{Code: "liw", EnglishName: "Col"},
			{Code: "max", EnglishName: "North Moluccan Malay"},
			{Code: "meo", EnglishName: "Kedah Malay"},
			{Code: "mfa", EnglishName: "Pattani Malay"},
			{Code: "mfb", EnglishName: "Bangka"},
			{Code: "min", EnglishName: "Minangkabau"},
			{Code: "mqg", EnglishName: "Kota Bangun Kutai Malay"},
			{Code: "msi", EnglishName: "Sabah Malay"},
			{Code: "mui", EnglishName: "Musi"},
			{Code: "orn", EnglishName: "Orang Kanaq"},
			{Code: "ors", EnglishName: "Orang Seletar"},
			{Code: "pel", EnglishName: "Pekal"},
			{Code: "pse", EnglishName: "Central Malay"},
			{Code: "tmw", EnglishName: "Temuan"},
			{Code: "urk", EnglishName: "Urak Lawoi'"},
			{Code: "vkk", EnglishName: "Kaur"},
			{Code: "vkt", EnglishName: "Tenggarong Kutai Malay"},
			{Code: "xmm", EnglishName: "Manado Malay"},
			{Code: "zlm", EnglishName: "Malay (individual language)"},
			{Code: "zmi", EnglishName: "Negeri Sembilan Malay"},
			{Code: "zsm", EnglishName: "Standard Malay"}}}
	Malayalam = Language{
		Codes:       []Code{"ml", "mal"},
		EnglishName: "Malayalam",
		Scope:       Individual,
		NativeNames: []string{"മലയാളം", "Malayāļã"}}
	Maltese = Language{
		Codes:       []Code{"mt", "mlt"},
		EnglishName: "Maltese",
		Scope:       Individual,
		NativeNames: []string{"Malti"}}
	Manx = Language{
		Codes:       []Code{"gv", "glv"},
		EnglishName: "Manx",
		Scope:       Individual,
		NativeNames: []string{"Gaelg", "Gailck"}}
	Maori = Language{
		Codes:       []Code{"mi", "mri", "mao"},
		EnglishName: "Maori",
		Scope:       Individual,
		NativeNames: []string{"reo Māori"}}
	Marathi = Language{
		Codes:       []Code{"mr", "mar"},
		EnglishName: "Marathi",
		Scope:       Individual,
		NativeNames: []string{"मराठी", "Marāṭhī", "Maharashtran"}}
	Marshallese = Language{
		Codes:       []Code{"mh", "mah"},
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
			{Code: "dty", EnglishName: "Dotyali"},
			{Code: "npi", EnglishName: "Nepali (individual language)"}}}
	Norwegian = Language{
		Codes:       []Code{"no", "nor"},
		EnglishName: "Norwegian",
		Type:        Living,
		Scope:       Macro,
		Regions:     []Region{Europe, Eurasia},
		Family:      IndoEuropean,
		Scripts:     []Script{LatinScript},
		NativeNames: []string{"Norsk"},
		Variants: []Variant{
			{Code: "nno", EnglishName: "Norwegian Nynorsk"},
			{Code: "nob", EnglishName: "Norwegian Bokmål"}}}
	Occitan = Language{
		Codes:       []Code{"oc", "oci"},
		EnglishName: "Occitan",
		Scope:       Individual,
	}
	Ojibwa = Language{
		Codes:       []Code{"oj", "oji"},
		EnglishName: "Ojibwa",
		Scope:       Macro,
		NativeNames: []string{"ᐊᓂᔑᓈᐯᒧᐎᓐ", "Anishinaabemowin"},
		Variants: []Variant{
			{Code: "ciw", EnglishName: "Chippewa"},
			{Code: "ojb", EnglishName: "Northwestern Ojibwa"},
			{Code: "ojc", EnglishName: "Central Ojibwa"},
			{Code: "ojg", EnglishName: "Eastern Ojibwa"},
			{Code: "ojs", EnglishName: "Severn Ojibwa"},
			{Code: "ojw", EnglishName: "Western Ojibwa"},
			{Code: "otw", EnglishName: "Ottawa"}}}
	Oriya = Language{
		Codes:       []Code{"or", "ori"},
		EnglishName: "Oriya",
		Scope:       Macro,
		NativeNames: []string{"ଓଡ଼ିଆ", "Odia"},
		Variants: []Variant{
			{Code: "ory", EnglishName: "Odia"},
			{Code: "spv", EnglishName: "Sambalpuri"}}}
	Oromo = Language{
		Codes:       []Code{"om", "orm"},
		EnglishName: "Oromo",
		Scope:       Macro,
		NativeNames: []string{"afaan Oromoo"},
		Variants: []Variant{
			{Code: "gax", EnglishName: "Borana-Arsi-Guji Oromo"},
			{Code: "gaz", EnglishName: "West Central Oromo"},
			{Code: "hae", EnglishName: "Eastern Oromo"},
			{Code: "orc", EnglishName: "Orma"}}}
	Ossetian = Language{
		Codes:       []Code{"os", "oss"},
		EnglishName: "Ossetian",
		Scope:       Individual,
		NativeNames: []string{"Ossetic", "дигорон Ӕвзаг", "digoron Ævzag"}}
	Pali = Language{
		Codes:       []Code{"pi", "pli"},
		EnglishName: "Pali",
		Scope:       Individual,
	}
	Pashto = Language{
		Codes:       []Code{"ps", "pus"},
		EnglishName: "Pashto",
		Type:        Living,
		Scope:       Macro,
		Family:      IndoEuropean,
		Regions:     []Region{MiddleEast, CentralAsia, Eurasia, Asia},
		Scripts:     []Script{PashtoScript},
		NativeNames: []string{"Pushto", "پښتو", "Pax̌tow"},
		Variants: []Variant{
			{Code: "pbt", EnglishName: "Southern Pashto"},
			{Code: "pbu", EnglishName: "Northern Pashto"},
			{Code: "pst", EnglishName: "Central Pashto"}}}
	Persian = Language{
		Codes:       []Code{"fa", "fas", "per"},
		EnglishName: "Persian",
		Type:        Living,
		Scope:       Macro,
		Family:      IndoEuropean,
		Regions:     []Region{MiddleEast, CentralAsia, Eurasia},
		Scripts:     []Script{PersianScript},
		NativeNames: []string{"فارسی", "Fārsiy"},
		Variants: []Variant{
			{Code: "pes", EnglishName: "Iranian Persian"},
			{Code: "prs", EnglishName: "Dari"}}}
	Polish = Language{
		Codes:       []Code{"pl", "pol"},
		EnglishName: "Polish",
		Scope:       Individual,
		NativeNames: []string{"Polski"}}
	Portuguese = Language{
		Codes:       []Code{"pt", "por"},
		EnglishName: "Portuguese",
		Scope:       Individual,
		NativeNames: []string{"Português"}}
	Punjabi = Language{
		Codes:       []Code{"pa", "pan"},
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
		Codes:       []Code{"ro", "ron", "rum"},
		EnglishName: "Romanian",
		Scope:       Individual,
		NativeNames: []string{"Românã", "Moldavian", " Moldovan"},
		Variants: []Variant{
			{Code: "mo", EnglishName: "Moldavian"},
			{Code: "mol", EnglishName: "Moldavian"}}}
	Romansh = Language{
		Codes:       []Code{"rm", "roh"},
		EnglishName: "Romansh",
		Scope:       Individual,
		NativeNames: []string{"Rumantsch", "Rumàntsch", "Romauntsch", "Romontsch"}}
	Rundi = Language{
		Codes:       []Code{"rn", "run"},
		EnglishName: "Rundi",
		Scope:       Individual,
		NativeNames: []string{"Ikirundi"}}
	Russian = Language{
		Codes:       []Code{"ru", "rus"},
		EnglishName: "Russian",
		Scope:       Individual,
		NativeNames: []string{"Русский язык", "Russkiĭ âzyk"}}
	NorthernSami = Language{
		Codes:       []Code{"se", "sme"},
		EnglishName: "Northern Sami",
		Scope:       Individual,
		NativeNames: []string{"Davvisámegiella"}}
	Samoan = Language{
		Codes:       []Code{"sm", "smo"},
		EnglishName: "Samoan",
		Scope:       Individual,
		NativeNames: []string{"gagana Sāmoa", "Sāmoa"}}
	Sango = Language{
		Codes:       []Code{"sg", "sag"},
		EnglishName: "Sango",
		Scope:       Individual,
		NativeNames: []string{"yângâ tî Sängö"}}
	Sanskrit = Language{
		Codes:       []Code{"sa", "san"},
		EnglishName: "Sanskrit",
		Type:        Ancient,
		Scope:       Macro,
		Family:      IndoEuropean,
		Scripts:     []Script{Devanagari, Brahmi},
		Regions:     []Region{Asia, India, Eurasia},
		NativeNames: []string{"संस्कृतम्", "Saṃskṛtam"},
		Variants: []Variant{
			{Code: "cls", EnglishName: "Classical Sanskrit"},
			{Code: "vsn", EnglishName: "Vedic Sanskrit"}}}
	Sardinian = Language{
		Codes:       []Code{"sc", "srd"},
		EnglishName: "Sardinian",
		Scope:       Macro,
		NativeNames: []string{"Sardu"},
		Variants: []Variant{
			{Code: "sdc", EnglishName: "Sassarese Sardinian"},
			{Code: "sdn", EnglishName: "Gallurese Sardinian"},
			{Code: "src", EnglishName: "Logudorese Sardinian"},
			{Code: "sro", EnglishName: "Campidanese Sardinian"}}}
	Serbian = Language{
		Codes:       []Code{"sr", "srp"},
		EnglishName: "Serbian",
		Scope:       Individual,
		NativeNames: []string{"Српски", "Srpski"}}
	Shona = Language{
		Codes:       []Code{"sn", "sna"},
		EnglishName: "Shona",
		Scope:       Individual,
		NativeNames: []string{"chiShona"}}
	Sindhi = Language{
		Codes:       []Code{"sd", "snd"},
		EnglishName: "Sindhi",
		Scope:       Individual,
		NativeNames: []string{"سنڌي", "सिन्धी", "Sindhī"}}
	Sinhala = Language{
		Codes:       []Code{"si", "sin"},
		EnglishName: "Sinhala",
		Scope:       Individual,
		NativeNames: []string{"Sinhalese", "සිංහල", "Sinhala"}}
	Slovak = Language{
		Codes:       []Code{"sk", "slk", "slo"},
		EnglishName: "Slovak",
		Scope:       Individual,
		NativeNames: []string{"Slovenčina"}}
	Slovenian = Language{
		Codes:       []Code{"sl", "slv"},
		EnglishName: "Slovenian",
		Scope:       Individual,
		NativeNames: []string{"Slovenščina"}}
	Somali = Language{
		Codes:       []Code{"so", "som"},
		EnglishName: "Somali",
		Scope:       Individual,
		NativeNames: []string{"Soomaali", "𐒈𐒝𐒑𐒛𐒐𐒘", "سٝومالِ"}}
	SouthernSotho = Language{
		Codes:       []Code{"st", "sot"},
		EnglishName: "Southern Sotho",
		Scope:       Individual,
		NativeNames: []string{"Sesotho"}}
	Spanish = Language{
		Codes:       []Code{"es", "spa"},
		EnglishName: "Spanish",
		Scope:       Individual,
		NativeNames: []string{"Castilian", "Español", "Castellano"}}
	Sundanese = Language{
		Codes:       []Code{"su", "sun"},
		EnglishName: "Sundanese",
		Scope:       Individual,
		NativeNames: []string{"basa Sunda", "ᮘᮞ ᮞᮥᮔ᮪ᮓ", "بَاسَا سُوْندَا"}}
	Swahili = Language{
		Codes:       []Code{"sw", "swa"},
		EnglishName: "Swahili",
		Type:        Living,
		Scope:       Macro,
		Scripts:     []Script{LatinScript, ArabicScript},
		Family:      NigerCongo,
		Regions:     []Region{Africa},
		NativeNames: []string{"Kiswahili", "كِسوَحِيلِ"},
		Variants: []Variant{
			{Code: "swc", EnglishName: "Congo Swahili"},
			{Code: "swh", EnglishName: "Swahili (individual language)"}}}
	Swati = Language{
		Codes:       []Code{"ss", "ssw"},
		EnglishName: "Swati",
		Scope:       Individual,
		NativeNames: []string{"SiSwati"}}
	Swedish = Language{
		Codes:       []Code{"sv", "swe"},
		EnglishName: "Swedish",
		Scope:       Individual,
		NativeNames: []string{"Svenska"}}
	Tagalog = Language{
		Codes:       []Code{"tl", "tgl"},
		EnglishName: "Tagalog",
		Scope:       Individual,
		NativeNames: []string{"Wikang Tagalog"}}
	Tahitian = Language{
		Codes:       []Code{"ty", "tah"},
		EnglishName: "Tahitian",
		Scope:       Individual,
		NativeNames: []string{"reo Tahiti", "Reo Mā`ohi"}}
	Tajik = Language{
		Codes:       []Code{"tg", "tgk"},
		EnglishName: "Tajik",
		Scope:       Individual,
		NativeNames: []string{"Тоҷикӣ", "Tojikī"}}
	Tamil = Language{
		Codes:       []Code{"ta", "tam"},
		EnglishName: "Tamil",
		Scope:       Individual,
		NativeNames: []string{"தமிழ்", "Tamiḻ"}}
	Tatar = Language{
		Codes:       []Code{"tt", "tat"},
		EnglishName: "Tatar",
		Scope:       Individual,
		NativeNames: []string{"Татар теле", "Tatar tele", "تاتار تئلئ"}}
	Telugu = Language{
		Codes:       []Code{"te", "tel"},
		EnglishName: "Telugu",
		Scope:       Individual,
		NativeNames: []string{"తెలుగు"}}
	Thai = Language{
		Codes:       []Code{"th", "tha"},
		EnglishName: "Thai",
		Scope:       Individual,
		NativeNames: []string{"ภาษาไทย", "Phasa Thai"}}
	Tibetan = Language{
		Codes:       []Code{"bo", "bod", "tib"},
		EnglishName: "Tibetan",
		Scope:       Individual,
		NativeNames: []string{"བོད་སྐད་", "Bodskad", "ལྷ་སའི་སྐད་", "Lhas'iskad"}}
	Tigrinya = Language{
		Codes:       []Code{"ti", "tir"},
		EnglishName: "Tigrinya",
		Scope:       Individual,
		NativeNames: []string{"ትግርኛ", "Təgrəñña"}}
	Tonga = Language{
		Codes:       []Code{"to", "ton"},
		EnglishName: "Tonga",
		Scope:       Individual,
		NativeNames: []string{"Tonga Islands", "lea faka-Tonga"}}
	Tsonga = Language{
		Codes:       []Code{"ts", "tso"},
		EnglishName: "Tsonga",
		Scope:       Individual,
	}
	Tswana = Language{
		Codes:       []Code{"tn", "tsn"},
		EnglishName: "Tswana",
		Scope:       Individual,
		NativeNames: []string{"Setswana", "Sechuana"}}
	Turkish = Language{
		Codes:       []Code{"tr", "tur"},
		EnglishName: "Turkish",
		Scope:       Individual,
		NativeNames: []string{"Türkçe"}}
	Turkmen = Language{
		Codes:       []Code{"tk", "tuk"},
		EnglishName: "Turkmen",
		Scope:       Individual,
		NativeNames: []string{"Türkmençe", "Түркменче", "تۆرکمنچه"}}
	Twi = Language{
		Codes:       []Code{"tw", "twi"},
		EnglishName: "Twi",
		Scope:       Individual,
	}
	Uighur = Language{
		Codes:       []Code{"ug", "uig"},
		EnglishName: "Uighur",
		Scope:       Individual,
		NativeNames: []string{"Uyghur", "ئۇيغۇر تىلى", "Уйғур тили", "Uyƣur tili"}}
	Ukrainian = Language{
		Codes:       []Code{"uk", "ukr"},
		EnglishName: "Ukrainian",
		Scope:       Individual,
		NativeNames: []string{"Українська", "Ukraїnska"}}
	Urdu = Language{
		Codes:       []Code{"ur", "urd"},
		EnglishName: "Urdu",
		Scope:       Individual,
		NativeNames: []string{"اُردُو", "Urduw"}}
	Uzbek = Language{
		Codes:       []Code{"uz", "uzb"},
		EnglishName: "Uzbek",
		Type:        Living,
		Scope:       Macro,
		Scripts:     []Script{Cyrillic, LatinScript, ArabicScript},
		Family:      Turkic,
		Regions:     []Region{Asia, CentralAsia},
		NativeNames: []string{"Ózbekça", "ўзбекча", "ئوزبېچه"},
		Variants: []Variant{
			{Code: "uzn", EnglishName: "Northern Uzbek"},
			{Code: "uzs", EnglishName: "Southern Uzbek"}}}
	Venda = Language{
		Codes:       []Code{"ve", "ven"},
		EnglishName: "Venda",
		Scope:       Individual,
		NativeNames: []string{"Tshivenḓa"}}
	Vietnamese = Language{
		Codes:       []Code{"vi", "vie"},
		EnglishName: "Vietnamese",
		Scope:       Individual,
		NativeNames: []string{"Tiếng Việt"}}
	Volapük = Language{
		Codes:       []Code{"vo", "vol"},
		EnglishName: "Volapük",
		Scope:       Individual,
	}
	Walloon = Language{
		Codes:       []Code{"wa", "wln"},
		EnglishName: "Walloon",
		Scope:       Individual,
		NativeNames: []string{"Walon"}}
	Welsh = Language{
		Codes:       []Code{"cy", "cym", "wel"},
		EnglishName: "Welsh",
		Scope:       Individual,
		NativeNames: []string{"Cymraeg"}}
	Wolof = Language{
		Codes:       []Code{WO, WOL},
		EnglishName: "Wolof",
		Scope:       Individual,
		NativeNames: []string{"وࣷلࣷفْ"}}
	Xhosa = Language{
		Codes:       []Code{"xh", "xho"},
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
)
