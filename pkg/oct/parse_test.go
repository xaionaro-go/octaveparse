package oct

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	type testCase struct {
		Name           string
		Sample         string
		ExpectedResult string
	}

	for _, testCase := range []testCase{
		{
			Name:           "simple",
			Sample:         "a*x**2 + b*x + c",
			ExpectedResult: `{"LHS":{"LHS":{"LHS":"a","RHS":{"LHS":"x","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"b","RHS":"x","Op":"*"},"Op":"+"},"RHS":"c","Op":"+"}`,
		},
		{
			Name: "big-matrix-complex",
			//         0         1         2         3         4         5         6         7         8         9        10        11        12        13        14        15        16        17        18        19        20        21        22        23        24        25        26        27
			//         01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567
			Sample:         "Matrix([[(30*2**(1/3)*21**(2/3)*b**2*x0**2 + 8*2**(1/3)*21**(2/3)*c**2 - 42*c*x0**2*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(1/3) - 84**(1/3)*x0**4*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(2/3))/(30*x0**4*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(1/3))], [21**(2/3)*(-84*21**(1/3)*c*x0**2*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 + sqrt(3)*I) + 42**(2/3)*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(2/3)*(1 + sqrt(3)*I)**2 - 168*2**(1/3)*(15*b**2*x0**2 + 4*c**2))/(1260*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 + sqrt(3)*I))], [21**(2/3)*(-84*21**(1/3)*c*x0**2*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 - sqrt(3)*I) + 42**(2/3)*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(2/3)*(1 - sqrt(3)*I)**2 - 168*2**(1/3)*(15*b**2*x0**2 + 4*c**2))/(1260*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 - sqrt(3)*I))]])",
			ExpectedResult: `[[{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":30,"RHS":{"LHS":2,"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":21,"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":8,"RHS":{"LHS":2,"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":21,"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":42,"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":180,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":432,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":5,"RHS":{"Base":3},"Op":"*"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":16875,"RHS":{"LHS":"M2","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":5400,"RHS":"M2","Op":"*"},"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":12960,"RHS":"M2","Op":"*"},"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":3780,"RHS":{"LHS":"b","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":3456,"RHS":{"LHS":"b","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":2880,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"c","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2560,"RHS":{"LHS":"c","RHS":6,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":84,"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":180,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":432,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":5,"RHS":{"Base":3},"Op":"*"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":16875,"RHS":{"LHS":"M2","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":5400,"RHS":"M2","Op":"*"},"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":12960,"RHS":"M2","Op":"*"},"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":3780,"RHS":{"LHS":"b","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":3456,"RHS":{"LHS":"b","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":2880,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"c","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2560,"RHS":{"LHS":"c","RHS":6,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":30,"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":180,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":432,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":5,"RHS":{"Base":3},"Op":"*"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":16875,"RHS":{"LHS":"M2","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":5400,"RHS":"M2","Op":"*"},"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":12960,"RHS":"M2","Op":"*"},"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":3780,"RHS":{"LHS":"b","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":3456,"RHS":{"LHS":"b","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":{"LHS":2880,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"c","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2560,"RHS":{"LHS":"c","RHS":6,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"Op":"/"}],[{"LHS":{"LHS":{"LHS":21,"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"RHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":84,"RHS":{"LHS":21,"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":1125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2007,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":315,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"Base":3},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":28,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":3,"Op":"**"},"Op":"*"},"RHS":{"LHS":27,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":223,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":35,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":1,"RHS":{"LHS":{"Base":3},"RHS":{},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":42,"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":1125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2007,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":315,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"Base":3},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":28,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":3,"Op":"**"},"Op":"*"},"RHS":{"LHS":27,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":223,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":35,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":1,"RHS":{"LHS":{"Base":3},"RHS":{},"Op":"*"},"Op":"+"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":168,"RHS":{"LHS":2,"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":1260,"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":1125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2007,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":315,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"Base":3},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":28,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":3,"Op":"**"},"Op":"*"},"RHS":{"LHS":27,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":223,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":35,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":1,"RHS":{"LHS":{"Base":3},"RHS":{},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"/"}],[{"LHS":{"LHS":{"LHS":21,"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"RHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":84,"RHS":{"LHS":21,"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":1125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2007,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":315,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"Base":3},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":28,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":3,"Op":"**"},"Op":"*"},"RHS":{"LHS":27,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":223,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":35,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":1,"RHS":{"LHS":{"Base":3},"RHS":{},"Op":"*"},"Op":"-"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":42,"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":1125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2007,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":315,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"Base":3},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":28,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":3,"Op":"**"},"Op":"*"},"RHS":{"LHS":27,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":223,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":35,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":2,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":1,"RHS":{"LHS":{"Base":3},"RHS":{},"Op":"*"},"Op":"-"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":168,"RHS":{"LHS":2,"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":1260,"RHS":{"LHS":"x0","RHS":4,"Op":"**"},"Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":1125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":1125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":2007,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":315,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"Base":3},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"*"},"RHS":{"Base":{"LHS":{"LHS":{"LHS":28,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":15,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":4,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":3,"Op":"**"},"Op":"*"},"RHS":{"LHS":27,"RHS":{"LHS":{"LHS":{"LHS":{"LHS":{"LHS":0,"RHS":{"LHS":125,"RHS":"M2","Op":"*"},"Op":"-"},"RHS":{"LHS":{"LHS":{"LHS":125,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":"c","Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":223,"RHS":{"LHS":"c","RHS":3,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":{"LHS":35,"RHS":"c","Op":"*"},"RHS":{"LHS":{"LHS":{"LHS":3,"RHS":{"LHS":"b","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":"x0","RHS":2,"Op":"**"},"Op":"*"},"RHS":{"LHS":5,"RHS":{"LHS":"c","RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"Op":"*"},"Op":"-"},"RHS":2,"Op":"**"},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":12,"Op":"**"},"Op":"/"}},"Op":"*"},"Op":"+"},"RHS":{"LHS":"x0","RHS":6,"Op":"**"},"Op":"/"},"RHS":{"LHS":1,"RHS":3,"Op":"/"},"Op":"**"},"Op":"*"},"RHS":{"LHS":1,"RHS":{"LHS":{"Base":3},"RHS":{},"Op":"*"},"Op":"-"},"Op":"*"},"Op":"/"}]]`,
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			p := NewParser()
			p.AddSym("x0", "M2", "a", "b", "c", "x")
			result, err := p.Parse(testCase.Sample)
			require.NoError(t, err)
			b, err := json.Marshal(result)
			require.NoError(t, err)
			require.Equal(
				t,
				testCase.ExpectedResult,
				string(b),
			)
		})
	}
}