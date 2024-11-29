package state

import (
	"github.com/hashicorp/go-memdb"
	"github.com/myback/oca/agent/configentry"
	"github.com/myback/oca/agent/structs"
)

// GetSamenessGroup returns a SamenessGroupConfigEntry from the state
// store using the provided parameters.
func (s *Store) GetSamenessGroup(ws memdb.WatchSet,
	name string,
	overrides map[configentry.KindName]structs.ConfigEntry,
	partition string) (uint64, *structs.SamenessGroupConfigEntry, error) {
	tx := s.db.ReadTxn()
	defer tx.Abort()

	return getSamenessGroupConfigEntryTxn(tx, ws, name, overrides, partition)
}
