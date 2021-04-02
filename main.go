package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strings"
)

var testPtr *int
var verbosePtr *bool
var ishex bool

var qrCharset = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")
var qrCharsetLen = 45

func betterBase45Encode(s []byte) string {
  if *verbosePtr {
    fmt.Println("BetterEncode:", s)
  }

  // Go through the list two bytes at a time
  firstlist := ""
  for i := 1; i < len(s); i += 2 {
    v := int(s[i-1])*256 + int(s[i])
    for j := 0; j < 3; j++ {
      firstlist = firstlist + string(qrCharset[(v%qrCharsetLen)])
      v = v / qrCharsetLen
    }
  }

  // If odd even of bytes, deal with last byte separately
  if len(s)%2 == 1 {
    v := int(s[len(s)-1])
    for j := 0; j < 2; j++ {
      firstlist = firstlist + string(qrCharset[(v%qrCharsetLen)])
      v = v / qrCharsetLen
    }
  }

  if *verbosePtr {
    fmt.Println("encode:", firstlist)
  }
  return firstlist
}

func betterBase45Decode(s string) []byte {
  if *verbosePtr {
    fmt.Println("decode:", s)
  }

  // Go through the list three bytes at a time
  firstlist := ""
  for i := 2; i < len(s); i += 3 {
    v := 0
    for j := 0; j < 3; j++ {
      v = v*45 + bytes.IndexByte(qrCharset, s[i-j])
    }
    firstlist = firstlist + string(rune(v/256)) + string(rune(v%256))
  }

  // Take care of last two bytes if they exist
  if len(s)%3 > 0 {
    v := 0
    i := len(s) - 1
    for j := 0; j < 2; j++ {
      v = v*45 + bytes.IndexByte(qrCharset, s[i-j])
    }
    firstlist = firstlist + string(rune(v))
  }

  if *verbosePtr {
    fmt.Println("encode:", firstlist)
  }
  return []byte(firstlist)
}

func base45Encode(s []byte) string {
  if *verbosePtr {
    fmt.Println("encode:", s)
  }

  encodeArray := s
  if *verbosePtr {
    fmt.Println("encodeArray:", encodeArray)
  }

  firstlist := []int{}
  for _, x := range encodeArray {
    firstlist = append(firstlist, int(x))
  }
  if *verbosePtr {
    fmt.Println("firstlist:", firstlist)
  }

  secondlist := [][]int{}
  for i := 1; i < len(firstlist); i += 2 {
    secondlist = append(secondlist, []int{firstlist[i-1], firstlist[i]})
  }
  if len(firstlist)%2 == 1 {
    secondlist = append(secondlist, []int{firstlist[len(firstlist)-1]})
  }
  if *verbosePtr {
    fmt.Println("secondlist:", secondlist)
  }

  thirdlist := []int{}
  for _, x := range secondlist {
    if len(x) > 1 {
      thirdlist = append(thirdlist, x[0]*256+x[1])
    } else {
      thirdlist = append(thirdlist, x[0])
    }
  }
  if *verbosePtr {
    fmt.Println("thirdlist:", thirdlist)
  }

  fourthlist := [][]int{}
  fifthlist := []int{}
  for _, x := range secondlist {
    sublist := []int{}
    if len(x) > 1 {
      v := x[0]*256 + x[1]
      for i := 0; i < 3; i++ {
        sublist = append(sublist, v%qrCharsetLen)
        fifthlist = append(fifthlist, v%qrCharsetLen)
        v = v / qrCharsetLen
      }
    } else {
      v := x[0]
      for i := 0; i < 2; i++ {
        sublist = append(sublist, v%qrCharsetLen)
        fifthlist = append(fifthlist, v%qrCharsetLen)
        v = v / qrCharsetLen
      }
    }
    fourthlist = append(fourthlist, sublist)
  }
  if *verbosePtr {
    fmt.Println("fourthlist:", fourthlist)
    fmt.Println("fifthlist:", fifthlist)
  }

  sixthlist := []string{}
  for _, x := range fifthlist {
    sixthlist = append(sixthlist, string(qrCharset[x]))
  }
  if *verbosePtr {
    fmt.Println("sixthlist:", sixthlist)
  }

  thestring := ""
  for _, x := range sixthlist {
    thestring = thestring + x
  }
  return thestring
}

