package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/bitly/go-simplejson"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/stateleveldb"
)

func main() {
	dbProvider := stateleveldb.NewVersionedDBProvider1("./stateLeveldb")
	defer dbProvider.Close()
	db, _ := dbProvider.GetDBHandle("tbchannel")
	defer db.Close()
	iter, _ := db.GetStateRangeScanIterator("tradetrain", "", "")
	defer iter.Close()

	kv := make(map[string]interface{})
	for {
		queryResult, _ := iter.Next()
		if queryResult == nil {
			break
		}
		vkv := queryResult.(*statedb.VersionedKV)
		key := bytes.ReplaceAll(bytes.Trim([]byte(vkv.Key), "\x00"), []byte{0}, []byte{95})
		res, _ := simplejson.NewJson(vkv.Value)
		kv[string(key)] = res
	}
	jsonStr, _ := json.Marshal(kv)
	ioutil.WriteFile("kv.json", jsonStr, 0666)
}
