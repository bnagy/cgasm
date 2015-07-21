# cgasm

## About

`cgasm` is a standalone, offline terminal-based tool with no dependencies that gives me x86 assembly documentation. It is pronounced "SeekAzzem".

https://storify.com/thegrugq/once-upon-a-time-in-the-valley

## Installation

You should follow the [instructions](https://golang.org/doc/install) to
install Go, if you haven't already done so. Then:
```bash
$ go get github.com/bnagy/cgasm
```

The binary is standalone - it's statically linked and the data is compiled in.
You can cross compile this to any architecture that go supports.

## Acknowledgements

- nologic for doing the hard work on [idaref](https://github.com/nologic/idaref)
- @brnocrist for being the first to point me at it

## Usage

What was that AES thing...?
```
velleity:~ ben$ cgasm -f aes
Fuzzy matches for "aes" (12):
AESENC - Perform One Round of an AES Encryption Flow
VAESIMC -> AESIMC - Perform the AES InvMixColumn Transformation
AESKEYGENASSIST - AES Round Key Generation Assist
AESIMC - Perform the AES InvMixColumn Transformation
VAESDECLAST -> AESDECLAST - Perform Last Round of an AES Decryption Flow
AESDEC - Perform One Round of an AES Decryption Flow
VAESENC -> AESENC - Perform One Round of an AES Encryption Flow
VAESDEC -> AESDEC - Perform One Round of an AES Decryption Flow
VAESENCLAST -> AESENCLAST - Perform Last Round of an AES Encryption Flow
AESENCLAST - Perform Last Round of an AES Encryption Flow
VAESKEYGENASSIST -> AESKEYGENASSIST - AES Round Key Generation Assist
AESDECLAST - Perform Last Round of an AES Decryption Flow
```

Default output is a summary
```
velleity:~ ben$ cgasm aesenc
AESENC - Perform One Round of an AES Encryption Flow

Description:
This instruction performs a single round of an AES encryption flow using a round
key from the second source operand, operating on 128-bit data (state) from the
first source operand, and store the result in the destination operand. Use the
AESENC instruction for all but the last encryption rounds. For the last encryption
round, use the AESENCCLAST instruction. 128-bit Legacy SSE version: The first
source operand and the destination operand are the same and must be an XMM register.
The second source operand can be an XMM register or a 128-bit memory location.
Bits (VLMAX1:128) of the corresponding YMM destination register remain unchanged.
VEX.128 encoded version: The first source operand and the destination operand
are XMM registers. The second source operand can be an XMM register or a 128-bit
memory location. Bits (VLMAX-1:128) of the destination YMM register are zeroed.
```

Or go full nerd
```
velleity:~ ben$ cgasm -v aesenc

AESENC - Perform One Round of an AES Encryption Flow:
| Opcode/Instruction                   | Op/En| 64/32-bit Mode| CPUID Feature Flag    | Description
| 66 0F 38 DC /r AESENC xmm1, xmm2/m128| RM   | V/V           | AES                   | Perform one round of an AES encryption
|                                      |      |               |                       | flow, operating on a 128-bit data (state)
|                                      |      |               |                       | from xmm1 with a 128-bit round key from
|                                      |      |               |                       | xmm2/m128.
| VEX.NDS.128.66.0F38.WIG DC /r VAESENC| RVM  | V/V           | Both AES and AVX flags| Perform one round of an AES encryption
| xmm1, xmm2, xmm3/m128                |      |               |                       | flow, operating on a 128-bit data (state)
|                                      |      |               |                       | from xmm2 with a 128-bit round key from
|                                      |      |               |                       | the xmm3/m128; store the result in xmm1.

Instruction Operand Encoding:
| Op/En| Operand 1       | Operand2     | Operand3     | Operand4
| RM   | ModRM:reg (r, w)| ModRM:r/m (r)| NA           | NA
| RVM  | ModRM:reg (w)   | VEX.vvvv (r) | ModRM:r/m (r)| NA

Description:
This instruction performs a single round of an AES encryption flow using a round
key from the second source operand, operating on 128-bit data (state) from the
first source operand, and store the result in the destination operand. Use the
AESENC instruction for all but the last encryption rounds. For the last encryption
round, use the AESENCCLAST instruction. 128-bit Legacy SSE version: The first
source operand and the destination operand are the same and must be an XMM register.
The second source operand can be an XMM register or a 128-bit memory location.
Bits (VLMAX1:128) of the corresponding YMM destination register remain unchanged.
VEX.128 encoded version: The first source operand and the destination operand
are XMM registers. The second source operand can be an XMM register or a 128-bit
memory location. Bits (VLMAX-1:128) of the destination YMM register are zeroed.

Operation:

AESENC
STATE <- SRC1;
RoundKey <- SRC2;
STATE <- ShiftRows( STATE );
STATE <- SubBytes( STATE );
STATE <- MixColumns( STATE );
DEST[127:0] <- STATE XOR RoundKey;
DEST[VLMAX-1:128] (Unmodified)
VAESENC
STATE <- SRC1;
RoundKey <- SRC2;
STATE <- ShiftRows( STATE );
STATE <- SubBytes( STATE );
STATE <- MixColumns( STATE );
DEST[127:0] <- STATE XOR RoundKey;
DEST[VLMAX-1:128] <- 0

Intel C/C++ Compiler Intrinsic Equivalent:
| (V)AESENC:| __m128i _mm_aesenc (__m128i, __m128i)

SIMD Floating-Point Exceptions:
None

Other Exceptions:
See Exceptions Type 4.
```

## License

GPLv2, see LICENSE.md for details

## TODO

Nothing. No other features. Ever.

## Contributing

I. Will. Cut. You.
