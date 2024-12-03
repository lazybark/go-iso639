package iso639

var (
	Abkhazian = Language{
		Codes:       []string{"ab", "abk"},
		EnglishName: "Abkhazian",
		Scope:       ScopeIndividual,
		Regions:     []Region{RegionEurasia, RegionGeorgia},
		Scripts:     []Script{ScriptLatin, ScriptCyrillic, ScriptGeorgian},
		Family:      FamilyNorthwestCaucasian,
		NativeNames: []string{"Аҧсуа", "Apsua", "აფსუა", "Abkhaz"}}
	Afar = Language{
		Codes:       []string{"aa", "aar"},
		EnglishName: "Afar",
		Scope:       ScopeIndividual,
		Regions:     []Region{RegionAfrica},
		Type:        TypeLiving,
		Scripts:     []Script{ScriptLatin},
		Family:      FamilyAfroasiatic,
		NativeNames: []string{"Qafar af"}}
	Afrikaans = Language{
		Codes:       []string{"af", "afr"},
		EnglishName: "Afrikaans",
		Scope:       ScopeIndividual,
	}
	Akan = Language{
		Codes:       []string{"ak", "aka"},
		EnglishName: "Akan",
		Scope:       ScopeMacro,
		NativeNames: []string{"Ákán"},
		Variants: []Variant{
			{Code: "fat", EnglishName: "Fanti"},
		}}
	Albanian = Language{
		Codes:       []string{"sq", "sqi", "alb"},
		EnglishName: "Albanian",
		Scope:       ScopeMacro,
		NativeNames: []string{"Shqip"},
		Variants: []Variant{
			{Code: "aae", EnglishName: "Arbëreshë Albanian"},
			{Code: "aat", EnglishName: "Arvanitika Albanian"},
			{Code: "aln", EnglishName: "Gheg Albanian"},
			{Code: "als", EnglishName: "Tosk Albanian"},
		}}
	Arabic = Language{
		Codes:       []string{"ar", "ara"},
		EnglishName: "Arabic",
		Scope:       ScopeMacro,
		NativeNames: []string{"اَلْعَرَبِيَّةُ", "al-ʿarabiyyah"},
		Variants: []Variant{
			{Code: "aao", EnglishName: "Algerian Saharan Arabic"},
			{Code: "abh", EnglishName: "Tajiki Arabic"},
			{Code: "abv", EnglishName: "Baharna Arabic"},
			{Code: "acm", EnglishName: "Mesopotamian Arabic"},
			{Code: "acq", EnglishName: "Ta'izzi-Adeni Arabic"},
			{Code: "acw", EnglishName: "Hijazi Arabic"},
			{Code: "acx", EnglishName: "Omani Arabic"},
			{Code: "acy", EnglishName: "Cypriot Arabic"},
			{Code: "adf", EnglishName: "Dhofari Arabic"},
			{Code: "aeb", EnglishName: "Tunisian Arabic"},
			{Code: "aec", EnglishName: "Saidi Arabic"},
			{Code: "afb", EnglishName: "Gulf Arabic"},
			{Code: "apc", EnglishName: "Levantine Arabic"},
			{Code: "apd", EnglishName: "Sudanese Arabic"},
			{Code: "arb", EnglishName: "Standard Arabic"},
			{Code: "arq", EnglishName: "Algerian Arabic"},
			{Code: "ars", EnglishName: "Najdi Arabic"},
			{Code: "ary", EnglishName: "Moroccan Arabic"},
			{Code: "arz", EnglishName: "Egyptian Arabic"},
			{Code: "auz", EnglishName: "Uzbeki Arabic"},
			{Code: "avl", EnglishName: "Eastern Egyptian Bedawi Arabic"},
			{Code: "ayh", EnglishName: "Hadrami Arabic"},
			{Code: "ayl", EnglishName: "Libyan Arabic"},
			{Code: "ayn", EnglishName: "Sanaani Arabic"},
			{Code: "ayp", EnglishName: "North Mesopotamian Arabic"},
			{Code: "pga", EnglishName: "Sudanese Creole Arabic"},
			{Code: "shu", EnglishName: "Chadian Arabic"},
			{Code: "ssh", EnglishName: "Shihhi Arabic"},
		}}
	Aragonese = Language{
		Codes:       []string{"an", "arg"},
		EnglishName: "Aragonese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Aragonés"}}
	Amharic = Language{
		Codes:       []string{"am", "amh"},
		EnglishName: "Amharic",
		Scope:       ScopeIndividual,
		NativeNames: []string{"አማርኛ", "Amarəñña"}}
	Armenian = Language{
		Codes:       []string{"hy", "hye", "arm"},
		EnglishName: "Armenian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Հայերեն", "Hayeren"}}
	Assamese = Language{
		Codes:       []string{"as", "asm"},
		EnglishName: "Assamese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"অসমীয়া", "Ôxômiya"}}
	Avaric = Language{
		Codes:       []string{"av", "ava"},
		EnglishName: "Avaric",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Авар мацӏ", "اوار ماض", "Avar maz"}}
	Avestan = Language{
		Codes:       []string{"ae", "ave"},
		EnglishName: "Avestan",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Upastawakaēna"}}
	Aymara = Language{
		Codes:       []string{"ay", "aym"},
		EnglishName: "Aymara",
		Scope:       ScopeMacro,
		NativeNames: []string{"Aymara"},
		Variants: []Variant{
			{Code: "ayc", EnglishName: "Southern Aymara"},
			{Code: "ayr", EnglishName: "Central Aymara"}}}
	Azerbaijani = Language{
		Codes:       []string{"az", "aze"},
		EnglishName: "Azerbaijani",
		Scope:       ScopeMacro,
		NativeNames: []string{"Azərbaycan dili", "آذربایجان دیلی", "Азәрбајҹан дили"},
		Variants: []Variant{
			{Code: "azj", EnglishName: "North Azerbaijani"},
			{Code: "azb", EnglishName: "South Azerbaijani"}}}
	Bambara = Language{
		Codes:       []string{"bm", "bam"},
		EnglishName: "Bambara",
		Scope:       ScopeIndividual,
		NativeNames: []string{"بَمَنَنكَن", "Bamanankan"}}
	Bashkir = Language{
		Codes:       []string{"ba", "bak"},
		EnglishName: "Bashkir",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Башҡорт теле", "Başqort tele"}}
	Basque = Language{
		Codes:       []string{"eu", "eus", "baq"},
		EnglishName: "Basque",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Euskara"}}
	Belarusian = Language{
		Codes:       []string{"be", "bel"},
		EnglishName: "Belarusian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Беларуская мова", "Belaruskaâ mova"}}
	Bengali = Language{
		Codes:       []string{"bn", "ben"},
		EnglishName: "Bengali",
		Scope:       ScopeIndividual,
		NativeNames: []string{"বাংলা", "Bāŋlā"}}
	Bislama = Language{
		Codes:       []string{"bi", "bis"},
		EnglishName: "Bislama",
		Scope:       ScopeIndividual,
	}
	Bosnian = Language{
		Codes:       []string{"bs", "bos"},
		EnglishName: "Bosnian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Босански", "Bosanski"}}
	Breton = Language{
		Codes:       []string{"br", "bre"},
		EnglishName: "Breton",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Brezhoneg"}}
	Bulgarian = Language{
		Codes:       []string{"bg", "bul"},
		EnglishName: "Bulgarian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Български", "Bulgarski"}}
	Burmese = Language{
		Codes:       []string{"my", "mya", "bur"},
		EnglishName: "Burmese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"မြန်မာစာ", "Mrãmācā"}}
	Catalan = Language{
		Codes:       []string{"ca", "cat"},
		EnglishName: "Catalan",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Valencian", "Català", "Valencià"}}
	Chamorro = Language{
		Codes:       []string{"ch", "cha"},
		EnglishName: "Chamorro",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Finu' Chamoru"}}
	Chechen = Language{
		Codes:       []string{"ce", "che"},
		EnglishName: "Chechen",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Нохчийн мотт", "Noxçiyn mott", "Chechnyan", "Chechnian"}}
	Chichewa = Language{
		Codes:       []string{"ny", "nya"},
		EnglishName: "Chichewa",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Chewa", "Nyanja", "Chinyanja"}}
	Chinese = Language{
		Codes:       []string{"zh", "zho", "chi"},
		EnglishName: "Chinese",
		Scope:       ScopeMacro,
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
		Codes:       []string{"cu", "chu"},
		EnglishName: "Church Slavonic",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Славе́нскїй ѧ҆зы́къ"}}
	Chuvash = Language{
		Codes:       []string{"cv", "chv"},
		EnglishName: "Chuvash",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Чăвашла", "Çăvaşla"}}
	Cornish = Language{
		Codes:       []string{"kw", "cor"},
		EnglishName: "Cornish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Kernowek"}}
	Corsican = Language{
		Codes:       []string{"co", "cos"},
		EnglishName: "Corsican",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Corsu"}}
	Cree = Language{
		Codes:       []string{"cr", "cre"},
		EnglishName: "Cree",
		Scope:       ScopeMacro,
		NativeNames: []string{"ᓀᐦᐃᔭᐁᐧᐃᐧᐣ", "Nehiyawewin"},
		Variants: []Variant{
			{Code: "crj", EnglishName: "Southern East Cree"},
			{Code: "crk", EnglishName: "Plains Cree"},
			{Code: "crl", EnglishName: "Northern East Cree"},
			{Code: "crm", EnglishName: "Moose Cree"},
			{Code: "csw", EnglishName: "Swampy Cree"},
			{Code: "cwd", EnglishName: "Woods Cree"}}}
	Croatian = Language{
		Codes:       []string{"hr", "hrv"},
		EnglishName: "Croatian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Hrvatski", "Crovatian"}}
	Czech = Language{
		Codes:       []string{"cs", "ces", "cze"},
		EnglishName: "Czech",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Čeština", "Czechian"}}
	Danish = Language{
		Codes:       []string{"da", "dan"},
		EnglishName: "Danish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Dansk"}}
	Divehi = Language{
		Codes:       []string{"dv", "div"},
		EnglishName: "Divehi",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Dhivehi", "Maldivian", "ދިވެހި "}}
	Dutch = Language{
		Codes:       []string{"nl", "nld", "dut"},
		EnglishName: "Dutch",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Flemish", "Nederlands"},
		Variants: []Variant{
			{Code: "vls", EnglishName: "West Flemish"}}}
	Dzongkha = Language{
		Codes:       []string{"dz", "dzo"},
		EnglishName: "Dzongkha",
		Scope:       ScopeIndividual,
		NativeNames: []string{"རྫོང་ཁ་", "Bhutanese"}}
	English = Language{
		Codes:       []string{"en", "eng"},
		EnglishName: "English",
		Scope:       ScopeIndividual,
	}
	Esperanto = Language{
		Codes:       []string{"eo", "epo"},
		EnglishName: "Esperanto",
		Scope:       ScopeIndividual,
	}
	Estonian = Language{
		Codes:       []string{"et", "est"},
		EnglishName: "Estonian",
		Scope:       ScopeMacro,
		NativeNames: []string{"Eesti keel"},
		Variants: []Variant{
			{Code: "vro", EnglishName: "Võro"},
			{Code: "ekk", EnglishName: "Standard Estonian"}}}
	Ewe = Language{
		Codes:       []string{"ee", "ewe"},
		EnglishName: "Ewe",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Èʋegbe"}}
	Faroese = Language{
		Codes:       []string{"fo", "fao"},
		EnglishName: "Faroese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Føroyskt"}}
	Fijian = Language{
		Codes:       []string{"fj", "fij"},
		EnglishName: "Fijian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Na Vosa Vakaviti"}}
	Finnish = Language{
		Codes:       []string{"fi", "fin"},
		EnglishName: "Finnish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Suomi"}}
	French = Language{
		Codes:       []string{"fr", "fra", "fre"},
		EnglishName: "French",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Français"}}
	WesternFrisian = Language{
		Codes:       []string{"fy", "fry"},
		EnglishName: "Western Frisian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Frysk", "West Frisian", "Frisian", "Fries"}}
	Fulah = Language{
		Codes:       []string{"ff", "ful"},
		EnglishName: "Fulah",
		Scope:       ScopeMacro,
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
		Codes:       []string{"gd", "gla"},
		EnglishName: "Gaelic",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Scottish Gaelic", "Gàidhlig"}}
	Galician = Language{
		Codes:       []string{"gl", "glg"},
		EnglishName: "Galician",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Galego"}}
	Ganda = Language{
		Codes:       []string{"lg", "lug"},
		EnglishName: "Ganda",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Luganda"}}
	Georgian = Language{
		Codes:       []string{"ka", "kat"},
		EnglishName: "Georgian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ქართული", "Kharthuli"}}
	German = Language{
		Codes:       []string{"de", "deu"},
		EnglishName: "German",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Deutsch"}}
	Greek = Language{
		Codes:       []string{"el", "ell"},
		EnglishName: "Greek",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Νέα Ελληνικά", "Néa Ellêniká"}}
	Greenlandic = Language{
		Codes:       []string{"kl", "kal"},
		EnglishName: "Greenlandic",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Kalaallisut"}}
	Guarani = Language{
		Codes:       []string{"gn", "grn"},
		EnglishName: "Guarani",
		Scope:       ScopeMacro,
		NativeNames: []string{"Avañe'ẽ"},
		Variants: []Variant{
			{Code: "gnw", EnglishName: "Western Bolivian Guaraní"},
			{Code: "gug", EnglishName: "Paraguayan Guaraní"},
			{Code: "gui", EnglishName: "Eastern Bolivian Guaraní"},
			{Code: "gun", EnglishName: "Mbyá Guaraní"},
			{Code: "nhd", EnglishName: "Chiripá"}}}
	Gujarati = Language{
		Codes:       []string{"gu", "guj"},
		EnglishName: "Gujarati",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ગુજરાતી", "Gujarātī"}}
	Haitian = Language{
		Codes:       []string{"ht", "hat"},
		EnglishName: "Haitian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Haitian Creole", "Kreyòl ayisyen"}}
	Hausa = Language{
		Codes:       []string{"ha", "hau"},
		EnglishName: "Hausa",
		Scope:       ScopeIndividual,
		NativeNames: []string{"هَرْشٜن هَوْس", "halshen Hausa"}}
	Hebrew = Language{
		Codes:       []string{"he", "heb"},
		EnglishName: "Hebrew",
		Scope:       ScopeIndividual,
		NativeNames: []string{"עברית", "Ivrit"}}
	Herero = Language{
		Codes:       []string{"hz", "her"},
		EnglishName: "Herero",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Otjiherero"}}
	Hindi = Language{
		Codes:       []string{"hi", "hin"},
		EnglishName: "Hindi",
		Scope:       ScopeIndividual,
		NativeNames: []string{"हिन्दी", "Hindī"}}
	HiriMotu = Language{
		Codes:       []string{"ho", "hmo"},
		EnglishName: "Hiri Motu",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Police Motu"}}
	Hungarian = Language{
		Codes:       []string{"hu", "hun"},
		EnglishName: "Hungarian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Magyar nyelv"}}
	Icelandic = Language{
		Codes:       []string{"is", "isl"},
		EnglishName: "Icelandic",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Íslenska"}}
	Ido = Language{
		Codes:       []string{"io", "ido"},
		EnglishName: "Ido",
		Scope:       ScopeIndividual,
	}
	Igbo = Language{
		Codes:       []string{"ig", "ibo"},
		EnglishName: "Igbo",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ásụ̀sụ́ Ìgbò"}}
	Indonesian = Language{
		Codes:       []string{"id", "ind"},
		EnglishName: "Indonesian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"bahasa Indonesia"}}
	Interlingua = Language{
		Codes:       []string{"ia", "ina"},
		EnglishName: "Interlingua",
		Scope:       ScopeIndividual,
	}
	Interlingue = Language{
		Codes:       []string{"ie", "ile"},
		EnglishName: "Interlingue",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Occidental"}}
	Inuktitut = Language{
		Codes:       []string{"iu", "iku"},
		EnglishName: "Inuktitut",
		Scope:       ScopeMacro,
		NativeNames: []string{"ᐃᓄᒃᑎᑐᑦ"},
		Variants: []Variant{
			{Code: "ike", EnglishName: "Eastern Canadian Inuktitut"},
			{Code: "ikt", EnglishName: "Inuinnaqtun"}}}
	Inupiaq = Language{
		Codes:       []string{"ik", "ipk"},
		EnglishName: "Inupiaq",
		Scope:       ScopeMacro,
		NativeNames: []string{"Iñupiaq"},
		Variants: []Variant{
			{Code: "esi", EnglishName: "North Alaskan Inupiatun"},
			{Code: "esk", EnglishName: "Northwest Alaska Inupiatun"}}}
	Irish = Language{
		Codes:       []string{"ga", "gle"},
		EnglishName: "Irish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Gaeilge"}}
	Italian = Language{
		Codes:       []string{"it", "ita"},
		EnglishName: "Italian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Italiano"}}
	Japanese = Language{
		Codes:       []string{"ja", "jpn"},
		EnglishName: "Japanese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"日本語", "Nihongo"}}
	Javanese = Language{
		Codes:       []string{"jv", "jav"},
		EnglishName: "Javanese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ꦧꦱꦗꦮ", "basa Jawa"}}
	Kannada = Language{
		Codes:       []string{"kn", "kan"},
		EnglishName: "Kannada",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ಕನ್ನಡ", "Kannaḍa"}}
	Kanuri = Language{
		Codes:       []string{"kr", "kau"},
		EnglishName: "Kanuri",
		Scope:       ScopeMacro,
		NativeNames: []string{"كَنُرِيِه", "Kànùrí"},
		Variants: []Variant{
			{Code: "kby", EnglishName: "Manga Kanuri"},
			{Code: "knc", EnglishName: "Central Kanuri"},
			{Code: "krt", EnglishName: "Tumari Kanuri"}}}
	Kashmiri = Language{
		Codes:       []string{"ks", "kas"},
		EnglishName: "Kashmiri",
		Scope:       ScopeIndividual,
		NativeNames: []string{"कॉशुर", "كأشُر", "Kosher"}}
	Kazakh = Language{
		Codes:       []string{"kk", "kaz"},
		EnglishName: "Kazakh",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Қазақша", "Qazaqşa"}}
	CentralKhmer = Language{
		Codes:       []string{"km", "khm"},
		EnglishName: "Central Khmer",
		NativeNames: []string{"ខេមរភាសា", "Khémôrôphéasa", "Khmer", "Cambodian"}}
	Kikuyu = Language{
		Codes:       []string{"ki", "kik"},
		EnglishName: "Kikuyu",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Gikuyu", "Gĩgĩkũyũ"}}
	Kinyarwanda = Language{
		Codes:       []string{"rw", "kin"},
		EnglishName: "Kinyarwanda",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Ikinyarwanda"}}
	Kirghiz = Language{
		Codes:       []string{"ky", "kir"},
		EnglishName: "Kirghiz",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Kyrgyz", "Кыргызча", "Kırgızça"}}
	Komi = Language{
		Codes:       []string{"kv", "kom"},
		EnglishName: "Komi",
		Scope:       ScopeMacro,
		NativeNames: []string{"Коми кыв", "Zyran", "Zyrian"},
		Variants: []Variant{
			{Code: "koi", EnglishName: "Komi-Permyak"},
			{Code: "kpv", EnglishName: "Komi-Zyrian"}}}
	Kongo = Language{
		Codes:       []string{"kg", "kon"},
		EnglishName: "Kongo",
		Scope:       ScopeMacro,
		NativeNames: []string{"Kikongo"},
		Variants: []Variant{
			{Code: "kng", EnglishName: "Koongo"},
			{Code: "kwy", EnglishName: "San Salvador Kongo"},
			{Code: "ldi", EnglishName: "Laari"}}}
	Korean = Language{
		Codes:       []string{"ko", "kor"},
		EnglishName: "Korean",
		Scope:       ScopeIndividual,
		NativeNames: []string{"한국어", "조선말", "Hangugeo", "Chosŏnmal"}}
	Kuanyama = Language{
		Codes:       []string{"kj", "kua"},
		EnglishName: "Kuanyama",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Kwanyama"}}
	Kurdish = Language{
		Codes:       []string{"ku", "kur"},
		EnglishName: "Kurdish",
		Scope:       ScopeMacro,
		NativeNames: []string{"Kurdî", "کوردی"},
		Variants: []Variant{
			{Code: "ckb", EnglishName: "Central Kurdish"},
			{Code: "kmr", EnglishName: "Northern Kurdish"},
			{Code: "sdh", EnglishName: "Southern Kurdish"}}}
	Lao = Language{
		Codes:       []string{"lo", "lao"},
		EnglishName: "Lao",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ພາສາລາວ", "phasa Lao"}}
	Latin = Language{
		Codes:       []string{"la", "lat"},
		EnglishName: "Latin",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Latinum"}}
	Latvian = Language{
		Codes:       []string{"lv", "lav"},
		EnglishName: "Latvian",
		Scope:       ScopeMacro,
		NativeNames: []string{"Latviski"},
		Variants: []Variant{
			{Code: "ltg", EnglishName: "Latgalian"},
			{Code: "lvs", EnglishName: "Standard Latvian"}}}
	Limburgan = Language{
		Codes:       []string{"li", "lim"},
		EnglishName: "Limburgan",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Limburger", "Limburgish", "Lèmburgs"}}
	Lingala = Language{
		Codes:       []string{"ln", "lin"},
		EnglishName: "Lingala",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Lingála"}}
	Lithuanian = Language{
		Codes:       []string{"lt", "lit"},
		EnglishName: "Lithuanian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Lietuviškai"}}
	LubaKatanga = Language{
		Codes:       []string{"lu", "lub"},
		EnglishName: "Luba-Katanga",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Kiluba", "Luba-Shaba"}}
	Luxembourgish = Language{
		Codes:       []string{"lb", "ltz"},
		EnglishName: "Luxembourgish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Letzeburgesch", "Lëtzebuergesch", "Luxembourgian"}}
	Macedonian = Language{
		Codes:       []string{"mk", "mkd", "mac"},
		EnglishName: "Macedonian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Македонски", "Makedonski"}}
	Malagasy = Language{
		Codes:       []string{"mg", "mlg"},
		EnglishName: "Malagasy",
		Scope:       ScopeMacro,
		NativeNames: []string{"مَلَغَسِ"},
		Variants: []Variant{
			{Code: "bhr", EnglishName: "Bara Malagasy"},
			{Code: "bmm", EnglishName: "Northern Betsimisaraka Malagasy"},
			{Code: "bzc", EnglishName: "Southern Betsimisaraka Malagasy"},
			{Code: "msh", EnglishName: "Masikoro Malagasy"},
			{Code: "plt", EnglishName: "Plateau Malagasy"},
			{Code: "skg", EnglishName: "Sakalava Malagasy"},
			{Code: "tdx", EnglishName: "Tandroy-Mahafaly Malagasy"},
			{Code: "tkg", EnglishName: "Tesaka Malagasy"},
			{Code: "txy", EnglishName: "Tanosy Malagasy"},
			{Code: "xmv", EnglishName: "Antankarana Malagasy"},
			{Code: "xmw", EnglishName: "Tsimihety Malagasy"}}}
	Malay = Language{
		Codes:       []string{"ms", "msa"},
		EnglishName: "Malay",
		Scope:       ScopeMacro,
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
		Codes:       []string{"ml", "mal"},
		EnglishName: "Malayalam",
		Scope:       ScopeIndividual,
		NativeNames: []string{"മലയാളം", "Malayāļã"}}
	Maltese = Language{
		Codes:       []string{"mt", "mlt"},
		EnglishName: "Maltese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Malti"}}
	Manx = Language{
		Codes:       []string{"gv", "glv"},
		EnglishName: "Manx",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Gaelg", "Gailck"}}
	Maori = Language{
		Codes:       []string{"mi", "mri", "mao"},
		EnglishName: "Maori",
		Scope:       ScopeIndividual,
		NativeNames: []string{"reo Māori"}}
	Marathi = Language{
		Codes:       []string{"mr", "mar"},
		EnglishName: "Marathi",
		Scope:       ScopeIndividual,
		NativeNames: []string{"मराठी", "Marāṭhī", "Maharashtran"}}
	Marshallese = Language{
		Codes:       []string{"mh", "mah"},
		EnglishName: "Marshallese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"kajin M̧ajel‌̧", "Ebon"}}
	Mongolian = Language{
		Codes:       []string{"mn", "mon"},
		EnglishName: "Mongolian",
		Scope:       ScopeMacro,
		NativeNames: []string{"ᠮᠣᠩᠭᠣᠯ ᠬᠡᠯᠡ", "Монгол хэл", "Mongol xel"},
		Variants: []Variant{
			{Code: "khk", EnglishName: "Halh Mongolian"},
			{Code: "mvf", EnglishName: "Peripheral Mongolian"}}}
	Nauru = Language{
		Codes:       []string{"na", "nau"},
		EnglishName: "Nauru",
		Scope:       ScopeIndividual,
		NativeNames: []string{"dorerin Naoe", "Nauruan"}}
	Navajo = Language{
		Codes:       []string{"nv", "nav"},
		EnglishName: "Navajo",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Navaho", "Diné bizaad", "Naabeehó bizaad"}}
	NorthNdebele = Language{
		Codes:       []string{"nd", "nde"},
		EnglishName: "North Ndebele",
		Scope:       ScopeIndividual,
		NativeNames: []string{"isiNdebele"}}
	SouthNdebele = Language{
		Codes:       []string{"nr", "nbl"},
		EnglishName: "South Ndebele",
		Scope:       ScopeIndividual,
		NativeNames: []string{"isiNdebele"}}
	Ndonga = Language{
		Codes:       []string{"ng", "ndo"},
		EnglishName: "Ndonga",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Oshiwambo"}}
	Nepali = Language{
		Codes:       []string{"ne", "nep"},
		EnglishName: "Nepali",
		Scope:       ScopeMacro,
		NativeNames: []string{"नेपाली", "Nepālī"},
		Variants: []Variant{
			{Code: "dty", EnglishName: "Dotyali"},
			{Code: "npi", EnglishName: "Nepali (individual language)"}}}
	Norwegian = Language{
		Codes:       []string{"no", "nor"},
		EnglishName: "Norwegian",
		Scope:       ScopeMacro,
		NativeNames: []string{"Norsk"},
		Variants: []Variant{
			{Code: "nno", EnglishName: "Norwegian Nynorsk"},
			{Code: "nob", EnglishName: "Norwegian Bokmål"}}}
	Occitan = Language{
		Codes:       []string{"oc", "oci"},
		EnglishName: "Occitan",
		Scope:       ScopeIndividual,
	}
	Ojibwa = Language{
		Codes:       []string{"oj", "oji"},
		EnglishName: "Ojibwa",
		Scope:       ScopeMacro,
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
		Codes:       []string{"or", "ori"},
		EnglishName: "Oriya",
		Scope:       ScopeMacro,
		NativeNames: []string{"ଓଡ଼ିଆ", "Odia"},
		Variants: []Variant{
			{Code: "ory", EnglishName: "Odia"},
			{Code: "spv", EnglishName: "Sambalpuri"}}}
	Oromo = Language{
		Codes:       []string{"om", "orm"},
		EnglishName: "Oromo",
		Scope:       ScopeMacro,
		NativeNames: []string{"afaan Oromoo"},
		Variants: []Variant{
			{Code: "gax", EnglishName: "Borana-Arsi-Guji Oromo"},
			{Code: "gaz", EnglishName: "West Central Oromo"},
			{Code: "hae", EnglishName: "Eastern Oromo"},
			{Code: "orc", EnglishName: "Orma"}}}
	Ossetian = Language{
		Codes:       []string{"os", "oss"},
		EnglishName: "Ossetian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Ossetic", "дигорон Ӕвзаг", "digoron Ævzag"}}
	Pali = Language{
		Codes:       []string{"pi", "pli"},
		EnglishName: "Pali",
		Scope:       ScopeIndividual,
	}
	Pashto = Language{
		Codes:       []string{"ps", "pus"},
		EnglishName: "Pashto",
		Scope:       ScopeMacro,
		NativeNames: []string{"Pushto", "پښتو", "Pax̌tow"},
		Variants: []Variant{
			{Code: "pbt", EnglishName: "Southern Pashto"},
			{Code: "pbu", EnglishName: "Northern Pashto"},
			{Code: "pst", EnglishName: "Central Pashto"}}}
	Persian = Language{
		Codes:       []string{"fa", "fas", "per"},
		EnglishName: "Persian",
		Scope:       ScopeMacro,
		NativeNames: []string{"فارسی", "Fārsiy"},
		Variants: []Variant{
			{Code: "pes", EnglishName: "Iranian Persian"},
			{Code: "prs", EnglishName: "Dari"}}}
	Polish = Language{
		Codes:       []string{"pl", "pol"},
		EnglishName: "Polish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Polski"}}
	Portuguese = Language{
		Codes:       []string{"pt", "por"},
		EnglishName: "Portuguese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Português"}}
	Punjabi = Language{
		Codes:       []string{"pa", "pan"},
		EnglishName: "Punjabi",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ਪੰਜਾਬੀ", "Panjabi", "پنجابی", "Pãjābī"}}
	Quechua = Language{
		Codes:       []string{"qu", "que"},
		EnglishName: "Quechua",
		Scope:       ScopeMacro,
		NativeNames: []string{"Runa simi", "kichwa simi", "Nuna shimi"},
		Variants: []Variant{
			{Code: "qub", EnglishName: "Huallaga Huánuco Quechua"},
			{Code: "qud", EnglishName: "Calderón Highland Quichua"},
			{Code: "quf", EnglishName: "Lambayeque Quechua"},
			{Code: "qug", EnglishName: "Chimborazo Highland Quichua"},
			{Code: "quh", EnglishName: "South Bolivian Quechua"},
			{Code: "quk", EnglishName: "Chachapoyas Quechua"},
			{Code: "qul", EnglishName: "North Bolivian Quechua"},
			{Code: "qup", EnglishName: "Southern Pastaza Quechua"},
			{Code: "qur", EnglishName: "Yanahuanca Pasco Quechua"},
			{Code: "qus", EnglishName: "Santiago del Estero Quichua"},
			{Code: "quw", EnglishName: "Tena Lowland Quichua"},
			{Code: "qux", EnglishName: "Yauyos Quechua"},
			{Code: "quy", EnglishName: "Ayacucho Quechua"},
			{Code: "quz", EnglishName: "Cusco Quechua"},
			{Code: "qva", EnglishName: "Ambo-Pasco Quechua"},
			{Code: "qvc", EnglishName: "Cajamarca Quechua"},
			{Code: "qve", EnglishName: "Eastern Apurímac Quechua"},
			{Code: "qvh", EnglishName: "Huamalíes-Dos de Mayo Huánuco Quechua"},
			{Code: "qvi", EnglishName: "Imbabura Highland Quichua"},
			{Code: "qvj", EnglishName: "Loja Highland Quichua"},
			{Code: "qvl", EnglishName: "Cajatambo North Lima Quechua"},
			{Code: "qvm", EnglishName: "Margos-Yarowilca-Lauricocha Quechua"}, {
				Code: "qvn", EnglishName: "North Junín Quechua"},
			{Code: "qvo", EnglishName: "Napo Lowland Quechua"},
			{Code: "qvp", EnglishName: "Pacaraos Quechua"},
			{Code: "qvs", EnglishName: "San Martín Quechua"},
			{Code: "qvw", EnglishName: "Huaylla Wanca Quechua"},
			{Code: "qvz", EnglishName: "Northern Pastaza Quichua"},
			{Code: "qwa", EnglishName: "Corongo Ancash Quechua"},
			{Code: "qwc", EnglishName: "Classical Quechua"},
			{Code: "qwh", EnglishName: "Huaylas Ancash Quechua"},
			{Code: "qws", EnglishName: "Sihuas Ancash Quechua"},
			{Code: "qxa", EnglishName: "Chiquián Ancash Quechua"},
			{Code: "qxc", EnglishName: "Chincha Quechua"},
			{Code: "qxh", EnglishName: "Panao Huánuco Quechua"},
			{Code: "qxl", EnglishName: "Salasaca Highland Quichua"},
			{Code: "qxn", EnglishName: "Northern Conchucos Ancash Quechua"},
			{Code: "qxo", EnglishName: "Southern Conchucos Ancash Quechua"},
			{Code: "qxp", EnglishName: "Puno Quechua"},
			{Code: "qxr", EnglishName: "Cañar Highland Quichua"},
			{Code: "qxt", EnglishName: "Santa Ana de Tusi Pasco Quechua"},
			{Code: "qxu", EnglishName: "Arequipa-La Unión Quechua"},
			{Code: "qxw", EnglishName: "Jauja Wanca Quechua"}}}
	Romanian = Language{
		Codes:       []string{"ro", "ron", "rum"},
		EnglishName: "Romanian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Românã", "Moldavian", " Moldovan"},
		Variants: []Variant{
			{Code: "mo", EnglishName: "Moldavian"},
			{Code: "mol", EnglishName: "Moldavian"}}}
	Romansh = Language{
		Codes:       []string{"rm", "roh"},
		EnglishName: "Romansh",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Rumantsch", "Rumàntsch", "Romauntsch", "Romontsch"}}
	Rundi = Language{
		Codes:       []string{"rn", "run"},
		EnglishName: "Rundi",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Ikirundi"}}
	Russian = Language{
		Codes:       []string{"ru", "rus"},
		EnglishName: "Russian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Русский язык", "Russkiĭ âzyk"}}
	NorthernSami = Language{
		Codes:       []string{"se", "sme"},
		EnglishName: "Northern Sami",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Davvisámegiella"}}
	Samoan = Language{
		Codes:       []string{"sm", "smo"},
		EnglishName: "Samoan",
		Scope:       ScopeIndividual,
		NativeNames: []string{"gagana Sāmoa", "Sāmoa"}}
	Sango = Language{
		Codes:       []string{"sg", "sag"},
		EnglishName: "Sango",
		Scope:       ScopeIndividual,
		NativeNames: []string{"yângâ tî Sängö"}}
	Sanskrit = Language{
		Codes:       []string{"sa", "san"},
		EnglishName: "Sanskrit",
		Scope:       ScopeMacro,
		NativeNames: []string{"संस्कृतम्", "Saṃskṛtam"},
		Variants: []Variant{
			{Code: "cls", EnglishName: "Classical Sanskrit"},
			{Code: "vsn", EnglishName: "Vedic Sanskrit"}}}
	Sardinian = Language{
		Codes:       []string{"sc", "srd"},
		EnglishName: "Sardinian",
		Scope:       ScopeMacro,
		NativeNames: []string{"Sardu"},
		Variants: []Variant{
			{Code: "sdc", EnglishName: "Sassarese Sardinian"},
			{Code: "sdn", EnglishName: "Gallurese Sardinian"},
			{Code: "src", EnglishName: "Logudorese Sardinian"},
			{Code: "sro", EnglishName: "Campidanese Sardinian"}}}
	Serbian = Language{
		Codes:       []string{"sr", "srp"},
		EnglishName: "Serbian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Српски", "Srpski"}}
	Shona = Language{
		Codes:       []string{"sn", "sna"},
		EnglishName: "Shona",
		Scope:       ScopeIndividual,
		NativeNames: []string{"chiShona"}}
	Sindhi = Language{
		Codes:       []string{"sd", "snd"},
		EnglishName: "Sindhi",
		Scope:       ScopeIndividual,
		NativeNames: []string{"سنڌي", "सिन्धी", "Sindhī"}}
	Sinhala = Language{
		Codes:       []string{"si", "sin"},
		EnglishName: "Sinhala",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Sinhalese", "සිංහල", "Sinhala"}}
	Slovak = Language{
		Codes:       []string{"sk", "slk", "slo"},
		EnglishName: "Slovak",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Slovenčina"}}
	Slovenian = Language{
		Codes:       []string{"sl", "slv"},
		EnglishName: "Slovenian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Slovenščina"}}
	Somali = Language{
		Codes:       []string{"so", "som"},
		EnglishName: "Somali",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Soomaali", "𐒈𐒝𐒑𐒛𐒐𐒘", "سٝومالِ"}}
	SouthernSotho = Language{
		Codes:       []string{"st", "sot"},
		EnglishName: "Southern Sotho",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Sesotho"}}
	Spanish = Language{
		Codes:       []string{"es", "spa"},
		EnglishName: "Spanish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Castilian", "Español", "Castellano"}}
	Sundanese = Language{
		Codes:       []string{"su", "sun"},
		EnglishName: "Sundanese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"basa Sunda", "ᮘᮞ ᮞᮥᮔ᮪ᮓ", "بَاسَا سُوْندَا"}}
	Swahili = Language{
		Codes:       []string{"sw", "swa"},
		EnglishName: "Swahili",
		Scope:       ScopeMacro,
		NativeNames: []string{"Kiswahili", "كِسوَحِيلِ"},
		Variants: []Variant{
			{Code: "swc", EnglishName: "Congo Swahili"},
			{Code: "swh", EnglishName: "Swahili (individual language)"}}}
	Swati = Language{
		Codes:       []string{"ss", "ssw"},
		EnglishName: "Swati",
		Scope:       ScopeIndividual,
		NativeNames: []string{"SiSwati"}}
	Swedish = Language{
		Codes:       []string{"sv", "swe"},
		EnglishName: "Swedish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Svenska"}}
	Tagalog = Language{
		Codes:       []string{"tl", "tgl"},
		EnglishName: "Tagalog",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Wikang Tagalog"}}
	Tahitian = Language{
		Codes:       []string{"ty", "tah"},
		EnglishName: "Tahitian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"reo Tahiti", "Reo Mā`ohi"}}
	Tajik = Language{
		Codes:       []string{"tg", "tgk"},
		EnglishName: "Tajik",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Тоҷикӣ", "Tojikī"}}
	Tamil = Language{
		Codes:       []string{"ta", "tam"},
		EnglishName: "Tamil",
		Scope:       ScopeIndividual,
		NativeNames: []string{"தமிழ்", "Tamiḻ"}}
	Tatar = Language{
		Codes:       []string{"tt", "tat"},
		EnglishName: "Tatar",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Татар теле", "Tatar tele", "تاتار تئلئ"}}
	Telugu = Language{
		Codes:       []string{"te", "tel"},
		EnglishName: "Telugu",
		Scope:       ScopeIndividual,
		NativeNames: []string{"తెలుగు"}}
	Thai = Language{
		Codes:       []string{"th", "tha"},
		EnglishName: "Thai",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ภาษาไทย", "Phasa Thai"}}
	Tibetan = Language{
		Codes:       []string{"bo", "bod", "tib"},
		EnglishName: "Tibetan",
		Scope:       ScopeIndividual,
		NativeNames: []string{"བོད་སྐད་", "Bodskad", "ལྷ་སའི་སྐད་", "Lhas'iskad"}}
	Tigrinya = Language{
		Codes:       []string{"ti", "tir"},
		EnglishName: "Tigrinya",
		Scope:       ScopeIndividual,
		NativeNames: []string{"ትግርኛ", "Təgrəñña"}}
	Tonga = Language{
		Codes:       []string{"to", "ton"},
		EnglishName: "Tonga",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Tonga Islands", "lea faka-Tonga"}}
	Tsonga = Language{
		Codes:       []string{"ts", "tso"},
		EnglishName: "Tsonga",
		Scope:       ScopeIndividual,
	}
	Tswana = Language{
		Codes:       []string{"tn", "tsn"},
		EnglishName: "Tswana",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Setswana", "Sechuana"}}
	Turkish = Language{
		Codes:       []string{"tr", "tur"},
		EnglishName: "Turkish",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Türkçe"}}
	Turkmen = Language{
		Codes:       []string{"tk", "tuk"},
		EnglishName: "Turkmen",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Türkmençe", "Түркменче", "تۆرکمنچه"}}
	Twi = Language{
		Codes:       []string{"tw", "twi"},
		EnglishName: "Twi",
		Scope:       ScopeIndividual,
	}
	Uighur = Language{
		Codes:       []string{"ug", "uig"},
		EnglishName: "Uighur",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Uyghur", "ئۇيغۇر تىلى", "Уйғур тили", "Uyƣur tili"}}
	Ukrainian = Language{
		Codes:       []string{"uk", "ukr"},
		EnglishName: "Ukrainian",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Українська", "Ukraїnska"}}
	Urdu = Language{
		Codes:       []string{"ur", "urd"},
		EnglishName: "Urdu",
		Scope:       ScopeIndividual,
		NativeNames: []string{"اُردُو", "Urduw"}}
	Uzbek = Language{
		Codes:       []string{"uz", "uzb"},
		EnglishName: "Uzbek",
		Scope:       ScopeMacro,
		NativeNames: []string{"Ózbekça", "ўзбекча", "ئوزبېچه"},
		Variants: []Variant{
			{Code: "uzn", EnglishName: "Northern Uzbek"},
			{Code: "uzs", EnglishName: "Southern Uzbek"}}}
	Venda = Language{
		Codes:       []string{"ve", "ven"},
		EnglishName: "Venda",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Tshivenḓa"}}
	Vietnamese = Language{
		Codes:       []string{"vi", "vie"},
		EnglishName: "Vietnamese",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Tiếng Việt"}}
	Volapük = Language{
		Codes:       []string{"vo", "vol"},
		EnglishName: "Volapük",
		Scope:       ScopeIndividual,
	}
	Walloon = Language{
		Codes:       []string{"wa", "wln"},
		EnglishName: "Walloon",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Walon"}}
	Welsh = Language{
		Codes:       []string{"cy", "cym", "wel"},
		EnglishName: "Welsh",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Cymraeg"}}
	Wolof = Language{
		Codes:       []string{"wo", "wol"},
		EnglishName: "Wolof",
		Scope:       ScopeIndividual,
		NativeNames: []string{"وࣷلࣷفْ"}}
	Xhosa = Language{
		Codes:       []string{"xh", "xho"},
		EnglishName: "Xhosa",
		Scope:       ScopeIndividual,
		NativeNames: []string{"isiXhosa"}}
	SichuanYi = Language{
		Codes:       []string{"ii", "iii"},
		EnglishName: "Sichuan Yi",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Nuosu", "ꆈꌠꉙ", "Nuosuhxop"}}
	Yoruba = Language{
		Codes:       []string{"yo", "yor"},
		EnglishName: "Yoruba",
		Scope:       ScopeIndividual,
		NativeNames: []string{"èdè Yorùbá"}}
	Yiddish = Language{
		Codes:       []string{"yi", "yid"},
		EnglishName: "Yiddish",
		Scope:       ScopeMacro,
		NativeNames: []string{"ייִדיש", "Yidiš"},
		Variants: []Variant{
			{Code: "ydd", EnglishName: "Eastern Yiddish"},
			{Code: "yih", EnglishName: "Western Yiddish"}}}
	Zhuang = Language{
		Codes:       []string{"za", "zha"},
		EnglishName: "Zhuang",
		Scope:       ScopeMacro,
		NativeNames: []string{"Chuang", "話僮", "Vahcuengh"},
		Variants: []Variant{
			{Code: "zch", EnglishName: "Central Hongshuihe Zhuang"},
			{Code: "zeh", EnglishName: "Eastern Hongshuihe Zhuang"},
			{Code: "zgb", EnglishName: "Guibei Zhuang"},
			{Code: "zgm", EnglishName: "Minz Zhuang"},
			{Code: "zgn", EnglishName: "Guibian Zhuang"},
			{Code: "zhd", EnglishName: "Dai Zhuang"},
			{Code: "zhn", EnglishName: "Nong Zhuang"},
			{Code: "zlj", EnglishName: "Liujiang Zhuang"},
			{Code: "zln", EnglishName: "Lianshan Zhuang"},
			{Code: "zlq", EnglishName: "Liuqian Zhuang"},
			{Code: "zqe", EnglishName: "Qiubei Zhuang"},
			{Code: "zyb", EnglishName: "Yongbei Zhuang"},
			{Code: "zyg", EnglishName: "Yang Zhuang"},
			{Code: "zyj", EnglishName: "Youjiang Zhuang"},
			{Code: "zyn", EnglishName: "Yongnan Zhuang"},
			{Code: "zzj", EnglishName: "Zuojiang Zhuang"},
		}}
	Zulu = Language{
		Codes:       []string{"zu", "zul"},
		EnglishName: "Zulu",
		Scope:       ScopeIndividual,
		NativeNames: []string{"isiZulu"}}
	AncientGreek = Language{
		Codes:       []string{"grc"},
		EnglishName: "Ancient Greek",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Ἀρχαία ἑλληνικὴ", "Archaía ellēnikḗ"}}
	Zapotec = Language{
		Codes:       []string{"zap"},
		EnglishName: "Zapotec",
		Scope:       ScopeIndividual,
		NativeNames: []string{"Didxsaj"}}
	Blissymbols = Language{
		Codes:       []string{"zbl"},
		EnglishName: "Blissymbols",
		Scope:       ScopeIndividual,
	}
)
