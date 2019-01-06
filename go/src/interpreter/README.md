## Writing An Interpreter In Go

https://interpreterbook.com/

## notes
- source code -> tokens -> abstract syntax tree (AST)
  - first transformation: lexical snalysis, lexing, lexer
  - second: parsing? parser
- define tokens in token package which is required by lexing
- REPL
  - "Read Eval Print Loop"
  - console, interactive mode etc...