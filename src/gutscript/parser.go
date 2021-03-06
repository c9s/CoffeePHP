
//line src/gutscript/parser.y:2
package gutscript
import __yyfmt__ "fmt"
//line src/gutscript/parser.y:2
		
// vim:et:sw=4:ai:si:ts=4:sts=4:

import "fmt"
import "gutscript/ast"

var regs = make([]int, 26)
var base int

func debug(msg string, vals ...interface{}) {
    if GutsDebug > 0 {
        fmt.Print(msg)
        for _, val := range vals {
            fmt.Printf(" %#v",val)
        }
        fmt.Println("\n")
    }
}


//line src/gutscript/parser.y:26
type GutsSymType struct{
	yys int
    typ TokenType
    val ast.Node
    line int
    pos  int
}

const T_DOT = 57346
const T_IDENTIFIER = 57347
const T_FLOATING = 57348
const T_NUMBER = 57349
const T_STRING = 57350
const T_ONELINE_COMMENT = 57351
const T_COMMENT = 57352
const T_EOF = 57353
const T_EQUAL = 57354
const T_LT_EQUAL = 57355
const T_GT_EQUAL = 57356
const T_INDENT = 57357
const T_OUTDENT = 57358
const T_NEWLINE = 57359
const T_NEW = 57360
const T_CLONE = 57361
const T_IF = 57362
const T_IN = 57363
const T_ELSEIF = 57364
const T_ELSE = 57365
const T_FOR = 57366
const T_SAY = 57367
const T_SPACE = 57368
const T_ECHO = 57369
const T_FOREACH = 57370
const T_TRY = 57371
const T_CATCH = 57372
const T_CLASS = 57373
const T_IS = 57374
const T_DOES = 57375
const T_EXTENDS = 57376
const T_FUNCTION_PROTOTYPE = 57377
const T_FUNCTION_GLYPH = 57378
const T_RANGE_OPERATOR = 57379
const T_CONST = 57380
const T_RETURN = 57381
const T_BREAK = 57382
const T_CONTINUE = 57383
const T_THROW = 57384
const T_NS_SEPARATOR = 57385
const T_NAMESPACE = 57386
const T_CALLSTART = 57387
const T_CALLEND = 57388
const T_OBJECT_OPERATOR = 57389
const UMINUS = 57390
const T_BOOLEAN_OR = 57391
const T_BOOLEAN_AND = 57392

var GutsToknames = []string{
	"T_DOT",
	"T_IDENTIFIER",
	"T_FLOATING",
	"T_NUMBER",
	"T_STRING",
	"T_ONELINE_COMMENT",
	"T_COMMENT",
	"T_EOF",
	"T_EQUAL",
	"T_LT_EQUAL",
	"T_GT_EQUAL",
	"T_INDENT",
	"T_OUTDENT",
	"T_NEWLINE",
	"T_NEW",
	"T_CLONE",
	"T_IF",
	"T_IN",
	"T_ELSEIF",
	"T_ELSE",
	"T_FOR",
	"T_SAY",
	"T_SPACE",
	"T_ECHO",
	"T_FOREACH",
	"T_TRY",
	"T_CATCH",
	"T_CLASS",
	"T_IS",
	"T_DOES",
	"T_EXTENDS",
	"T_FUNCTION_PROTOTYPE",
	"T_FUNCTION_GLYPH",
	"T_RANGE_OPERATOR",
	"T_CONST",
	"T_RETURN",
	"T_BREAK",
	"T_CONTINUE",
	"T_THROW",
	"T_NS_SEPARATOR",
	"T_NAMESPACE",
	"T_CALLSTART",
	"T_CALLEND",
	"T_OBJECT_OPERATOR",
	" (",
	" )",
	" |",
	" ^",
	" &",
	" +",
	" -",
	" *",
	" /",
	" %",
	" !",
	"UMINUS",
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
}
var GutsStatenames = []string{}

const GutsEofCode = 1
const GutsErrCode = 2
const GutsMaxDepth = 200

//line src/gutscript/parser.y:523
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 79
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 267

