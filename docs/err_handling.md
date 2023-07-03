# Critical Error Handling

## Dropped cross-chain packages 

Since the cross-chain packages are relayed to BNB Beacon Chain grouped by `OracleSequence`. 

The `OracleSequence` is increased one by one in BNB Beacon Chain, so there is not possible to
skip `OracleSequence`. But it is possible that some packages are dropped by relayers.

The error can't be detected immediately util BNB Beacon Chain finds that there are some 
`PackageSequence` missing.

When that happens, the relayers needs to add the missing packages to the next group 
manually so that every thing will be back to normal and BNB Beacon Chain will not complain
missing `PackageSequence` any longer.