package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

// Block Chain
type Block struct {
	// data                   string
	MerkleTree             *Node
	Nonce                  string
	Previous_block_address *Block
	Previous_Block_hash    string
	Block_number           int
	Block_hash             string
}

type Blockchain struct {
	Genesis_block *Block
	Last_block    *Block
}

func New_Block(merkelTree *Node, Nonce string) *Block {
	new_block := new(Block)
	new_block.MerkleTree = merkelTree
	new_block.Nonce = Nonce

	return new_block
}

func Mine_Block(new_block *Block) bool {
	number := 0
	threshold := 1000
	for {
		if threshold < number {
			return false
		}
		if strconv.Itoa(number) == new_block.Nonce {
			return true
		} else {
			number++
		}
	}
}

func Add_To_Blockchain(Previous_block_address *Block, new_block *Block) *Block {

	if Mine_Block(new_block) {
		if Previous_block_address == nil {
			new_block.Previous_block_address = nil
			new_block.Previous_Block_hash = ""
			new_block.Block_number = 1

		} else {
			new_block.Previous_block_address = Previous_block_address
			new_block.Previous_Block_hash = Previous_block_address.Block_hash
			new_block.Block_number = Previous_block_address.Block_number + 1

		}
		new_block.Block_hash = Calculate_Hash(new_block)
		return new_block
	}
	return Previous_block_address
}

func Display_Blocks(Last_block *Block) {
	list := Last_block
	if list == nil {
		println("No blocks in the Blockchain")
	}
	for list != nil {
		fmt.Println("-------------------- Block", list.Block_number, "--------------------")
		fmt.Println("Transaction (Merkle Tree) : ")
		DisplayMerkelTree(list.MerkleTree)
		fmt.Println("Nonce                     : ", list.Nonce)
		fmt.Println("Block Hash                : ", list.Block_hash)
		fmt.Println("Previous block Hash       : ", list.Previous_Block_hash)
		fmt.Println("")
		list = list.Previous_block_address
	}
	println("")
	println("Blockchain end")
}

func Create_Blockchain() *Blockchain {
	temp_block_chain := new(Blockchain)
	temp_block_chain.Genesis_block = nil
	temp_block_chain.Last_block = nil
	return temp_block_chain
}
func Calculate_Hash(block *Block) string {
	str2 := strconv.Itoa(block.Block_number)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(block.MerkleTree.hash+block.Nonce+str2+block.Previous_Block_hash)))
}

func Verify_Chain(Last_block *Block) {
	list := Last_block

	if list == nil {
		println("No blocks in the Blockchain")
	}
	for list != nil {
		if list.Previous_block_address != nil {
			if list.Previous_Block_hash == Calculate_Hash(list.Previous_block_address) {
				list = list.Previous_block_address
			} else {
				fmt.Println("Block", list.Block_number-1, "is tempered")
				break
			}
		} else {
			fmt.Println("All blocks are verified and not tempered")
			break
		}
	}
	println("")

}

func Change_Block(block *Block) {
	Change_Block_Menu()
	user_input := ""
	fmt.Print("Enter number : ")
	fmt.Scanln(&user_input)

	if user_input == "1" {
		fmt.Print("Enter new Block number : ")
		fmt.Scanln(&user_input)
		Int, err := strconv.Atoi(user_input)
		if err == nil {
			block.Block_number = Int
		} else {
			fmt.Println("Should be integer")
		}

	} else if user_input == "2" {
		fmt.Print("Enter Nonce: ")
		fmt.Scanln(&user_input)
		block.Nonce = user_input

	} else if user_input == "3" {
		fmt.Print("Enter New Hash: ")
		fmt.Scanln(&user_input)
		block.Previous_Block_hash = user_input
	} //else if user_input == "4" {
	// 	fmt.Print("Enter New Data: ")
	// 	fmt.Scanln(&user_input)
	// 	block.data = user_input
	// }

}