var GutsAct = []int{

	6, 8, 112, 85, 74, 99, 88, 5, 119, 129,
	120, 118, 94, 42, 4, 72, 45, 46, 86, 49,
	82, 98, 87, 50, 101, 48, 37, 38, 39, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 53, 78, 79, 91, 48, 104, 70, 67,
	75, 121, 40, 41, 123, 37, 38, 39, 76, 49,
	122, 123, 84, 127, 34, 95, 33, 28, 29, 30,
	31, 32, 116, 89, 80, 47, 37, 38, 39, 14,
	81, 35, 36, 100, 68, 27, 27, 14, 75, 3,
	96, 108, 102, 125, 124, 109, 86, 97, 37, 38,
	39, 111, 110, 106, 44, 92, 117, 37, 38, 39,
	35, 36, 52, 51, 34, 93, 33, 28, 29, 30,
	31, 32, 126, 115, 114, 117, 128, 37, 38, 39,
	130, 35, 36, 18, 1, 69, 34, 2, 33, 28,
	29, 30, 31, 32, 11, 34, 9, 33, 28, 29,
	30, 31, 32, 35, 36, 37, 38, 39, 37, 38,
	39, 7, 35, 36, 113, 103, 105, 33, 28, 29,
	30, 31, 32, 43, 26, 25, 90, 43, 26, 25,
	77, 10, 35, 36, 17, 73, 83, 71, 20, 19,
	0, 0, 0, 0, 0, 0, 28, 29, 30, 31,
	32, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	35, 36, 0, 35, 36, 0, 15, 21, 26, 25,
	15, 0, 16, 0, 0, 0, 16, 14, 0, 13,
	0, 107, 22, 0, 0, 0, 24, 0, 0, 0,
	0, 0, 0, 23, 0, 0, 0, 0, 0, 0,
	0, 12, 0, 0, 0, 0, 0, 0, 0, 0,
	15, 0, 0, 0, 0, 0, 16,
}
var GutsPact = []int{

	212, -1000, -1000, 69, -1000, -1000, 95, -1000, -1000, 30,
	-1000, -1000, 172, -1000, 212, 172, 172, -1000, -1000, -1000,
	-1000, 11, 172, 108, 107, -1000, -1000, 212, 172, 172,
	172, 172, 172, 172, 172, 172, 172, 172, 172, 172,
	172, 72, 95, -29, 68, 86, 43, 172, -33, 172,
	64, 9, 23, -1000, 146, 146, 43, 43, 43, 143,
	115, 95, 95, 95, 95, 95, 64, 63, -1000, -1000,
	95, -16, 13, -43, -1000, 95, 56, 12, 100, -50,
	48, -1000, 212, -44, -1000, -1000, 19, 172, -1000, -1000,
	32, 98, -1000, -1000, 168, -1000, -1000, -1000, 91, -1000,
	172, -1000, -1000, -1000, 6, -57, -1000, -1000, -53, 14,
	-1000, 95, 44, -1000, -1000, -1000, 89, -1000, -10, 88,
	-1000, 172, 46, 6, -55, -1000, 95, -1000, 37, 172,
	95,
}
var GutsPgo = []int{

	0, 0, 189, 188, 14, 1, 187, 186, 3, 4,
	185, 184, 89, 181, 180, 176, 166, 165, 164, 161,
	146, 144, 7, 137, 134, 133, 124, 123, 2, 115,
	91, 91, 91,
}
var GutsR1 = []int{

	0, 24, 23, 22, 12, 12, 12, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 29, 29, 21, 20,
	20, 20, 19, 8, 8, 8, 7, 7, 6, 6,
	6, 5, 5, 13, 17, 17, 14, 14, 15, 15,
	16, 16, 28, 28, 28, 18, 18, 26, 26, 27,
	30, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 31,
	32, 25, 3, 2, 9, 10, 10, 10, 11,
}
var GutsR2 = []int{

	0, 1, 1, 3, 3, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 1, 2, 3, 4, 4,
	5, 4, 3, 3, 2, 1, 3, 1, 3, 2,
	0, 5, 5, 5, 4, 0, 2, 0, 2, 0,
	3, 1, 3, 2, 1, 1, 1, 4, 2, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 1, 1, 1, 1, 4,
	2, 1, 1, 1, 1, 3, 1, 0, 4,
}
var GutsChk = []int{

	-1000, -24, -23, -12, -4, -22, -1, -19, -5, -20,
	-13, -21, 39, 17, 15, 48, 54, -11, -25, -2,
	-3, 5, 20, 31, 24, 7, 6, 17, 53, 54,
	55, 56, 57, 52, 50, 67, 68, 12, 13, 14,
	22, 23, -1, 5, -12, -1, -1, 64, 35, 48,
	-1, 5, 5, -4, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -22, 16, 49,
	-1, -6, 48, -10, -9, -1, -22, -14, 34, 21,
	-22, 17, 36, -7, 49, -8, 5, 65, 49, 17,
	-15, 33, 5, -29, 62, 17, -22, -4, 65, 49,
	64, 5, -9, -17, 15, -16, 5, 63, -30, -1,
	-8, -1, -28, -18, -26, -27, 66, -5, 5, 65,
	63, 37, 16, 17, 5, 5, -1, 17, -28, 64,
	-1,
}
var GutsDef = []int{

	0, -2, 1, 2, 6, 7, 8, 9, 10, 11,
	12, 13, 0, 15, 0, 0, 0, 65, 66, 67,
	68, 71, 0, 0, 0, 73, 72, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 14, 71, 0, 0, 64, 0, 30, 77,
	0, 37, 0, 4, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 0, 0, 3, 51,
	22, 0, 0, 0, 76, 74, 0, 39, 0, 0,
	0, 21, 0, 0, 29, 27, 25, 0, 78, 19,
	35, 0, 36, 18, 0, 20, 7, 32, 0, 28,
	0, 24, 75, 33, 0, 38, 41, 16, 0, 0,
	26, 23, 0, 44, 45, 46, 0, 49, 0, 0,
	17, 0, 0, 43, 48, 40, 50, 34, 42, 0,
	47,
}
var GutsTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 58, 3, 3, 3, 57, 52, 3,
	48, 49, 55, 53, 65, 54, 3, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	68, 64, 67, 3, 66, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 62, 3, 63, 51, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 50,
}
var GutsTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 59, 60, 61,
}
var GutsTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var GutsDebug = 100

