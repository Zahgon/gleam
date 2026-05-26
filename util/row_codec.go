package util

// NOTE: THIS FILE WAS PRODUCED BY THE
// TRUEPACK CODE GENERATION TOOL (github.com/glycerine/truepack)
// DO NOT EDIT

import (
	"github.com/glycerine/truepack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Row) DecodeMsg(dc *msgp.Reader) (err error) { _ = "STUB: not implemented"; return nil }

// -- templateDecodeMsg starts here--

// First fill all the encoded fields, then
// treat the remaining, missing fields, as Nil.

//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zgensym_56fd93edf47ccbd5_3, missingFieldsLeft2zgensym_56fd93edf47ccbd5_3, msgp.ShowFound(found2zgensym_56fd93edf47ccbd5_3[:]), decodeMsgFieldOrder2zgensym_56fd93edf47ccbd5_3)

//missing fields need handling

// tell the reader to only give us Nils
// until further notice.

// filled all the empty fields!

//fmt.Printf("switching on curField: '%v'\n", curField2zgensym_56fd93edf47ccbd5_3)

// -- templateDecodeMsg ends here --

// fields of Row
var decodeMsgFieldOrder2zgensym_56fd93edf47ccbd5_3 = []string{"K__slc", "V__slc", "T__i64"}

var decodeMsgFieldSkip2zgensym_56fd93edf47ccbd5_3 = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Row) fieldsNotEmpty(isempty []bool) uint32 { _ = "STUB: not implemented"; return 0 }

// string, omitempty

// string, omitempty

// number, omitempty

// EncodeMsg implements msgp.Encodable
func (z *Row) EncodeMsg(en *msgp.Writer) (err error) { _ = "STUB: not implemented"; return nil }

// honor the omitempty tags

// map header

// write "K__slc"

// write "V__slc"

// write "T__i64"

// MarshalMsg implements msgp.Marshaler
func (z *Row) MarshalMsg(b []byte) (o []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// honor the omitempty tags

// string "K__slc"

// string "V__slc"

// string "T__i64"

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Row) UnmarshalMsg(bts []byte) (o []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *Row) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// -- templateUnmarshalMsg starts here--

// First fill all the encoded fields, then
// treat the remaining, missing fields, as Nil.

//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zgensym_56fd93edf47ccbd5_9, missingFieldsLeft8zgensym_56fd93edf47ccbd5_9, msgp.ShowFound(found8zgensym_56fd93edf47ccbd5_9[:]), unmarshalMsgFieldOrder8zgensym_56fd93edf47ccbd5_9)

//missing fields need handling

// set bts to contain just mnil (0xc0)

// filled all the empty fields!

//fmt.Printf("switching on curField: '%v'\n", curField8zgensym_56fd93edf47ccbd5_9)

// -- templateUnmarshalMsg ends here --

// fields of Row
var unmarshalMsgFieldOrder8zgensym_56fd93edf47ccbd5_9 = []string{"K__slc", "V__slc", "T__i64"}

var unmarshalMsgFieldSkip8zgensym_56fd93edf47ccbd5_9 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Row) Msgsize() (s int) { _ = "STUB: not implemented"; return 0 }
