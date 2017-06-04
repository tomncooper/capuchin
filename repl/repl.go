package repl

import (
	"bufio"
	"capuchin/lexer"
	"capuchin/token"
	"fmt"
	"io"
)

// PROMPT is the value that is shown at the beginning of each REPL line
const PROMPT = ">>"

// Start will read in the input source until it encounters a new line character
// and then pass the that string to the Lexer. It will then print out each
// token produced by the Lexer.
func Start(input io.Reader, ouput io.Writer) {

	scanner := bufio.NewScanner(input)

	for {

		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
