// Code generated by goyacc -p expr -o pkg/logql/expr.y.go pkg/logql/expr.y. DO NOT EDIT.

//line pkg/logql/expr.y:2
package logql

import __yyfmt__ "fmt"

//line pkg/logql/expr.y:2

import (
	"github.com/prometheus/prometheus/pkg/labels"
	"time"
)

//line pkg/logql/expr.y:10
type exprSymType struct {
	yys                   int
	Expr                  Expr
	Filter                labels.MatchType
	Grouping              *grouping
	Labels                []string
	LogExpr               LogSelectorExpr
	LogRangeExpr          *logRange
	Matcher               *labels.Matcher
	Matchers              []*labels.Matcher
	RangeAggregationExpr  SampleExpr
	RangeOp               string
	Selector              []*labels.Matcher
	VectorAggregationExpr SampleExpr
	MetricExpr            SampleExpr
	VectorOp              string
	BinOpExpr             SampleExpr
	binOp                 string
	str                   string
	duration              time.Duration
	LiteralExpr           *literalExpr
	BinOpModifier         BinOpOptions
	LabelParser           struct{ op, param string }
}

const IDENTIFIER = 57346
const STRING = 57347
const NUMBER = 57348
const DURATION = 57349
const MATCHERS = 57350
const LABELS = 57351
const EQ = 57352
const RE = 57353
const NRE = 57354
const OPEN_BRACE = 57355
const CLOSE_BRACE = 57356
const OPEN_BRACKET = 57357
const CLOSE_BRACKET = 57358
const COMMA = 57359
const DOT = 57360
const PIPE_MATCH = 57361
const PIPE_EXACT = 57362
const OPEN_PARENTHESIS = 57363
const CLOSE_PARENTHESIS = 57364
const BY = 57365
const WITHOUT = 57366
const COUNT_OVER_TIME = 57367
const RATE = 57368
const SUM = 57369
const AVG = 57370
const MAX = 57371
const MIN = 57372
const COUNT = 57373
const STDDEV = 57374
const STDVAR = 57375
const BOTTOMK = 57376
const TOPK = 57377
const BYTES_OVER_TIME = 57378
const BYTES_RATE = 57379
const BOOL = 57380
const JSON = 57381
const REGEXP = 57382
const LOGFMT = 57383
const PIPE = 57384
const OR = 57385
const AND = 57386
const UNLESS = 57387
const CMP_EQ = 57388
const NEQ = 57389
const LT = 57390
const LTE = 57391
const GT = 57392
const GTE = 57393
const ADD = 57394
const SUB = 57395
const MUL = 57396
const DIV = 57397
const MOD = 57398
const POW = 57399

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"STRING",
	"NUMBER",
	"DURATION",
	"MATCHERS",
	"LABELS",
	"EQ",
	"RE",
	"NRE",
	"OPEN_BRACE",
	"CLOSE_BRACE",
	"OPEN_BRACKET",
	"CLOSE_BRACKET",
	"COMMA",
	"DOT",
	"PIPE_MATCH",
	"PIPE_EXACT",
	"OPEN_PARENTHESIS",
	"CLOSE_PARENTHESIS",
	"BY",
	"WITHOUT",
	"COUNT_OVER_TIME",
	"RATE",
	"SUM",
	"AVG",
	"MAX",
	"MIN",
	"COUNT",
	"STDDEV",
	"STDVAR",
	"BOTTOMK",
	"TOPK",
	"BYTES_OVER_TIME",
	"BYTES_RATE",
	"BOOL",
	"JSON",
	"REGEXP",
	"LOGFMT",
	"PIPE",
	"OR",
	"AND",
	"UNLESS",
	"CMP_EQ",
	"NEQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"POW",
}
var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16

//line pkg/logql/expr.y:206

//line yacctab:1
var exprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 3,
	1, 2,
	22, 2,
	43, 2,
	44, 2,
	45, 2,
	46, 2,
	48, 2,
	49, 2,
	50, 2,
	51, 2,
	52, 2,
	53, 2,
	54, 2,
	55, 2,
	56, 2,
	57, 2,
	-2, 0,
	-1, 54,
	43, 2,
	44, 2,
	45, 2,
	46, 2,
	48, 2,
	49, 2,
	50, 2,
	51, 2,
	52, 2,
	53, 2,
	54, 2,
	55, 2,
	56, 2,
	57, 2,
	-2, 0,
}

const exprPrivate = 57344

const exprLast = 277

