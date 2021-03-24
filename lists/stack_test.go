package lists

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	// Arrange
	sut := NewStack()
	expectedResult := 1
	// Act
	sut.Push('c')
	sut.Push(expectedResult)
	actualResult := sut.Peek()
	// Assert
	if sut.Size() != 2 {
		t.Errorf("Expected size is %v but got %v.", 2, sut.Size())
	}
	if expectedResult != actualResult {
		t.Errorf("Expected value is %v but got %v.", expectedResult, actualResult)
	}
}

func TestStack_Pop_ShouldReturnValue(t *testing.T) {
	// Arrange
	sut := NewStack()
	expectedResult := 1
	sut.Push(expectedResult)
	// Act
	actualResult := sut.Pop()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected value is %v but got %v.", expectedResult, actualResult)
	}
}

func TestStack_Pop_ShouldRemoveValue(t *testing.T) {
	// Arrange
	sut := NewStack()
	sut.Push(1)
	expectedResult := 0
	// Act
	sut.Pop()
	actualResult := sut.Size()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected size is %v but got %v.", expectedResult, actualResult)
	}
}

func TestStack_Peek_ShouldReturnLastPushedValue(t *testing.T) {
	// Arrange
	sut := NewStack()
	expectedResult := "I am the last value"
	sut.Push(1)
	sut.Push(2)
	sut.Push(expectedResult)
	// Act
	actualResult := sut.Peek()
	// Assert
	if expectedResult != actualResult {
		t.Errorf("Expected value is %v but got %v.", expectedResult, actualResult)
	}
}

func TestStack_Peek_ShouldNotPopValue(t *testing.T) {
	// Arrange
	sut := NewStack()
	sut.Push(1)
	sut.Push(2)
	sut.Push(3)
	expectedResult := 3
	// Act
	sut.Peek()
	actualResult := sut.Size()
	// Assert
	if actualResult != expectedResult {
		t.Errorf("Expected size is %v but got %v.", expectedResult, actualResult)
	}
}
