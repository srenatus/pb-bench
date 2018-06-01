# PB 1.0.0 vs 1.1.0 benchmark ![travis](https://api.travis-ci.org/srenatus/pb-bench.svg?branch=master)

With the giant message that is used here for benchmarking, a rather severe regression seems to be present.

(Alternatively, I'm doing something wrong in the benchmark, which is at least just as likely.)

Use `make all` to run the benchmarks for v1.0.0, v1.1.0 and master (ref 3a3da3 at the moment), and have some benchstat calls compare the results:

```console
$ make all
cd v1.0.0 && make -f ../Makefile bench | grep ^Bench | tee ../v1.0.0.out
BenchmarkRunMsgMarshal-8                    1000           1387356 ns/op          495933 B/op       5540 allocs/op
BenchmarkRunMsgMarshal-8                    1000           1373843 ns/op          495932 B/op       5540 allocs/op
...
```

The crucial part of that seems to be differences between v1.0.0 and v1.1.0:

```
v1.0.0 (old) vs v1.1.0 (new)
name                  old time/op    new time/op    delta
RunMsgMarshal-8         1.43ms ± 6%    5.59ms ± 2%  +290.08%  (p=0.000 n=9+10)
RunStructMarshal-8      1.52ms ± 4%    6.78ms ± 2%  +347.25%  (p=0.000 n=10+10)
RunMsgUnmarshal-8       2.55ms ± 3%    2.39ms ± 2%    -6.22%  (p=0.000 n=10+10)
RunStructUnmarshal-8    2.66ms ± 1%    2.58ms ± 3%    -3.01%  (p=0.000 n=10+10)
name                  old alloc/op   new alloc/op   delta
RunMsgMarshal-8          496kB ± 0%    1318kB ± 0%  +165.77%  (p=0.000 n=10+10)
RunStructMarshal-8       500kB ± 0%    1595kB ± 0%  +219.33%  (p=0.000 n=10+10)
RunMsgUnmarshal-8        649kB ± 0%     582kB ± 0%   -10.26%  (p=0.000 n=9+10)
RunStructUnmarshal-8     675kB ± 0%     604kB ± 0%   -10.47%  (p=0.000 n=10+10)
name                  old allocs/op  new allocs/op  delta
RunMsgMarshal-8          5.54k ± 0%    47.20k ± 0%  +751.93%  (p=0.000 n=10+10)
RunStructMarshal-8       5.76k ± 0%    57.27k ± 0%  +894.39%  (p=0.000 n=10+10)
RunMsgUnmarshal-8        23.7k ± 0%     18.1k ± 0%   -23.47%  (p=0.000 n=10+10)
RunStructUnmarshal-8     24.8k ± 0%     18.9k ± 0%   -23.64%  (p=0.000 n=9+10)
```
