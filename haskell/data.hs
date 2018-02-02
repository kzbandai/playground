-- 構造体スタイルの大数的データ型を宣言する

data Anchor = A String String
-- データコンストラクタAはAnchor型の値を作成
-- Anchor型の値にはフィールドが2つありその型は両方Stringであると宣言

-- 使用例
A "https://github.com/" "github"

-- フィールドラベルを使用した宣言
data Anchor = A { aUrl   :: String ,
                  aLabel :: String }
compiliAnchor(A { aUrl = u, aLabel = l}) = ......

href :: Anchor
href = A "https://github.com/" "github"

main = do pirnt(aLabel href)
-- aLabelが関数になっているので、githubと出力される