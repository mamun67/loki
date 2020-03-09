package logql

import (
	"bytes"
	"fmt"
	"regexp"
	"regexp/syntax"

	"github.com/prometheus/prometheus/pkg/labels"
)

// LineFilter is a function to filter logs.
type LineFilter interface {
	Filter(line []byte) bool
}

type notFilter struct {
	LineFilter
}

func (n notFilter) Filter(line []byte) bool {
	return !n.LineFilter.Filter(line)
}

// newNotFilter creates a new filter which matches only if the base filter doesn't match.
// If the base filter is a `or` it will recursively simplify with `and` operations.
func newNotFilter(base LineFilter) LineFilter {
	// not(a|b) = not(a) and not(b)
	if or, ok := base.(orFilter); ok {
		return newAndFilter(newNotFilter(or.left), newNotFilter(or.right))
	}
	return notFilter{LineFilter: base}
}

type andFilter struct {
	left  LineFilter
	right LineFilter
}

// newAndFilter creates a new filter which matches only if left and right matches.
func newAndFilter(left LineFilter, right LineFilter) LineFilter {
	return andFilter{
		left:  left,
		right: right,
	}
}

func (a andFilter) Filter(line []byte) bool {
	return a.left.Filter(line) && a.right.Filter(line)
}

type orFilter struct {
	left  LineFilter
	right LineFilter
}

// newOrFilter creates a new filter which matches only if left and right matches.
func newOrFilter(left LineFilter, right LineFilter) LineFilter {
	return orFilter{
		left:  left,
		right: right,
	}
}

// chainOrFilter is a syntax sugar to chain multiple `or` filters. (1 or many)
func chainOrFilter(curr, new LineFilter) LineFilter {
	if curr == nil {
		return new
	}
	return newOrFilter(curr, new)
}

func (a orFilter) Filter(line []byte) bool {
	return a.left.Filter(line) || a.right.Filter(line)
}

type regexpFilter struct {
	*regexp.Regexp
}

// newRegexpFilter creates a new line filter for a given regexp.
// If match is false the filter is the negation of the regexp.
func newRegexpFilter(re string, match bool) (LineFilter, error) {
	reg, err := regexp.Compile(re)
	if err != nil {
		return nil, err
	}
	f := regexpFilter{reg}
	if match {
		return f, nil
	}
	return newNotFilter(f), nil
}

func (r regexpFilter) Filter(line []byte) bool {
	return r.Match(line)
}

type literalFilter []byte

func (l literalFilter) Filter(line []byte) bool {
	return bytes.Contains(line, l)
}

func (l literalFilter) String() string {
	return string(l)
}

// newFilter creates a new line filter from a match string and type.
func newFilter(match string, mt labels.MatchType) (LineFilter, error) {
	switch mt {
	case labels.MatchRegexp:
		return parseRegexpFilter(match, true)
	case labels.MatchNotRegexp:
		return parseRegexpFilter(match, false)
	case labels.MatchEqual:
		return literalFilter(match), nil
	case labels.MatchNotEqual:
		return newNotFilter(literalFilter(match)), nil

	default:
		return nil, fmt.Errorf("unknown matcher: %v", match)
	}
}

// parseRegexpFilter parses a regexp and attempt to simplify it with only literal filters.
// If not possible it will returns the original regexp filter.
func parseRegexpFilter(re string, match bool) (LineFilter, error) {
	reg, err := syntax.Parse(re, syntax.Perl)
	if err != nil {
		return nil, err
	}
	reg = reg.Simplify()

	// attempt to improve regex with tricks
	f, ok := simplify(reg)
	if !ok {
		return newRegexpFilter(re, match)
	}
	if match {
		return f, nil
	}
	return newNotFilter(f), nil
}

// simplify a regexp expression by replacing it, when possible, with a succession of literal filters.
// For example `(foo|bar)` will be replaced by  `literalFilter(foo) or literalFilter(bar)`
func simplify(reg *syntax.Regexp) (LineFilter, bool) {
	switch reg.Op {
	case syntax.OpAlternate:
		return simplifyAlternate(reg)
	case syntax.OpConcat:
		return simplifyConcat(reg, nil)
	case syntax.OpCapture:
		clearCapture(reg)
		return simplify(reg)
	case syntax.OpLiteral:
		return literalFilter([]byte(string(reg.Rune))), true
	}
	return nil, false
}

