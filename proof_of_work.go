package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
)

const difficulty = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	return &ProofOfWork{b, target}
}

func (w *ProofOfWork) GetHash(trail int) [32]byte {
	return sha256.Sum256(bytes.Join([][]byte{
		w.block.PrevBlock,
		w.block.Data,
		IntToHex(w.block.Timestamp),
		IntToHex(int64(difficulty)),
		IntToHex(int64(trail)),
	}, []byte{}))
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func (w *ProofOfWork) Mine() (int, []byte) {
	trial := 0
	var hashInt big.Int
	var hash [32]byte
	log.Printf("Mining block with data=%s", w.block.Data)
	for trial < math.MaxInt64 {
		hash = w.GetHash(trial)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(w.target) == -1 {
			break
		} else {
			trial++
		}
	}
	return trial, hash[:]
}

func (w *ProofOfWork) Validate() bool {
	var hashInt big.Int
	hash := w.GetHash(w.block.Trail)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(w.target) == -1
}
