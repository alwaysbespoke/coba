package v1

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	files, err := os.ReadDir("./test-messages")
	if err != nil {
		t.Fatal(fmt.Errorf("failed to load test messages: %w", err))
	}

	for _, file := range files {
		filePath := fmt.Sprintf("./test-messages/%s", file.Name())
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatal(fmt.Errorf("failed to read test message: %w", err))
		}

		message := &SIPMessage{}
		err = message.Unmarshal(fileBytes)

		assert.Nil(t, err)
	}
}
