### Potential Redis Schema
    gomarket:orders:{symbol}:uOrderId # INCR
    gomarket:transactions:{symbol}:uOrderId # hash
    gomarket:orders:{symbol}:sell # list
    gomarket:orders:{symbol}:buy # list
    gomarket:{symbol}:ask
    gomarket:{symbol}:bid
    gomarket:{symbol}:price # Is this relavent?

### Other
Then have the bid price be a couple cents lower than the ask so it costs more to buy the shares than you get by selling them. Put a small commission on every trade, and let 'er go'.


### Symbols

    gomarket:symbols # Set of all symbols
    gomarket:symbols:JOBR # Hash 
    gomarket:symbols:JOBR:listed # Bit
    gomarket:sellorders:JOBR
    gomarket:sellorders:JOBR:123456
