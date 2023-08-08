package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type KVArgs struct {
	K int
	V int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")   
    if err != nil {
        log.Fatal("dialing:", err)
    }
   
	// put
	var putReply1 string
	kv1 := &KVArgs {
		K: 101,
		V: 1000,
	}    
   
    err = client.Call("KVStoreService.Put", kv1, &putReply1)
    if err != nil {
        log.Fatal(err)
	}	

	var putReply2 string
	kv2 := &KVArgs {
		K: 102,
		V: 2000,
	}    
   
    err = client.Call("KVStoreService.Put", kv2, &putReply2)
    if err != nil {
        log.Fatal(err)
    }

	var putReply3 string
	kv3 := &KVArgs {
		K: 103,
		V: 3000,
	}    
   
    err = client.Call("KVStoreService.Put", kv3, &putReply3)
    if err != nil {
        log.Fatal(err)
    }
    
	// get with key: 103
	var getReply3 int
	err = client.Call("KVStoreService.Get", kv3.K, &getReply3)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("Get with k: %d, v: %d \n", kv3.K, getReply3)

	// count
	var countReply int
	err = client.Call("KVStoreService.Count", "Count", &countReply)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Number of records: %d\n", countReply)

	// del with key: 103
	var delReply bool
	err = client.Call("KVStoreService.Del", kv2.K, &delReply)
    if err != nil {
        log.Fatal(err)
    }

	if delReply {
		fmt.Printf("Found and deleted the key: %d", kv2.K)
	} else {
		fmt.Printf("Not found the key: %d", kv2.K)
	}
}