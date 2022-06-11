Bit Magic
---
Bit magic has following operators:
- & (AND)
- | (OR)
- ^ (XOR)
- ~ (NOT)
- << (Left Shift)
- >> (Right shift)

## Operators and their uses

### AND (&)
In this, only both set bits produces 1. i.e.
```
0 & 0 = 0
1 & 0 = 0
0 & 1 = 0
1 & 1 = 1
```

Hence, this has following uses:
- check if last bit is set (is 1), by AND'ing  anything with 1. If set, it produces 1 else 0

### OR (|)
In this, either set bits produces 1. i.e.
```
0 | 0 = 0
1 | 0 = 1
0 | 1 = 1
1 | 1 = 1
```

Hence, this has following uses:
- check if both bit is unset, by OR'ing  anything with 0. If unset, it produces 0 else 1

### Left Shift (<<)
It shifts 1 bit towards left and inserts 0 at the rightmost place.
i.e left shift(000101) = 001010 (1 has shifted to left, 0 inserted at rightmost)


### Right Shift (<<)
It shifts 1 bit towards right and inserts 0 at the leftmost place.
i.e right shift(000101) = 001010 (1 has shifted to right, 0 inserted at leftmost)

### XOR (^)

XOR has many uses in bitwise problems because of following properties:

a ^ b = b ^ a
a ^ a = 0 (most used)
a ^ 0 = a
(a ^ b) ^ c = a ^ (b ^ c)

## Useful logics
1. To count set bits in a number, keep shifting right and AND'ing by 1.
The set bits will produce only 1, count them
2. To check Kth bit is set in num, shift Left 1 (k-1) times and AND with number.
if 1, its set, else not
3. if n & (n-1), the right most set bit will become 0. (Brian Kerningam algo).
Use this to count set bits (keep reducing 1s by doing n & n-1 and counting them)
4. To check if number is power of 2, keep dividing the number by 2. If you get 
an odd number greater than 1, it is not power of 2. else it is.
5. In any number that is power of 2 (4,8,16), they all have  only 1 bit set. Use
Brian Kerningam algo to check if a number has only 1 set bit.
6. To find odd appearing number in an array (eg [4,4,2,3,3], 2 is odd), use XOR.
(4 ^ 4) ^ 2 ^ (3 ^ 3) = (0 ^ 2 ^ 0) = 2
