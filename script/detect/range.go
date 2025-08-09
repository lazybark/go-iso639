package detect

import "github.com/lazybark/go-iso639"

type ScriptRange struct {
	Start  rune          // Start of the Unicode range
	End    rune          // End of the Unicode range
	Name   string        // Script name, e.g., "Latin", "Cyrillic", "Arabic"
	Script iso639.Script // Script linked to iso639 package
}

var scriptRanges = []ScriptRange{
	{0x0041, 0x005A, "Latin", iso639.LatinScript},                     // A-Z
	{0x0061, 0x007A, "Latin", iso639.LatinScript},                     // a-z
	{0x00C0, 0x00D6, "Latin", iso639.LatinScript},                     // À-Ö
	{0x00D8, 0x00F6, "Latin", iso639.LatinScript},                     // Ø-ö
	{0x00F8, 0x00FF, "Latin", iso639.LatinScript},                     // ø-ÿ
	{0x0100, 0x01BF, "Latin", iso639.LatinScript},                     // Ā-ɏ
	{0x01C0, 0x024F, "Latin", iso639.LatinScript},                     // ǀ-ɏ
	{0x1E00, 0x1EFF, "Latin", iso639.LatinScript},                     // Ḁ-ỿ
	{0x0250, 0x02AF, "Latin", iso639.LatinScript},                     // ɐ-ʯ
	{0x0370, 0x03FF, "Greek", iso639.GreekScript},                     // Ͱ-Ͽ
	{0x0400, 0x04FF, "Cyrillic", iso639.CyrillicScript},               // Ѐ-ӿ
	{0x0500, 0x052F, "Cyrillic", iso639.CyrillicScript},               // Ԁ-ԯ
	{0x0530, 0x058F, "Armenian", iso639.ArmenianScript},               // ԰-Տ
	{0x0590, 0x05FF, "Hebrew", iso639.HebrewScript},                   // ׀-׿
	{0x0600, 0x06FF, "Arabic", iso639.ArabicScript},                   // ؀-ۿ
	{0x0700, 0x074F, "Syriac", iso639.SyriacScript},                   // ܀-ݏ
	{0x0780, 0x07BF, "Thaana", iso639.ThaanaScript},                   // ހ-޿
	{0x0900, 0x097F, "Devanagari", iso639.DevanagariScript},           // ऀ-ॿ
	{0x0980, 0x09FF, "Bengali", iso639.BengaliAssameseScript},         // ঀ-৿
	{0x0A00, 0x0A7F, "Gurmukhi", iso639.GurmukhiScript},               // ਀-੿
	{0x0A80, 0x0AFF, "Gujarati", iso639.GujaratiScript},               // ઀-૿
	{0x0B00, 0x0B7F, "Oriya", iso639.OdiaScript},                      // ଀-୿
	{0x0B80, 0x0BFF, "Tamil", iso639.TamilScript},                     // ஀-௿
	{0x0C00, 0x0C7F, "Telugu", iso639.TeluguScript},                   // ఀ-౿
	{0x0C80, 0x0CFF, "Kannada", iso639.KannadaScript},                 // ಀ-೿
	{0x0D00, 0x0D7F, "Malayalam", iso639.MalayalamScript},             // ഀ-ൿ
	{0x0D80, 0x0DFF, "Sinhala", iso639.SinhalaScript},                 // ඀-෿
	{0x0E00, 0x0E7F, "Thai", iso639.ThaiScript},                       // ก-฿
	{0x0E80, 0x0EFF, "Lao", iso639.LaoScript},                         // ຀-໿
	{0x0F00, 0x0FFF, "Tibetan", iso639.TibetanScript},                 // ༀ-࿿
	{0x1000, 0x109F, "Myanmar", iso639.BurmeseScript},                 // က-႟
	{0x10A0, 0x10FF, "Georgian", iso639.GeorgianScript},               // Ⴀ-ჿ
	{0x1100, 0x11FF, "Hangul", iso639.HangulScript},                   // ᄀ-퟿
	{0x1200, 0x137F, "Ethiopic", iso639.GeezScript},                   // ሀ-፿
	{0x13A0, 0x13FF, "Cherokee", iso639.CherokeeScript},               // Ꭰ-᏿
	{0x1400, 0x167F, "Canadian Aboriginal", iso639.CanadianSyllabics}, // ᐀-ᙿ
	{0x1680, 0x169F, "Ogham", iso639.OghamScript},                     //  -᚟
	{0x16A0, 0x16FF, "Runic", iso639.AngloSaxonRunes},                 // ᚠ-᛿
	{0x1780, 0x17FF, "Khmer", iso639.KhmerScript},                     // ក-៿
	{0x1800, 0x18AF, "Mongolian", iso639.MongolianScript},             // ᠀-᢯
	{0x1900, 0x194F, "Limbu", iso639.LimbuScript},                     // ᤀ-᥏
	{0x1950, 0x197F, "Tai Le", iso639.TaiLeScript},                    // ᥐ-᥿
	{0x19E0, 0x19FF, "Khmer Symbols", iso639.KhmerScript},             // ហ-ឿ
	{0x1E00, 0x1EFF, "Latin", iso639.LatinScript},
	{0x1F00, 0x1FFF, "Greek", iso639.GreekScript},
	{0x2800, 0x28FF, "Braille Patterns", iso639.BrailleScript},
	{0x3040, 0x309F, "Hiragana", iso639.HiraganaScript},
	{0x30A0, 0x30FF, "Katakana", iso639.KatakanaScript},
	{0x3100, 0x312F, "Bopomofo", iso639.BopomofoScript},
	{0x3130, 0x318F, "Hangul", iso639.HangulScript},
	{0x3190, 0x319F, "Kanbun", iso639.KanbunScript},
	{0x31A0, 0x31BF, "Bopomofo", iso639.BopomofoScript},
	{0xAC00, 0xD7AF, "Hangul", iso639.HangulScript},
	{0xFE70, 0xFEFF, "Arabic", iso639.ArabicScript},
	{0x10300, 0x1032F, "Old Italic", iso639.OldItalicScript},
	{0x10330, 0x1034F, "Gothic", iso639.GothicScript},
	{0x10380, 0x1039F, "Ugaritic", iso639.UgariticScript},
	{0x10400, 0x1044F, "Deseret", iso639.DeseretScript},
	{0x10450, 0x1047F, "Shavian", iso639.ShavianScript},
	{0x10480, 0x104AF, "Osmanya", iso639.OsmanyaScript},
	{0x10800, 0x1083F, "Cypriot Syllabary", iso639.CypriotScript},
}

func GetDominantScript(text string) string {
	scriptCount := make(map[string]int)

	for _, char := range text {
		for _, script := range scriptRanges {
			if char >= script.Start && char <= script.End {
				scriptCount[script.Name]++

				break
			}
		}
	}

	if len(scriptCount) == 0 {
		return ""
	}

	// Find the most frequent script.
	dominantScript := ""
	maxCount := 0

	for script, count := range scriptCount {
		if count > maxCount {
			dominantScript = script
			maxCount = count
		}
	}

	return dominantScript
}
