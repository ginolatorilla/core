// # Copyright © 2025 Gino Latorilla
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package core

import (
	"context"
	"os/exec"

	"github.com/stretchr/testify/mock"
)

// TODO: Add more functions of exec.Cmd to the Exec interface as needed.

// Exec is an interface for the exec.Cmd type from the os/exec package.
type Exec interface {
	Run() error              // Run starts the command. It should return an exec.ExitError if the underlying command fails.
	Output() ([]byte, error) // Output is similar to Run, but it returns the output of the command.
}

// Executor is the function signature of exec.Command.
type Executor func(ctx context.Context, cmd string, args ...string) Exec

// Command creates a new Exec instance using exec.Command.
func Command(cmd string, args ...string) Exec {
	return exec.Command(cmd, args...)
}

// CommandContext creates a new Exec instance using exec.CommandContext.
func CommandContext(ctx context.Context, cmd string, args ...string) Exec {
	return exec.CommandContext(ctx, cmd, args...)
}

// MockExec is a mock implementation of the Exec interface.
//
// Use this mock to setup the "backend" behaviour of the command.
type MockExec struct {
	mock.Mock
}

func (m *MockExec) Run() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockExec) Output() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// MockExecutor mocks the Executor function signature.
type MockExecutor struct {
	mock.Mock
}

// Do forwards the call to the MockExecutor and returns the MockExec.
//
// Because the mocked target is not an object, we need to use this function to handle the call.
func (m *MockExecutor) Do(ctx context.Context, cmd string, args ...string) Exec {
	var mockArgs []any

	mockArgs = append(mockArgs, ctx, cmd)
	for _, arg := range args {
		mockArgs = append(mockArgs, arg)
	}

	callArgs := m.Called(mockArgs...)
	return callArgs.Get(0).(Exec)
}

// Executor returns the Do function as an Executor.
func (m *MockExecutor) Executor() Executor {
	return m.Do
}
