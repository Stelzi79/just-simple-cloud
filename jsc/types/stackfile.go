package types

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/k0kubun/pp/v3"

	"github.com/Stelzi79/just-simple-cloud/enums"
)

type Stack struct {
	DependsOn []string        `json:"dependsOn,omitempty"`
	Type      enums.StackType `json:"type"`
	Source    string          `json:"src"`
}

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
func (stackFile *StackFile) Validate(filePath string) (*StackFile, error) {
	if stackFile.Name == "" {
		stackFile.Name = filepath.Base(filepath.Dir(filePath))
	}
	// Validate stacks
	err, shouldReturn := ValidateStacks(&stackFile.Stacks)
	if shouldReturn {
		return nil, err
	}
	// Validate infra stacks
	err, shouldReturn = ValidateStacks(&stackFile.Infra)
	if shouldReturn {
		return nil, err
	}

	return stackFile, nil
}

func ValidateStacks(stacks *map[string]Stack) (error, bool) {
	for stackName, stack := range *stacks {

		// set default type if not specified
		if stack.Type.Value == "" {
			stack.Type = enums.Default
			(*stacks)[stackName] = stack
		}

		// Validate source if not specified
		if stack.Source == "" {
			switch stack.Type {
			case enums.Folder:
				stack.Source = filepath.Join(BASE_PATH, stackName, STACK_FILE_NAME)
				(*stacks)[stackName] = stack
			case enums.GitRepo:
				fallthrough
			default:
				return errors.New("Source is required for stack of type " + stack.Type.Value + " (" + stackName + ")"), false
			}
		} else {
			// Clean source path
			stack.Source = strings.TrimSpace(stack.Source)
			if !strings.Contains(stack.Source, BASE_PATH) {
				stack.Source = filepath.Join(BASE_PATH, stack.Source)
			}
			(*stacks)[stackName] = stack
		}

	}
	return nil, false
}

// IsRootStack checks if the stack file is of type .rootstack
func (stackFile *StackFile) IsRootStack() bool {
	return stackFile.Type == `.rootstack`
}

// PrettyPrint returns a pretty-printed string representation of the StackFile
func (stackFile *StackFile) PrettyPrint() string {
	return "\n" + strings.Trim(pp.Sprint(stackFile), "&types.")
}
