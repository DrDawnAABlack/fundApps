# fundApps Courier Kata Application

Hello! This document is intended to enable you to understand the thought processes behind the software committed. As such I have included some of my thoughts and assumptions whilst writing this initial implementation, some thoughts on what I didn't get to finish or might do differently. I have tried to be as honest and comprehensive as possible. 

## TL;DR
Run `go build main/go` from the `fundApps` directory.  
Run `go test ./...` to run my tests  
Run `./main test_data/name_of_test.txt` to run from the cmd line with a text input test.  
Add further test files if you like. I have assumed the following input format:  

`parcel_id1 | d1 | d2 | d3`  
`parcel_id2 | d1 | d2 | d3`  
`speedy`  

This lists two parcels where:  
`parcel_id` is a unique identifier  
`d1,d2,d3` are the three parcel dimensions (assuming a cuboid)  
`speedy` indicates speedy shipping. Omit if not required.  

The output is printed to a text file and echoed to the console. It will be of the format:  

`parcel_id1 | cost of parcel 1`   
`parcel_id2 | cost of parcel 2`   
`Total = 106.00` // total cost of parcels due to size   
`Total with speedy shipping = 212.00`   

## Assumptions, Decisions, Refactors

### Float Dimensions
The parcel dimensions should have been in `float` format. Whilst a simple refactor, the time taken would have prevented me from completing additional parts of the test. So I made the decision to leave them as integers for now. 

### Speedy Shipping
Having decided on an input format for the parcels and dimensions, speedy shippping presented a challenge. The decision was whether to append an `s` to each input row, `parcel_id1 | d1 | d2 | d3 | s` for speedy shipping, or to include one entry at the end of the list of parcels. I settled on the latter. Whilst I realised that being able to select some packages as speedy shipping and some not (so essentially an order can be shipped in two parts) I felt that this approach would require lots of change to the parser (or it would end up very clumsy). The output would also be either harder to read or require a refactor. These changes, whilst worth doing, require time and would prevent me moving forward with the test, so I chose the simpler option of providing one entry `speedy` at the end of the order to indicate that all packages require speedy shipping.

I am also not entirely comfortable with the string-matching approach for speedy shipping, but assuming that the input files are auto-generated by some front-end software, it feels safe enough.
 
### The Parser
I have not tested the parser exhaustively. Writing parsers, whilst conceptually simple, are not often a day-to-day requirement in industry, and so it's always been a while since you last wrote one. Hence they suck up a fair bit of time and contribute fairly little to the test. My parser is not particularly robust and has not been exhaustively tested. Can I lean of the imaginary front-end to provide me with correct data?

### Discounts
Sadly, I did not get very far into finishing this. It occurs to me that the efficient solution to this problem requires a relational database. The sorting I am trying to achieve is simple given SQL queries. Realising this, I thought it best to get my code in than spend much more time implementing a solution. It has interested me, though, and I may well have a go at finishing the code prior to interview as I have to leave a problem unsolved :)

 
