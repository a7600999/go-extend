// Copyright (C) 2019  Qi Yin <qiyin@thinkeridea.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package benchmark

import (
	"testing"

	"github.com/thinkeridea/go-extend/exstrings"
)

func BenchmarkStandardLibraryStringToBytes(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		p := []byte(s)
		_ = p

	}
}

func BenchmarkExstringsStringToBytes(b *testing.B) {
	s := "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
	for i := 0; i < b.N; i++ {
		p := exstrings.Bytes(s)
		_ = p
	}
}
