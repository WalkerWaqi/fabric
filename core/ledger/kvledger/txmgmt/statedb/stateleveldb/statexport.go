package stateleveldb

import (
	"github.com/hyperledger/fabric/common/ledger/util/leveldbhelper"
)

func NewVersionedDBProvider1(dbPath string) *VersionedDBProvider {
	dbProvider := leveldbhelper.NewProvider(&leveldbhelper.Conf{DBPath: dbPath})
	return &VersionedDBProvider{dbProvider}
}
