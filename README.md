# base45
Specification of base45 encoding

Here you can find both the internet draft and example code in go.

All changes to the draft should be in draft-faltstrom-base45.xml.

```
$ ./base45 -encode 'Hello!'
encode: Hello!
encodeArray: [H e l l o !]
firstlist: [72 101 108 108 111 33]
secondlist: [[72 101] [108 108] [111 33]]
thirdlist: [18533 27756 28449]
fourthlist: [[38 6 9] [36 31 13] [9 2 14]]
fifthlist: [38 6 9 36 31 13 9 2 14]
sixthlist: [% 6 9   V D 9 2 E]
Encoded string: %69 VD92E
```

```
$ ./base45 -decode '%69 VD92E'
decode: %69 VD92E
firstlist: [% 6 9   V D 9 2 E]
secondlist: [38 6 9 36 31 13 9 2 14]
thirdlist: [[38 6 9] [36 31 13] [9 2 14]]
fourthlist: [18533 27756 28449]
fifthlist: [[72 101] [108 108] [111 33]]
sixthlist: [72 101 108 108 111 33]
seventhlist: [H e l l o !]
Decoded string: Hello!
```
