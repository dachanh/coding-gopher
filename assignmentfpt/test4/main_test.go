package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	if !reflect.DeepEqual(Solution([]string{"pim", "pom"}, []string{"999999999", "777888999"}, "88999"), "pom") {
		t.Fatalf("FAILED TEST 1, Expected %v but result %v", "pom", Solution([]string{"pim", "pom"}, []string{"999999999", "777888999"}, "88999"))
	}
	if !reflect.DeepEqual(Solution([]string{"sander", "army", "ann", "michael"}, []string{"123456789", "234567890", "789123456", "123123123"}, "1"), "ann") {
		t.Fatalf("FAILED TEST 2, Expected %v but result %v", "ann", Solution([]string{"sander", "army", "ann", "michael"}, []string{"123456789", "234567890", "789123456", "123123123"}, "1"))
	}
	if !reflect.DeepEqual(Solution([]string{"adam", "eva", "leo"}, []string{"121212121", "111111111", "444555666"}, "112"), "NO CONTACT") {
		t.Fatalf("FAILED TEST 3, Expected %v but result %v", "NO CONTACT", Solution([]string{"adam", "eva", "leo"}, []string{"121212121", "111111111", "444555666"}, "112"))
	}
}
