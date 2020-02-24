// Copyright 2019 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package generated implements field arithmetics with field modulus q =
//
// 322682568289525361762046250237368232446223696123443138893952234340997414728466624440898096951488632464052441876554825788797264920515711
//
// Code generated by goff DO NOT EDIT
// Element07 are assumed to be in Montgomery form in all methods
package generated

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
	"math/bits"
	"sync"

	"unsafe"
)

// Element07 represents a field element stored on 7 words (uint64)
// Element07 are assumed to be in Montgomery form in all methods
type Element07 [7]uint64

// Element07Limbs number of 64 bits words needed to represent Element07
const Element07Limbs = 7

// Element07Bits number bits needed to represent Element07
const Element07Bits = 447

// SetUint64 z = v, sets z LSB to v (non-Montgomery form) and convert z to Montgomery form
func (z *Element07) SetUint64(v uint64) *Element07 {
	z[0] = v
	z[1] = 0
	z[2] = 0
	z[3] = 0
	z[4] = 0
	z[5] = 0
	z[6] = 0
	return z.ToMont()
}

// Set z = x
func (z *Element07) Set(x *Element07) *Element07 {
	z[0] = x[0]
	z[1] = x[1]
	z[2] = x[2]
	z[3] = x[3]
	z[4] = x[4]
	z[5] = x[5]
	z[6] = x[6]
	return z
}

// SetZero z = 0
func (z *Element07) SetZero() *Element07 {
	z[0] = 0
	z[1] = 0
	z[2] = 0
	z[3] = 0
	z[4] = 0
	z[5] = 0
	z[6] = 0
	return z
}

// SetOne z = 1 (in Montgomery form)
func (z *Element07) SetOne() *Element07 {
	z[0] = 4808574769336276738
	z[1] = 603692077527013820
	z[2] = 17058623858472768205
	z[3] = 6464587670255442682
	z[4] = 1334920108916464395
	z[5] = 6230098516080116469
	z[6] = 2067752269020542685
	return z
}

// Neg z = q - x
func (z *Element07) Neg(x *Element07) *Element07 {
	if x.IsZero() {
		return z.SetZero()
	}
	var borrow uint64
	z[0], borrow = bits.Sub64(16042456689041413247, x[0], 0)
	z[1], borrow = bits.Sub64(8921525998091268897, x[1], borrow)
	z[2], borrow = bits.Sub64(9917432144473167513, x[2], borrow)
	z[3], borrow = bits.Sub64(5991078201727054466, x[3], borrow)
	z[4], borrow = bits.Sub64(8555911982396543610, x[4], borrow)
	z[5], borrow = bits.Sub64(6108322778814717573, x[5], borrow)
	z[6], _ = bits.Sub64(8189495902344504465, x[6], borrow)
	return z
}

// Div z = x*y^-1 mod q
func (z *Element07) Div(x, y *Element07) *Element07 {
	var yInv Element07
	yInv.Inverse(y)
	z.Mul(x, &yInv)
	return z
}

// Equal returns z == x
func (z *Element07) Equal(x *Element07) bool {
	return (z[6] == x[6]) && (z[5] == x[5]) && (z[4] == x[4]) && (z[3] == x[3]) && (z[2] == x[2]) && (z[1] == x[1]) && (z[0] == x[0])
}

// IsZero returns z == 0
func (z *Element07) IsZero() bool {
	return (z[6] | z[5] | z[4] | z[3] | z[2] | z[1] | z[0]) == 0
}

// field modulus stored as big.Int
var _element07ModulusBigInt big.Int
var onceelement07Modulus sync.Once

func element07ModulusBigInt() *big.Int {
	onceelement07Modulus.Do(func() {
		_element07ModulusBigInt.SetString("322682568289525361762046250237368232446223696123443138893952234340997414728466624440898096951488632464052441876554825788797264920515711", 10)
	})
	return &_element07ModulusBigInt
}

