package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	data1 := 1
	data2 := 2
	data3 := 3

	testList := New(&data1)
	testList.Append(New(&data2))
	testList.Append(New(&data3))

	assert.Equal(t, testList.Length(), 3)
	assert.Equal(t, testList.Last().Data(), &data3)

	testList0 := New[int](nil)
	testList0.Append(New(&data3))
	assert.Equal(t, *testList0.Data(), 3)
}

func TestAppend(t *testing.T) {
	data1 := 1
	data2 := 2
	data3 := 3
	data4 := 4
	data5 := 5
	data6 := 6

	testList0 := New(&data1)
	testList0.Append(New(&data2))
	testList0.Append(New(&data3))

	testList1 := New(&data4)
	testList1.Append(New(&data5))
	testList1.Append(New(&data6))

	testList0.Append(testList1)
	assert.Equal(t, testList0.Length(), 6)

	probeList := testList0
	assert.Equal(t, *(probeList.Data()), 1)
	for v := 2; v <= 6; v++ {
		probeList = probeList.Next()
		assert.Equal(t, *(probeList.Data()), v)
	}

}
