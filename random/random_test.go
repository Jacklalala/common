package random

import (
	"fmt"
	"testing"
)

func TestValue(t *testing.T) {
	var value1 = Value(8)
	var value2 = ValueHex(8)
	fmt.Println("value1 : ", value1)
	fmt.Println("value2 : ", value2)
}
