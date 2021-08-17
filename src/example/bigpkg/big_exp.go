package bigpkg

import (
	"encoding/base64"
	"fmt"
	"math/big"
)

func NewFromInt64(i int64) {
	num := big.NewInt(i)
	fmt.Printf("Big int: %v\n", num)
}

func NewFromBytes(b []byte) {
	z := new(big.Int)
	z.SetBytes(b)
	fmt.Printf("Big int: %v\n", z)
}

func Base64ToInt(s string) (*big.Int, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetBytes(data)
	return i, nil
}

func NewFromString(s string, base int) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(s, base)
	if !ok {
		fmt.Println("SetString: error")
		return nil
	}
	//fmt.Println(n)
	return n
}

type Point struct {
	X *big.Int
	Y *big.Int
}

// k*g
// k < 115792089237316195423570985008687907852837564279074904382605163141518161494337
// g (55066263022277343669578718895168534326250603453777594175500187360389116729240, 32670510020758816978083085130507043184471273380659243275938904335757337482424)
func CalKPoint(k *big.Int, point *Point) *Point {
	n := NewFromString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	// k %n == 0
	if new(big.Int).Mod(k, n).Cmp(new(big.Int)) == 0 {

	}
	// k < 0
	if new(big.Int).Cmp(k) > 0 {
		return CalKPoint(k.Neg(k), PointNeg(point))
	}
	addend := point
	var result *Point
	// k > 0
	for new(big.Int).Cmp(k) < 0 {
		// k %2 == 1
		if k.Mod(k, new(big.Int).SetInt64(2)).Cmp(new(big.Int).SetInt64(1)) == 0 {
			result = PointAdd(result, addend)
		}
		addend = PointAdd(addend, addend)
		// k /= 2
		k = k.Div(k, new(big.Int).SetInt64(2))
	}
	return result
}

func PointNeg(point *Point) *Point {
	p := NewFromString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	y_neg := point.Y.Neg(point.Y)
	y_neg_p := y_neg.Mod(y_neg, p)
	return &Point{
		X: point.X,
		Y: y_neg_p,
	}
}

func PointAdd(point1 *Point, point2 *Point) *Point {
	if point1 == nil {
		return point2
	}
	if point2 == nil {
		return point1
	}
	x1, y1 := point1.X, point1.Y
	x2, y2 := point2.X, point2.Y
	if x1.Cmp(x2) == 0 && y1.Cmp(y2) != 0 {
		return nil
	}
	if x1.Cmp(x2) == 0 {
		//         m = (3 * x1 * x1 + curve.a) * inverse_mod(2 * y1, curve.p)
	} else {
		//         m = (y1 - y2) * inverse_mod(x1 - x2, curve.p)
	}
	//  x3 = m * m - x1 - x2
	//    y3 = y1 + m * (x3 - x1)
	//    result = (x3 % curve.p,
	//              -y3 % curve.p)
	return &Point{}
}
