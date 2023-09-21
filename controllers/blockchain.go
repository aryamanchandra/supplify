package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	models "github.com/aryamanchandra/supplify/models"
	"github.com/gin-gonic/gin"
)

type Block struct {
	Pos       int
	Data      models.ProductCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type Supplychain struct {
	Blocks []*Block
}

var SupplyChain *Supplychain

func NewProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in creating product"})
		return
	}

	h := md5.New()
	io.WriteString(h, product.Brand+product.Seller+product.MFD+product.UPC)
	product.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(product, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in saving the data"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (bc *Supplychain) AddBlock(data models.ProductCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
	}
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, models.ProductCheckout{IsGenesis: true})
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

func GetSupplychain(c *gin.Context) {
	jbytes, err := json.MarshalIndent(SupplyChain.Blocks, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, string(jbytes))
}

func WriteBlock(c *gin.Context) {
	var checkoutItem models.ProductCheckout
	if err := c.ShouldBindJSON(&checkoutItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write block"})
		return
	}

	SupplyChain.AddBlock(checkoutItem)
	resp, err := json.MarshalIndent(checkoutItem, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write block"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(b.Pos) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *Block, checkoutItem models.ProductCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}
