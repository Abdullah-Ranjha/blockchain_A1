package main

import (
	"./ledger" 
	"log"
)

func main() {
	ledger.CreateNewBlock("Block 1")          
	ledger.CreateNewBlock("Block 2")         
	ledger.ShowAllLedgerBlocks()                

	ledger.UpdateBlockContent(1, "Changed Second Block") 
	log.Println("After changes:")
	ledger.ShowAllLedgerBlocks()                  
}
