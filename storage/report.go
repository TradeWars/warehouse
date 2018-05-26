package storage

import (
	"github.com/globalsign/mgo/bson"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (mgr *Manager) ensureReportCollection() (err error) {
	mgr.reports = mgr.db.C("reports")

	return
}

// ReportCreate creates a report in the database
func (mgr *Manager) ReportCreate(report types.Report) (id bson.ObjectId, err error) {
	report.ID = bson.NewObjectId()
	return report.ID, mgr.reports.Insert(report)
}

// ReportArchive sets archive status on a report
func (mgr *Manager) ReportArchive(id bson.ObjectId, archived bool) (err error) {
	return mgr.reports.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"archived": archived}})
}

// ReportGetList returns a list of reports based on search parameters
func (mgr *Manager) ReportGetList() (result []types.Report, err error) {
	err = mgr.reports.Find(bson.M{}).All(&result)
	return
}

// ReportGet returns a specific report given an id
func (mgr *Manager) ReportGet(id bson.ObjectId) (result types.Report, err error) {
	err = mgr.reports.Find(bson.M{"_id": id}).One(&result)
	return
}
