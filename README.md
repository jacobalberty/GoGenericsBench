This is a collection of various sorts implemnted both as []int only and as constraints.Ordered generics.
I split the generics into their own source file so they aren't ordered in the output. I run the benchmarks with the following command:
`gotip test -bench=. -benchmem | sort` to sort the output to have them grouped together.

Most of these use `constraints.Ordered` the generic counting sort uses `constraints.Integer` due to limitations in what a counting sort can do.
