package main

/* ------------------------------- Imports --------------------------- */

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"

	"github.com/marcelloh/fastdb"
)

type KVStoreService struct {
	Store *fastdb.DB
}

type KVArgs struct {
	K int
	V int
}

// Put
func (p *KVStoreService) Put(ars *KVArgs, reply *string) error {	
	value := []byte(strconv.Itoa(ars.V))
	p.Store.Set("KVStore", ars.K, value)
	fmt.Printf("Put with k: %d, v: %s \n", ars.K, string(value))
	return nil
}

// Get
func (p *KVStoreService) Get(k int, reply *int) error {
	value, ok := p.Store.Get("KVStore", k)
	if !ok {
		fmt.Println("Error during read the data")
	}

	tmp := string(value)
	val, err := strconv.Atoi(tmp)

	if err != nil {
		fmt.Println("Error during conversion")		
	}

	*reply = val
	
	return nil
}

// Del
func (p *KVStoreService) Del(k int, reply *bool) error {
	ok, err := p.Store.Del("KVStore", k)
	if err != nil {
		log.Fatal(err)
	}

	*reply = ok
	return nil
}

// Count
func (p *KVStoreService) Count(k string, reply *int) error {
	dbRecords, err := p.Store.GetAll("KVStore")
	if err != nil {
		log.Fatal(err)
	}

	*reply = len(dbRecords)
	return nil
}

/* -------------------------- Methods/Functions ---------------------- */

func main() {
	store, err := fastdb.Open(":memory:", 100)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = store.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	kvStoreSrv := &KVStoreService {
		Store: store,
	}

	rpc.RegisterName("KVStoreService", kvStoreSrv)
	listener, err := net.Listen("tcp", "localhost:1234")
	fmt.Println("Listening at localhost:1234")

	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
