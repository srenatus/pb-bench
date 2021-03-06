all: bench-v1.0.0 bench-v1.1.0 bench-r-3a3da3 bench-r-1325a0 benchstat
	@echo "v1.0.0 (old) vs v1.1.0 (new)"
	@benchstat v1.0.0.out v1.1.0.out
	@echo "\n\nv1.0.0 (old) vs 3a3da3 (new)"
	@benchstat v1.0.0.out r-3a3da3.out
	@echo "\n\nv1.1.0 (old) vs 3a3da3 (new)"
	@benchstat v1.1.0.out r-3a3da3.out
	@echo "\n\nv1.0.0 (old) vs 1325a0 (new)"
	@benchstat v1.0.0.out r-1325a0.out
	@echo "\n\nv1.1.0 (old) vs 1325a0 (new)"
	@benchstat v1.1.0.out r-1325a0.out

bench:
	go test -v -count=10 -benchmem -bench .

benchstat:
	go get -u golang.org/x/perf/cmd/benchstat

bench-%:
	cd $* && make -f ../Makefile bench | grep ^Bench | tee ../$*.out
