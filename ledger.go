package ledger

import "log"


type LedgerBlock struct {
	Content string
}
var LedgerChain []LedgerBlock

func ShowAllLedgerBlocks() {
	for _, block := range LedgerChain {
		log.Println(block.Content)
	}
}

func CreateNewBlock(content string) {
	blockToAdd := LedgerBlock{Content: content}
	LedgerChain = append(LedgerChain, blockToAdd)
}
func UpdateBlockContent(blockIndex int, newContent string) {
	if blockIndex < 0 || blockIndex >= len(LedgerChain) {
		log.Println("Block index is out of bounds")
		return
	}
	LedgerChain[blockIndex].Content = newContent
}
