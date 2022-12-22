package input

import (
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
	f, err := os.ReadFile(tf.filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	lines := strings.Split(string(f), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		if err := handler(line); err != nil {
			return fmt.Errorf("failed to handle line: %w", err)
		}
	}
	return nil
}

func (tf *TXTFile) ReadByLineEx(_ context.Context, handler func(i int, line string) error) error {
	f, err := os.ReadFile(tf.filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	lines := strings.Split(string(f), "\n")

	for index, line := range lines {
		if err := handler(index, line); err != nil {
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

func (tf *TXTFile) ReadByBlockEx(ctx context.Context, sep func(i int, line string) bool, handler func(lines []string) error) error {
	fd, err := os.ReadFile(tf.filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	lines := strings.Split(string(fd), "\n")
	var offset int
	for i := 0; i < len(lines); i++ {
		if sep(i, lines[i]) {
			if err := handler(lines[offset:i]); err != nil {
				return fmt.Errorf("failed to handle block: %w", err)
			}
			offset = i
		}
	}
	if err := handler(lines[offset:]); err != nil {
		return fmt.Errorf("failed to handle block: %w", err)
	}
	return nil
}
