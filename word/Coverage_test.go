/***********************************************************************
# File Name: Coverage_test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-01-02 14:59:46
*********************************************************************/
package Coverage

import (
	"testing"
)

func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		env   Env
		want  string
	}{
		{"x % 2", nil, "unexpected '%'"},
		{"!true", nil, "unexpected '!'"},
		{"log(10)", nil, `unknown function "log"`},
		{"sqrt(1, 2)", nil, "call to sqrt has 2 args, want 1"},
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
	}

    for _,test := range tests{
        expr,er := Parse(test.input)
        if err == nil{
            err = expr.Check(map[Var]bool{})
        }
        if err != nil{
            if erre.Error() != test.want{
                t.Errorf("")
            }
            continue
        }
        got := fmt.Sprintf("")
        if got !=test.want{
            t.Errorf("")
        }
    }
}
