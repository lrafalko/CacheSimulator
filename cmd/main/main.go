package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Define Instruction to put in the cache
type instruction struct {
	tag             uint64
	dataBits        int
	setIndexBits    int
	instructionType string
	time            int
}

type cache struct {
	dataMap     map[int][]*instruction
	linesPerSet int
}

func parseAddress(addr uint64, s int, b int) (uint64, int, int) {

	// do a bit mask using the number of block bits and then shift the address b bits to the right
	blockMask := int(math.Pow(2, float64(b))) - 1
	blockOffsetReturn := addr & uint64(blockMask)
	addr = addr >> b

	setMask := int(math.Pow(2, float64(s))) - 1
	setOffsetReturn := addr & uint64(setMask)
	addr = addr >> s
	tagReturn := addr

	return tagReturn, int(setOffsetReturn), int(blockOffsetReturn)
}

func createCache(sets int, lines int) *cache {

	// init the cachemap
	cache := cache{}
	cache.dataMap = make(map[int][]*instruction, sets)

	// init the lines in the sets
	i := 0
	for i < sets {

		cache.dataMap[i] = make([]*instruction, 0, lines)
	}

	cache.linesPerSet = lines
	return &cache
}

// tok is a string containing one full address trace from valgrind
// s are the number of set bits
// b are the number of byte offset bits
func parseInstruction(tok string, s int, b int) (*instruction, error) {
	// return nil if the current trace is not a data parseInstruction
	if tok[0] != ' ' {

		return nil, nil
	}

	// create new scanner to parse the string
	scanner := bufio.NewScanner(strings.NewReader(tok))
	scanner.Split(bufio.ScanWords)

	// parse each piece of the line

	scanner.Scan()
	instruction_type := scanner.Text()
	// make sure the instruction is stipped
	instruction_type = strings.TrimSpace(instruction_type)
	scanner.Scan()
	memAddrStr := scanner.Text()
	// Split the memory adress string into two parts using split
	splitStrs := strings.Split(memAddrStr, ",")

	if len(splitStrs) != 2 {
		return nil, errors.New("Len of split Strs is not 2")
	}

	// parse the information from the memory address
	// convert the string into an int
	memAddr, err := strconv.ParseUint(splitStrs[0], 16, 64)
	fmt.Println(memAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	tag, set, byteOffset := parseAddress(memAddr, s, b)

	// fill instruction with information
	instruct := instruction{
		dataBits:        byteOffset,
		tag:             tag,
		setIndexBits:    set,
		instructionType: instruction_type,
	}

	return &instruct, nil
}
func main() {

	// create flags to get the command line arguments

	var setIndexbits int
	var linesFlag int
	var blockBits int
	var traceFile string

	flag.IntVar(&setIndexbits, "s", 1, "Usage")
	flag.IntVar(&linesFlag, "E", 1, "Usage")
	flag.IntVar(&blockBits, "b", 1, "Usage")
	flag.StringVar(&traceFile, "t", "", "Usage")

	flag.Parse()

	workdir, err := os.Getwd()
	var filePath string = workdir + traceFile

	fp, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Can't open trace at given file location")
	}

	scanner := bufio.NewScanner(fp)
	// init a new cace struct
	// keep track of instruction order
	var time int = 0
	var instruc *instruction
	// keep track of misses, hits, evictions
	var (
		misses    int = 0
		hits      int = 0
		evictions int = 0
	)
	println(time, instruc, misses, hits, evictions)
	// init the cache
	cache := createCache(int(math.Pow(float64(2), float64(setIndexbits))), linesFlag)
	for scanner.Scan() {
		// get bytes from line
		if len(scanner.Text()) == 0 {
			// Skip empty line
			continue
		}
		instruc, err = parseInstruction(scanner.Text(), setIndexbits, blockBits)

		if err != nil {
			log.Fatal(err.Error())
			return
		}

		if instruc != nil {

			fmt.Printf("Tag: %d, set: %d, offset: %d\n", instruc.tag, instruc.setIndexBits, instruc.dataBits)
		}

		// add the instruction to the cache

	}
}
