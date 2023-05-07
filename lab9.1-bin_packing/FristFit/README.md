## FFBinPacking

The `FFBinPacking` function implements the First Fit bin packing algorithm. It takes four arguments:

- `n` (int) - the number of items to pack
- `BinS` (int) - the maximum size of each bin
- `s` (slice of int) - the sizes of the items to pack
- `Bin` (slice of int) - the output array that contains the bin number for each item

It uses a nested loop to iterate through each item and each bin. For each item, it tries to place it in the first bin that has enough remaining space. If it finds a suitable bin, it assigns the bin number to the corresponding element in the `Bin` array and updates the `used` array to reflect the new space usage. If no suitable bin is found, the item is left unpacked.

Finally, the function returns the number of bins used.

## printObjectAndBin

The `printObjectAndBin` function takes three arguments:

- `n` (int) - the number of items to pack
- `s` (slice of int) - the sizes of the items to pack
- `Bin` (slice of int) - the bin number for each item

It simply iterates through each item and prints its size and corresponding bin number.

## printBinj

The `printBinj` function takes four arguments:

- `n` (int) - the number of items to pack
- `B` (int) - the bin number to print
- `s` (slice of int) - the sizes of the items to pack
- `Bin` (slice of int) - the bin number for each item

It prints the item indices and sizes for all items that belong to the given bin.

## PackAndPrint

The `PackAndPrint` function takes four arguments:

- `n` (int) - the number of items to pack
- `BinSize` (int) - the maximum size of each bin
- `s` (slice of int) - the sizes of the items to pack
- `Bin` (slice of int) - the output array that contains the bin number for each item

It first prints the item sizes using a loop. Then it calls `FFBinPacking` to pack the items into bins and prints the results using `printObjectAndBin` and `printBinj`.

## sortObjects

The `sortObjects` function implements a bubble sort algorithm to sort the item sizes in descending order. It takes two arguments:

- `n` (int) - the number of items to sort
- `x` (slice of int) - the items to sort

It iterates through each pair of adjacent items and swaps them if they are in the wrong order. The loop is repeated until no more swaps are needed.

## main

The `main` function simply initializes the input parameters and calls `PackAndPrint` twice: once with the original item sizes, and once with the sizes sorted in descending order.
