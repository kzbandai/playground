gcc_options = -std=c++17 -Wall --pedantic-errors

build:
	/usr/bin/g++ $(gcc_options) -include all.h -o ./build/program ./main.cpp

header:
	/usr/bin/g++ $(gcc_options) -x c++-header -o ./all.h.gch ./all.h

run: header build
	./build/program

measure:
	time /usr/bin/g++ $(gcc_options) -include ./all.h -o ./build/program ./main.cpp

clean :
	rm -f ./build/program
	rm -f ./all.h.gch

.PHONY: build header run measure clean