// Inverse z = x^-1 mod q
// Algorithm 16 in "Efficient Software-Implementation of Finite Fields with Applications to Cryptography"
// if x == 0, sets and returns z = x
func (z *Element07) Inverse(x *Element07) *Element07 {
	if x.IsZero() {
		return z.Set(x)
	}

	// initialize u = q
	var u = Element07{
		16042456689041413247,
		8921525998091268897,
		9917432144473167513,
		5991078201727054466,
		8555911982396543610,
		6108322778814717573,
		8189495902344504465,
	}

	// initialize s = r^2
	var s = Element07{
		11394061373368433203,
		6151568944091229183,
		13234826941241376623,
		15235843581831928969,
		9210085638756949795,
		3757868771183409087,
		7192684189326524933,
	}

	// r = 0
	r := Element07{}

	v := *x

	var carry, borrow, t, t2 uint64
	var bigger, uIsOne, vIsOne bool

	for !uIsOne && !vIsOne {
		for v[0]&1 == 0 {

			// v = v >> 1
			t2 = v[6] << 63
			v[6] >>= 1
			t = t2
			t2 = v[5] << 63
			v[5] = (v[5] >> 1) | t
			t = t2
			t2 = v[4] << 63
			v[4] = (v[4] >> 1) | t
			t = t2
			t2 = v[3] << 63
			v[3] = (v[3] >> 1) | t
			t = t2
			t2 = v[2] << 63
			v[2] = (v[2] >> 1) | t
			t = t2
			t2 = v[1] << 63
			v[1] = (v[1] >> 1) | t
			t = t2
			v[0] = (v[0] >> 1) | t

			if s[0]&1 == 1 {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 16042456689041413247, 0)
				s[1], carry = bits.Add64(s[1], 8921525998091268897, carry)
				s[2], carry = bits.Add64(s[2], 9917432144473167513, carry)
				s[3], carry = bits.Add64(s[3], 5991078201727054466, carry)
				s[4], carry = bits.Add64(s[4], 8555911982396543610, carry)
				s[5], carry = bits.Add64(s[5], 6108322778814717573, carry)
				s[6], _ = bits.Add64(s[6], 8189495902344504465, carry)

			}

			// s = s >> 1
			t2 = s[6] << 63
			s[6] >>= 1
			t = t2
			t2 = s[5] << 63
			s[5] = (s[5] >> 1) | t
			t = t2
			t2 = s[4] << 63
			s[4] = (s[4] >> 1) | t
			t = t2
			t2 = s[3] << 63
			s[3] = (s[3] >> 1) | t
			t = t2
			t2 = s[2] << 63
			s[2] = (s[2] >> 1) | t
			t = t2
			t2 = s[1] << 63
			s[1] = (s[1] >> 1) | t
			t = t2
			s[0] = (s[0] >> 1) | t

		}
		for u[0]&1 == 0 {

			// u = u >> 1
			t2 = u[6] << 63
			u[6] >>= 1
			t = t2
			t2 = u[5] << 63
			u[5] = (u[5] >> 1) | t
			t = t2
			t2 = u[4] << 63
			u[4] = (u[4] >> 1) | t
			t = t2
			t2 = u[3] << 63
			u[3] = (u[3] >> 1) | t
			t = t2
			t2 = u[2] << 63
			u[2] = (u[2] >> 1) | t
			t = t2
			t2 = u[1] << 63
			u[1] = (u[1] >> 1) | t
			t = t2
			u[0] = (u[0] >> 1) | t

			if r[0]&1 == 1 {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 16042456689041413247, 0)
				r[1], carry = bits.Add64(r[1], 8921525998091268897, carry)
				r[2], carry = bits.Add64(r[2], 9917432144473167513, carry)
				r[3], carry = bits.Add64(r[3], 5991078201727054466, carry)
				r[4], carry = bits.Add64(r[4], 8555911982396543610, carry)
				r[5], carry = bits.Add64(r[5], 6108322778814717573, carry)
				r[6], _ = bits.Add64(r[6], 8189495902344504465, carry)

			}

			// r = r >> 1
			t2 = r[6] << 63
			r[6] >>= 1
			t = t2
			t2 = r[5] << 63
			r[5] = (r[5] >> 1) | t
			t = t2
			t2 = r[4] << 63
			r[4] = (r[4] >> 1) | t
			t = t2
			t2 = r[3] << 63
			r[3] = (r[3] >> 1) | t
			t = t2
			t2 = r[2] << 63
			r[2] = (r[2] >> 1) | t
			t = t2
			t2 = r[1] << 63
			r[1] = (r[1] >> 1) | t
			t = t2
			r[0] = (r[0] >> 1) | t

		}

		// v >= u
		bigger = !(v[6] < u[6] || (v[6] == u[6] && (v[5] < u[5] || (v[5] == u[5] && (v[4] < u[4] || (v[4] == u[4] && (v[3] < u[3] || (v[3] == u[3] && (v[2] < u[2] || (v[2] == u[2] && (v[1] < u[1] || (v[1] == u[1] && (v[0] < u[0])))))))))))))

		if bigger {

			// v = v - u
			v[0], borrow = bits.Sub64(v[0], u[0], 0)
			v[1], borrow = bits.Sub64(v[1], u[1], borrow)
			v[2], borrow = bits.Sub64(v[2], u[2], borrow)
			v[3], borrow = bits.Sub64(v[3], u[3], borrow)
			v[4], borrow = bits.Sub64(v[4], u[4], borrow)
			v[5], borrow = bits.Sub64(v[5], u[5], borrow)
			v[6], _ = bits.Sub64(v[6], u[6], borrow)

			// r >= s
			bigger = !(r[6] < s[6] || (r[6] == s[6] && (r[5] < s[5] || (r[5] == s[5] && (r[4] < s[4] || (r[4] == s[4] && (r[3] < s[3] || (r[3] == s[3] && (r[2] < s[2] || (r[2] == s[2] && (r[1] < s[1] || (r[1] == s[1] && (r[0] < s[0])))))))))))))

			if bigger {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 16042456689041413247, 0)
				s[1], carry = bits.Add64(s[1], 8921525998091268897, carry)
				s[2], carry = bits.Add64(s[2], 9917432144473167513, carry)
				s[3], carry = bits.Add64(s[3], 5991078201727054466, carry)
				s[4], carry = bits.Add64(s[4], 8555911982396543610, carry)
				s[5], carry = bits.Add64(s[5], 6108322778814717573, carry)
				s[6], _ = bits.Add64(s[6], 8189495902344504465, carry)

			}

			// s = s - r
			s[0], borrow = bits.Sub64(s[0], r[0], 0)
			s[1], borrow = bits.Sub64(s[1], r[1], borrow)
			s[2], borrow = bits.Sub64(s[2], r[2], borrow)
			s[3], borrow = bits.Sub64(s[3], r[3], borrow)
			s[4], borrow = bits.Sub64(s[4], r[4], borrow)
			s[5], borrow = bits.Sub64(s[5], r[5], borrow)
			s[6], _ = bits.Sub64(s[6], r[6], borrow)

		} else {

			// u = u - v
			u[0], borrow = bits.Sub64(u[0], v[0], 0)
			u[1], borrow = bits.Sub64(u[1], v[1], borrow)
			u[2], borrow = bits.Sub64(u[2], v[2], borrow)
			u[3], borrow = bits.Sub64(u[3], v[3], borrow)
			u[4], borrow = bits.Sub64(u[4], v[4], borrow)
			u[5], borrow = bits.Sub64(u[5], v[5], borrow)
			u[6], _ = bits.Sub64(u[6], v[6], borrow)

			// s >= r
			bigger = !(s[6] < r[6] || (s[6] == r[6] && (s[5] < r[5] || (s[5] == r[5] && (s[4] < r[4] || (s[4] == r[4] && (s[3] < r[3] || (s[3] == r[3] && (s[2] < r[2] || (s[2] == r[2] && (s[1] < r[1] || (s[1] == r[1] && (s[0] < r[0])))))))))))))

			if bigger {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 16042456689041413247, 0)
				r[1], carry = bits.Add64(r[1], 8921525998091268897, carry)
				r[2], carry = bits.Add64(r[2], 9917432144473167513, carry)
				r[3], carry = bits.Add64(r[3], 5991078201727054466, carry)
				r[4], carry = bits.Add64(r[4], 8555911982396543610, carry)
				r[5], carry = bits.Add64(r[5], 6108322778814717573, carry)
				r[6], _ = bits.Add64(r[6], 8189495902344504465, carry)

			}

			// r = r - s
			r[0], borrow = bits.Sub64(r[0], s[0], 0)
			r[1], borrow = bits.Sub64(r[1], s[1], borrow)
			r[2], borrow = bits.Sub64(r[2], s[2], borrow)
			r[3], borrow = bits.Sub64(r[3], s[3], borrow)
			r[4], borrow = bits.Sub64(r[4], s[4], borrow)
			r[5], borrow = bits.Sub64(r[5], s[5], borrow)
			r[6], _ = bits.Sub64(r[6], s[6], borrow)

		}
		uIsOne = (u[0] == 1) && (u[6]|u[5]|u[4]|u[3]|u[2]|u[1]) == 0
		vIsOne = (v[0] == 1) && (v[6]|v[5]|v[4]|v[3]|v[2]|v[1]) == 0
	}

	if uIsOne {
		z.Set(&r)
	} else {
		z.Set(&s)
	}

	return z
}

