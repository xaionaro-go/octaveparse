<p xmlns:dct="http://purl.org/dc/terms/" xmlns:vcard="http://www.w3.org/2001/vcard-rdf/3.0#">
  <a rel="license"
     href="http://creativecommons.org/publicdomain/zero/1.0/">
    <img src="http://i.creativecommons.org/p/zero/1.0/88x31.png" style="border-style: none;" alt="CC0" />
  </a>
  <br />
  To the extent possible under law,
  <a rel="dct:publisher"
     href="https://github.com/xaionaro/">
    <span property="dct:title">Dmitrii Okunev</span></a>
  has waived all copyright and related or neighboring rights to
  "<span property="dct:title">An Octave formulas parser</span>.
This work is published from:
<span property="vcard:Country" datatype="dct:ISO3166"
      content="IE" about="https://github.com/xaionaro-go/octaveparse">
  Ireland</span>".
</p>


# About

Convert formulas from Octave to Go or/and C

# Quick start

Getting some crazy formula in Octave:
```sh
octave:50> pkg load symbolic
octave:51> f = a * (x**2) + b * x + c
f = (sym)

     2
  a⋅x  + b⋅x + c

octave:52> E2(x) = simplify((int(f**3, x, -x0, x0)) / (2 * x0)) - M2
E2(x) = (symfun)

         3   6         4 ⎛       2⎞
        a ⋅x₀    3⋅a⋅x₀ ⋅⎝a⋅c + b ⎠    3       2 ⎛       2⎞
  -M₂ + ────── + ────────────────── + c  + c⋅x₀ ⋅⎝a⋅c + b ⎠
          7              5

octave:53> sympref display flat
octave:54> simplify(solve(E2(x),a))
ans = (sym) Matrix([[(30*2**(1/3)*21**(2/3)*b**2*x0**2 + 8*2**(1/3)*21**(2/3)*c**2 - 42*c*x0**2*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(1/3) - 84**(1/3)*x0**4*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(2/3))/(30*x0**4*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(1/3))], [21**(2/3)*(-84*21**(1/3)*c*x0**2*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 + sqrt(3)*I) + 42**(2/3)*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(2/3)*(1 + sqrt(3)*I)**2 - 168*2**(1/3)*(15*b**2*x0**2 + 4*c**2))/(1260*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 + sqrt(3)*I))], [21**(2/3)*(-84*21**(1/3)*c*x0**2*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 - sqrt(3)*I) + 42**(2/3)*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(2/3)*(1 - sqrt(3)*I)**2 - 168*2**(1/3)*(15*b**2*x0**2 + 4*c**2))/(1260*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 - sqrt(3)*I))]])  (3x1 matrix)
```

Converting it to Go:
```sh
go run github.com/xaionaro-go/octaveparse/cmd/octaveparse -syms b,c,x0,M2 <<EOF
Matrix([[(30*2**(1/3)*21**(2/3)*b**2*x0**2 + 8*2**(1/3)*21**(2/3)*c**2 - 42*c*x0**2*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(1/3) - 84**(1/3)*x0**4*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(2/3))/(30*x0**4*((-1125*M2 + 180*b**2*c*x0**2 + 432*c**3 + 5*sqrt(3)*x0**6*sqrt((16875*M2**2 - 5400*M2*b**2*c*x0**2 - 12960*M2*c**3 + 3780*b**6*x0**6 + 3456*b**4*c**2*x0**4 + 2880*b**2*c**4*x0**2 + 2560*c**6)/x0**12))/x0**6)**(1/3))], [21**(2/3)*(-84*21**(1/3)*c*x0**2*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 + sqrt(3)*I) + 42**(2/3)*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(2/3)*(1 + sqrt(3)*I)**2 - 168*2**(1/3)*(15*b**2*x0**2 + 4*c**2))/(1260*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 + sqrt(3)*I))], [21**(2/3)*(-84*21**(1/3)*c*x0**2*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 - sqrt(3)*I) + 42**(2/3)*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(2/3)*(1 - sqrt(3)*I)**2 - 168*2**(1/3)*(15*b**2*x0**2 + 4*c**2))/(1260*x0**4*((-1125*M2 + 1125*b**2*c*x0**2 + 2007*c**3 - 315*c*(3*b**2*x0**2 + 5*c**2) + sqrt(3)*x0**6*sqrt((28*(15*b**2*x0**2 + 4*c**2)**3 + 27*(-125*M2 + 125*b**2*c*x0**2 + 223*c**3 - 35*c*(3*b**2*x0**2 + 5*c**2))**2)/x0**12))/x0**6)**(1/3)*(1 - sqrt(3)*I))]])
EOF
```

