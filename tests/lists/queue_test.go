package lists

import (
	"datastructures/src/lists"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	// Arrange
	sut := lists.NewQueue()
	expectedResult := 1
	// Act
	sut.Enqueue(expectedResult)
	sut.Enqueue('c')
	actualResult := sut.Front()
	// Assert
	if sut.Size() != 2 {
		t.Errorf("Expected size is %v but got %v.", 2, sut.Size())
	}
	if expectedResult != actualResult {
		t.Errorf("Expected value is %v but got %v.", expectedResult, actualResult)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	// Arrange
	sut := lists.NewQueue()
	expectedResult := 1
	// Act
	sut.Enqueue(expectedResult)
	sut.Enqueue('c')
	sut.Enqueue('a')
	sut.Enqueue("Hello")
	actualResult := sut.Dequeue()
	// Assert
	if sut.Size() != 3 {
		t.Errorf("Expected size is %v but got %v.", 2, sut.Size())
	}
	if expectedResult != actualResult {
		t.Errorf("Expected value is %v but got %v.", expectedResult, actualResult)
	}
}

func TestQueue_Front_ShouldReturnFirstEnqueuedValue(t *testing.T) {
	// Arrange
	sut := lists.NewQueue()
	expectedResult := "I am the first value"
	sut.Enqueue(expectedResult)
	sut.Enqueue(1)
	sut.Enqueue(2)
	// Act
	actualResult := sut.Front()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected value is %v but got %v.", expectedResult, actualResult)
	}
}

func TestQueue_Front_ShouldNotPopValue(t *testing.T) {
	// Arrange
	sut := lists.NewQueue()
	sut.Enqueue(1)
	sut.Enqueue(2)
	sut.Enqueue(3)
	expectedResult := 3
	// Act
	sut.Front()
	actualResult := sut.Size()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected size is %v but got %v.", expectedResult, actualResult)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	// Arrange
	sut := lists.NewQueue()
	expectedResult := true
	// Act
	actualResult := sut.IsEmpty()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected list to be empty but it is not.")
	}
}

func TestQueue_IsEmpty_WhenStackIsNotEmpty_ShouldReturnFalse(t *testing.T) {
	// Arrange
	sut := lists.NewQueue()
	sut.Enqueue(1)
	expectedResult := false
	// Act
	actualResult := sut.IsEmpty()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected list to be not empty but it is not.")
	}
}