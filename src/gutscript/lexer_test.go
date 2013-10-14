package gutscript

// vim:list:

import "testing"

var lextests = []struct {
	input    string
	expected TokenType
	cnt      int
}{
	{"a = 102", T_IDENTIFIER, 1},
	{"a = 102\n", T_NEWLINE, 1},
	{"a = 102\nb = 333\n", T_NEWLINE, 2},
	{"abc123 = 100", T_IDENTIFIER, 1},
	{"af = 102", T_NUMBER, 1},
	{"af = 102", '=', 1},
	{"af = 3.1415926", T_FLOATING, 1},
	{"af = bk", T_IDENTIFIER, 2},
	{"b = 3 * -10", T_IDENTIFIER, 1},
	{"// oneline comment", T_ONELINE_COMMENT, 1},
	{"/* comment */", T_COMMENT, 1},
	{`a = "string content"`, T_STRING, 1},
	{`a = "string content contains quote \""`, T_STRING, 1},
	{`if a > 3
    a = 10
    b = 20
`, T_IF, 1},
	{`if a > 3
    a = 10
    b = 20
else
    a = 1
`, T_ELSE, 1},
	{`if a > 3
    a = 10
    b = 20
elseif a < 10
    b = 2
else
    a = 1
`, T_ELSEIF, 1},
	{`foo :: ->
    a = 10
    b = 20
`, T_FUNCTION_PROTOTYPE, 1},
}

func expectLexItems(items chan *GutsSymType, expectedItems []TokenType) {
	for {
		item := <-items
		if item == nil {
			break
		}
		fmt.Printf("Got token %s: '%s'\n", GetTokenName(int(item.typ)), item.val)
		if item.typ == T_EOF || item.typ == eof {
			break
		}
	}

}

func expectLexInput(t *testing.T, input string, typ TokenType, cnt int) {
	lexer := GutsLex{
		input: input,
		start: 0,
		pos:   0,
		items: make(chan *GutsSymType, 100),
	}
	go lexer.run()

	var found int = 0
	var item *GutsSymType
	for {
		item = <-lexer.items
		if item == nil {
			break
		}
		t.Logf("Got token %s: %s", GetTokenName(int(item.typ)), item.val)
		if item.typ == typ {
			found++
		}
		if item.typ == T_EOF || item.typ == eof {
			break
		}
	}
	if found != cnt {
		t.Fatalf("Expecting token %s x %d", GetTokenName(int(typ)), cnt)
	}
	lexer.close()
}

func BenchmarkLexer(b *testing.B) {
	input := `
	a = 102
	foo = 200
	pi = 3.1415926
	// oneline comment
	`
	var item *GutsSymType
	for i := 0; i < b.N; i++ {
		lexer := GutsLex{
			input: input,
			start: 0,
			pos:   0,
			items: make(chan *GutsSymType, 100),
		}
		go lexer.run()
		for {
			item = <-lexer.items
			if item == nil || item.typ == T_EOF {
				break
			}
		}
		lexer.close()
	}
}

func TestLexer(t *testing.T) {
	for _, test := range lextests {
		t.Log("Testing lex from ", test.input)
		expectLexInput(t, test.input, test.expected, test.cnt)
	}
}
