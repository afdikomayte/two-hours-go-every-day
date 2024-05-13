package method_test

import (
	"testing"

	"github.com/afdikomayte/two-hours-code-every-day/method"
	"github.com/stretchr/testify/assert"
)


func TestMethod(t *testing.T){

  //make obj student
  student1 := method.Student{}
  student1.SetName("afdikomayte", 0, 12)

  assert.Equal(t, "afdikomayte", student1.GetName())

}


