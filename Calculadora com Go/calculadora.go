package calculator

import "testing"
import "errors"

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero is not allowed")
    }
    return float64(a) / float64(b), nil
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add test failed, expected %d but got %d", expected, result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(10, 5)
    expected := 5
    if result != expected {
        t.Errorf("Subtract test failed, expected %d but got %d", expected, result)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(2, 3)
    expected := 6
    if result != expected {
        t.Errorf("Multiply test failed, expected %d but got %d", expected, result)
    }
}

func TestDivide(t *testing.T) {
    t.Run("Valid Division", func(t *testing.T) {
        result, err := Divide(10, 5)
        expected := 2.0
        if err != nil {
            t.Errorf("Divide test failed, expected no error but got %v", err)
        }
        if result != expected {
            t.Errorf("Divide test failed, expected %f but got %f", expected, result)
        }
    })
    
    t.Run("Invalid Division", func(t *testing.T) {
        _, err := Divide(10, 0)
        if err == nil {
            t.Errorf("Divide test failed, expected error but got nil")
        }
    })
}
