package practice

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var errDivisorZero = errors.New("0으로 나눌 수 없습니다")

func sum(num1, num2 int) int {
	return num1 + num2
}

func divide(dividend, divisor int) (float32, error) {
	if divisor == 0 {
		return 0, errDivisorZero
	}
	return float32(dividend / divisor), nil
}

func generateRandomID() string {
	return uuid.New().String()
}

func TestPractice(t *testing.T) {
	t.Run("두 숫자를 더하면 합이 나온다", func(t *testing.T) {
		actual := sum(1, 2)
		expected := 3
		assert.Equal(t, expected, actual)
		//TODO actual과 expected가 같은지 검증해주세요.
	})

	t.Run("두 숫자를 더하면 합이 나온다", func(t *testing.T) {
		actual := sum(1, 2)
		expected := float32(3)
		assert.EqualValues(t, expected, actual)
		//TODO actual과 expected가 같은지 검증해주세요.
	})

	t.Run("두 숫자를 나눗셈 할 수 있다.", func(t *testing.T) {
		actual, err := divide(10, 2)
		assert.EqualValues(t, 5, actual)
		if err != nil {
			t.Fail()
		}
		//TODO err가 발생하지 않았음을 검증해주세요.
		//TODO 결과로 나온 숫자가 5가 맞는지 검증해주세요.
	})

	t.Run("0으로 나누기를 할 수 없다.", func(t *testing.T) {
		actual, err := divide(10, 0)
		assert.EqualValues(t, 0, actual)
		assert.Equal(t, err, errDivisorZero)
		//TODO errDivisorZero가 맞는지 검증해주세요.
		//TODO actual의 값을 검증해주세요.
	})

	t.Run("uuid가 생성된다.", func(t *testing.T) {
		uuid := generateRandomID()
		assert.NotEmpty(t, uuid)
		//TODO uuid가 생성되었는지 검증해주세요.
	})
}