var exprAct = [...]int{

	62, 4, 45, 136, 58, 3, 98, 38, 53, 55,
	2, 68, 54, 30, 31, 32, 39, 40, 43, 44,
	41, 42, 33, 34, 35, 36, 37, 38, 14, 33,
	34, 35, 36, 37, 38, 11, 35, 36, 37, 38,
	85, 87, 86, 6, 63, 64, 148, 17, 18, 21,
	22, 24, 25, 23, 26, 27, 28, 29, 19, 20,
	144, 133, 61, 101, 63, 64, 99, 94, 96, 97,
	88, 145, 105, 145, 15, 16, 147, 106, 146, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120, 104, 11, 103, 60, 122, 93, 134,
	140, 127, 100, 139, 95, 135, 131, 132, 66, 138,
	31, 32, 39, 40, 43, 44, 41, 42, 33, 34,
	35, 36, 37, 38, 84, 91, 65, 83, 57, 126,
	59, 102, 125, 124, 142, 127, 143, 90, 11, 123,
	92, 121, 149, 137, 59, 46, 6, 67, 10, 150,
	17, 18, 21, 22, 24, 25, 23, 26, 27, 28,
	29, 19, 20, 39, 40, 43, 44, 41, 42, 33,
	34, 35, 36, 37, 38, 9, 13, 15, 16, 69,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 47, 8, 5, 12, 47, 130, 47,
	7, 56, 130, 50, 1, 0, 0, 50, 128, 50,
	48, 49, 128, 89, 48, 49, 48, 49, 50, 89,
	0, 0, 50, 0, 0, 48, 49, 0, 141, 48,
	49, 47, 129, 52, 0, 0, 0, 52, 51, 52,
	0, 50, 51, 0, 51, 0, 0, 0, 48, 49,
	0, 0, 0, 51, 0, 0, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 52, 0, 0, 0, 0, 51,
}
var exprPact = [...]int{

	22, -1000, -30, 229, -1000, -1000, 22, -1000, -1000, -1000,
	-1000, 126, 75, 41, -1000, 120, 102, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, 122, -1000, -1000, -1000, -1000,
	-1000, -1000, 1, 48, 197, -30, 123, 84, -1000, 57,
	81, 125, 74, 72, 51, -1000, -1000, 22, -1000, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, -1000, -1000, -1000, -1000, 136, -1000, -1000,
	-1000, -1000, 140, -1000, 134, 128, 127, 124, 210, 195,
	81, 39, 82, 22, 139, 139, 66, 117, 117, -18,
	-18, -50, -50, -50, -50, -23, -23, -23, -23, -23,
	-23, -1000, -1000, -1000, -1000, -1000, -1000, 98, -1000, -1000,
	-1000, 191, 206, 21, 22, 38, 56, -1000, 54, -1000,
	-1000, -1000, -1000, 24, -1000, 138, -1000, -1000, 21, -1000,
	-1000,
}
var exprPgo = [...]int{

	0, 204, 9, 2, 0, 3, 5, 1, 6, 4,
	201, 200, 196, 195, 194, 176, 175, 148, 147, 145,
}
var exprR1 = [...]int{

	0, 1, 2, 2, 7, 7, 7, 7, 7, 6,
	6, 6, 6, 6, 6, 8, 8, 8, 8, 8,
	11, 14, 14, 14, 14, 14, 3, 3, 3, 3,
	13, 13, 13, 10, 10, 9, 9, 9, 9, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 18, 18, 17, 17, 17, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 12, 12,
	12, 12, 5, 5, 4, 4, 19, 19, 19,
}
var exprR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 2, 3, 3, 2, 2, 3, 3, 3, 2,
	4, 4, 5, 5, 6, 7, 1, 1, 1, 1,
	3, 3, 3, 1, 3, 3, 3, 3, 3, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 0, 1, 1, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 4, 4, 2, 2, 3,
}
var exprChk = [...]int{

	-1000, -1, -2, -6, -7, -13, 21, -11, -14, -16,
	-17, 13, -12, -15, 6, 52, 53, 25, 26, 36,
	37, 27, 28, 31, 29, 30, 32, 33, 34, 35,
	43, 44, 45, 52, 53, 54, 55, 56, 57, 46,
	47, 50, 51, 48, 49, -3, -19, 2, 19, 20,
	12, 47, 42, -7, -6, -2, -10, 2, -9, 4,
	21, 21, -4, 23, 24, 6, 6, -18, 38, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, 5, 2, 39, 41, 40, 22, 22,
	14, 2, 17, 14, 10, 47, 11, 12, -8, -6,
	21, -7, 6, 21, 21, 21, -2, -2, -2, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, 5, -9, 5, 5, 5, 5, -3, 2, 22,
	7, -6, -8, 22, 17, -7, -5, 4, -5, 5,
	2, 22, -4, -7, 22, 17, 22, 22, 22, 4,
	-4,
}
var exprDef = [...]int{

	0, -2, 1, -2, 3, 9, 0, 4, 5, 6,
	7, 0, 0, 0, 56, 0, 0, 68, 69, 70,
	71, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	54, 54, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 54, 54, 54, 54, 0, 11, 14, 26, 27,
	28, 29, 0, 3, -2, 0, 0, 0, 33, 0,
	0, 0, 0, 0, 0, 57, 58, 0, 55, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 10, 13, 76, 77, 0, 8, 12,
	30, 31, 0, 32, 0, 0, 0, 0, 0, 0,
	0, 3, 56, 0, 0, 0, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
	53, 78, 34, 35, 36, 37, 38, 0, 19, 20,
	15, 0, 0, 21, 0, 3, 0, 72, 0, 16,
	18, 17, 23, 3, 22, 0, 74, 75, 24, 73,
	25,
}
var exprTok1 = [...]int{

	1,
}
var exprTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57,
}
var exprTok3 = [...]int{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := exprPact[state]
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && exprChk[exprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || exprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := exprExca[i]
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		token = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = exprTok3[i+0]
		if token == char {
			token = exprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:71
		{
			exprlex.(*lexer).expr = exprDollar[1].Expr
		}
	case 2:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:74
		{
			exprVAL.Expr = exprDollar[1].LogExpr
		}
	case 3:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:75
		{
			exprVAL.Expr = exprDollar[1].MetricExpr
		}
	case 4:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:79
		{
			exprVAL.MetricExpr = exprDollar[1].RangeAggregationExpr
		}
	case 5:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:80
		{
			exprVAL.MetricExpr = exprDollar[1].VectorAggregationExpr
		}
	case 6:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:81
		{
			exprVAL.MetricExpr = exprDollar[1].BinOpExpr
		}
	case 7:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:82
		{
			exprVAL.MetricExpr = exprDollar[1].LiteralExpr
		}
	case 8:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:83
		{
			exprVAL.MetricExpr = exprDollar[2].MetricExpr
		}
	case 9:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:87
		{
			exprVAL.LogExpr = newMatcherExpr(exprDollar[1].Selector)
		}
	case 10:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:88
		{
			exprVAL.LogExpr = NewFilterExpr(exprDollar[1].LogExpr, exprDollar[2].Filter, exprDollar[3].str)
		}
	case 11:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line pkg/logql/expr.y:89
		{
			exprVAL.LogExpr = newParserExpr(exprDollar[1].LogExpr, exprDollar[2].LabelParser.op, exprDollar[2].LabelParser.param)
		}
	case 12:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:90
		{
			exprVAL.LogExpr = exprDollar[2].LogExpr
		}
	case 15:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line pkg/logql/expr.y:96
		{
			exprVAL.LogRangeExpr = newLogRange(exprDollar[1].LogExpr, exprDollar[2].duration)
		}
	case 16:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:97
		{
			exprVAL.LogRangeExpr = addFilterToLogRangeExpr(exprDollar[1].LogRangeExpr, exprDollar[2].Filter, exprDollar[3].str)
		}
	case 17:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:98
		{
			exprVAL.LogRangeExpr = exprDollar[2].LogRangeExpr
		}
	case 20:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:103
		{
			exprVAL.RangeAggregationExpr = newRangeAggregationExpr(exprDollar[3].LogRangeExpr, exprDollar[1].RangeOp)
		}
	case 21:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:107
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].MetricExpr, exprDollar[1].VectorOp, nil, nil)
		}
	case 22:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line pkg/logql/expr.y:108
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[4].MetricExpr, exprDollar[1].VectorOp, exprDollar[2].Grouping, nil)
		}
	case 23:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line pkg/logql/expr.y:109
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].MetricExpr, exprDollar[1].VectorOp, exprDollar[5].Grouping, nil)
		}
	case 24:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line pkg/logql/expr.y:111
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].MetricExpr, exprDollar[1].VectorOp, nil, &exprDollar[3].str)
		}
	case 25:
		exprDollar = exprS[exprpt-7 : exprpt+1]
