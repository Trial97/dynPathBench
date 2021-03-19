## Some containers benchmarking

In order to improve navigable map we benchmarked:
- https://github.com/Jeffail/gabs
- https://github.com/akshaybharambe14/ijson
- https://github.com/bitly/go-simplejson
- the NavigableMap from [utils](https://github.com/cgrates/cgrates/blob/d4d602001c67dd301ff0f2e8858d71ad77d7ecb6/utils/navigablemap.go#L27)
- the OrderedNavigableMap from [utils](https://github.com/cgrates/cgrates/blob/d4d602001c67dd301ff0f2e8858d71ad77d7ecb6/utils/orderednavigablemap.go#L37)
- a new implentation in node.go

Ignored the following because they do not support Set:
- https://github.com/antonholmquist/jason
- https://github.com/chrhlnd/dynjson


## Results
1. For this path `Field1[*raw][0].Field2[0].Field3[*new]`
    - Get:
    ```
    goos: linux
    goarch: amd64
    pkg: github.com/Trial97/github.com/Trial97/dynPathBench
    cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
    BenchmarkGabsField
    BenchmarkGabsField-8                 	10571704	       115.4 ns/op	      16 B/op	       1 allocs/op
    BenchmarkIJSONField
    BenchmarkIJSONField-8                	14814273	        78.35 ns/op	       0 B/op	       0 allocs/op
    BenchmarkOrderdNavigableMapField
    BenchmarkOrderdNavigableMapField-8   	 9967938	       110.1 ns/op	       0 B/op	       0 allocs/op
    BenchmarkNavigableMapField
    BenchmarkNavigableMapField-8         	10827984	       117.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkNodeField
    BenchmarkNodeField-8                 	13504286	        88.17 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSimpleJSONField
    BenchmarkSimpleJSONField-8           	 3026115	       463.7 ns/op	     112 B/op	       7 allocs/op
    PASS
    ok  	github.com/Trial97/github.com/Trial97/dynPathBench	25.403s
    ```

    - Set
    ```
    goos: linux
    goarch: amd64
    pkg: github.com/Trial97/github.com/Trial97/dynPathBench
    cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
    BenchmarkGabsSet
    BenchmarkGabsSet-8                 	 7471309	       163.8 ns/op	      32 B/op	       2 allocs/op
    BenchmarkIJSONSet
    BenchmarkIJSONSet-8                	 4803643	       265.8 ns/op	      16 B/op	       1 allocs/op
    BenchmarkOrderdNavigableMapSet
    BenchmarkOrderdNavigableMapSet-8   	 2050528	       581.5 ns/op	     256 B/op	       5 allocs/op
    BenchmarkNavigableMapSet
    BenchmarkNavigableMapSet-8         	 2232222	       550.1 ns/op	     256 B/op	       5 allocs/op
    BenchmarkNodeSet
    BenchmarkNodeSet-8                 	 8660870	       135.1 ns/op	      16 B/op	       1 allocs/op
    BenchmarkSimpleJSONSet
    BenchmarkSimpleJSONSet-8           	 5793783	       190.9 ns/op	      16 B/op	       1 allocs/op
    PASS
    ok  	github.com/Trial97/github.com/Trial97/dynPathBench	9.142s
    ```

2. For random generated path
    - Get
    ```
    goos: linux
    goarch: amd64
    pkg: github.com/Trial97/github.com/Trial97/dynPathBench
    cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
    BenchmarkGabsField
    BenchmarkGabsField-8                 	     200	   5575869 ns/op	  160000 B/op	   10000 allocs/op
    BenchmarkIJSONField
    BenchmarkIJSONField-8                	     181	   6803418 ns/op	       0 B/op	       0 allocs/op
    BenchmarkOrderdNavigableMapField
    BenchmarkOrderdNavigableMapField-8   	     150	   9650566 ns/op	       0 B/op	       0 allocs/op
    BenchmarkNavigableMapField
    BenchmarkNavigableMapField-8         	     128	   9529681 ns/op	       0 B/op	       0 allocs/op
    BenchmarkNodeField
    BenchmarkNodeField-8                 	     171	   8136261 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSimpleJSONField
    BenchmarkSimpleJSONField-8           	     129	   8201091 ns/op	 1361153 B/op	   85072 allocs/op
    PASS
    ok  	github.com/Trial97/github.com/Trial97/dynPathBench	16.115s
    ```


    - Set
    ```
    goos: linux
    goarch: amd64
    pkg: github.com/Trial97/github.com/Trial97/dynPathBench
    cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
    BenchmarkGabsSet
    BenchmarkGabsSet-8                 	     158	   6673167 ns/op	  487702 B/op	   20951 allocs/op
    BenchmarkIJSONSet
    BenchmarkIJSONSet-8                	     123	   9146072 ns/op	  375452 B/op	   11222 allocs/op
    BenchmarkOrderdNavigableMapSet
    BenchmarkOrderdNavigableMapSet-8   	      81	  13964574 ns/op	  673490 B/op	   22106 allocs/op
    BenchmarkNavigableMapSet
    BenchmarkNavigableMapSet-8         	     147	   7199784 ns/op	  500241 B/op	   21022 allocs/op
    BenchmarkNodeSet
    BenchmarkNodeSet-8                 	     140	   7153473 ns/op	  343107 B/op	   11682 allocs/op
    BenchmarkSimpleJSONSet
    BenchmarkSimpleJSONSet-8           	     162	   7209049 ns/op	  323591 B/op	   10928 allocs/op
    PASS
    ok  	github.com/Trial97/github.com/Trial97/dynPathBench	9.084s
    ```