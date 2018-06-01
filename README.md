# PB 1.0.0 vs 1.1.0 benchmark

With the giant message that is used here for benchmarking, a rather severe regression seems to be present.

(Alternatively, I'm doing something wrong in the benchmark, which is at least just as likely.)

To reproduce this, there's three branches, each vendoring a different tag/branch of google/protobuf, and including a `pb.go` file generated with the vendored version:

- v1.0.0: `git checkout v100`
- v1.1.0: `git checkout v110`
- master: `git checkout r-3a3da3`

Then run `go test -v -count=10 -benchmem -bench .` (or `make bench`), store the result, and conclude with running benchstat on the results.
