package types

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/k0kubun/pp/v3"
)

type Stack struct {
	DependsOn []string `json:"dependsOn,omitempty"`
	Type      string   `json:"type,omitempty"`
	Source    string   `json:"src,omitempty"`
}

type StackType int

const (
	folder StackType = iota
	gitRepo
)

type StackFile struct {
	Name   string           `json:"name"`
	Type   string           `json:"type"`
	Stacks map[string]Stack `json:"stacks"`
	Infra  map[string]Stack `json:"infra,omitempty"`
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
	return s.Type == `.rootstack`
}

func (s *StackFile) PrettyPrint() string {
	return "\n" + strings.Trim(pp.Sprint(s), "&types.")
}
