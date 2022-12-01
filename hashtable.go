package main

import "fmt"

const arraysize = 50

// elements''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''

// #Table structure
type HashTable struct {
	array [arraysize]*Bucket
}

// ------------------------------------------------
// bucket structure
type Bucket struct {
	head   *BucketNode
	length int
}

// ------------------------------------------------
// bucket node structure
type BucketNode struct {
	key  string
	next *BucketNode
}

//*** elements ***''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''

// functions''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''

// init
func Init0() *HashTable {
	create := &HashTable{}
	for i := range create.array {
		create.array[i] = &Bucket{}
	}
	return create
}

// ------------------------------------------------
//  #fn

func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	index := sum % arraysize
	return index

}

//*** functions ***''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
// methods''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''

// Insertion
func (h *HashTable) Insertion(key1 string) {
	index := Hash(key1)
	h.array[index].insert(key1)
}

// ------------------------------------------------
// search
func (h *HashTable) Search(key2 string) bool {
	index := Hash(key2)
	return h.array[index].search(key2)
}

// ------------------------------------------------
// Deletion
func (h *HashTable) Deletion(key3 string) {
	index := Hash(key3)
	h.array[index].delet(key3)
}

// ------------------------------------------------
// Insertion in bucket
func (b *Bucket) insert(key1 string) {

	if b.search(key1) == true {
		fmt.Println("fuck")
		return
	}
	if b.search(key1) == false {
		newNode := &BucketNode{key: key1}
		newNode.next = b.head
		b.head = newNode
		return
	}
	// nextNode := BucketNode{key: key1}
	// add := b.head

	// // head is nil
	// if b.length == 0 {
	// 	b.head = &nextNode
	// 	b.length++
	// 	return
	// }
	// // head have value
	// // addig value to next node
	// for i := 0; i < b.length; i++ {
	// 	if add.next == nil {
	// 		add.next = &nextNode
	// 		b.length++
	// 		return
	// 	}
	// 	add = add.next
	// }

}

// ------------------------------------------------
// search in bucket
func (b *Bucket) search(key2 string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key2 {
			fmt.Println("Input[", key2, "]found in position", Hash(key2))
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// ------------------------------------------------
// Deletion in bucket
func (b *Bucket) delet(key3 string) {

	currentNode := b.head
	if currentNode.key == key3 {
		currentNode = currentNode.next
		return
	}

	for currentNode.next != nil {
		if currentNode.next.key == key3 {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode.next = currentNode.next.next
	}
	println("User doesnot exist")
}

func InputSlice(arrSlice []string, err error) []string {
	var value string

	if err != nil {
		return arrSlice
	}
	number, err := fmt.Scanf("%s", &value)

	if number == 1 {
		arrSlice = append(arrSlice, value)
	}
	return InputSlice(arrSlice, err)
}

// *** methods ***”””””””””””””””””””””””””””””””””””””””
func main() {
	var count, choice, option int
	var inputValue string
	mytable := Init0()

	list := []string{
		"apple",
		"orange",
		"grape",
		"mango",
	}

	for _, v := range list {
		mytable.Insertion(v)
	}

	count = 1
	for i := 0; i < count; i++ {
		// user menu
		fmt.Println("choose an option")
		fmt.Println("1 : for adding inputs")
		fmt.Println("2 : for search")
		fmt.Println("3 : for deleting")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter  inputs")
			fmt.Scan(&inputValue)
			fmt.Println("")
			mytable.Insertion(inputValue)

		case 2:
			fmt.Println("Enter name for search")
			fmt.Scan(&inputValue)
			fmt.Println("")
			if !mytable.Search(inputValue) {
				fmt.Println("illaaaa ")
				fmt.Println("")
			}

		case 3:
			fmt.Println("Enter name for delete")
			fmt.Scan(&inputValue)
			fmt.Println("")
			mytable.Deletion(inputValue)

		default:
			fmt.Println("You entered wrong option number.....")
			fmt.Println("enter numbers from option")
			fmt.Println("")

		}
		fmt.Println("")
		fmt.Println(" if you want to continue enter 9 else 0")
		fmt.Scan(&option)
		if option == 9 {
			count++
			fmt.Println("")
			fmt.Println("")
		} else if option == 0 {
			println("")

			fmt.Println("program exit surcessfully")
		} else {
			fmt.Println("sorry you entered a wrong number .... program exit surcessfully")

		}
	}

}
