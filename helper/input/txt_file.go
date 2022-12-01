package input

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func NewTXTFile(filename string) *TXTFile {
	return &TXTFile{filename: filename}
}

type TXTFile struct {
	filename string
}

func (tf *TXTFile) ReadByLine(_ context.Context, handler func(line string) error) error {
	f, err := os.Open(tf.filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := handler(scanner.Text()); err != nil {
			return fmt.Errorf("failed to handle line: %w", err)
		}
	}
	return nil
}

func (tf *TXTFile) ReadByBlock(ctx context.Context, separator string, handler func(block []string) error) error {
	fd, err := os.ReadFile(tf.filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	if err := handler(strings.Split(string(fd), separator)); err != nil {
		return fmt.Errorf("failed to handle block: %w", err)
	}
	return nil
}
