# adventofcode

Here are my solutions of [advent of code](https://adventofcode.com/) written in Go.

When I benchmark my code I generate it like this:

```sh
go test -count 5 -run '^$' -bench . -memprofile=v1.mem.pprof -cpuprofile=v1.cpu.pprof > v1.txt
```

And for every iteration on the code I call the same method but with a different version. When I want to compare, I use benchstat:

```sh
benchstat v1.txt v2.txt
```