package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	//"github.com/ha-wk/Assmnt1/models"
	//"fmt"
)

func create_hash() {
	str := Block_info.PrevBlkHash

	hmd5 := md5.Sum([]byte(str))
	hsha1 := sha1.Sum([]byte(str))
	hsha2 := sha256.Sum256([]byte(str))

	// fmt.Printf("   MD5: %x\n", hmd5)
	// fmt.Printf("  SHA1: %x\n", hsha1)
	// fmt.Printf("SHA256: %x\n", hsha2)
}
