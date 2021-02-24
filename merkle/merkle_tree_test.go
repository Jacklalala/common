package trie_test

import (
	"common/crypto/hash"
	"common/merkle"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	fmt.Println("test merkle tree")
	var hashes []hash.Hash
	hx, _ := hex.DecodeString("765116a5d819943920cec5dkjsc92jdgh2dfad475a406d730eb5680a856baf003")
	h := hash.NewHash(hx)
	hashes = append(hashes, h)
	merkleRoot, _ := trie.GetMerkleRoot(hashes)
	fmt.Println(merkleRoot.HexString())

	hx, _ = hex.DecodeString("12fbbbca6d41fa262e610e26af488f4ce9a8b7f9dd47025d03b5a33fdc7a0d66")
	h = hash.NewHash(hx)
	hashes = append(hashes, h)
	merkleRoot, _ = trie.GetMerkleRoot(hashes)
	fmt.Println(merkleRoot.HexString())

	hx, _ = hex.DecodeString("9fef5218557442a89ee05a736413f0e9a48cd97fab0d560b5709d8c820f6c2e2")
	h = hash.NewHash(hx)
	hashes = append(hashes, h)
	merkleRoot, _ = trie.GetMerkleRoot(hashes)
	fmt.Println(merkleRoot.HexString())

}
