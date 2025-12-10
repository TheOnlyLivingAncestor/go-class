# Read secret device register

Consider a custom RF device that has `Channel Control Register` known as `CHAN_CTRL`. The register contains four fields, each of them has one-byte length.

Write a small function that accepts one argument: `CHAN_CTRL`, which has `uint32` type. The function should parse the register and return each the value of each field. The field order in the register is the following:

| **Bytes** | **1**   | **2**   | **3**    | **4**    |
|-----------|---------|---------|----------|----------|
| Fields    | RX_PCODE | TX_PCODE | RX_CHAN | TX_CHAN |

You have to return the fields in this order: `TX_CHAN`, `RX_CHAN`, `RX_PCODE`, `TX_PCODE`.

Example: if the parameter is `0x82ABBA19`, then `RX_PCODE = 0x82`, `TX_PCODE = 0xAB`, `RX_CHAN = 0xBA` and `TX_CHAN = 0x19`.

Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.
