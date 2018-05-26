package storage

func (mgr *Manager) ensureReportCollection() (err error) {
	mgr.reports = mgr.db.C("reports")

	return
}
