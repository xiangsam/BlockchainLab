/*
 * @Author: Samrito
 * @Date: 2021-10-24 14:49:40
 * @LastEditors: Samrito
 * @LastEditTime: 2021-10-24 15:37:18
 */
package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBits = 20

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			IntToHex(pow.block.Time),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// get nonce and Hash by pow
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for {
		hash = sha256.Sum256(pow.prepareData(nonce))
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("%x\n\n", hash)
			return nonce, hash[:]
		} else {
			nonce += 1
		}
	}
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	var isValid bool
	hash := sha256.Sum256(pow.prepareData(pow.block.Nonce))
	hashInt.SetBytes(hash[:])
	isValid = (hashInt.Cmp(pow.target) == -1)
	return isValid
}