// SetRandom sets z to a random element < q
func (z *Element07) SetRandom() *Element07 {
	bytes := make([]byte, 448)
	io.ReadFull(rand.Reader, bytes)
	z[0] = binary.BigEndian.Uint64(bytes[0:64])
	z[1] = binary.BigEndian.Uint64(bytes[64:128])
	z[2] = binary.BigEndian.Uint64(bytes[128:192])
	z[3] = binary.BigEndian.Uint64(bytes[192:256])
	z[4] = binary.BigEndian.Uint64(bytes[256:320])
	z[5] = binary.BigEndian.Uint64(bytes[320:384])
	z[6] = binary.BigEndian.Uint64(bytes[384:448])
	z[6] %= 8189495902344504465

	// if z > q --> z -= q
	if !(z[6] < 8189495902344504465 || (z[6] == 8189495902344504465 && (z[5] < 6108322778814717573 || (z[5] == 6108322778814717573 && (z[4] < 8555911982396543610 || (z[4] == 8555911982396543610 && (z[3] < 5991078201727054466 || (z[3] == 5991078201727054466 && (z[2] < 9917432144473167513 || (z[2] == 9917432144473167513 && (z[1] < 8921525998091268897 || (z[1] == 8921525998091268897 && (z[0] < 16042456689041413247))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 16042456689041413247, 0)
		z[1], b = bits.Sub64(z[1], 8921525998091268897, b)
		z[2], b = bits.Sub64(z[2], 9917432144473167513, b)
		z[3], b = bits.Sub64(z[3], 5991078201727054466, b)
		z[4], b = bits.Sub64(z[4], 8555911982396543610, b)
		z[5], b = bits.Sub64(z[5], 6108322778814717573, b)
		z[6], _ = bits.Sub64(z[6], 8189495902344504465, b)
	}

	return z
}

// Add z = x + y mod q
func (z *Element07) Add(x, y *Element07) *Element07 {
	var carry uint64

	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], carry = bits.Add64(x[3], y[3], carry)
	z[4], carry = bits.Add64(x[4], y[4], carry)
	z[5], carry = bits.Add64(x[5], y[5], carry)
	z[6], _ = bits.Add64(x[6], y[6], carry)

	// if z > q --> z -= q
	if !(z[6] < 8189495902344504465 || (z[6] == 8189495902344504465 && (z[5] < 6108322778814717573 || (z[5] == 6108322778814717573 && (z[4] < 8555911982396543610 || (z[4] == 8555911982396543610 && (z[3] < 5991078201727054466 || (z[3] == 5991078201727054466 && (z[2] < 9917432144473167513 || (z[2] == 9917432144473167513 && (z[1] < 8921525998091268897 || (z[1] == 8921525998091268897 && (z[0] < 16042456689041413247))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 16042456689041413247, 0)
		z[1], b = bits.Sub64(z[1], 8921525998091268897, b)
		z[2], b = bits.Sub64(z[2], 9917432144473167513, b)
		z[3], b = bits.Sub64(z[3], 5991078201727054466, b)
		z[4], b = bits.Sub64(z[4], 8555911982396543610, b)
		z[5], b = bits.Sub64(z[5], 6108322778814717573, b)
		z[6], _ = bits.Sub64(z[6], 8189495902344504465, b)
	}
	return z
}

// AddAssign z = z + x mod q
func (z *Element07) AddAssign(x *Element07) *Element07 {
	var carry uint64

	z[0], carry = bits.Add64(z[0], x[0], 0)
	z[1], carry = bits.Add64(z[1], x[1], carry)
	z[2], carry = bits.Add64(z[2], x[2], carry)
	z[3], carry = bits.Add64(z[3], x[3], carry)
	z[4], carry = bits.Add64(z[4], x[4], carry)
	z[5], carry = bits.Add64(z[5], x[5], carry)
	z[6], _ = bits.Add64(z[6], x[6], carry)

	// if z > q --> z -= q
	if !(z[6] < 8189495902344504465 || (z[6] == 8189495902344504465 && (z[5] < 6108322778814717573 || (z[5] == 6108322778814717573 && (z[4] < 8555911982396543610 || (z[4] == 8555911982396543610 && (z[3] < 5991078201727054466 || (z[3] == 5991078201727054466 && (z[2] < 9917432144473167513 || (z[2] == 9917432144473167513 && (z[1] < 8921525998091268897 || (z[1] == 8921525998091268897 && (z[0] < 16042456689041413247))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 16042456689041413247, 0)
		z[1], b = bits.Sub64(z[1], 8921525998091268897, b)
		z[2], b = bits.Sub64(z[2], 9917432144473167513, b)
		z[3], b = bits.Sub64(z[3], 5991078201727054466, b)
		z[4], b = bits.Sub64(z[4], 8555911982396543610, b)
		z[5], b = bits.Sub64(z[5], 6108322778814717573, b)
		z[6], _ = bits.Sub64(z[6], 8189495902344504465, b)
	}
	return z
}

// Double z = x + x mod q, aka Lsh 1
func (z *Element07) Double(x *Element07) *Element07 {
	var carry uint64

	z[0], carry = bits.Add64(x[0], x[0], 0)
	z[1], carry = bits.Add64(x[1], x[1], carry)
	z[2], carry = bits.Add64(x[2], x[2], carry)
	z[3], carry = bits.Add64(x[3], x[3], carry)
	z[4], carry = bits.Add64(x[4], x[4], carry)
	z[5], carry = bits.Add64(x[5], x[5], carry)
	z[6], _ = bits.Add64(x[6], x[6], carry)

	// if z > q --> z -= q
	if !(z[6] < 8189495902344504465 || (z[6] == 8189495902344504465 && (z[5] < 6108322778814717573 || (z[5] == 6108322778814717573 && (z[4] < 8555911982396543610 || (z[4] == 8555911982396543610 && (z[3] < 5991078201727054466 || (z[3] == 5991078201727054466 && (z[2] < 9917432144473167513 || (z[2] == 9917432144473167513 && (z[1] < 8921525998091268897 || (z[1] == 8921525998091268897 && (z[0] < 16042456689041413247))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 16042456689041413247, 0)
		z[1], b = bits.Sub64(z[1], 8921525998091268897, b)
		z[2], b = bits.Sub64(z[2], 9917432144473167513, b)
		z[3], b = bits.Sub64(z[3], 5991078201727054466, b)
		z[4], b = bits.Sub64(z[4], 8555911982396543610, b)
		z[5], b = bits.Sub64(z[5], 6108322778814717573, b)
		z[6], _ = bits.Sub64(z[6], 8189495902344504465, b)
	}
	return z
}

// Sub  z = x - y mod q
func (z *Element07) Sub(x, y *Element07) *Element07 {
	var b uint64
	z[0], b = bits.Sub64(x[0], y[0], 0)
	z[1], b = bits.Sub64(x[1], y[1], b)
	z[2], b = bits.Sub64(x[2], y[2], b)
	z[3], b = bits.Sub64(x[3], y[3], b)
	z[4], b = bits.Sub64(x[4], y[4], b)
	z[5], b = bits.Sub64(x[5], y[5], b)
	z[6], b = bits.Sub64(x[6], y[6], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 16042456689041413247, 0)
		z[1], c = bits.Add64(z[1], 8921525998091268897, c)
		z[2], c = bits.Add64(z[2], 9917432144473167513, c)
		z[3], c = bits.Add64(z[3], 5991078201727054466, c)
		z[4], c = bits.Add64(z[4], 8555911982396543610, c)
		z[5], c = bits.Add64(z[5], 6108322778814717573, c)
		z[6], _ = bits.Add64(z[6], 8189495902344504465, c)
	}
	return z
}

// SubAssign  z = z - x mod q
func (z *Element07) SubAssign(x *Element07) *Element07 {
	var b uint64
	z[0], b = bits.Sub64(z[0], x[0], 0)
	z[1], b = bits.Sub64(z[1], x[1], b)
	z[2], b = bits.Sub64(z[2], x[2], b)
	z[3], b = bits.Sub64(z[3], x[3], b)
	z[4], b = bits.Sub64(z[4], x[4], b)
	z[5], b = bits.Sub64(z[5], x[5], b)
	z[6], b = bits.Sub64(z[6], x[6], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 16042456689041413247, 0)
		z[1], c = bits.Add64(z[1], 8921525998091268897, c)
		z[2], c = bits.Add64(z[2], 9917432144473167513, c)
		z[3], c = bits.Add64(z[3], 5991078201727054466, c)
		z[4], c = bits.Add64(z[4], 8555911982396543610, c)
		z[5], c = bits.Add64(z[5], 6108322778814717573, c)
		z[6], _ = bits.Add64(z[6], 8189495902344504465, c)
	}
	return z
}

// Exp z = x^e mod q
func (z *Element07) Exp(x Element07, e uint64) *Element07 {
	if e == 0 {
		return z.SetOne()
	}

	z.Set(&x)

	l := bits.Len64(e) - 2
	for i := l; i >= 0; i-- {
		z.Square(z)
		if e&(1<<uint(i)) != 0 {
			z.MulAssign(&x)
		}
	}
	return z
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func (z *Element07) FromMont() *Element07 {

	// the following lines implement z = z * 1
	// with a modified CIOS montgomery multiplication
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 1339912018327196801
		C := madd0(m, 16042456689041413247, z[0])
		C, z[0] = madd2(m, 8921525998091268897, z[1], C)
		C, z[1] = madd2(m, 9917432144473167513, z[2], C)
		C, z[2] = madd2(m, 5991078201727054466, z[3], C)
		C, z[3] = madd2(m, 8555911982396543610, z[4], C)
		C, z[4] = madd2(m, 6108322778814717573, z[5], C)
		C, z[5] = madd2(m, 8189495902344504465, z[6], C)
		z[6] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 1339912018327196801
		C := madd0(m, 16042456689041413247, z[0])
		C, z[0] = madd2(m, 8921525998091268897, z[1], C)
		C, z[1] = madd2(m, 9917432144473167513, z[2], C)
		C, z[2] = madd2(m, 5991078201727054466, z[3], C)
		C, z[3] = madd2(m, 8555911982396543610, z[4], C)
		C, z[4] = madd2(m, 6108322778814717573, z[5], C)
		C, z[5] = madd2(m, 8189495902344504465, z[6], C)
		z[6] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 1339912018327196801
		C := madd0(m, 16042456689041413247, z[0])
		C, z[0] = madd2(m, 8921525998091268897, z[1], C)
		C, z[1] = madd2(m, 9917432144473167513, z[2], C)
		C, z[2] = madd2(m, 5991078201727054466, z[3], C)
		C, z[3] = madd2(m, 8555911982396543610, z[4], C)
		C, z[4] = madd2(m, 6108322778814717573, z[5], C)
		C, z[5] = madd2(m, 8189495902344504465, z[6], C)
		z[6] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 1339912018327196801
		C := madd0(m, 16042456689041413247, z[0])
		C, z[0] = madd2(m, 8921525998091268897, z[1], C)
		C, z[1] = madd2(m, 9917432144473167513, z[2], C)
		C, z[2] = madd2(m, 5991078201727054466, z[3], C)
		C, z[3] = madd2(m, 8555911982396543610, z[4], C)
		C, z[4] = madd2(m, 6108322778814717573, z[5], C)
		C, z[5] = madd2(m, 8189495902344504465, z[6], C)
		z[6] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 1339912018327196801
		C := madd0(m, 16042456689041413247, z[0])
		C, z[0] = madd2(m, 8921525998091268897, z[1], C)
		C, z[1] = madd2(m, 9917432144473167513, z[2], C)
		C, z[2] = madd2(m, 5991078201727054466, z[3], C)
		C, z[3] = madd2(m, 8555911982396543610, z[4], C)
		C, z[4] = madd2(m, 6108322778814717573, z[5], C)
		C, z[5] = madd2(m, 8189495902344504465, z[6], C)
		z[6] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 1339912018327196801
		C := madd0(m, 16042456689041413247, z[0])
		C, z[0] = madd2(m, 8921525998091268897, z[1], C)
		C, z[1] = madd2(m, 9917432144473167513, z[2], C)
		C, z[2] = madd2(m, 5991078201727054466, z[3], C)
		C, z[3] = madd2(m, 8555911982396543610, z[4], C)
		C, z[4] = madd2(m, 6108322778814717573, z[5], C)
		C, z[5] = madd2(m, 8189495902344504465, z[6], C)
		z[6] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 1339912018327196801
		C := madd0(m, 16042456689041413247, z[0])
		C, z[0] = madd2(m, 8921525998091268897, z[1], C)
		C, z[1] = madd2(m, 9917432144473167513, z[2], C)
		C, z[2] = madd2(m, 5991078201727054466, z[3], C)
		C, z[3] = madd2(m, 8555911982396543610, z[4], C)
		C, z[4] = madd2(m, 6108322778814717573, z[5], C)
		C, z[5] = madd2(m, 8189495902344504465, z[6], C)
		z[6] = C
	}

	// if z > q --> z -= q
	if !(z[6] < 8189495902344504465 || (z[6] == 8189495902344504465 && (z[5] < 6108322778814717573 || (z[5] == 6108322778814717573 && (z[4] < 8555911982396543610 || (z[4] == 8555911982396543610 && (z[3] < 5991078201727054466 || (z[3] == 5991078201727054466 && (z[2] < 9917432144473167513 || (z[2] == 9917432144473167513 && (z[1] < 8921525998091268897 || (z[1] == 8921525998091268897 && (z[0] < 16042456689041413247))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 16042456689041413247, 0)
		z[1], b = bits.Sub64(z[1], 8921525998091268897, b)
		z[2], b = bits.Sub64(z[2], 9917432144473167513, b)
		z[3], b = bits.Sub64(z[3], 5991078201727054466, b)
		z[4], b = bits.Sub64(z[4], 8555911982396543610, b)
		z[5], b = bits.Sub64(z[5], 6108322778814717573, b)
		z[6], _ = bits.Sub64(z[6], 8189495902344504465, b)
	}
	return z
}

// ToMont converts z to Montgomery form
// sets and returns z = z * r^2
func (z *Element07) ToMont() *Element07 {
	var rSquare = Element07{
		11394061373368433203,
		6151568944091229183,
		13234826941241376623,
		15235843581831928969,
		9210085638756949795,
		3757868771183409087,
		7192684189326524933,
	}
	return z.MulAssign(&rSquare)
}

// ToRegular returns z in regular form (doesn't mutate z)
func (z Element07) ToRegular() Element07 {
	return *z.FromMont()
}

// String returns the string form of an Element07 in Montgomery form
func (z *Element07) String() string {
	var _z big.Int
	return z.ToBigIntRegular(&_z).String()
}

// ToBigInt returns z as a big.Int in Montgomery form
func (z *Element07) ToBigInt(res *big.Int) *big.Int {
	bits := (*[7]big.Word)(unsafe.Pointer(z))
	return res.SetBits(bits[:])
}

// ToBigIntRegular returns z as a big.Int in regular form
func (z Element07) ToBigIntRegular(res *big.Int) *big.Int {
	z.FromMont()
	bits := (*[7]big.Word)(unsafe.Pointer(&z))
	return res.SetBits(bits[:])
}

// SetBigInt sets z to v (regular form) and returns z in Montgomery form
func (z *Element07) SetBigInt(v *big.Int) *Element07 {
	z.SetZero()

	zero := big.NewInt(0)
	q := element07ModulusBigInt()

	// copy input
	vv := new(big.Int).Set(v)

	// while v < 0, v+=q
	for vv.Cmp(zero) == -1 {
		vv.Add(vv, q)
	}
	// while v > q, v-=q
	for vv.Cmp(q) == 1 {
		vv.Sub(vv, q)
	}
	// if v == q, return 0
	if vv.Cmp(q) == 0 {
		return z
	}
	// v should
	vBits := vv.Bits()
	for i := 0; i < len(vBits); i++ {
		z[i] = uint64(vBits[i])
	}
	return z.ToMont()
}

// SetString creates a big.Int with s (in base 10) and calls SetBigInt on z
func (z *Element07) SetString(s string) *Element07 {
	x, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("Element07.SetString failed -> can't parse number in base10 into a big.Int")
	}
	return z.SetBigInt(x)
}

// Mul z = x * y mod q
func (z *Element07) Mul(x, y *Element07) *Element07 {

	var t [7]uint64
	var c [3]uint64
	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, y[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd1(v, y[1], c[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd1(v, y[2], c[1])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd1(v, y[3], c[1])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd1(v, y[4], c[1])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd1(v, y[5], c[1])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd1(v, y[6], c[1])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, y[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, y[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, y[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 4
		v := x[4]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, y[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 5
		v := x[5]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, y[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 6
		v := x[6]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], z[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], z[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], z[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], z[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		c[2], z[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, y[6], c[1], t[6])
		z[6], z[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	if !(z[6] < 8189495902344504465 || (z[6] == 8189495902344504465 && (z[5] < 6108322778814717573 || (z[5] == 6108322778814717573 && (z[4] < 8555911982396543610 || (z[4] == 8555911982396543610 && (z[3] < 5991078201727054466 || (z[3] == 5991078201727054466 && (z[2] < 9917432144473167513 || (z[2] == 9917432144473167513 && (z[1] < 8921525998091268897 || (z[1] == 8921525998091268897 && (z[0] < 16042456689041413247))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 16042456689041413247, 0)
		z[1], b = bits.Sub64(z[1], 8921525998091268897, b)
		z[2], b = bits.Sub64(z[2], 9917432144473167513, b)
		z[3], b = bits.Sub64(z[3], 5991078201727054466, b)
		z[4], b = bits.Sub64(z[4], 8555911982396543610, b)
		z[5], b = bits.Sub64(z[5], 6108322778814717573, b)
		z[6], _ = bits.Sub64(z[6], 8189495902344504465, b)
	}
	return z
}

// MulAssign z = z * x mod q
func (z *Element07) MulAssign(x *Element07) *Element07 {

	var t [7]uint64
	var c [3]uint64
	{
		// round 0
		v := z[0]
		c[1], c[0] = bits.Mul64(v, x[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd1(v, x[1], c[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd1(v, x[2], c[1])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd1(v, x[3], c[1])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd1(v, x[4], c[1])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd1(v, x[5], c[1])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd1(v, x[6], c[1])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 1
		v := z[1]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 2
		v := z[2]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 3
		v := z[3]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 4
		v := z[4]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 5
		v := z[5]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], t[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		t[6], t[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}
	{
		// round 6
		v := z[6]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 1339912018327196801
		c[2] = madd0(m, 16042456689041413247, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], z[0] = madd2(m, 8921525998091268897, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], z[1] = madd2(m, 9917432144473167513, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], z[2] = madd2(m, 5991078201727054466, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], z[3] = madd2(m, 8555911982396543610, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		c[2], z[4] = madd2(m, 6108322778814717573, c[2], c[0])
		c[1], c[0] = madd2(v, x[6], c[1], t[6])
		z[6], z[5] = madd3(m, 8189495902344504465, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	if !(z[6] < 8189495902344504465 || (z[6] == 8189495902344504465 && (z[5] < 6108322778814717573 || (z[5] == 6108322778814717573 && (z[4] < 8555911982396543610 || (z[4] == 8555911982396543610 && (z[3] < 5991078201727054466 || (z[3] == 5991078201727054466 && (z[2] < 9917432144473167513 || (z[2] == 9917432144473167513 && (z[1] < 8921525998091268897 || (z[1] == 8921525998091268897 && (z[0] < 16042456689041413247))))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 16042456689041413247, 0)
		z[1], b = bits.Sub64(z[1], 8921525998091268897, b)
		z[2], b = bits.Sub64(z[2], 9917432144473167513, b)
		z[3], b = bits.Sub64(z[3], 5991078201727054466, b)
		z[4], b = bits.Sub64(z[4], 8555911982396543610, b)
		z[5], b = bits.Sub64(z[5], 6108322778814717573, b)
		z[6], _ = bits.Sub64(z[6], 8189495902344504465, b)
	}
	return z
}

// Square z = x * x mod q
func (z *Element07) Square(x *Element07) *Element07 {
	return z.Mul(x, x)
}