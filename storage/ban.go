package storage

func (mgr *Manager) ensureBanCollection() (err error) {
	mgr.bans = mgr.db.C("bans")

	return
}
