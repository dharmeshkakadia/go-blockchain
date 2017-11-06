package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte // technically should be separate struct
	PrevBlock []byte
	Hash      []byte
	Trail     int // also called Nonce
}

// simple for test, not used anymore
func (b *Block) SetHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{
		b.PrevBlock,
		b.Data,
		[]byte(strconv.FormatInt(b.Timestamp, 10))}, // convert timestamp to []byte
		[]byte{})) // separator
	b.Hash = hash[:] // convert to slice
}

func NewBlock(data string, prevBlock []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlock, []byte{}, 0}
	block.Trail, block.Hash = NewProofOfWork(block).Mine()
	log.Printf("Generated new block")
	return block
}

func (block *Block) print() {
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("Trail: %d\n", block.Trail)
	fmt.Printf("Prev. hash: %x\n", block.PrevBlock)
	fmt.Println()
}

func InitBlock() *Block {
	return NewBlock("Gensis", []byte{})
}
