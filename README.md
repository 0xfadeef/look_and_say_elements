## Look-and-say sequence Elements
[![Go Reference](https://pkg.go.dev/badge/github.com/0xfadeef/look_and_say_elements.svg)](https://pkg.go.dev/github.com/0xfadeef/look_and_say_elements)

 Look-and-say sequence starts with arbitrary string of digits; to generate the next member, iterate over the digits
 and write the count of consecutive repetitions of a digit followed by the digit itself. For example, the member
 after `1113213211` is going to be `31131211131221`. Original look-and-say sequence starts with `1` (then `11`,
 `21`, `1211`, `111221`, etc), but it can start with any other number.

 The sequence has several notable properties, one of which is *digits presense limitation*: no digits other than 1,
 2, and 3 appear in the sequence, unless the seed string contains such a digit or any digit repeats more than 3
 times in a row.
      
 Another cool property is that *every sequence eventually splits*. What this means is that at some point the disjoint
 parts of the current sequence member start to evolve independently of each other (not affecting adjacent parts), so
 that transforming each part separately and then putting them back together in order will produce the next member
 of a sequence. Even more fascinating is that these parts are not actually random!
 [John Conway](https://mathshistory.st-andrews.ac.uk/Biographies/Conway) proved that they are members of a set
 consisting of 92 elements; Conway named them **atomic elements**. Every sequence eventually evolves into
 a combination of atomic elements.

 This module exposes the array of atomic elements. Each member of the array is a structure, which holds the "atomic
 number" of the element, its mnemonic name (corresponding chemical element), the "formula" (corresponding string of
 digits) and the slice of atomic numbers of the elements this element evolves to. The array is ordered so that the
 atomic number of an element equals to its array index: `Elements[n].Number == n`.
