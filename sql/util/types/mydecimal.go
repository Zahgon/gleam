// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// constant values.
const (
	ten0 = 1
	ten1 = 10
	ten2 = 100
	ten3 = 1000
	ten4 = 10000
	ten5 = 100000
	ten6 = 1000000
	ten7 = 10000000
	ten8 = 100000000
	ten9 = 1000000000

	maxWordBufLen = 9 // A MyDecimal holds 9 words.
	digitsPerWord = 9 // A word holds 9 digits.
	wordSize      = 4 // A word is 4 bytes int32.
	digMask       = ten8
	wordBase      = ten9
	wordMax       = wordBase - 1
	notFixedDec   = 31

	MaxFraction = 30
	DivFracIncr = 4
)

var (
	wordBufLen = 9
	powers10   = [10]int32{ten0, ten1, ten2, ten3, ten4, ten5, ten6, ten7, ten8, ten9}
	dig2bytes  = [10]int{0, 1, 1, 2, 2, 3, 3, 4, 4, 4}
	fracMax    = [8]int32{
		900000000,
		990000000,
		999000000,
		999900000,
		999990000,
		999999000,
		999999900,
		999999990,
	}
	zeroMyDecimal = MyDecimal{}
)

// add adds a and b and carry, returns the sum and new carry.
func add(a, b, carry int32) (int32, int32) { _ = "STUB: not implemented"; return 0, 0 }

// sub subtracts b and carry from a, returns the diff and new carry.
func sub(a, b, carry int32) (int32, int32) { _ = "STUB: not implemented"; return 0, 0 }

// sub2 subtracts b and carry from a, returns the diff and new carry.
// the new carry may be 2.
func sub2(a, b, carry int32) (int32, int32) { _ = "STUB: not implemented"; return 0, 0 }

