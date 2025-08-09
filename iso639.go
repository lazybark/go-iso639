package iso639

import (
	"slices"
	"strings"
	"sync"
)

type Code string
type Type string
type Scope string
type Script string
type Family string
type Region string

// Pre-computed indexes for O(1) lookups.
var (
	codeIndex     map[string]*Language
	nameIndex     map[string]*Language
	scriptIndex   map[Script][]*Language
	familyIndex   map[Family][]*Language
	regionIndex   map[Region][]*Language
	typeIndex     map[Type][]*Language
	scopeIndex    map[Scope][]*Language
	indexInitOnce sync.Once
)

// initIndexes builds lookup indexes for O(1) performance.
func initIndexes() {
	codeIndex = make(map[string]*Language)
	nameIndex = make(map[string]*Language)
	scriptIndex = make(map[Script][]*Language)
	familyIndex = make(map[Family][]*Language)
	regionIndex = make(map[Region][]*Language)
	typeIndex = make(map[Type][]*Language)
	scopeIndex = make(map[Scope][]*Language)

	for key := range LanguageMap {
		lang := LanguageMap[key]
		langPtr := &lang

		// Index all codes (case-insensitive).
		for _, code := range lang.Codes {
			codeIndex[strings.ToLower(string(code))] = langPtr
		}

		// Index English and native names (case-insensitive).
		nameIndex[strings.ToLower(lang.EnglishName)] = langPtr
		for _, nativeName := range lang.NativeNames {
			nameIndex[strings.ToLower(nativeName)] = langPtr
		}

		// Index variant codes and names.
		for _, variant := range lang.Variants {
			codeIndex[strings.ToLower(string(variant.Code))] = langPtr
			nameIndex[strings.ToLower(variant.EnglishName)] = langPtr

			for _, variantNativeName := range variant.NativeNames {
				nameIndex[strings.ToLower(variantNativeName)] = langPtr
			}
		}

		// Index by type, scope, family.
		typeIndex[lang.Type] = append(typeIndex[lang.Type], langPtr)
		scopeIndex[lang.Scope] = append(scopeIndex[lang.Scope], langPtr)
		familyIndex[lang.Family] = append(familyIndex[lang.Family], langPtr)

		// Index by scripts and regions.
		for _, script := range lang.Scripts {
			scriptIndex[script] = append(scriptIndex[script], langPtr)
		}

		for _, region := range lang.Regions {
			regionIndex[region] = append(regionIndex[region], langPtr)
		}
	}
}

// ensureInitialized makes sure indexes are built (thread-safe).
func ensureInitialized() {
	indexInitOnce.Do(initIndexes)
}

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
// Case-insensitive lookup with O(1) performance.
func ByCode(code string) *Language {
	ensureInitialized()
	return codeIndex[strings.ToLower(code)]
}

// ByCodeIgnoreCase is an alias for ByCode (which is already case-insensitive).
// Kept for backward compatibility.
func ByCodeIgnoreCase(code string) *Language {
	return ByCode(code)
}

// ByName fetches a language by its English or native name or by variant name.
// Case-insensitive lookup with O(1) performance.
func ByName(name string) *Language {
	ensureInitialized()
	return nameIndex[strings.ToLower(name)]
}

// ByNameIgnoreCase is an alias for ByName (which is already case-insensitive).
// Kept for backward compatibility.
func ByNameIgnoreCase(name string) *Language {
	return ByName(name)
}

// ByType fetches all languages that have a specific type.
func ByType(t Type) []*Language {
	ensureInitialized()
	return typeIndex[t]
}

// ByScope fetches all languages that have a specific scope.
func ByScope(scope Scope) []*Language {
	ensureInitialized()
	return scopeIndex[scope]
}

// ByFamily fetches all languages that belong to a specific family.
func ByFamily(family string) []*Language {
	ensureInitialized()
	return familyIndex[Family(family)]
}

// ByScript fetches all languages that use a specific script.
func ByScript(script string) []*Language {
	ensureInitialized()
	return scriptIndex[Script(script)]
}

// ByRegion fetches all languages spoken in a specific region.
func ByRegion(region string) []*Language {
	ensureInitialized()
	return regionIndex[Region(region)]
}

// GetAllLanguages returns all defined languages as a slice of pointers.
func GetAllLanguages() []*Language {
	ensureInitialized()
	langs := make([]*Language, 0, len(LanguageMap))

	for key := range LanguageMap {
		lang := LanguageMap[key]
		langs = append(langs, &lang)
	}

	return langs
}

// IsValidCode checks if a code exists without returning the language.
func IsValidCode(code string) bool {
	return ByCode(code) != nil
}

// HasScript checks if a language supports a specific script.
func HasScript(lang *Language, script Script) bool {
	return slices.Contains(lang.Scripts, script)
}
