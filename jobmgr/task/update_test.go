package task

import (
	"github.com/stretchr/testify/assert"
	pb_task "peloton/api/task"
	"testing"
)

func TestIsErrorState(t *testing.T) {
	assert.Equal(t, true, isUnexpected(pb_task.TaskState_FAILED))
	assert.Equal(t, true, isUnexpected(pb_task.TaskState_LOST))

	assert.Equal(t, false, isUnexpected(pb_task.TaskState_KILLED))
	assert.Equal(t, false, isUnexpected(pb_task.TaskState_LAUNCHING))
	assert.Equal(t, false, isUnexpected(pb_task.TaskState_RUNNING))
	assert.Equal(t, false, isUnexpected(pb_task.TaskState_SUCCEEDED))
	assert.Equal(t, false, isUnexpected(pb_task.TaskState_INITIALIZED))
}
