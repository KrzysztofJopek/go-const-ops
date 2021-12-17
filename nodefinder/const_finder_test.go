package nodefinder

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConstBlock(t *testing.T) {

	tests := []struct {
		name        string
		inputFile   string
		inputPos    int
		expectedErr error
	}{
		{
			name:        "Success",
			inputFile:   "package test; const ( one int = 5; two int = 6 )",
			inputPos:    15,
			expectedErr: nil,
		},
		{
			name:        "Not found, too small pos",
			inputFile:   "package test; const ( one int = 5; two int = 6 )",
			inputPos:    0,
			expectedErr: ErrNodeNotFound,
		},
		{
			name:        "Not found, too big pos",
			inputFile:   "package test; const ( one int = 5; two int = 6 )",
			inputPos:    50,
			expectedErr: ErrNodeNotFound,
		},
		{
			name:        "Not found, no const block",
			inputFile:   "package test; ",
			inputPos:    6,
			expectedErr: ErrNodeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, fmt.Sprintf("%s-constBlock.go", tt.name), tt.inputFile, 0)
			fmt.Println(err)
			if !assert.Nil(err) {
				return
			}
			_, err = GetConstBlock(file, token.Pos(tt.inputPos))
			if tt.expectedErr != nil {
				assert.ErrorIs(err, tt.expectedErr)
			} else {
				assert.Nil(err)
			}

		})
	}
}