func Menu() {
	println("-----------------------------")
	println("1) Display Blocks")
	println("2) Add new block")
	println("3) Verify Chain")
	println("4) Add temporary Blocks to chain")
	println("5) Change Block")
	println("9) Exit")
	println("-----------------------------")

}

func Change_Block_Menu() {
	fmt.Println("1) Change Block Number")
	fmt.Println("2) Change Nonce")
	fmt.Println("3) Change Previous Block Hash")
	fmt.Println("4) Change Data")
}

// //////////////////////////////////////////////////////////////////
// Merkel Tree
type Node struct {
	hash  string
	left  *Node
	right *Node
}

func getLeft(n *Node) *Node {
	return n.left
}
func setLeft(n *Node, x *Node) {
	n.left = x
}
func getRight(n *Node) *Node {
	return n.right
}
func setRight(n *Node, x *Node) {
	n.right = x
}
func getHash(n *Node) string {
	return n.hash
}
func setHash(n *Node, x string) {
	n.hash = x
}

func generateTree(dataBlocks []string) *Node {

	var arr1 = make([]*Node, len(dataBlocks))

	for i := 0; i < len(dataBlocks); i++ {

		nodeObj := new(Node)
		setLeft(nodeObj, nil)
		setRight(nodeObj, nil)
		setHash(nodeObj, fmt.Sprintf("%x", sha256.Sum256([]byte(dataBlocks[i]))))
		arr1[i] = nodeObj
	}

	return buildTree(arr1)
}
func buildTree(children []*Node) *Node {

	var parents = make([]*Node, len(children))

	for len(children) != 1 {
		var index = 0
		var length = len(children)
		var i = 0
		for index < length {
			leftChild := children[index]
			rightChild := new(Node)

			if (index + 1) < length {
				rightChild = children[index+1]
			} else {
				nodeObj := new(Node)
				setLeft(nodeObj, nil)
				setRight(nodeObj, nil)
				setHash(nodeObj, getHash(leftChild))
				rightChild = nodeObj
			}
			var parentHash = fmt.Sprintf("%x", sha256.Sum256([]byte(getHash(leftChild)+getHash(rightChild))))
			nodeObj := new(Node)
			setLeft(nodeObj, leftChild)
			setRight(nodeObj, rightChild)
			setHash(nodeObj, parentHash)

			parents[i] = nodeObj
			i++
			index += 2
		}
		children = parents[0:i]

		parents = parents[0:0]
		parents = parents[0:len(children)]

	}
	return children[0]
}
func DisplayMerkelTree(root *Node) {
	if root == nil {
		return
	}

	if getLeft(root) == nil && getRight(root) == nil {
		fmt.Println(getHash(root))
	}
	queue := make([]*Node, 0)
	// Push queue
	queue = append(queue, root)
	queue = append(queue, nil)

	for !(len(queue) == 0) {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			fmt.Println(getHash(node))
		} else {
			fmt.Println()
			if !(len(queue) == 0) {
				queue = append(queue, nil)
			}
		}

		if node != nil && getLeft(node) != nil {
			queue = append(queue, getLeft(node))
		}

		if node != nil && getRight(node) != nil {
			queue = append(queue, getRight(node))
		}

	}

}

func Get_Transactions() *Node {
	user_input := 0
	fmt.Print("How many transactions you want to store in block : ")
	fmt.Scanln(&user_input)
	transactions := ""
	var dataBlocks = make([]string, user_input)
	for i := 0; i < user_input; i++ {
		fmt.Print("Enter Transaction", i+1, ":")
		fmt.Scanln(&transactions)
		dataBlocks[i] = transactions
	}
	return generateTree(dataBlocks)
}
func Get_Random_Transactions() *Node {
	user_input := 2
	transactions := ""
	var dataBlocks = make([]string, user_input)
	for i := 0; i < user_input; i++ {
		transactions = "Fake-transaction" + strconv.Itoa(rand.Intn(1000000))
		dataBlocks[i] = transactions
	}
	return generateTree(dataBlocks)
}
