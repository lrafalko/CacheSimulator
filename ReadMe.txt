
Class: CS4541
Assignment: Assignment 2 Cache Simulator
Author Lance Rafalko

Summary:

I decided to use GoLang for this assignment as it is a great language for developing
command line tools. The entry point for this project is the main function in main.go.
I use two different structs in this assignment to store important information. Each
instruction is parsed into an instruction struct. These instruction structs are then
added to a cache in the cache struct. The cache is implimented using a simple hashmap
where the key is an integer and the value is an array of pointers to instructions.
The main function calls the update cache function with a newly parsed instruction which
adds the instruction to the cache. The update cache function automatically keeps track
of the number of hits, misses, and evictions and records these numbers in the cache
struct itself.

How to run
Navigate to the base directory of this project which contains this file and the go.mod file
. Run the command: make build

Then run the executable ./main -t traces/<trace file> -b <number> -s <number> - E <number> 

The program will then print out the number of hits, misses and evictions

Credits
I used the GoLang documentation at the below URL to understand how the flag, bufio, and
path packages work.

https://go.dev/doc/


