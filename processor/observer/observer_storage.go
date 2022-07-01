package observer

import (
	"fmt"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

// ObserverStorage save the ondeck tx in item to key value store , in case bifrost restart
type ObserverStorage struct {
	db *leveldb.DB
}

const (
	OnDeckTxKey = "ondeck-tx"
)

// NewObserverStorage create a new instance of LevelDBScannerStorage
func NewObserverStorage(path string) (*ObserverStorage, error) {
	// no directory given, use in memory store
	if path == "" {
		storage := storage.NewMemStorage()
		db, err := leveldb.Open(storage, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to in memory open level db: %w", err)
		}
		return &ObserverStorage{db: db}, nil
	}
	levelDbFolder := filepath.Join(path, "observer")
	db, err := leveldb.OpenFile(levelDbFolder, nil)
	if err != nil {
		return nil, fmt.Errorf("fail to open level db %s: %w", levelDbFolder, err)
	}
	return &ObserverStorage{db: db}, nil
}

// // GetOnDeckTxs retrieve the ondeck tx from key value store
// func (s *ObserverStorage) GetOnDeckTxs() ([]types.TxIn, error) {
// 	buf, err := s.db.Get([]byte(OnDeckTxKey), nil)
// 	if err != nil {
// 		if errors.Is(err, leveldb.ErrNotFound) {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("fail to get ondeck tx from key value store: %w", err)
// 	}
// 	var result []types.TxIn
// 	if err := json.Unmarshal(buf, &result); err != nil {
// 		return nil, fmt.Errorf("fail to unmarshal ondeck tx: %w", err)
// 	}
// 	return result, nil
// }

// // SetOnDeckTxs save the ondeck tx to key value store
// func (s *ObserverStorage) SetOnDeckTxs(ondeck []types.TxIn) error {
// 	buf, err := json.Marshal(ondeck)
// 	if err != nil {
// 		return fmt.Errorf("fail to marshal ondeck tx to json: %w", err)
// 	}
// 	return s.db.Put([]byte(OnDeckTxKey), buf, nil)
// }

func (s *ObserverStorage) Close() error {
	return s.db.Close()
}