// fixWordCntError limits word count in wordBufLen, and returns overflow or truncate error.
func fixWordCntError(wordsInt, wordsFrac int) (newWordsInt int, newWordsFrac int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

/*
countLeadingZeroes returns the number of leading zeroes that can be removed from fraction.

@param   i    start index
@param   word value to compare against list of powers of 10
*/
func countLeadingZeroes(i int, word int32) int { _ = "STUB: not implemented"; return 0 }

/*
countTrailingZeros returns the number of trailing zeroes that can be removed from fraction.

@param   i    start index
@param   word  value to compare against list of powers of 10
*/
func countTrailingZeroes(i int, word int32) int { _ = "STUB: not implemented"; return 0 }

func digitsToWords(digits int) int { _ = "STUB: not implemented"; return 0 }

// MyDecimal represents a decimal value.
type MyDecimal struct {
	// The number of *decimal* digits before the point.
	digitsInt int8

	// The number of decimal digits after the point.
	digitsFrac int8

	// result fraction digits.
	resultFrac int8

	negative bool

	// An array of int32 words.
	// A word is an int32 value can hold 9 digits.(0 <= word < wordBase)
	wordBuf [maxWordBufLen]int32
}

// IsNegative returns whether a decimal is negative.
func (d *MyDecimal) IsNegative() bool {
	_ = "STUB: not implemented"

	// String returns the decimal string representation rounded to resultFrac.
	return false
}

func (d *MyDecimal) String() string { _ = "STUB: not implemented"; return "" }

func (d *MyDecimal) stringSize() int {
	_ = "STUB: not implemented"
	// sign, zero integer and dot.
	return 0
}

func (d *MyDecimal) removeLeadingZeros() (wordIdx int, digitsInt int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// ToString converts decimal to its printable string representation without rounding.
//
//	RETURN VALUE
//
//	    str       - result string
//	    errCode   - eDecOK/eDecTruncate/eDecOverflow
func (d *MyDecimal) ToString() (str []byte) { _ = "STUB: not implemented"; return nil }

/* symbol 0 before digital point */

// FromString parses decimal from string.
func (d *MyDecimal) FromString(str []byte) error { _ = "STUB: not implemented"; return nil }

// TODO: need a way to check if there is at least one digit.

// Shift shifts decimal digits in given number (with rounding if it need), shift > 0 means shift to left shift,
// shift < 0 means right shift. In fact it is multiplying on 10^shift.
//
// RETURN
//
//	eDecOK          OK
//	eDecOverflow    operation lead to overflow, number is untoched
//	eDecTruncated   number was rounded to fit into buffer
func (d *MyDecimal) Shift(shift int) error { _ = "STUB: not implemented"; return nil }

/* index of first non zero digit (all indexes from 0) */

/* index of position after last decimal digit */

/* index of digit position just after point */

/* new point position */

/* number of digits in result */

/* cat off fraction part to allow new number to fit in our buffer */

/*
   We lost all digits (they will be shifted out of buffer), so we can
   just return 0.
*/

/*
   Calculate left/right shift to align decimal digits inside our bug
   digits correctly.
*/

/*
   If number is shifted and correctly aligned in buffer we can finish.
*/

/* already shifted as it should be */

/* if new 'decimal front' is in first digit, we do not need move digits */

/* need to move digits */

/* move left */

/* move right */

/*
   If there are gaps then fill them with 0.

   Only one of following 'for' loops will work because wordIdxBegin <= wordIdxEnd.
*/

/* We don't want negative new_point below */

/*
digitBounds returns bounds of decimal digits in the number.

	start - index (from 0 ) of first decimal digits.
	end   - index of position just after last decimal digit.
*/
func (d *MyDecimal) digitBounds() (start, end int) { _ = "STUB: not implemented"; return 0, 0 }

/* find non-zero digit from number beginning */

/* find non-zero decimal digit from number beginning */

/* find non-zero digit at the end */

/* find non-zero decimal digit from the end */

/*
doMiniLeftShift does left shift for alignment of data in buffer.

	shift   number of decimal digits on which it should be shifted
	beg/end bounds of decimal digits (see digitsBounds())

NOTE

	Result fitting in the buffer should be garanted.
	'shift' have to be from 1 to digitsPerWord-1 (inclusive)
*/
func (d *MyDecimal) doMiniLeftShift(shift, beg, end int) { _ = "STUB: not implemented"; return }

/*
doMiniRightShift does right shift for alignment of data in buffer.

	shift   number of decimal digits on which it should be shifted
	beg/end bounds of decimal digits (see digitsBounds())

NOTE

	Result fitting in the buffer should be garanted.
	'shift' have to be from 1 to digitsPerWord-1 (inclusive)
*/
func (d *MyDecimal) doMiniRightShift(shift, beg, end int) { _ = "STUB: not implemented"; return }

// Round rounds the decimal to "frac" digits.
//
//	to     - result buffer. d == to is allowed
//	frac   - to what position after fraction point to round. can be negative!
//	mode   - round to nearest even or truncate
//
// NOTES
//
//	scale can be negative !
//	one TRUNCATED error (line XXX below) isn't treated very logical :(
//
// RETURN VALUE
//
//	eDecOK/eDecTruncated
func (d *MyDecimal) Round(to *MyDecimal, frac int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// wordsFracTo is the number of fraction words in buffer.

// Do increment.

// If any word after scale is not zero, do increment.
// e.g ceiling 3.0001 to scale 1, gets 3.1

// the first digit after scale.
// If first digit after scale is 5 and round even, do incre if digit at scale is odd.

/* TODO - fix this code as it won't work for CEILING mode */

/*
   In case we're rounding e.g. 1.5e9 to 2.0e9, the decimal words inside
   the buffer are as follows.

   Before <1, 5e8>
   After  <2, 5e8>

   Hence we need to set the 2nd field to 0.
   The same holds if we round 1.5e-9 to 2e-9.
*/

// Handle carry.

/* We cannot have more than 9 * 9 = 81 digits. */

/* making 'zero' with the proper scale */

/* Here we check 999.9 -> 1000 case when we need to increase intDigCnt */

// FromInt sets the decimal value from int64.
func (d *MyDecimal) FromInt(val int64) *MyDecimal { _ = "STUB: not implemented"; return nil }

// FromUint sets the decimal value from uint64.
func (d *MyDecimal) FromUint(val uint64) *MyDecimal { _ = "STUB: not implemented"; return nil }

// ToInt returns int part of the decimal, returns the result and errcode.
func (d *MyDecimal) ToInt() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

/*
   Attention: trick!
   we're calculating -|from| instead of |from| here
   because |LONGLONG_MIN| > LONGLONG_MAX
   so we can convert -9223372036854775808 correctly
*/

/*
   the decimal is bigger than any possible integer
   return border integer depending on the sign
*/

/* boundary case: 9223372036854775808 */

// ToUint returns int part of the decimal, returns the result and errcode.
func (d *MyDecimal) ToUint() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// FromFloat64 creates a decimal from float64 value.
func (d *MyDecimal) FromFloat64(f float64) error { _ = "STUB: not implemented"; return nil }

// ToFloat64 converts decimal to float64 value.
func (d *MyDecimal) ToFloat64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

/*
ToBin converts decimal to its binary fixed-length representation
two representations of the same length can be compared with memcmp
with the correct -1/0/+1 result

	  PARAMS
			precision/frac - if precision is 0, internal value of the decimal will be used,
			then the encoded value is not memory comparable.

	  NOTE
	    the buffer is assumed to be of the size decimalBinSize(precision, frac)

	  RETURN VALUE
	  	bin     - binary value
	    errCode - eDecOK/eDecTruncate/eDecOverflow

	  DESCRIPTION
	    for storage decimal numbers are converted to the "binary" format.

	    This format has the following properties:
	      1. length of the binary representation depends on the {precision, frac}
	      as provided by the caller and NOT on the digitsInt/digitsFrac of the decimal to
	      convert.
	      2. binary representations of the same {precision, frac} can be compared
	      with memcmp - with the same result as DecimalCompare() of the original
	      decimals (not taking into account possible precision loss during
	      conversion).

	    This binary format is as follows:
	      1. First the number is converted to have a requested precision and frac.
	      2. Every full digitsPerWord digits of digitsInt part are stored in 4 bytes
	         as is
	      3. The first digitsInt % digitesPerWord digits are stored in the reduced
	         number of bytes (enough bytes to store this number of digits -
	         see dig2bytes)
	      4. same for frac - full word are stored as is,
	         the last frac % digitsPerWord digits - in the reduced number of bytes.
	      5. If the number is negative - every byte is inversed.
	      5. The very first bit of the resulting byte array is inverted (because
	         memcmp compares unsigned bytes, see property 2 above)

	    Example:

	      1234567890.1234

	    internally is represented as 3 words

	      1 234567890 123400000

	    (assuming we want a binary representation with precision=14, frac=4)
	    in hex it's

	      00-00-00-01  0D-FB-38-D2  07-5A-EF-40

	    now, middle word is full - it stores 9 decimal digits. It goes
	    into binary representation as is:


	      ...........  0D-FB-38-D2 ............

	    First word has only one decimal digit. We can store one digit in
	    one byte, no need to waste four:

	                01 0D-FB-38-D2 ............

	    now, last word. It's 123400000. We can store 1234 in two bytes:

	                01 0D-FB-38-D2 04-D2

	    So, we've packed 12 bytes number in 7 bytes.
	    And now we invert the highest bit to get the final result:

	                81 0D FB 38 D2 04 D2

	    And for -1234567890.1234 it would be

	                7E F2 04 C7 2D FB 2D
*/
func (d *MyDecimal) ToBin(precision, frac int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// xIntFrom part

// wordsInt + wordsFrac part.

// xFracFrom part

// PrecisionAndFrac returns the internal precision and frac number.
func (d *MyDecimal) PrecisionAndFrac() (precision, frac int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// FromBin Restores decimal from its binary fixed-length representation.
func (d *MyDecimal) FromBin(bin []byte, precision, frac int) (binSize int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// decimalBinSize returns the size of array to hold a binary representation of a decimal.
func decimalBinSize(precision, frac int) int { _ = "STUB: not implemented"; return 0 }

func readWord(b []byte, size int) int32 { _ = "STUB: not implemented"; return 0 }

func writeWord(b []byte, word int32, size int) { _ = "STUB: not implemented"; return }

// Compare compares one decimal to another, returns -1/0/1.
func (d *MyDecimal) Compare(to *MyDecimal) int { _ = "STUB: not implemented"; return 0 }

// DecimalAdd adds two decimals, sets the result to 'to'.
func DecimalAdd(from1, from2, to *MyDecimal) error { _ = "STUB: not implemented"; return nil }

// DecimalSub subs one decimal from another, sets the result to 'to'.
func DecimalSub(from1, from2, to *MyDecimal) error { _ = "STUB: not implemented"; return nil }

func doSub(from1, from2, to *MyDecimal) (cmp int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// from2 is negative too.

/* ensure that always idx1 > idx2 (and wordsInt1 >= wordsInt2) */

/* part 1 - max(frac) ... min (frac) */

/* part 2 - min(frac) ... wordsInt2 */

/* part 3 - wordsInt2 ... wordsInt1 */

func doAdd(from1, from2, to *MyDecimal) error { _ = "STUB: not implemented"; return nil }

/* yes, there is */

/* safety */

/* part 1 - max(frac) ... min (frac) */

/* part 2 - min(frac) ... min(digitsInt) */

/* part 3 - min(digitsInt) ... max(digitsInt) */

func maxDecimal(precision, frac int, to *MyDecimal) { _ = "STUB: not implemented"; return }

/* get 9 99 999 ... */

/*
DecimalMul multiplies two decimals.

	    from1, from2 - factors
	    to      - product

	RETURN VALUE
	  E_DEC_OK/E_DEC_TRUNCATED/E_DEC_OVERFLOW;

	NOTES
	  in this implementation, with wordSize=4 we have digitsPerWord=9,
	  and 63-digit number will take only 7 words (basically a 7-digit
	  "base 999999999" number).  Thus there's no need in fast multiplication
	  algorithms, 7-digit numbers can be multiplied with a naive O(n*n)
	  method.

	  XXX if this library is to be used with huge numbers of thousands of
	  digits, fast multiplication must be implemented.
*/
func DecimalMul(from1, from2, to *MyDecimal) error { _ = "STUB: not implemented"; return nil }

/* Now we have to check for -0.000 case */

/* We got decimal zero */

// DecimalDiv does division of two decimals.
//
// from1    - dividend
// from2    - divisor
// to       - quotient
// fracIncr - increment of fraction
func DecimalDiv(from1, from2, to *MyDecimal, fracIncr int) error {
	_ = "STUB: not implemented"
	return nil
}

/*
DecimalMod does modulus of two decimals.

	    from1   - dividend
	    from2   - divisor
	    to      - modulus

	RETURN VALUE
	  E_DEC_OK/E_DEC_TRUNCATED/E_DEC_OVERFLOW/E_DEC_DIV_ZERO;

	NOTES
	  see do_div_mod()

	DESCRIPTION
	  the modulus R in    R = M mod N

	 is defined as

	   0 <= |R| < |M|
	   sign R == sign M
	   R = M - k*N, where k is integer

	 thus, there's no requirement for M or N to be integers
*/
func DecimalMod(from1, from2, to *MyDecimal) error { _ = "STUB: not implemented"; return nil }

func doDivMod(from1, from2, to, mod *MyDecimal, fracIncr int) error {
	_ = "STUB: not implemented"
	return nil
}

/* removing all the leading zeros */

/* short-circuit everything: from2 == 0 */

/* short-circuit everything: from1 == 0 */

/* let's fix fracIncr, taking into account frac1,frac2 increase */

// we're calculating N1 % N2.
// The result will have
// digitsFrac=max(frac1, frac2), as for subtraction
// digitsInt=from2.digitsInt

/* removing end zeroes */

/*
   calculating norm2 (normalized from2.wordBuf[start2]) - we need from2.wordBuf[start2] to be large
   (at least > DIG_BASE/2), but unlike Knuth's Alg. D we don't want to
   normalize input numbers (as we don't make a copy of the divisor).
   Thus we normalize first dec1 of buf2 only, and we'll normalize tmp1[start1]
   on the fly for the purpose of guesstimation only.
   It's also faster, as we're saving on normalization of from2.
*/

// main loop

/* short-circuit, if possible */

/* D3: make a guess */

/* remove normalization */

/* D4: multiply and subtract */

/* D5: check the remainder */

/* D6: correct the guess */

/*
   now the result is in tmp1, it has
   digitsInt=prec1-frac1
   digitsFrac=max(frac1, frac2)
*/

// DecimalPeak returns the length of the encoded decimal.
func DecimalPeak(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// NewDecFromInt creates a MyDecimal from int.
func NewDecFromInt(i int64) *MyDecimal { _ = "STUB: not implemented"; return nil }

// NewDecFromFloatForTest creates a MyDecimal from float, as it returns no error, it should only be used in test.
func NewDecFromFloatForTest(f float64) *MyDecimal { _ = "STUB: not implemented"; return nil }

// NewDecFromStringForTest creates a MyDecimal from string, as it returns no error, it should only be used in test.
func NewDecFromStringForTest(s string) *MyDecimal { _ = "STUB: not implemented"; return nil }

// NewMaxOrMinDec returns the max or min value decimal for given precision and fraction.
func NewMaxOrMinDec(negative bool, prec, frac int) *MyDecimal {
	_ = "STUB: not implemented"
	return nil
}
