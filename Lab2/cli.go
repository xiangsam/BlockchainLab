/*
 * @Author: Samrito
 * @Date: 2021-10-26 19:07:21
 * @LastEditors: Samrito
 * @LastEditTime: 2021-10-26 19:45:13
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type BlockchainCLI struct {
	bc *Blockchain
}

func (cli *BlockchainCLI) help() {
	fmt.Println("help:")
	fmt.Println("	newblock -data [you data] % add a block to blockchain")
	fmt.Println("	listblocks % list all blocks in blockchain")
}

func (cli *BlockchainCLI) Run() {
	if len(os.Args) < 2 {
		cli.help()
		os.Exit(1)
	}
	newBlockCMD := flag.NewFlagSet("newblock", flag.ExitOnError)
	listBlocksCMD := flag.NewFlagSet("listblocks", flag.ExitOnError)
	newBlockdata := newBlockCMD.String("data", "", "block data")
	switch os.Args[1] {
	case "newblock":
		err := newBlockCMD.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listblocks":
		err := listBlocksCMD.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.help()
	}
	if newBlockCMD.Parsed() {
		if *newBlockdata == "" {
			newBlockCMD.Usage()
			os.Exit(1)
		} else {
			cli.newBlock(*newBlockdata)
		}
	}
	if listBlocksCMD.Parsed() {
		cli.listBlock()
	}
}
func (cli *BlockchainCLI) newBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("\nSuccess!")
}
func (cli *BlockchainCLI) listBlock() {
	bci := cli.bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevHash) == 0 {
			break
		}
	}
}
