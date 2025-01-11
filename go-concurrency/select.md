# SELECT statement
- The select statement is the glue that binds channels together; it’s how we’re able to compose channels together in a program to form larger abstractions.
 ```go
 var c1, c2 <-chan interface{}
var c3 chan<- interface{}
select {
case <- c1:
// Do something
case <- c2:
// Do something
case c3<- struct{}{}:
// Do something
 ```
- looks similar to switch block but they are not. Unlike switch blocks, case statements in a select block aren’t tested sequentially, and execution won’t automatically fall through if none of the criteria are met.
- Instead, all channel reads and writes are considered simultaneously
to see if any of them are ready: populated or closed channels in the case of reads, and channels that are not at capacity in the case of writes. If none of the channels are ready, the entire select statement **blocks**. Then when one the channels is ready, that operation will proceed, and its corresponding statements will execute.
