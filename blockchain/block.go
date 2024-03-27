package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Create a Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0} //Creation of Block without Hash
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Add a Block to a BlockChain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1] //Gets the previous Block of the BlockChain
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Create the first Block of the BlockChain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Initialize the BlockChain with "Genesis" block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