type GutsLexer interface {
	Lex(lval *GutsSymType) int
	Error(s string)
}

const GutsFlag = -1000

func GutsTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(GutsToknames) {
		if GutsToknames[c-4] != "" {
			return GutsToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func GutsStatname(s int) string {
	if s >= 0 && s < len(GutsStatenames) {
		if GutsStatenames[s] != "" {
			return GutsStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Gutslex1(lex GutsLexer, lval *GutsSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = GutsTok1[0]
		goto out
	}
	if char < len(GutsTok1) {
		c = GutsTok1[char]
		goto out
	}
	if char >= GutsPrivate {
		if char < GutsPrivate+len(GutsTok2) {
			c = GutsTok2[char-GutsPrivate]
			goto out
		}
	}
	for i := 0; i < len(GutsTok3); i += 2 {
		c = GutsTok3[i+0]
		if c == char {
			c = GutsTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = GutsTok2[1] /* unknown char */
	}
	if GutsDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", GutsTokname(c), uint(char))
	}
	return c
}

type GutsParser struct {
	Val GutsSymType
}

func (self *GutsParser) Parse(Gutslex GutsLexer) int {
	var Gutsn int
	var Gutslval GutsSymType
	var GutsVAL GutsSymType
	GutsS := make([]GutsSymType, GutsMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Gutsstate := 0
	Gutschar := -1
	Gutsp := -1
	goto Gutsstack

ret0:
	return 0

ret1:
	return 1

Gutsstack:
	/* put a state and value onto the stack */
	if GutsDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", GutsTokname(Gutschar), GutsStatname(Gutsstate))
	}

	Gutsp++
	if Gutsp >= len(GutsS) {
		nyys := make([]GutsSymType, len(GutsS)*2)
		copy(nyys, GutsS)
		GutsS = nyys
	}
	GutsS[Gutsp] = GutsVAL
	GutsS[Gutsp].yys = Gutsstate

Gutsnewstate:
	Gutsn = GutsPact[Gutsstate]
	if Gutsn <= GutsFlag {
		goto Gutsdefault /* simple state */
	}
	if Gutschar < 0 {
		Gutschar = Gutslex1(Gutslex, &Gutslval)
	}
	Gutsn += Gutschar
	if Gutsn < 0 || Gutsn >= GutsLast {
		goto Gutsdefault
	}
	Gutsn = GutsAct[Gutsn]
	if GutsChk[Gutsn] == Gutschar { /* valid shift */
		Gutschar = -1
		GutsVAL = Gutslval
		Gutsstate = Gutsn
		if Errflag > 0 {
			Errflag--
		}
		goto Gutsstack
	}

Gutsdefault:
	/* default state action */
	Gutsn = GutsDef[Gutsstate]
	if Gutsn == -2 {
		if Gutschar < 0 {
			Gutschar = Gutslex1(Gutslex, &Gutslval)
		}

		/* look through exception table */
		xi := 0
		for {
			if GutsExca[xi+0] == -1 && GutsExca[xi+1] == Gutsstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Gutsn = GutsExca[xi+0]
			if Gutsn < 0 || Gutsn == Gutschar {
				break
			}
		}
		Gutsn = GutsExca[xi+1]
		if Gutsn < 0 {
			goto ret0
		}
	}
	if Gutsn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Gutslex.Error("syntax error")
			Nerrs++
			if GutsDebug >= 1 {
				__yyfmt__.Printf("%s", GutsStatname(Gutsstate))
				__yyfmt__.Printf(" saw %s\n", GutsTokname(Gutschar))
				__yyfmt__.Printf(" Sym: %#v\n",GutsVAL)
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Gutsp >= 0 {
				Gutsn = GutsPact[GutsS[Gutsp].yys] + GutsErrCode
				if Gutsn >= 0 && Gutsn < GutsLast {
					Gutsstate = GutsAct[Gutsn] /* simulate a shift of "error" */
					if GutsChk[Gutsstate] == GutsErrCode {
						goto Gutsstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if GutsDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", GutsS[Gutsp].yys)
				}
				Gutsp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if GutsDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", GutsTokname(Gutschar))
			}
			if Gutschar == GutsEofCode {
				goto ret1
			}
			Gutschar = -1
			goto Gutsnewstate /* try again in the same state */
		}
	}

	/* reduction by production Gutsn */
	if GutsDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Gutsn, GutsStatname(Gutsstate))
	}

	Gutsnt := Gutsn
	Gutspt := Gutsp
	_ = Gutspt // guard against "declared and not used"

	Gutsp -= GutsR2[Gutsn]
	GutsVAL = GutsS[Gutsp+1]

	self.Val = GutsVAL

	/* consult goto table to find next state */
	Gutsn = GutsR1[Gutsn]
	Gutsg := GutsPgo[Gutsn]
	Gutsj := Gutsg + GutsS[Gutsp].yys + 1

	if Gutsj >= GutsLast {
		Gutsstate = GutsAct[Gutsg]
	} else {
		Gutsstate = GutsAct[Gutsj]
		if GutsChk[Gutsstate] != -Gutsn {
			Gutsstate = GutsAct[Gutsg]
		}
	}
	// dummy call; replaced with literal code
	switch Gutsnt {

	case 1:
		//line src/gutscript/parser.y:151
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:158
		{ 
	        GutsVAL.val = GutsS[Gutspt-0].val 
	    }
	case 3:
		//line src/gutscript/parser.y:164
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 4:
		//line src/gutscript/parser.y:168
		{
	            if stmts, ok := GutsS[Gutspt-2].val.(*ast.StatementList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-2].val
	            }
	      }
	case 5:
		//line src/gutscript/parser.y:174
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 6:
		//line src/gutscript/parser.y:176
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 7:
		//line src/gutscript/parser.y:184
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 8:
		//line src/gutscript/parser.y:185
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 9:
		//line src/gutscript/parser.y:186
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 10:
		//line src/gutscript/parser.y:187
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:188
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:189
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 13:
		//line src/gutscript/parser.y:190
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 14:
		//line src/gutscript/parser.y:191
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 15:
		//line src/gutscript/parser.y:192
		{  }
	case 18:
		//line src/gutscript/parser.y:201
		{  }
	case 19:
		//line src/gutscript/parser.y:207
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-2].val.(ast.Expr), GutsS[Gutspt-1].val.(*ast.StatementList))
	        }
	case 20:
		//line src/gutscript/parser.y:212
		{
	            GutsS[Gutspt-4].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-2].val.(ast.Expr),GutsS[Gutspt-1].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-4].val
	        }
	case 21:
		//line src/gutscript/parser.y:218
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-1].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 22:
		//line src/gutscript/parser.y:226
		{
	            debug("assignment", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignment(ast.CreateVariable(GutsS[Gutspt-2].val.(string)), GutsS[Gutspt-0].val)
	        }
	case 23:
		//line src/gutscript/parser.y:234
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-2].val.(string), "", GutsS[Gutspt-0].val}
	    }
	case 24:
		//line src/gutscript/parser.y:238
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), GutsS[Gutspt-1].val.(string), nil}
	    }
	case 25:
		//line src/gutscript/parser.y:242
		{ 
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), "", nil}
	    }
	case 26:
		//line src/gutscript/parser.y:248
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.FunctionParam) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.FunctionParam))
	            GutsVAL.val = params
	        }
	      }
	case 27:
		//line src/gutscript/parser.y:254
		{
	            GutsVAL.val = []ast.FunctionParam{GutsS[Gutspt-0].val.(ast.FunctionParam)}
	        }
	case 28:
		//line src/gutscript/parser.y:260
		{
	        GutsVAL.val = GutsS[Gutspt-1].val
	    }
	case 29:
		//line src/gutscript/parser.y:263
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 30:
		//line src/gutscript/parser.y:266
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 31:
		//line src/gutscript/parser.y:273
		{ 
	        debug("function declaration", GutsS[Gutspt-4].val, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementList))
	    }
	case 32:
		//line src/gutscript/parser.y:279
		{
	        var stmts = ast.StatementList{}
	        stmts.Append(GutsS[Gutspt-0].val)
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), &stmts)
	    }
	case 33:
		//line src/gutscript/parser.y:288
		{
	            // debug("class", $2, $3, $4)
            var cls = ast.CreateClass(GutsS[Gutspt-3].val.(string)).(ast.Class)
	
	            // decl extends
            if GutsS[Gutspt-2].val != nil {
	                // debug("extends", $3)
                cls.SetSuper(GutsS[Gutspt-2].val.(string))
	            }
	
	            // decl does
            if GutsS[Gutspt-1].val != nil {
	                if infs, ok := GutsS[Gutspt-1].val.([]string) ; ok {
	                    cls.SetInterfaces(infs)
	                } else {
	                    panic(fmt.Errorf("Can not cast interface"))
	                }
	            }
	
	            // class body
            if GutsS[Gutspt-0].val != nil {
	                cls.Body = GutsS[Gutspt-0].val.(*ast.StatementList)
	            }
	            GutsVAL.val = cls
	            debug("class",GutsVAL.val)
	        }
	case 34:
		//line src/gutscript/parser.y:318
		{ 
	            GutsVAL.val = GutsS[Gutspt-2].val
	        }
	case 35:
		//line src/gutscript/parser.y:321
		{ GutsVAL.val = nil }
	case 36:
		//line src/gutscript/parser.y:325
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 37:
		//line src/gutscript/parser.y:326
		{ GutsVAL.val = nil }
	case 38:
		//line src/gutscript/parser.y:330
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 39:
		//line src/gutscript/parser.y:331
		{ GutsVAL.val = nil }
	case 40:
		//line src/gutscript/parser.y:336
		{
	            var interfaceList = GutsS[Gutspt-2].val.([]string)
	            interfaceList = append(interfaceList, GutsS[Gutspt-0].val.(string))
	            GutsVAL.val = interfaceList
	        }
	case 41:
		//line src/gutscript/parser.y:342
		{
	            GutsVAL.val = []string{ GutsS[Gutspt-0].val.(string) }
	        }
	case 42:
		//line src/gutscript/parser.y:349
		{ 
	            if stmts, ok := GutsS[Gutspt-2].val.(*ast.StatementList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-2].val
	            }
	      }
	case 43:
		//line src/gutscript/parser.y:356
		{  
	            GutsVAL.val = GutsS[Gutspt-1].val
	      }
	case 44:
		//line src/gutscript/parser.y:360
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 45:
		//line src/gutscript/parser.y:368
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 46:
		//line src/gutscript/parser.y:369
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 47:
		//line src/gutscript/parser.y:374
		{ 
	            member := ast.CreateClassMember(GutsS[Gutspt-2].val.(string))
	            member.SetValue(GutsS[Gutspt-0].val)
	            GutsVAL.val = member
	      }
	case 48:
		//line src/gutscript/parser.y:379
		{ 
	            GutsVAL.val = ast.CreateClassMember(GutsS[Gutspt-0].val.(string))
	        }
	case 49:
		//line src/gutscript/parser.y:385
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 51:
		//line src/gutscript/parser.y:391
		{
	            if node, ok := GutsS[Gutspt-1].val.(ast.Expr) ; ok {
	                node.Parenthesis = true
	                GutsVAL.val = node
	            } else {
	                panic(" type casting to ast.Expr failed.")
	            }
	            // $$ = $2
        }
	case 52:
		//line src/gutscript/parser.y:401
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 53:
		//line src/gutscript/parser.y:405
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 54:
		//line src/gutscript/parser.y:409
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 55:
		//line src/gutscript/parser.y:413
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 56:
		//line src/gutscript/parser.y:417
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 57:
		//line src/gutscript/parser.y:421
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 58:
		//line src/gutscript/parser.y:425
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 59:
		//line src/gutscript/parser.y:429
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 60:
		//line src/gutscript/parser.y:433
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 61:
		//line src/gutscript/parser.y:437
		{
	            GutsVAL.val = ast.CreateExpr(T_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 62:
		//line src/gutscript/parser.y:441
		{
	            GutsVAL.val = ast.CreateExpr(T_LT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 63:
		//line src/gutscript/parser.y:445
		{
	            GutsVAL.val = ast.CreateExpr(T_GT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 64:
		//line src/gutscript/parser.y:449
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 65:
		//line src/gutscript/parser.y:452
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 66:
		//line src/gutscript/parser.y:456
		{ 
	            GutsVAL.val = GutsS[Gutspt-0].val
	        }
	case 67:
		//line src/gutscript/parser.y:460
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 68:
		//line src/gutscript/parser.y:464
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 70:
		//line src/gutscript/parser.y:476
		{
	            // $$ = ast.CreateObjectVariable("this")
        }
	case 71:
		//line src/gutscript/parser.y:482
		{ 
	            // $$ = ast.UnaryExpr{0, $1}
            GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string)) 
	        }
	case 72:
		//line src/gutscript/parser.y:487
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 73:
		//line src/gutscript/parser.y:491
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 74:
		//line src/gutscript/parser.y:495
		{
	        GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	    }
	case 75:
		//line src/gutscript/parser.y:500
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.Node) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.Node))
	            GutsVAL.val = params
	        }
	    }
	case 76:
		//line src/gutscript/parser.y:507
		{
	        // create the expr list
        GutsVAL.val = []ast.Node{GutsS[Gutspt-0].val}
	    }
	case 77:
		//line src/gutscript/parser.y:512
		{
	        GutsVAL.val = []ast.Node{}
	    }
	case 78:
		//line src/gutscript/parser.y:518
		{
	        GutsVAL.val = ast.CreateFunctionCall(GutsS[Gutspt-3].val.(string), GutsS[Gutspt-1].val.([]ast.Node))
	    }
	}
	goto Gutsstack /* stack new state and value */
}
