package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      ProductCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type ProductCheckout struct {
	ProductID    string `json:"product_id"`
	Buyer        string `json:"buyer"`
	CheckoutDate string `json:"checkout_date"`
	isGenesis    bool   `json:"is_genesis"`
}

type Product struct {
	ID           string `json:"id"`
	Category     string `json:"category"`
	Title        string `json:"title"`
	Brand        string `json:"brand"`
	Seller       string `json:"seller"`
	UPC          string `json:"upc"`
	MFD          string `json:"mfd"`
	Cost         string `json:"cost"`
	Availability bool   `json:"availability"`
}

type Supplychain struct {
	blocks []*Block
}

var SupplyChain *Supplychain

func newProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error in creating product"))
		return
	}

	h := md5.New()
	io.WriteString(h, product.Brand+product.Seller+product.MFD+product.UPC)
	product.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(product, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error in saving the data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (bc *Supplychain) AddBlock(data ProductCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1]

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, ProductCheckout{isGenesis: true})
}

func NewSupplychain() *Supplychain {
	return &Supplychain{[]*Block{GenesisBlock()}}
}

func validBlock(block, prevBlock *Block) bool {

	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if !block.validateHash(block.Hash) {
		return false
	}

	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}
	return true
}

func getSupplychain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(SupplyChain.blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	io.WriteString(w, string(jbytes))
}

func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutItem ProductCheckout
	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not write Block: %v", err)
		w.Write([]byte("could not write block"))
		return
	}

	SupplyChain.AddBlock(checkoutItem)
	resp, err := json.MarshalIndent(checkoutItem, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not write block"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(b.Pos) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *Block, checkoutItem ProductCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

func main() {

	SupplyChain = NewSupplychain()

	r := mux.NewRouter()
	r.HandleFunc("/", getSupplychain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newProduct).Methods("POST")
	go func() {

		for _, block := range SupplyChain.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}

	}()
	log.Printf("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
