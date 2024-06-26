package execext

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func CommandContextStream(ctx context.Context, name string, args ...string) error {
	color.Cyan("$ %v %v\n", name, strings.Join(args, " "))

	var stopChan = make(chan os.Signal, 2)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	cmd := exec.CommandContext(ctx, name, args...)
	{
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		go func(scanner *bufio.Scanner) {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}(bufio.NewScanner(stdout))
	}
	{
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return err
		}
		go func(scanner *bufio.Scanner) {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}(bufio.NewScanner(stderr))
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
