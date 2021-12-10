/*
 * @Author: Samrito
 * @Date: 2021-10-24 00:19:03
 * @LastEditors: Samrito
 * @LastEditTime: 2021-10-26 19:39:14
 */
package main

/*
func main() {
	t := time.Now()

	block := NewBlock("Genesis Block", []byte{})

	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	fmt.Printf("Time: %s\n", time.Unix(block.Time, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)

	fmt.Println("Time using: ", time.Since(t))
}
func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for _, block := range bc.blocks {
		pow := NewProofOfWork(block)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
*/
func main() {
	bc := NewBlockchain()
	cli := BlockchainCLI{bc}
	cli.Run()
}
