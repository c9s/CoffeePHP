
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

//line src/gutscript/parser.y:422
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 64
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 241

var GutsAct = []int{

	7, 81, 72, 6, 99, 90, 41, 25, 24, 86,
	82, 84, 92, 70, 40, 47, 43, 44, 36, 37,
	38, 89, 48, 78, 36, 37, 38, 83, 52, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
	41, 25, 24, 100, 64, 65, 68, 75, 73, 15,
	66, 26, 74, 39, 80, 16, 33, 26, 32, 27,
	28, 29, 30, 31, 94, 76, 36, 37, 38, 77,
	4, 91, 34, 35, 36, 37, 38, 14, 34, 35,
	87, 14, 88, 15, 73, 82, 93, 96, 50, 16,
	3, 97, 98, 36, 37, 38, 49, 51, 36, 37,
	38, 101, 95, 67, 33, 42, 32, 27, 28, 29,
	30, 31, 33, 85, 32, 27, 28, 29, 30, 31,
	34, 35, 36, 37, 38, 36, 37, 38, 34, 35,
	1, 33, 46, 32, 27, 28, 29, 30, 31, 27,
	28, 29, 30, 31, 2, 47, 12, 34, 35, 10,
	8, 11, 34, 35, 17, 71, 79, 69, 9, 20,
	19, 45, 32, 27, 28, 29, 30, 31, 29, 30,
	31, 18, 25, 24, 0, 0, 34, 35, 0, 34,
	35, 14, 0, 5, 0, 0, 21, 0, 0, 0,
	23, 18, 25, 24, 0, 0, 0, 22, 0, 0,
	0, 14, 0, 0, 0, 13, 21, 0, 0, 0,
	23, 0, 0, 0, 15, 0, 0, 22, 0, 0,
	16, 0, 0, 0, 0, 13, 0, 0, 0, 0,
	0, 0, 0, 0, 15, 0, 0, 0, 0, 0,
	16,
}
var GutsPact = []int{

	166, -1000, -1000, 40, -1000, -1000, -1000, 81, -1000, -1000,
	36, -1000, -1000, 35, 166, 35, 35, -1000, 97, -1000,
	-1000, 35, 91, 83, -1000, -1000, 186, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 22,
	81, -33, 34, 54, 12, 35, -35, 35, 62, -1000,
	26, -1000, 113, 113, 12, 12, 12, 86, 110, 81,
	81, 81, 81, 81, 35, 66, -1000, -1000, 81, -13,
	5, -38, -1000, 81, -1000, -53, 62, -1000, 66, -44,
	-1000, -1000, 7, 35, -1000, -1000, 1, -1000, -1000, 80,
	-1000, 35, -1000, -1000, -1000, -59, 6, -1000, 81, -1000,
	35, 81,
}
var GutsPgo = []int{

	0, 0, 160, 159, 70, 158, 157, 156, 1, 2,
	155, 154, 90, 151, 151, 151, 151, 150, 149, 146,
	3, 144, 130, 113, 102,
}
var GutsR1 = []int{

	0, 22, 21, 20, 12, 12, 12, 4, 4, 4,
	4, 4, 4, 4, 4, 23, 23, 19, 18, 18,
	18, 17, 8, 8, 8, 7, 7, 6, 6, 6,
	5, 13, 14, 14, 15, 15, 16, 16, 24, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 2, 9,
	10, 10, 10, 11,
}
var GutsR2 = []int{

	0, 1, 1, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 3, 4, 3, 5,
	4, 3, 3, 2, 1, 3, 1, 3, 2, 0,
	5, 2, 2, 0, 2, 0, 3, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 0, 4,
}
var GutsChk = []int{

	-1000, -22, -21, -12, -4, 17, -20, -1, -17, -5,
	-18, -13, -19, 39, 15, 48, 54, -11, 5, -2,
	-3, 20, 31, 24, 7, 6, 17, 53, 54, 55,
	56, 57, 52, 50, 66, 67, 12, 13, 14, 17,
	-1, 5, -12, -1, -1, 64, 35, 48, -1, 5,
	5, -4, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, 22, 23, 16, 49, -1, -6,
	48, -10, -9, -1, -20, 21, -1, -20, 36, -7,
	49, -8, 5, 65, 49, -23, 62, -20, -20, 65,
	49, 64, 5, -9, 63, -24, -1, -8, -1, 63,
	37, -1,
}
var GutsDef = []int{

	0, -2, 1, 2, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 0, 0, 0, 0, 53, 54, 55,
	56, 0, 0, 0, 58, 57, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	14, 54, 0, 0, 52, 0, 29, 62, 0, 31,
	0, 4, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 0, 0, 3, 39, 21, 0,
	0, 0, 61, 59, 18, 0, 0, 20, 0, 0,
	28, 26, 24, 0, 63, 17, 0, 19, 30, 0,
	27, 0, 23, 60, 15, 0, 0, 25, 22, 16,
	0, 38,
}
var GutsTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 58, 3, 3, 3, 57, 52, 3,
	48, 49, 55, 53, 65, 54, 3, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	67, 64, 66, 3, 3, 3, 3, 3, 3, 3,
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
		//line src/gutscript/parser.y:147
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:154
		{ 
	        GutsVAL.val = GutsS[Gutspt-0].val 
	    }
	case 3:
		//line src/gutscript/parser.y:160
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 4:
		//line src/gutscript/parser.y:164
		{
	            if stmts, ok := GutsS[Gutspt-2].val.(*ast.StatementList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-2].val
	            }
	      }
	case 5:
		//line src/gutscript/parser.y:171
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 6:
		//line src/gutscript/parser.y:176
		{ }
	case 7:
		//line src/gutscript/parser.y:180
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 8:
		//line src/gutscript/parser.y:181
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 9:
		//line src/gutscript/parser.y:182
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 10:
		//line src/gutscript/parser.y:183
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:184
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:185
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 13:
		//line src/gutscript/parser.y:186
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 14:
		//line src/gutscript/parser.y:187
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 17:
		//line src/gutscript/parser.y:196
		{  }
	case 18:
		//line src/gutscript/parser.y:202
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-1].val.(ast.Expr), GutsS[Gutspt-0].val.(*ast.StatementList))
	        }
	case 19:
		//line src/gutscript/parser.y:207
		{
	            GutsS[Gutspt-4].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-1].val.(ast.Expr),GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-4].val
	        }
	case 20:
		//line src/gutscript/parser.y:213
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 21:
		//line src/gutscript/parser.y:221
		{
	            debug("assignment", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignStatement(ast.CreateVariable(GutsS[Gutspt-2].val.(string)), GutsS[Gutspt-0].val)
	        }
	case 22:
		//line src/gutscript/parser.y:229
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-2].val.(string), "", GutsS[Gutspt-0].val}
	    }
	case 23:
		//line src/gutscript/parser.y:233
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), GutsS[Gutspt-1].val.(string), nil}
	    }
	case 24:
		//line src/gutscript/parser.y:237
		{ 
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), "", nil}
	    }
	case 25:
		//line src/gutscript/parser.y:243
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.FunctionParam) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.FunctionParam))
	            GutsVAL.val = params
	        }
	      }
	case 26:
		//line src/gutscript/parser.y:249
		{
	            GutsVAL.val = []ast.FunctionParam{GutsS[Gutspt-0].val.(ast.FunctionParam)}
	        }
	case 27:
		//line src/gutscript/parser.y:255
		{
	        GutsVAL.val = GutsS[Gutspt-1].val
	    }
	case 28:
		//line src/gutscript/parser.y:258
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 29:
		//line src/gutscript/parser.y:261
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 30:
		//line src/gutscript/parser.y:267
		{ 
	        debug("function declaration", GutsS[Gutspt-4].val, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementList))
	    }
	case 31:
		//line src/gutscript/parser.y:276
		{
	        GutsVAL.val = ast.CreateClassStatement(GutsS[Gutspt-0].val.(string))
	        // $2.(string)     $3 (extend list)   $4 (interface list)
    }
	case 32:
		//line src/gutscript/parser.y:283
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 33:
		//line src/gutscript/parser.y:284
		{ GutsVAL.val = nil }
	case 34:
		//line src/gutscript/parser.y:288
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 35:
		//line src/gutscript/parser.y:289
		{ GutsVAL.val = nil }
	case 36:
		//line src/gutscript/parser.y:294
		{
	            var interfaceList = GutsS[Gutspt-2].val.([]string)
	            interfaceList = append(interfaceList, GutsS[Gutspt-0].val.(string))
	            GutsVAL.val = interfaceList
	        }
	case 37:
		//line src/gutscript/parser.y:300
		{
	            GutsVAL.val = []string{ GutsS[Gutspt-0].val.(string) }
	        }
	case 39:
		//line src/gutscript/parser.y:310
		{
	            if node, ok := GutsS[Gutspt-1].val.(ast.Expr) ; ok {
	                node.Parenthesis = true
	                GutsVAL.val = node
	            } else {
	                panic(" type casting to ast.Expr failed.")
	            }
	            // $$ = $2
        }
	case 40:
		//line src/gutscript/parser.y:320
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 41:
		//line src/gutscript/parser.y:324
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 42:
		//line src/gutscript/parser.y:328
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 43:
		//line src/gutscript/parser.y:332
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 44:
		//line src/gutscript/parser.y:336
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 45:
		//line src/gutscript/parser.y:340
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 46:
		//line src/gutscript/parser.y:344
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 47:
		//line src/gutscript/parser.y:348
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 48:
		//line src/gutscript/parser.y:352
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 49:
		//line src/gutscript/parser.y:356
		{
	            GutsVAL.val = ast.CreateExpr(T_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 50:
		//line src/gutscript/parser.y:360
		{
	            GutsVAL.val = ast.CreateExpr(T_LT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 51:
		//line src/gutscript/parser.y:364
		{
	            GutsVAL.val = ast.CreateExpr(T_GT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 52:
		//line src/gutscript/parser.y:368
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 53:
		//line src/gutscript/parser.y:371
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 54:
		//line src/gutscript/parser.y:374
		{ 
	            // $$ = ast.UnaryExpr{0, $1}
            GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string)) 
	        }
	case 55:
		//line src/gutscript/parser.y:378
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 56:
		//line src/gutscript/parser.y:381
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 57:
		//line src/gutscript/parser.y:386
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 58:
		//line src/gutscript/parser.y:390
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 59:
		//line src/gutscript/parser.y:394
		{
	        GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	    }
	case 60:
		//line src/gutscript/parser.y:399
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.Node) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.Node))
	            GutsVAL.val = params
	        }
	    }
	case 61:
		//line src/gutscript/parser.y:406
		{
	        // create the expr list
        GutsVAL.val = []ast.Node{GutsS[Gutspt-0].val}
	    }
	case 62:
		//line src/gutscript/parser.y:411
		{
	        GutsVAL.val = []ast.Node{}
	    }
	case 63:
		//line src/gutscript/parser.y:417
		{
	        GutsVAL.val = ast.CreateFunctionCall(GutsS[Gutspt-3].val.(string), GutsS[Gutspt-1].val.([]ast.Node))
	    }
	}
	goto Gutsstack /* stack new state and value */
}
