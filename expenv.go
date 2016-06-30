package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	argFile := flag.String("f", "", "Input file, otherwise stdin")
	argInPlace := flag.Bool("i", false, "Replace file inplace")
	flag.Parse()

	var input io.ReadCloser
	var output io.WriteCloser
	var err error

	defer func() {
		if output != nil {
			if err := output.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			}
		}
	}()

	defer func() {
		if input != nil {
			if err := input.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			}
		}
	}()

	if *argFile != "" {
		input, err = os.Open(*argFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error input file: %s\n", err)
			os.Exit(1)
		}
	} else {
		input = ReadNopCloser(os.Stdin)
	}

	if *argInPlace && *argFile != "" {
		file, err := ioutil.TempFile("", "expenv")
		output = FileReplaceWriteCloser(file, *argFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error output file: %s\n", err)
			os.Exit(1)
		}
	} else {
		output = WriteNopCloser(os.Stdout)
	}

	err = process(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

}

// process reads line by line from reader, expands environment variables and writes lines to writer
func process(r io.Reader, w io.Writer) error {
	bufr := bufio.NewReader(r)
	for {
		line, err := bufr.ReadString('\n')
		line = expandEnv(line)
		_, writererr := io.WriteString(w, line)
		if writererr != nil {
			return err
		}
		if err != nil {
			break
		}
	}
	return nil
}

// Nop Closer for io.Reader
type readNopCloser struct {
	io.Reader
}

func (r readNopCloser) Close() error {
	return nil
}

func ReadNopCloser(r io.Reader) io.ReadCloser {
	return readNopCloser{r}
}

// Nop Closer for io.Writer
type writeNopCloser struct {
	io.Writer
}

func (r writeNopCloser) Close() error {
	return nil
}

func WriteNopCloser(w io.Writer) io.WriteCloser {
	return writeNopCloser{w}
}

// fileReplaceWriteCloser replaces the replacePath with given file on Close()
type fileReplaceWriteCloser struct {
	*os.File
	replacePath string
}

// Close closes the file and renames it to replace path. If an error occurs the file is removed.
func (f fileReplaceWriteCloser) Close() error {
	err := f.File.Close()
	if err != nil {
		os.Remove(f.File.Name())
		return err
	}
	err = os.Rename(f.File.Name(), f.replacePath)
	if err != nil {
		os.Remove(f.File.Name())
		return err
	}
	return nil
}

func FileReplaceWriteCloser(f *os.File, replacePath string) io.WriteCloser {
	return &fileReplaceWriteCloser{f, replacePath}
}

// Follow os.ExpandEnv's contract except for `$$` which is transformed to `$`
func expandEnv(s string) string {
	os.Setenv("EXPENV_DOLLAR", "$")
	return os.ExpandEnv(strings.Replace(s, "$$", "${EXPENV_DOLLAR}", -1))
}
