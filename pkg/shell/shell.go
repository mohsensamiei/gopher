package shell

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type Shell string

const (
	SH   Shell = "sh"
	BASH Shell = "bash"
	ZSH  Shell = "zsh"
)

func (shell Shell) CommandContextStream(ctx context.Context, cmd string, filters ...string) error {
	filtered := cmd
	sort.Slice(filters, func(i, j int) bool {
		return len(filters[i]) < len(filters[j])
	})
	for _, filter := range filters {
		filtered = strings.ReplaceAll(filtered, filter, "*****")
	}
	fmt.Printf("$ %v\n", filtered)

	command := exec.CommandContext(ctx, string(shell), "-c", cmd)
	{
		stdout, err := command.StdoutPipe()
		if err != nil {
			return err
		}
		go write(os.Stdout, bufio.NewReader(stdout))
	}
	{
		stderr, err := command.StderrPipe()
		if err != nil {
			return err
		}
		go write(os.Stderr, bufio.NewReader(stderr))
	}
	if err := command.Start(); err != nil {
		return err
	}
	if err := command.Wait(); err != nil {
		return err
	}
	return nil
}

type readLine interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

func write(output *os.File, reader readLine) {
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				_, _ = fmt.Fprintf(output, "  %v\n", err)
			}
			break
		}
		_, _ = fmt.Fprintf(output, "  %s\n", line)
	}
}
