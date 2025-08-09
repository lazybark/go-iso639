package iso639

import (
	"slices"
	"testing"
)

func TestByCode(t *testing.T) {
	tests := []struct {
		code     string
		expected string
		found    bool
	}{
		{"en", "English", true},
		{"zh", "Chinese", true},
		{"la", "Latin", true},
		{"notexists", "", false},
	}

	for _, test := range tests {
		lang := ByCode(test.code)
		found := lang != nil
		if found != test.found || (found && lang.EnglishName != test.expected) {
			t.Errorf("ByCode(%s) = %v, %v; want %v, %v", test.code, lang, found, test.expected, test.found)
		}
	}
}

func TestByName(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		found    bool
	}{
		{"English", "English", true},
		{"中文", "Chinese", true},
		{"Latinum", "Latin", true},
		{"Unknown", "", false},
	}

	for _, test := range tests {
		lang := ByName(test.name)
		found := lang != nil
		if found != test.found || (found && lang.EnglishName != test.expected) {
			t.Errorf("ByName(%s) = %v, %v; want %v, %v", test.name, lang, found, test.expected, test.found)
		}
	}
}

func TestByType(t *testing.T) {
	var living, ancient, constructed int

	for _, lang := range LanguageMap {
		switch lang.Type {
		case Living:
			living++
		case Ancient:
			ancient++
		case Constructed:
			constructed++
		}
	}

	tests := []struct {
		typ      Type
		expected int
	}{
		{Living, living},
		{typ: Ancient, expected: ancient},
		{Constructed, constructed},
	}

	for _, test := range tests {
		langs := ByType(test.typ)
		if len(langs) != test.expected {
			t.Errorf("ByType(%s) = %d; want %d", test.typ, len(langs), test.expected)
		}
	}
}

func TestByScope(t *testing.T) {
	var individual, macro int

	for _, lang := range LanguageMap {
		switch lang.Scope {
		case Individual:
			individual++
		case Macro:
			macro++
		}
	}

	tests := []struct {
		scope    Scope
		expected int
	}{
		{Individual, individual},
		{Macro, macro},
	}

	for _, test := range tests {
		langs := ByScope(test.scope)
		if len(langs) != test.expected {
			t.Errorf("ByScope(%s) = %d; want %d", test.scope, len(langs), test.expected)
		}
	}
}

func TestByFamily(t *testing.T) {
	var indoeuropean, sinoTibetan int

	for _, lang := range LanguageMap {
		switch lang.Family {
		case IndoEuropean:
			indoeuropean++
		case SinoTibetan:
			sinoTibetan++
		}
	}

	tests := []struct {
		family   string
		expected int
	}{
		{string(IndoEuropean), indoeuropean},
		{string(SinoTibetan), sinoTibetan},
		{"Unknown", 0},
	}

	for _, test := range tests {
		langs := ByFamily(test.family)
		if len(langs) != test.expected {
			t.Errorf("ByFamily(%s) = %d; want %d", test.family, len(langs), test.expected)
		}
	}
}

func TestByScript(t *testing.T) {
	var latin, chinese int

	for _, lang := range LanguageMap {
		if slices.Contains(lang.Scripts, LatinScript) {
			latin++
		}

		if slices.Contains(lang.Scripts, ChineseScript) {
			chinese++
		}
	}

	tests := []struct {
		script   string
		expected int
	}{
		{string(LatinScript), latin},
		{string(ChineseScript), chinese},
		{"Unknown", 0},
	}

	for _, test := range tests {
		langs := ByScript(test.script)
		if len(langs) != test.expected {
			t.Errorf("ByScript(%s) = %d; want %d", test.script, len(langs), test.expected)
		}
	}
}

func TestByRegion(t *testing.T) {
	var europe, asia int

	for _, lang := range LanguageMap {
		if slices.Contains(lang.Regions, Europe) {
			europe++
		}

		if slices.Contains(lang.Regions, Asia) {
			asia++
		}
	}

	tests := []struct {
		region   string
		expected int
	}{
		{string(Europe), europe},
		{string(Asia), asia},
		{"Unknown", 0},
	}

	for _, test := range tests {
		langs := ByRegion(test.region)
		if len(langs) != test.expected {
			t.Errorf("ByRegion(%s) = %d; want %d", test.region, len(langs), test.expected)
		}
	}
}

func TestGetAllLanguages(t *testing.T) {
	langs := GetAllLanguages()
	expectedCount := len(LanguageMap) // Count will change based on language data.

	if len(langs) != expectedCount {
		t.Errorf("GetAllLanguages() = %d; want %d", len(langs), expectedCount)
	}
}

func TestCaseInsensitiveSearch(t *testing.T) {
	tests := []struct {
		code     string
		expected string
	}{
		{"EN", "English"},
		{"en", "English"},
		{"En", "English"},
	}

	for _, test := range tests {
		lang := ByCodeIgnoreCase(test.code)
		if lang == nil || lang.EnglishName != test.expected {
			t.Errorf("ByCodeIgnoreCase(%s) failed", test.code)
		}
	}
}

func BenchmarkByCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByCode("en")
	}
}

func BenchmarkByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByName("English")
	}
}

func BenchmarkByNameIgnoreCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByNameIgnoreCase("english")
	}
}
