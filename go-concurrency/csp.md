# Communicating Sequencial Processes (CSP)

## What is the difference between concurrency and parallelism?
- concurrency and parallelism are not the same
 >Concurrency is a property of the code; parallelism is a property of the running
program.
 - Well, let’s think about that for second. If I write my code with the intent that two
chunks of the program will run in parallel, do I have any guarantee that will actually
happen when the program is run? What happens if I run the code on a machine with
only one core? Some of you may be thinking, It will run in parallel, but this isn’t true!
The chunks of our program may appear to be running in parallel, but really they’re
executing in a sequential manner faster than is distinguishable. The CPU context
switches to share time between different programs, and over a coarse enough granularity of time, the tasks appear to be running in parallel. If we were to run the same
binary on a machine with two cores, the program’s chunks might actually be running
in parallel.
- This reveals a few interesting and important things. **The first is that we do not write
parallel code, only concurrent code that we hope will be run in parallel.**