func base45Decode(s string) []byte {
  if *verbosePtr {
    fmt.Println("decode:", s)
  }

  firstlist := strings.Split(s, "")
  if *verbosePtr {
    fmt.Println("firstlist:", firstlist)
  }

  secondlist := []int{}
  for _, x := range firstlist {
    secondlist = append(secondlist, bytes.IndexByte(qrCharset, x[0]))
  }
  if *verbosePtr {
    fmt.Println("secondlist:", secondlist)
  }

  thirdlist := [][]int{}
  sublist := []int{}
  for _, x := range secondlist {
    sublist = append(sublist, x)
    if len(sublist)%3 == 0 {
      thirdlist = append(thirdlist, sublist)
      sublist = nil
    }
  }
  if sublist != nil {
    thirdlist = append(thirdlist, sublist)
  }
  if *verbosePtr {
    fmt.Println("thirdlist:", thirdlist)
  }

  fourthlist := []int{}
  for _, x := range thirdlist {
    i := x[0] + qrCharsetLen*x[1]
    if len(x) > 2 {
      i = i + qrCharsetLen*qrCharsetLen*x[2]
    }
    fourthlist = append(fourthlist, i)
  }
  if *verbosePtr {
    fmt.Println("fourthlist:", fourthlist)
  }

  fifthlist := [][]int{}
  sixthlist := []int{}
  for z, x := range fourthlist {
    sublist = nil
    if x >= 256 {
      sublist = append(sublist, x/256)
      sixthlist = append(sixthlist, x/256)
    } else if z+1 < len(fourthlist) {
      sublist = append(sublist, 0)
      sixthlist = append(sixthlist, 0)
    }      
    sublist = append(sublist, x%256)
    sixthlist = append(sixthlist, x%256)
    fifthlist = append(fifthlist, sublist)
  }
  if *verbosePtr {
    fmt.Println("fifthlist:", fifthlist)
    fmt.Println("sixthlist:", sixthlist)
  }

  seventhlist := []string{}
  for _, x := range sixthlist {
    seventhlist = append(seventhlist, string(rune(x)))
  }
  if *verbosePtr {
    fmt.Println("seventhlist:", seventhlist)
  }

  thestring := ""
  for _, x := range seventhlist {
    thestring = thestring + string(x[0])
  }
  return []byte(thestring)
}

func main() {
  ishex = false
  encodePtr := flag.String("e", "", "a string")
  testPtr = flag.Int("t", 0, "test 1 or 2")
  verbosePtr = flag.Bool("v", false, "Verbose yes/no")

  flag.Parse()

  numRun := 1
  if *testPtr > 0 {
    numRun = 100000
  }

  var theBytes []byte
  var err error
  a := *encodePtr
  if a[0:2] == "0x" {
    ishex = true
    theBytes, err = hex.DecodeString(a[2:])
    if err != nil {
      log.Fatal(err)
    }
  } else {
    theBytes = []byte(a)
  }

  if len(theBytes) > 0 {
    for i := 0; i < numRun; i++ {
      if *testPtr == 0 || *testPtr == 1 {
        encodedString := base45Encode(theBytes)
        decodedBytes := base45Decode(encodedString)
        if i == 0 {
          var decodedString string
          if ishex {
            decodedString = "0x" + string(hex.EncodeToString(decodedBytes))
          } else {
            decodedString = string(decodedBytes)
          }
          fmt.Println("Encoded string:", len(encodedString), encodedString)
          fmt.Println("Decoded string:", len(decodedBytes), decodedString)
        }
      }
      if *testPtr == 0 || *testPtr == 2 {
        betterEncodedString := betterBase45Encode(theBytes)
        betterDecodedBytes := betterBase45Decode(betterEncodedString)
        if i == 0 {
          var betterDecodedString string
          if ishex {
            betterDecodedString = "0x" + string(hex.EncodeToString(betterDecodedBytes))
          } else {
            betterDecodedString = string(betterDecodedBytes)
          }
          fmt.Println("Better Encoded string:", len(betterEncodedString), betterEncodedString)
          fmt.Println("Better Decoded string:", len(betterDecodedBytes), betterDecodedString)
        }
      }
    }
  }
}
