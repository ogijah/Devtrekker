package main

import (
	"reflect"
	"testing"
)

func TestExtractArray(t *testing.T) {
	input := "[3, 10, 5, 7]"
	expected := []int{3, 10, 5, 7}

	result := ExtractArray(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Extract Array FAILED. Expected %v, got %v", expected, result)
	}
}

func TestDeduplicateArray(t *testing.T) {
	array := []int{5, 3, 5, 2, 3, 7, 7, 3}
	expected := []int{5, 3, 2, 7}

	result := DeduplicateArray(array)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Deduplicate Array FAILED. Expected %v, got %v", expected, result)
	}
}

func TestSortArray(t *testing.T) {
	array := []int{5, 3, 2, 7}
	expected := []int{2, 3, 5, 7}

	result := SortArray(array)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sort Array FAILED. Expected %v, got %v", expected, result)
	}
}
