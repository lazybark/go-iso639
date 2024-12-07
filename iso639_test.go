package iso639

import "testing"

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
		lang, found := ByCode(test.code)
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
		lang, found := ByName(test.name)
		if found != test.found || (found && lang.EnglishName != test.expected) {
			t.Errorf("ByName(%s) = %v, %v; want %v, %v", test.name, lang, found, test.expected, test.found)
		}
	}
}

func TestByType(t *testing.T) {
	tests := []struct {
		typ      Type
		expected int
	}{
		{Living, 87},     // Count will change based on language data.
		{Ancient, 5},     // Count will change based on language data.
		{Constructed, 0}, // Count will change based on language data.
	}

	for _, test := range tests {
		langs := ByType(test.typ)
		if len(langs) != test.expected {
			t.Errorf("ByType(%s) = %d; want %d", test.typ, len(langs), test.expected)
		}
	}
}

func TestByScope(t *testing.T) {
	tests := []struct {
		scope    Scope
		expected int
	}{
		{Individual, 308}, // Count will change based on language data.
		{Macro, 307},      // Count will change based on language data.
	}

	for _, test := range tests {
		langs := ByScope(test.scope)
		if len(langs) != test.expected {
			t.Errorf("ByScope(%s) = %d; want %d", test.scope, len(langs), test.expected)
		}
	}
}

func TestByFamily(t *testing.T) {
	tests := []struct {
		family   string
		expected int
	}{
		{"Indo-European", 29}, // Count will change based on language data.
		{"Sino-Tibetan", 0},   // Count will change based on language data.
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
	tests := []struct {
		script   string
		expected int
	}{
		{"Latin", 79},   // Count will change based on language data.
		{"Chinese", 22}, // Count will change based on language data.
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
	tests := []struct {
		region   string
		expected int
	}{
		{"Europe", 13}, // Count will change based on language data.
		{"Asia", 80},   // Count will change based on language data.
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
