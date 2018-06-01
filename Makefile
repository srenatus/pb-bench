bench:
	go test -v -count=1 -benchmem -bench .

all: bench-v1.0.0 bench-v1.1.0 bench-ra3da3

bench-%:
	cd $* && make -f ../Makefile bench