//line pkg/logql/expr.y:112
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].MetricExpr, exprDollar[1].VectorOp, exprDollar[7].Grouping, &exprDollar[3].str)
		}
	case 26:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:116
		{
			exprVAL.Filter = labels.MatchRegexp
		}
	case 27:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:117
		{
			exprVAL.Filter = labels.MatchEqual
		}
	case 28:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:118
		{
			exprVAL.Filter = labels.MatchNotRegexp
		}
	case 29:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:119
		{
			exprVAL.Filter = labels.MatchNotEqual
		}
	case 30:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:123
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 31:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:124
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 32:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:125
		{
		}
	case 33:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:129
		{
			exprVAL.Matchers = []*labels.Matcher{exprDollar[1].Matcher}
		}
	case 34:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:130
		{
			exprVAL.Matchers = append(exprDollar[1].Matchers, exprDollar[3].Matcher)
		}
	case 35:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:134
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 36:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:135
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchNotEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 37:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:136
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 38:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:137
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchNotRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 39:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:143
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("or", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 40:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:144
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("and", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 41:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:145
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("unless", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 42:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:146
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("+", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 43:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:147
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("-", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 44:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:148
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("*", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 45:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:149
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("/", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 46:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:150
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("%", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 47:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:151
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("^", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 48:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:152
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("==", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 49:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:153
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("!=", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 50:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:154
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr(">", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 51:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:155
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr(">=", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 52:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:156
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("<", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 53:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:157
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("<=", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 54:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line pkg/logql/expr.y:161
		{
			exprVAL.BinOpModifier = BinOpOptions{}
		}
	case 55:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:162
		{
			exprVAL.BinOpModifier = BinOpOptions{ReturnBool: true}
		}
	case 56:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:166
		{
			exprVAL.LiteralExpr = mustNewLiteralExpr(exprDollar[1].str, false)
		}
	case 57:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line pkg/logql/expr.y:167
		{
			exprVAL.LiteralExpr = mustNewLiteralExpr(exprDollar[2].str, false)
		}
	case 58:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line pkg/logql/expr.y:168
		{
			exprVAL.LiteralExpr = mustNewLiteralExpr(exprDollar[2].str, true)
		}
	case 59:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:172
		{
			exprVAL.VectorOp = OpTypeSum
		}
	case 60:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:173
		{
			exprVAL.VectorOp = OpTypeAvg
		}
	case 61:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:174
		{
			exprVAL.VectorOp = OpTypeCount
		}
	case 62:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:175
		{
			exprVAL.VectorOp = OpTypeMax
		}
	case 63:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:176
		{
			exprVAL.VectorOp = OpTypeMin
		}
	case 64:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:177
		{
			exprVAL.VectorOp = OpTypeStddev
		}
	case 65:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:178
		{
			exprVAL.VectorOp = OpTypeStdvar
		}
	case 66:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:179
		{
			exprVAL.VectorOp = OpTypeBottomK
		}
	case 67:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:180
		{
			exprVAL.VectorOp = OpTypeTopK
		}
	case 68:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:184
		{
			exprVAL.RangeOp = OpRangeTypeCount
		}
	case 69:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:185
		{
			exprVAL.RangeOp = OpRangeTypeRate
		}
	case 70:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:186
		{
			exprVAL.RangeOp = OpRangeTypeBytes
		}
	case 71:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:187
		{
			exprVAL.RangeOp = OpRangeTypeBytesRate
		}
	case 72:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:192
		{
			exprVAL.Labels = []string{exprDollar[1].str}
		}
	case 73:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:193
		{
			exprVAL.Labels = append(exprDollar[1].Labels, exprDollar[3].str)
		}
	case 74:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:197
		{
			exprVAL.Grouping = &grouping{without: false, groups: exprDollar[3].Labels}
		}
	case 75:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:198
		{
			exprVAL.Grouping = &grouping{without: true, groups: exprDollar[3].Labels}
		}
	case 76:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line pkg/logql/expr.y:202
		{
			exprVAL.LabelParser = struct{ op, param string }{op: OpParserTypeJSON}
		}
	case 77:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line pkg/logql/expr.y:203
		{
			exprVAL.LabelParser = struct{ op, param string }{op: OpParserTypeLogfmt}
		}
	case 78:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:204
		{
			exprVAL.LabelParser = struct{ op, param string }{op: OpParserTypeRegexp, param: exprDollar[3].str}
		}
	}
	goto exprstack /* stack new state and value */
}
