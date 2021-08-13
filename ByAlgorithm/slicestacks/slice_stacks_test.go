package slicestacks
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Clear(t *testing.T) {
	stack := New()
	assert.Equal(t, stack.Size(), 0)
	assert.Equal(t, stack.IsEmpty(), true)

	stack.Push("hello")
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)

	stack.Clear()
	assert.Equal(t, stack.IsEmpty(), true)
	assert.Equal(t, stack.Size(), 0)
}
