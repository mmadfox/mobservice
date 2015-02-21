package props

import (
	"regexp"
)

func FindOs(term string) []Type {
	return search(mapos, term)
}

func FindPhone(term string) []Type {
	return search(mapmob, term)
}

func FindBrowser(term string) []Type {
	return search(mapbrows, term)
}

func FindTablet(term string) []Type {
	return search(maptab, term)
}

func search(handbook map[int]Type, term string) []Type {
	founded := make([]Type, 0)
	for _, val := range handbook {
		if regexp.MustCompile("(?i)"+term).MatchString(val.Name) == true {
			founded = append(founded, val)
		}
	}
	return founded
}
