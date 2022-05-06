package main

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	sStrList := strings.Split(s,"")
	tStrList := strings.Split(t,"")
	sort.Strings(sStrList)
	sort.Strings(tStrList)
	s = strings.Join(sStrList,"")
	t = strings.Join(tStrList,"")
	if s == t{
		return true
	}
	return  false
}

