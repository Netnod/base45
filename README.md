# base45
Specification of base45 encoding

Here you can find both the internet draft and example code in go.

All changes to the draft should be in draft-faltstrom-base45.xml.

```
$ ./base45 -encode 'Hello!!'
encode: Hello!!
encodeArray: [H e l l o ! !]
firstlist: [72 101 108 108 111 33 33]
secondlist: [[72 101] [108 108] [111 33] [33]]
thirdlist: [18533 27756 28449 33]
fourthlist: [[38 6 9] [36 31 13] [9 2 14] [33 0]]
fifthlist: [38 6 9 36 31 13 9 2 14 33 0]
sixthlist: [% 6 9   V D 9 2 E X 0]
Encoded string: %69 VD92EX0
encode: Hello!!
encode: %69 VD92EX0
Better Encoded string: %69 VD92EX0
```

```
$ ./base45 -decode '%69 VD92EX0'
decode: %69 VD92EX0
firstlist: [% 6 9   V D 9 2 E X 0]
secondlist: [38 6 9 36 31 13 9 2 14 33 0]
thirdlist: [[38 6 9] [36 31 13] [9 2 14] [33 0]]
fourthlist: [18533 27756 28449 33]
fifthlist: [[72 101] [108 108] [111 33] [33]]
sixthlist: [72 101 108 108 111 33 33]
seventhlist: [H e l l o ! !]
Decoded string: Hello!!
decode: %69 VD92EX0
encode: Hello!!
Better Decoded string: Hello!!
```
