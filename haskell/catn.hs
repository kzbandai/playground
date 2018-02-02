main = do cs <- getContents
          putStr $ numbering cs

-- numbering関数にStringを突っ込むとStringが返ってくる
numbering :: String -> String
numbering cs = unlines $ map format $ zipLineNumber $ lines cs

-- zipLineNumber関数に[String]を突っ込むと[(Int, String)]が返ってくる
zipLineNumber :: [String] -> [(Int, String)]
zipLineNumber xs = zip [1..] xs

-- format関数に(Int, String)を突っ込むとStringが返ってくる
format :: (Int, String) -> String
format (n, line) = rjust 6 (show n) ++ " " ++ line

-- rjust関数にIntとStringを突っ込むとStringが返ってくる
rjust :: Int -> String -> String
-- widthとsがIntとStringで右辺が処理内容
rjust width s = replicat (width - length s) '  ' ++ s