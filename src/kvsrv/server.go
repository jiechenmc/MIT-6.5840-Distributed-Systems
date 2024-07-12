package kvsrv

import (
	"log"
	"sync"
)

const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}

type KVServer struct {
	mu    sync.Mutex
	kvmap map[string]string
}

func (kv *KVServer) Get(args *GetArgs, reply *GetReply) {
	kv.mu.Lock()
	i, ok := kv.kvmap[args.Key]

	if ok {
		reply.Value = i
	} else {
		reply.Value = ""
	}
	kv.mu.Unlock()
}

func (kv *KVServer) Put(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	kv.kvmap[args.Key] = args.Value
	kv.mu.Unlock()
}

func (kv *KVServer) Append(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	i, ok := kv.kvmap[args.Key]

	if ok {
		kv.kvmap[args.Key] = kv.kvmap[args.Key] + args.Value
		reply.Value = i
	} else {
		kv.kvmap[args.Key] = args.Value
		reply.Value = ""
	}
	kv.mu.Unlock()
}

func StartKVServer() *KVServer {
	kv := new(KVServer)
	kv.kvmap = make(map[string]string)
	return kv
}
