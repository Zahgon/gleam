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

package printer

// Version information.
var (
	TiDBBuildTS = "None"
	TiDBGitHash = "None"
)

// PrintTiDBInfo prints the TiDB version information.
func PrintTiDBInfo() { _ = "STUB: not implemented"; return }

// PrintRawTiDBInfo prints the TiDB version information without log info.
func PrintRawTiDBInfo() { _ = "STUB: not implemented"; return }

// checkValidity checks whether cols and every data have the same length.
func checkValidity(cols []string, datas [][]string) bool { _ = "STUB: not implemented"; return false }

func getMaxColLen(cols []string, datas [][]string) []int { _ = "STUB: not implemented"; return nil }

func getPrintDivLine(maxColLen []int) []byte { _ = "STUB: not implemented"; return nil }

func getPrintCol(cols []string, maxColLen []int) []byte { _ = "STUB: not implemented"; return nil }

func getPrintRow(data []string, maxColLen []int) []byte { _ = "STUB: not implemented"; return nil }

func getPrintRows(datas [][]string, maxColLen []int) []byte { _ = "STUB: not implemented"; return nil }

// GetPrintResult gets a result with a formatted string.
func GetPrintResult(cols []string, datas [][]string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