The output is:
```go
[][]any{[]any{((((((((30 * math.Pow(2, (1 / 3))) * math.Pow(21, (2 / 3))) * math.Pow(b, 2)) * math.Pow(x0, 2)) + (((8 * math.Pow(2, (1 / 3))) * math.Pow(21, (2 / 3))) * math.Pow(c, 2))) - (((42 * c) * math.Pow(x0, 2)) * math.Pow((((((0 - (1125 * M2)) + (((180 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (432 * math.Pow(c, 3))) + (((5 * math.Sqrt(3)) * math.Pow(x0, 6)) * math.Sqrt(((((((((16875 * math.Pow(M2, 2)) - ((((5400 * M2) * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) - ((12960 * M2) * math.Pow(c, 3))) + ((3780 * math.Pow(b, 6)) * math.Pow(x0, 6))) + (((3456 * math.Pow(b, 4)) * math.Pow(c, 2)) * math.Pow(x0, 4))) + (((2880 * math.Pow(b, 2)) * math.Pow(c, 4)) * math.Pow(x0, 2))) + (2560 * math.Pow(c, 6))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (1 / 3)))) - ((math.Pow(84, (1 / 3)) * math.Pow(x0, 4)) * math.Pow((((((0 - (1125 * M2)) + (((180 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (432 * math.Pow(c, 3))) + (((5 * math.Sqrt(3)) * math.Pow(x0, 6)) * math.Sqrt(((((((((16875 * math.Pow(M2, 2)) - ((((5400 * M2) * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) - ((12960 * M2) * math.Pow(c, 3))) + ((3780 * math.Pow(b, 6)) * math.Pow(x0, 6))) + (((3456 * math.Pow(b, 4)) * math.Pow(c, 2)) * math.Pow(x0, 4))) + (((2880 * math.Pow(b, 2)) * math.Pow(c, 4)) * math.Pow(x0, 2))) + (2560 * math.Pow(c, 6))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (2 / 3)))) / ((30 * math.Pow(x0, 4)) * math.Pow((((((0 - (1125 * M2)) + (((180 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (432 * math.Pow(c, 3))) + (((5 * math.Sqrt(3)) * math.Pow(x0, 6)) * math.Sqrt(((((((((16875 * math.Pow(M2, 2)) - ((((5400 * M2) * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) - ((12960 * M2) * math.Pow(c, 3))) + ((3780 * math.Pow(b, 6)) * math.Pow(x0, 6))) + (((3456 * math.Pow(b, 4)) * math.Pow(c, 2)) * math.Pow(x0, 4))) + (((2880 * math.Pow(b, 2)) * math.Pow(c, 4)) * math.Pow(x0, 2))) + (2560 * math.Pow(c, 6))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (1 / 3))))},[]any{((math.Pow(21, (2 / 3)) * (((0 - (((((84 * math.Pow(21, (1 / 3))) * c) * math.Pow(x0, 2)) * math.Pow(((((((0 - (1125 * M2)) + (((1125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (2007 * math.Pow(c, 3))) - ((315 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))) + ((math.Sqrt(3) * math.Pow(x0, 6)) * math.Sqrt((((28 * math.Pow((((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2))), 3)) + (27 * math.Pow(((((0 - (125 * M2)) + (((125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (223 * math.Pow(c, 3))) - ((35 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))), 2))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (1 / 3))) * (1 + (math.Sqrt(3) * complex(0, 1))))) + (((math.Pow(42, (2 / 3)) * math.Pow(x0, 4)) * math.Pow(((((((0 - (1125 * M2)) + (((1125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (2007 * math.Pow(c, 3))) - ((315 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))) + ((math.Sqrt(3) * math.Pow(x0, 6)) * math.Sqrt((((28 * math.Pow((((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2))), 3)) + (27 * math.Pow(((((0 - (125 * M2)) + (((125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (223 * math.Pow(c, 3))) - ((35 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))), 2))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (2 / 3))) * math.Pow((1 + (math.Sqrt(3) * complex(0, 1))), 2))) - ((168 * math.Pow(2, (1 / 3))) * (((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2)))))) / (((1260 * math.Pow(x0, 4)) * math.Pow(((((((0 - (1125 * M2)) + (((1125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (2007 * math.Pow(c, 3))) - ((315 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))) + ((math.Sqrt(3) * math.Pow(x0, 6)) * math.Sqrt((((28 * math.Pow((((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2))), 3)) + (27 * math.Pow(((((0 - (125 * M2)) + (((125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (223 * math.Pow(c, 3))) - ((35 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))), 2))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (1 / 3))) * (1 + (math.Sqrt(3) * complex(0, 1)))))},[]any{((math.Pow(21, (2 / 3)) * (((0 - (((((84 * math.Pow(21, (1 / 3))) * c) * math.Pow(x0, 2)) * math.Pow(((((((0 - (1125 * M2)) + (((1125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (2007 * math.Pow(c, 3))) - ((315 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))) + ((math.Sqrt(3) * math.Pow(x0, 6)) * math.Sqrt((((28 * math.Pow((((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2))), 3)) + (27 * math.Pow(((((0 - (125 * M2)) + (((125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (223 * math.Pow(c, 3))) - ((35 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))), 2))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (1 / 3))) * (1 - (math.Sqrt(3) * complex(0, 1))))) + (((math.Pow(42, (2 / 3)) * math.Pow(x0, 4)) * math.Pow(((((((0 - (1125 * M2)) + (((1125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (2007 * math.Pow(c, 3))) - ((315 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))) + ((math.Sqrt(3) * math.Pow(x0, 6)) * math.Sqrt((((28 * math.Pow((((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2))), 3)) + (27 * math.Pow(((((0 - (125 * M2)) + (((125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (223 * math.Pow(c, 3))) - ((35 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))), 2))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (2 / 3))) * math.Pow((1 - (math.Sqrt(3) * complex(0, 1))), 2))) - ((168 * math.Pow(2, (1 / 3))) * (((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2)))))) / (((1260 * math.Pow(x0, 4)) * math.Pow(((((((0 - (1125 * M2)) + (((1125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (2007 * math.Pow(c, 3))) - ((315 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))) + ((math.Sqrt(3) * math.Pow(x0, 6)) * math.Sqrt((((28 * math.Pow((((15 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (4 * math.Pow(c, 2))), 3)) + (27 * math.Pow(((((0 - (125 * M2)) + (((125 * math.Pow(b, 2)) * c) * math.Pow(x0, 2))) + (223 * math.Pow(c, 3))) - ((35 * c) * (((3 * math.Pow(b, 2)) * math.Pow(x0, 2)) + (5 * math.Pow(c, 2))))), 2))) / math.Pow(x0, 12))))) / math.Pow(x0, 6)), (1 / 3))) * (1 - (math.Sqrt(3) * complex(0, 1)))))}}
```