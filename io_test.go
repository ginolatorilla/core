package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLinesFromFile(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		tempFile := filepath.Join(t.TempDir(), "text")
		assert.NoError(
			t,
			os.WriteFile(
				tempFile,
				[]byte(`Hello, World!
Good morning!`),
				os.ModePerm),
			"failed to write test file")

		expected := []string{"Hello, World!", "Good morning!"}
		lines, err := GetLinesFromFile(tempFile)
		require.NoError(t, err)
		assert.Len(t, lines, 2)
		for i := 0; i < len(lines); i++ {
			assert.Equal(t, expected[i], lines[i], "line %d should match", i)
		}
	})
	t.Run("file not found", func(t *testing.T) {
		t.Parallel()
		_, err := GetLinesFromFile("nonexistent.txt")
		assert.Error(t, err)
	})
}
