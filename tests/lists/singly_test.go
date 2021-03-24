package lists

import (
	"datastructures/src/lists"
	"reflect"
	"testing"
)

func TestSinglyLinkedList_Size(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
	sut.Prepend(1)
	sut.Append(2)
	sut.Append(3)
	sut.Prepend(1)
	sut.Append("hello")
	sut.RemoveTail()
	sut.RemoveHead()
	expectedResult := 3
	// Act
	actualResult := sut.Size()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected size is %v but got %v.", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_IsEmpty(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
	expectedResult := true
	// Act
	actualResult := sut.IsEmpty()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected list to be empty but it is not.")
	}
}

func TestSinglyLinkedList_IsEmpty_WhenListIsNotEmpty_ShouldReturnFalse(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
	sut.Append(1)
	expectedResult := false
	// Act
	actualResult := sut.IsEmpty()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected list to be not empty but it is not.")
	}
}

func TestSinglyLinkedList_Size_WhenListIsEmpty_ShouldReturnZero(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
	expectedResult := 0
	// Act
	actualResult := sut.Size()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected size is %v but got %v.", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_Append(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
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
	sut := lists.NewSinglyLinkedList()
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
	sut := lists.NewSinglyLinkedList()
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
	sut := lists.NewSinglyLinkedList()
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
	sut := lists.NewSinglyLinkedList()
	expectedResult := "Hello"
	sut.Append(expectedResult)
	// Act
	actualResult := sut.RemoveHead()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_Head(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
	expectedResult := "Hi"
	sut.Append(expectedResult)
	// Act
	actualResult := sut.Head()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_Tail(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
	expectedResult := "Hi"
	sut.Append("Hello")
	sut.Append(expectedResult)
	// Act
	actualResult := sut.Tail()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_Tail_WhenValueIsTheOnlyElement_ShouldReturnSaidValue(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
	expectedResult := "Hello"
	sut.Append(expectedResult)
	// Act
	actualResult := sut.Tail()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestSinglyLinkedList_RemoveTail(t *testing.T) {
	// Arrange
	sut := lists.NewSinglyLinkedList()
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
	sut := lists.NewSinglyLinkedList()
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
	sut := lists.NewSinglyLinkedList()
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
