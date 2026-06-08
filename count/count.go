package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	input  io.Reader
	output io.Writer
}
type option func(*Counter) error

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}

}
func WithOutput(output io.Writer) option {
	return func(c *Counter) error {
		if output == nil {
			return errors.New("nil output reader")
		}
		c.output = output
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *Counter) error {
		f, err := os.Open(args[0])
		if len(args) < 1 {
			return nil
		}
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func NewCounter(opts ...option) (*Counter, error) {
	c := &Counter{
		output: os.Stdout,
		input:  os.Stdin,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scan error:", err)
	}
	return lines
}
func Main() {
	c, err := NewCounter()
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Lines())
}

// func (C *Counter) Count() {
//     lines := 0
//     input := bufio.NewScanner(os.Stdin)
//     for input.Scan() {
//         lines++
//     }
//     if err := input.Err(); err != nil {
//         fmt.Fprintln(os.Stderr, "scan error:", err)
//         return
//     }
//     fmt.Println(lines)
// }
