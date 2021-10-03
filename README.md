# base45
Specification of base45 encoding

Here you can find both the internet draft and example code in go.

Licensed with a BSD-3 clause to Netnod.

All changes to the draft should be in draft-faltstrom-base45.xml.

```
$ ./base45 -e 'Hello!!'
Encoded string: 11 %69 VD92EX0
Better Encoded string: 11 %69 VD92EX0
```

```
$ ./base45 -d '%69 VD92EX0'
Decoded string: 7 Hello!!
Better Decoded string: 7 Hello!!
```

```
$ ./base45 -b '%69 VD92EX0'
Encoded string: 17 VV4:97Y+AHA7MY831
Decoded string: 11 %69 VD92EX0
Better Encoded string: 17 VV4:97Y+AHA7MY831
Better Decoded string: 11 %69 VD92EX0
```

```
$ ./base45 -v -b '%69 VD92EX0'
encode3: [37 54 57 32 86 68 57 50 69 88 48]
encodeArray: [37 54 57 32 86 68 57 50 69 88 48]
firstlist: [37 54 57 32 86 68 57 50 69 88 48]
secondlist: [[37 54] [57 32] [86 68] [57 50] [69 88] [48]]
thirdlist: [9526 14624 22084 14642 17752 48]
fourthlist: [[31 31 4] [44 9 7] [34 40 10] [17 10 7] [22 34 8] [3 1]]
fifthlist: [31 31 4 44 9 7 34 40 10 17 10 7 22 34 8 3 1]
sixthlist: [V V 4 : 9 7 Y + A H A 7 M Y 8 3 1]
decode: VV4:97Y+AHA7MY831
firstlist: [V V 4 : 9 7 Y + A H A 7 M Y 8 3 1]
secondlist: [31 31 4 44 9 7 34 40 10 17 10 7 22 34 8 3 1]
thirdlist: [[31 31 4] [44 9 7] [34 40 10] [17 10 7] [22 34 8] [3 1]]
fourthlist: [9526 14624 22084 14642 17752 48]
fifthlist: [[37 54] [57 32] [86 68] [57 50] [69 88] [48]]
sixthlist: [37 54 57 32 86 68 57 50 69 88 48]
seventhlist: [% 6 9   V D 9 2 E X 0]
Encoded string: 17 VV4:97Y+AHA7MY831
Decoded string: 11 %69 VD92EX0
BetterEncode: [37 54 57 32 86 68 57 50 69 88 48]
encode1: VV4:97Y+AHA7MY831
decode: VV4:97Y+AHA7MY831
Better Encoded string: 17 VV4:97Y+AHA7MY831
Better Decoded string: 11 %69 VD92EX0
```