// clearCapture removes capture operation as they are not used for filtering.
func clearCapture(regs ...*syntax.Regexp) {
	for _, r := range regs {
		if r.Op == syntax.OpCapture {
			*r = *r.Sub[0]
		}
	}
}

// simplifyAlternate simplifies, when possible, alternate regexp expressions such as:
// (foo|bar) or (foo|(bar|buzz)).
func simplifyAlternate(reg *syntax.Regexp) (LineFilter, bool) {
	clearCapture(reg.Sub...)
	// attempt to simplify the first leg
	f, ok := simplify(reg.Sub[0])
	if !ok {
		return nil, false
	}
	// merge the rest of the legs
	for i := 1; i < len(reg.Sub); i++ {
		f2, ok := simplify(reg.Sub[i])
		if !ok {
			return nil, false
		}
		f = newOrFilter(f, f2)
	}
	return f, true
}

// simplifyConcat attempt to simplify concat operations.
// Concat operations are either literal and star such as foo.* .*foo.* .*foo
// which is a literalFilter.
// Or a literal and alternates operation (see simplifyConcatAlternate), which represent a multiplication of alternates.
// Anything else is rejected.
func simplifyConcat(reg *syntax.Regexp, baseLiteral []byte) (LineFilter, bool) {
	clearCapture(reg.Sub...)
	// we support only simplication of concat operation with 3 sub expressions.
	// for instance .*foo.*bar contains 4 subs (.*+foo+.*+bar) and can't be simplified.
	if len(reg.Sub) > 3 {
		return nil, false
	}

	var curr LineFilter
	var ok bool
	literals := 0
	for _, sub := range reg.Sub {
		if sub.Op == syntax.OpLiteral {
			// only one literal is allowed.
			if literals != 0 {
				return nil, false
			}
			literals++
			baseLiteral = append(baseLiteral, []byte(string(sub.Rune))...)
			continue
		}
		// if we have an alternate we must also have a base literal to apply the concatenation with.
		if sub.Op == syntax.OpAlternate && baseLiteral != nil {
			if curr, ok = simplifyConcatAlternate(sub, baseLiteral, curr); !ok {
				return nil, false
			}
			continue
		}
		if sub.Op == syntax.OpStar && sub.Sub[0].Op == syntax.OpAnyCharNotNL {
			continue
		}
		return nil, false
	}

	// if we have a filter from concat alternates.
	if curr != nil {
		return curr, true
	}

	// if we have only a concat with literals.
	if baseLiteral != nil {
		return literalFilter(baseLiteral), true

	}

	return nil, false
}

// simplifyConcatAlternate simplifies concat alternate operations.
// A concat alternate is found when a concat operation has a sub alternate and is preceded by a literal.
// For instance bar|b|buzz is expressed as b(ar|(?:)|uzz) => b concat alternate(ar,(?:),uzz).
// (?:) being an OpEmptyMatch and b being the literal to concat all alternates (ar,(?:),uzz) with.
func simplifyConcatAlternate(reg *syntax.Regexp, literal []byte, curr LineFilter) (LineFilter, bool) {
	for _, alt := range reg.Sub {
		switch alt.Op {
		case syntax.OpEmptyMatch:
			curr = chainOrFilter(curr, literalFilter(literal))
		case syntax.OpLiteral:
			// concat the root literal with the alternate one.
			altBytes := []byte(string(alt.Rune))
			altLiteral := make([]byte, 0, len(literal)+len(altBytes))
			altLiteral = append(altLiteral, literal...)
			altLiteral = append(altLiteral, altBytes...)
			curr = chainOrFilter(curr, literalFilter(altLiteral))
		case syntax.OpConcat:
			f, ok := simplifyConcat(alt, literal)
			if !ok {
				return nil, false
			}
			curr = chainOrFilter(curr, f)
		case syntax.OpStar:
			if alt.Sub[0].Op != syntax.OpAnyCharNotNL {
				return nil, false
			}
			curr = chainOrFilter(curr, literalFilter(literal))
		default:
			return nil, false
		}
	}
	if curr != nil {
		return curr, true
	}
	return nil, false
}
