numberOfLines :: String -> Int
numberOfLines cs = length $ lines cs

-- 以上は以下と一致する
numberOfLines :: String -> Int
numberOfLines = length . lines

-- これを関数合成という
-- f(g(x))みたいな合成
-- 上で言うならlinesの値をlengthで処理したものがnumberOfLinesの値になる