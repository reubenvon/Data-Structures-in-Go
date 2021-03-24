package lists

import (
	"datastructures/src/lists"
	"reflect"
	"testing"
)

func TestSinglyLinkedList_Append(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	expectedResult := []interface{}{1, 2, "Hello", "World"}
	// Act
	for _, value := range expectedResult {
		sut.Append(value)
	}
	// Assert
	actualResult := sut.ToSlice()
	checkResults(t, expectedResult, actualResult)
}

func TestSinglyLinkedList_Prepend(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	testData := []interface{}{1, 2, "Hello", "World"}
	expectedResult := []interface{}{"World", "Hello", 2, 1}
	// Act
	for _, value := range testData {
		sut.Prepend(value)
	}
	// Assert
	actualResult := sut.ToSlice()
	checkResults(t, expectedResult, actualResult)
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	testData := []interface{}{1, 2, "Hello", "World"}
	expectedResult := []interface{}{"World", "Hello", 2, 1}
	for _, value := range testData {
		sut.Append(value)
	}
	// Act
	actualResult := sut.Reverse().ToSlice()
	// Assert
	checkResults(t, expectedResult, actualResult)
}

func TestSinglyLinkedList_RemoveHead(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	testData := []interface{}{1, 2, "Hello", "World"}
	expectedResult := []interface{}{2, "Hello", "World"}
	for _, value := range testData {
		sut.Append(value)
	}
	// Act
	sut.RemoveHead()
	actualResult := sut.ToSlice()
	// Assert
	checkResults(t, expectedResult, actualResult)
}

func TestSinglyLinkedList_RemoveHead_ShouldReturnData(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	expectedResult := "Hello"
	sut.Append(expectedResult)
	// Act
	actualResult := sut.RemoveHead()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_RemoveTail(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	testData := []interface{}{1, 2, "Hello", "World"}
	expectedResult := []interface{}{1, 2, "Hello"}
	for _, value := range testData {
		sut.Append(value)
	}
	// Act
	sut.RemoveTail()
	actualResult := sut.ToSlice()
	// Assert
	checkResults(t, expectedResult, actualResult)
}

func TestSinglyLinkedList_RemoveTail_ShouldReturnData(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	expectedResult := "Hi there, I'm tail"
	sut.Append(expectedResult)
	// Act
	actualResult := sut.RemoveTail()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_ToSlice(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList(nil)
	expectedResult := []interface{}{1}
	// Act
	sut.Append(1)
	// Assert
	actualResult := sut.ToSlice()
	if reflect.ValueOf(actualResult).Kind() != reflect.Slice {
		t.Errorf("Kind of actual result is not slice")
	}
	checkResults(t, expectedResult, actualResult)
}

func checkResults(t *testing.T, expected []interface{}, actual []interface{})  {
	expectedLength := len(expected)
	actualLength := len(actual)
	if expectedLength != actualLength {
		t.Errorf("Expected size is %d but got %d.", expectedLength, actualLength)
	}
	for i, got := range actual {
		if got != expected[i] {
			t.Errorf("Expected value is %v but got %v at index %d.", expected[i], got, i)
		}
	}
}
