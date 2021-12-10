/*
 * @Author: Samrito
 * @Date: 2021-10-24 00:19:03
 * @LastEditors: Samrito
 * @LastEditTime: 2021-10-24 15:58:45
 */
package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Time     int64
	Hash     []byte
	PrevHash []byte
	Data     []byte
	Nonce    int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte{}, prevHash, []byte(data), 0}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	//为Block生成hash，使用sha256.Sum256(data []byte)函数
	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()
	b.Nonce = nonce
	b.Hash = hash
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
