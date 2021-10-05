package main

import "testing"

func TestAddToOrderedSet(t *testing.T) {
	set := [6]uint8{0, 0, 0, 0, 0, 0}

	expected0 := [6]uint8{10, 0, 0, 0, 0, 0}
	addToOrderedSet(&set, 10)
	if set != expected0 {
		t.Fatalf("Should add to first item %v", set)
	}

	addToOrderedSet(&set, 10)
	if set != expected0 {
		t.Fatalf("Should not add the same value again %v", set)
	}

	expected1 := [6]uint8{10, 12, 0, 0, 0, 0}
	addToOrderedSet(&set, 12)
	if set != expected1 {
		t.Fatalf("Should add 12 %v", set)
	}

	expected2 := [6]uint8{10, 12, 15, 0, 0, 0}
	addToOrderedSet(&set, 15)
	if set != expected2 {
		t.Fatalf("Should add 15 %v", set)
	}

	expected3 := [6]uint8{10, 12, 15, 9, 0, 0}
	addToOrderedSet(&set, 9)
	if set != expected3 {
		t.Fatalf("Should add 9 %v", set)
	}

	expected4 := [6]uint8{10, 12, 15, 9, 8, 0}
	addToOrderedSet(&set, 8)
	if set != expected4 {
		t.Fatalf("Should add 8 %v", set)
	}

	addToOrderedSet(&set, 15)
	if set != expected4 {
		t.Fatalf("Should not add 15 again %v", set)
	}

	expected5 := [6]uint8{10, 12, 15, 9, 8, 100}
	addToOrderedSet(&set, 100)
	if set != expected5 {
		t.Fatalf("Should add 100 %v", set)
	}

	addToOrderedSet(&set, 50)
	if set != expected5 {
		t.Fatalf("Should not keep adding past the end %v", set)
	}
}

func TestRemoveFromOrderedSet(t *testing.T) {
	set := [6]uint8{43, 54, 23, 10, 50, 3}

	expected0 := [6]uint8{43, 54, 23, 10, 50, 3}
	removeFromOrderedSet(&set, 99)
	if set != expected0 {
		t.Fatalf("Should not remove nonexistent number %v", set)
	}

	expected1 := [6]uint8{43, 54, 23, 10, 50, 0}
	removeFromOrderedSet(&set, 3)
	if set != expected1 {
		t.Fatalf("Should be able to remove last number %v", set)
	}

	expected2 := [6]uint8{43, 54, 10, 50, 0, 0}
	removeFromOrderedSet(&set, 23)
	if set != expected2 {
		t.Fatalf("Should remove and shift all numbers %v", set)
	}

	removeFromOrderedSet(&set, 23)
	if set != expected2 {
		t.Fatalf("Should not remove nonexistent number %v", set)
	}

	expected3 := [6]uint8{54, 10, 50, 0, 0, 0}
	removeFromOrderedSet(&set, 43)
	if set != expected3 {
		t.Fatalf("Should be able to remove first number with blank after %v", set)
	}

	expected4 := [6]uint8{10, 50, 0, 0, 0, 0}
	removeFromOrderedSet(&set, 54)
	if set != expected4 {
		t.Fatalf("Should be able to remove first number with blank after %v", set)
	}

	expected5 := [6]uint8{50, 0, 0, 0, 0, 0}
	removeFromOrderedSet(&set, 10)
	if set != expected5 {
		t.Fatalf("Should be able to remove first number with blank after %v", set)
	}

	expected6 := [6]uint8{0, 0, 0, 0, 0, 0}
	removeFromOrderedSet(&set, 50)
	if set != expected6 {
		t.Fatalf("Should be able to remove first and also last number %v", set)
	}
}
