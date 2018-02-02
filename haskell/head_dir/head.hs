-- ./head < ../import_text.txt

main = do cs <- getContents
          putStr $ firstNLines 10 cs

firstNLines n = unlines . take n . lines