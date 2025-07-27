package core

import (
	"context"
	"testing"

	"os/exec"

	"github.com/stretchr/testify/assert"
)

func TestMockExecutor_SingleCommand(t *testing.T) {
	t.Parallel()

	var (
		mockExecutor MockExecutor
		mockExec     MockExec
	)
	mockExecutor.On("Do", context.TODO(), "echo", "Hello, World").
		Return(&mockExec)
	mockExec.On("Output").
		Return([]byte("hi from test"), nil)

	executor := mockExecutor.Executor()
	out, err := executor(context.TODO(), "echo", "Hello, World").Output()

	assert.NoError(t, err)
	assert.Equal(t, "hi from test", string(out))
	mockExecutor.AssertExpectations(t)
	mockExec.AssertExpectations(t)
}

func TestMockExecutor_MultipleCommands(t *testing.T) {
	t.Parallel()

	var (
		mockExecutor MockExecutor
		mockExec1    MockExec
		mockExec2    MockExec
	)
	mockExecutor.On("Do", context.TODO(), "echo", "Hello, World").
		Return(&mockExec1)
	mockExec1.On("Output").
		Return([]byte("hi from test"), nil)

	mockExecutor.On("Do", context.TODO(), "echo", "Good morning").
		Return(&mockExec2)
	mockExec2.On("Output").
		Return([]byte("good morning from test"), nil)

	executor := mockExecutor.Executor()
	out, err := executor(context.TODO(), "echo", "Hello, World").Output()

	assert.NoError(t, err)
	assert.Equal(t, "hi from test", string(out))
	mockExecutor.AssertCalled(t, "Do", context.TODO(), "echo", "Hello, World")
	mockExec1.AssertCalled(t, "Output")

	out, err = executor(context.TODO(), "echo", "Good morning").Output()

	assert.NoError(t, err)
	assert.Equal(t, "good morning from test", string(out))
	mockExecutor.AssertExpectations(t)
	mockExec2.AssertExpectations(t)
}

func TestMockExecutor_NonZeroExit(t *testing.T) {
	t.Parallel()

	var (
		mockExecutor MockExecutor
		mockExec     MockExec
	)
	mockExecutor.On("Do", context.TODO(), "something").
		Return(&mockExec)
	mockExec.On("Run").
		Return(&exec.ExitError{})

	executor := mockExecutor.Executor()
	err := executor(context.TODO(), "something").Run()

	assert.Error(t, err)
	assert.Equal(t, &exec.ExitError{}, err)
	mockExecutor.AssertExpectations(t)
	mockExec.AssertExpectations(t)
}

func TestCommand(t *testing.T) {
	t.Parallel()
	c := Command("go", "version")
	out, err := c.Output()
	assert.NoError(t, err)
	assert.NotEmpty(t, out)
}

func TestCommandContext(t *testing.T) {
	t.Parallel()
	c := CommandContext(context.TODO(), "go", "version")
	out, err := c.Output()
	assert.NoError(t, err)
	assert.NotEmpty(t, out)
}
