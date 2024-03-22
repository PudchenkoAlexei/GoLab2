package Lab_2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")
	result, err := PrefixToPostfix(str)

	if err != nil {
		fmt.Fprintln(ch.Output, err)
		return err
	}
	_, err = ch.Output.Write([]byte(result))

	if err != nil {
		return fmt.Errorf("failed to write output: %v", err)
	}
	return nil
}
