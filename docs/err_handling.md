# Critical Error Handling

## Dropped cross-chain packages 

Since the cross-chain packages are relayed to Binance Chain grouped by `OracleSequence`. 

The `OracleSequence` is increased one by one in Binance Chain, so there is not possible to
skip `OracleSequence`. But it is possible that some packages are dropped by relayers.

The error can't be detected immediately util Binance Chain finds that there are some 
`PackageSequence` missing.

When that happens, the relayers needs to add the missing packages to the next group 
manually so that every thing will be back to normal and Binance Chain will not complain
missing `PackageSequence` any longer.