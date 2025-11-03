package types

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/k0kubun/pp/v3"
)

type Stack struct {
	DependsOn []string `yaml:"dependsOn,omitempty"`
	Source    string   `yaml:"src,omitempty"`
}

type StackFile struct {
	Name   string           `yaml:"name"`
	Type   string           `yaml:"type"`
	Stacks map[string]Stack `yaml:"stacks"`
}

// NewStackFile reads and parses the stack file from the given path
func NewStackFile(stackFilePath string) (*StackFile, error) {
	var stackData StackFile

	data, err := os.ReadFile(stackFilePath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &stackData); err != nil {
		return nil, err
	}

	return stackData.Validate(stackFilePath)
}

// Validate checks if the stack file has valid content
func (s *StackFile) Validate(filePath string) (*StackFile, error) {
	if s.Name == "" {
		s.Name = filepath.Base(filepath.Dir(filePath))
	}

	return s, nil
}

// IsRootStack checks if the stack file is of type .rootstack
func (s *StackFile) IsRootStack() bool {
	return s.Type == ".rootstack"
}

func (s *StackFile) PrettyPrint() string {
	return "\n" + strings.Trim(pp.Sprint(s), "&types.")
}
