package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"google.golang.org/grpc/codes"
)

func Repository() (string, error) {
	file, err := os.ReadFile("go.mod")
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.Wrap(err, codes.NotFound)
		}
		return "", err
	}
	var repo string
	if _, err = fmt.Sscanf(string(file), "module %s\n", &repo); err != nil {
		return "", err
	}
	return strings.TrimSpace(repo), nil
}

func Registry(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.Wrap(err, codes.NotFound)
		}
		return "", err
	}
	var registry string
	for _, line := range strings.Split(string(file), "\n") {
		if _, err = fmt.Sscanf(line, "# registry %s\n", &registry); err == nil {
			break
		}
	}
	if registry == "" {
		return "", fmt.Errorf("registry does not set")
	}
	return strings.TrimSpace(registry), nil
}

func Services() ([]string, error) {
	dirs, err := os.ReadDir("deploy")
	if err != nil {
		return nil, err
	}
	var services []string
	for _, dir := range dirs {
		services = append(services, dir.Name())
	}
	return services, nil
}

func Commands() ([]string, error) {
	dirs, err := os.ReadDir("cmd")
	if err != nil {
		return nil, err
	}
	var commands []string
	for _, dir := range dirs {
		commands = append(commands, dir.Name())
	}
	return commands, nil
}

func Applications() ([]string, error) {
	dirs, err := os.ReadDir("internal/app")
	if err != nil {
		return nil, err
	}
	var applications []string
	for _, dir := range dirs {
		applications = append(applications, dir.Name())
	}
	return applications, nil
}
