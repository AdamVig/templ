package parser

import (
	"github.com/a-h/parse"
)

var childrenExpressionParser = parse.StringFrom(
	openBraceWithOptionalPadding,
	parse.OptionalWhitespace,
	parse.String("children..."),
	parse.OptionalWhitespace,
	closeBraceWithOptionalPadding,
)

var childrenExpression = parse.Func(func(in *parse.Input) (n Node, ok bool, err error) {
	start := in.Index()
	_, ok, err = childrenExpressionParser.Parse(in)
	if err != nil || !ok {
		if ok {
			in.Seek(start)
		}
		return
	}
	return &ChildrenExpression{
		Range: NewRange(in.PositionAt(start), in.Position()),
	}, true, nil
})
