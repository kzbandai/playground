goの覚書

- 関数の引数
	- 変数名の後ろに型名を書く
		- `add(x int, y int)`
	- 型が同じであれば最後の型を残して省略して記述できる
		- `add(x, y int)`
- 返り値は複数返せる

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

- 引数に名前をつけると返り値は省略できる
	- naked returnと呼ぶ

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

- varは宣言
	- 初期化子を与えることが出来る

```go
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```

- `:=`は暗黙的な型宣言

```go
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

- 基本型
	- [こちら](https://go-tour-jp.appspot.com/basics/11)
	- まとめて宣言できる
	- 初期化子を与えずに宣言するとゼロ値になる
		- 0, false, ""
- 定数
	- const
		- character, string, boolean, numeric
		- `:=`は使えない
- for
	- [こちら](https://go-tour-jp.appspot.com/flowcontrol/1)
	- whileはない
		- forの条件を省略した形
- if
	- ()は不要
- defer
	- 関数実行を呼び出し元の関数の終わりまで遅延させる
	- phpでいうclosure

```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
- 出力は`hello world`
- deferへ関数を複数渡すと、呼び出しはLIFOの順番で実行される
	- stackされる

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

- ポインタ
	- 変数Tのポインタは、`*T型`で、ゼロ値はnil
- make
	- スライスを作る
	- スライスはphpでいう配列
		- [こちら](https://go-tour-jp.appspot.com/moretypes/13)

- range
	- よくわからん
		- PHPのforeachやん
	- 「for ループに利用する range は、スライスや、マップ( map )をひとつずつ反復処理するために使います。スライスをrangeで繰り返す場合、rangeは反復毎に2つの変数を返します。 1つ目の変数はインデックス( index )で、2つ目はインデックスの場所の要素のコピーです。」

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

- exercise slices
	- phpでいう多次元配列？

```go
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sy := make([][]uint8, dy)

	for i := range sy {
		sy[i] = make([]uint8, dx)

		for k := range sy[i] {
			sy[i][k] = uint8((i + k) / 2)
		}
	}

	return sy
}

func main() {
	pic.Show(Pic)
}
```

- map
	- phpでいう連想配列？
	- キーと値をmapする
	- [ここ](https://go-tour-jp.appspot.com/moretypes/19)

```go
func WordCount(s string) map[string]int {
	t := strings.Fields(s)
	r := make(map[string]int)

	for _, v := range t {
		r[v]++
	}

	return r
}

func main() {
	wc.Test(WordCount)
}
```

- closure
	- 動きだけ見るとphpのyieldに近い？

```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

- 返り値の値がbindされるのでこれは成立する
	- 初期化されない

```go
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

- method
	- 型に定義する
	- methodは特別なレシーバを引数に取る
		- 自身のインスタンス？
	- ポインターreceiverは自己破壊的
		- ポインターを受け取るメソッドを作れる
		- `*`を消すとコピーを操作する
	- `&`はポインタを取り出す

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

```go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
```

- type assertion
	- phpでいうinstance ofですね
	- boolで確認しましょうね

```go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

- switchは文ではなく型

```go
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

- Stringer
	- phpでいうmagicメソッド？

```go
import "fmt"

type IPAddr [4]byte

func (ipa IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipa[0], ipa[1], ipa[2], ipa[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

```

- error
	- 組み込みinterface
	- エラーの状態はerror変数で表現する
	- [ここ](https://go-tour-jp.appspot.com/methods/19)

```go
type error interface {
    Error() string
}
```

- rot13Reader

```go
package main

import (
	"io"
	"os"
	"strings"
)

func rot13(c byte) byte {
	switch {
	case ('A' <= c && c <= 'Z'):
		return (c-'A'+13)%26 + 'A'
	case ('a' <= c && c <= 'z'):
		return (c-'a'+13)%26 + 'a'
	default:
		return c
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(s []byte) (n int, err error) {
	n, err = r.r.Read(s)

	if err != nil {
		return 0, err
	}

	for i := range s {
		s[i] = rot13(s[i])``
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```
