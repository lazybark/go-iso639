package iso639

import "slices"

type Code string
type Type string
type Scope string
type Script string
type Family string
type Region string

// Language represents an ISO 639 language.
type Language struct {
	Codes       []Code    // All ISO codes (ISO 639-1, ISO 639-2, ISO 639-3)
	EnglishName string    // English name of the primary language
	Type        Type      // Type of the language (e.g. Living, Ancient)
	Scope       Scope     // Scope of the language (e.g. Individual, Macrolanguage, Ancient)
	Scripts     []Script  // Writing script of the language (e.g. Latin, Cyrillic)
	Family      Family    // Language family (e.g. Indo-European, Sino-Tibetan)
	Regions     []Region  // Region where the language is spoken (e.g. Europe, Africa)
	NativeNames []string  // Native names of the language
	Variants    []Variant // Specific dialects or sub-languages
}

type Variant struct {
	Code        Code     // Unique ISO 639-3 code for the variant
	EnglishName string   // English name of the variant
	NativeNames []string // Native names of the variant
}

// ByCode fetches a language by its Alpha-1,2 or 3 code or by variant code.
func ByCode(code string) *Language {
	for _, lang := range LanguageMap {
		if slices.Contains(lang.Codes, Code(code)) {
			return &lang
		}

		for _, variant := range lang.Variants {
			if variant.Code == Code(code) {
				return &lang
			}
		}
	}

	return nil
}

// ByName fetches a language by its English or native name or by variant name.
func ByName(name string) *Language {
	for _, lang := range LanguageMap {
		if lang.EnglishName == name {
			return &lang
		}

		if slices.Contains(lang.NativeNames, name) {
			return &lang
		}

		for _, variant := range lang.Variants {
			if variant.EnglishName == name {
				return &lang
			}

			if slices.Contains(variant.NativeNames, name) {
				return &lang
			}
		}
	}

	return nil
}

// ByType fetches all languages that have a specific type.
func ByType(t Type) []Language {
	langs := make([]Language, 0)

	for _, lang := range LanguageMap {
		if lang.Type == t {
			langs = append(langs, lang)
		}
	}

	return langs
}

// ByScope fetches all languages that have a specific scope.
func ByScope(scope Scope) []Language {
	langs := make([]Language, 0)

	for _, lang := range LanguageMap {
		if lang.Scope == scope {
			langs = append(langs, lang)
		}
	}

	return langs
}

// ByFamily fetches all languages that belong to a specific family.
func ByFamily(family string) []Language {
	langs := make([]Language, 0)

	for _, lang := range LanguageMap {
		if lang.Family == Family(family) {
			langs = append(langs, lang)
		}
	}

	return langs
}

// ByScript fetches all languages that use a specific script.
func ByScript(script string) []Language {
	langs := make([]Language, 0)

	for _, lang := range LanguageMap {
		if slices.Contains(lang.Scripts, Script(script)) {
			langs = append(langs, lang)
		}
	}

	return langs
}

// ByRegion fetches all languages spoken in a specific region.
func ByRegion(region string) []Language {
	langs := make([]Language, 0)

	for _, lang := range LanguageMap {
		if slices.Contains(lang.Regions, Region(region)) {
			langs = append(langs, lang)
		}
	}

	return langs
}

// GetAllLanguages returns all defined languages as a slice.
func GetAllLanguages() []Language {
	langs := make([]Language, 0, len(LanguageMap))

	for _, lang := range LanguageMap {
		langs = append(langs, lang)
	}

	return langs
}
