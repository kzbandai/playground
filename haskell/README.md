Haskell勉強用git
=
## 参考文献
ふつうのHaskellプログラミング

http://www.amazon.co.jp/dp/4797336021

## コンパイルして実行
`ghc hoge.hs -o hoge`

`./hoge`

## コンパイルせずにいきなり実行
`runghc hoge.hs`

## $
$は括弧として解釈すれば良い

`print $ length $ lines cs`　なら　`print(length(lines cs))`　として解釈

## 型の宣言

```
firstNLines :: Int -> String -> String
firstNLines n cs = unlines $ take n $ lines cs
```

これは`関数名 :: 第一引数の型(n) -> 第二引数の型(cs) -> 返り値の型`

## 遅延評価と参照透明性

* 関数に先に展開していく
* 最後に計算する
* 透明なパイプで様々な参照先が繋がっていて、その中の一つの要素がするする動くイメージ

## letとwhereの違い
```
f n = let x = n + 1
          y = n + 2
          z = n + 3
      in x * y * z
```

letは式自体も値を持つ

```
where
  translate '\t' = replicate width ' '
  translate c    = [c]
```

whereは節なので値を持たない

## ポイントフリースタイル

* 書き直す

* 書き直しは脊髄反射のスピードでやれるとよい
