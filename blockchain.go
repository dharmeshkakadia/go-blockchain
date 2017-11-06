package blockchain

type Blockchain struct {
	blocks []*Block
}

func (c *Blockchain) Add(data string) {
	c.blocks = append(c.blocks, NewBlock(data, c.blocks[len(c.blocks)-1].Hash))
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{InitBlock()}}
}

func (c *Blockchain) Print() {
	for _, block := range c.blocks {
		block.print()
	}
}
