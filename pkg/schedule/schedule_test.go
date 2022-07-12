package schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	sch := Schedule()
	assert.Equal(t, sch, "2", "Test failed nasty like")
}
