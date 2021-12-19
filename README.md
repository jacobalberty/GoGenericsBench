This is a collection of various sorts implemnted both as []int only and as constraints.Ordered generics.
I split the generics into their own source file so they aren't ordered in the output. I run the benchmarks with the following command:
`gotip test -bench=. -benchmem | sort` to sort the output to have them grouped together.

the `BenchmarkCountingSortG` implementation is an awful version of the counting sort I wouldn't look to it as anything other than a bad port to generics.
The other sort implementations should all be textbook and give you a decent example of performance comparison.
I welcome any pull request to give me a sane implementation of the `csortg` function
