// Copyright 2015 PingCAP, Inc.
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

package util

// CalcPassword is the algorithm convert hashed password to auth string.
// See https://dev.mysql.com/doc/internals/en/secure-password-authentication.html
// SHA1( password ) XOR SHA1( "20-bytes random data from server" <concat> SHA1( SHA1( password ) ) )
func CalcPassword(scramble, sha1pwd []byte) []byte { _ = "STUB: not implemented"; return nil }

// scrambleHash = SHA1(scramble + SHA1(sha1pwd))
// inner Hash

// outer Hash

// token = scrambleHash XOR stage1Hash

// Sha1Hash is an util function to calculate sha1 hash.
func Sha1Hash(bs []byte) []byte { _ = "STUB: not implemented"; return nil }

// EncodePassword converts plaintext password to hashed hex string.
func EncodePassword(pwd string) string { _ = "STUB: not implemented"; return "" }

// DecodePassword converts hex string password to byte array.
func DecodePassword(pwd string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
