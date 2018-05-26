package storage

import (
	"time"

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
func (mgr *Manager) ReportGetList(pageSize, page int, archived, noRead bool, by, of string, from, to *time.Time) (result []types.Report, err error) {
	query := bson.M{}

	// todo
	// if archived {
	// 	query["archived"] = true
	// }
	// if noRead {
	// 	query["read"] = false
	// }
	// if by != "" {
	// 	query["by"] = by
	// }
	// if of != "" {
	// 	query["name"] = of
	// }
	// if from != nil {
	// 	//
	// }
	// if to != nil {
	// 	//
	// }

	if pageSize > 0 {
		err = mgr.reports.Find(query).Skip(pageSize * page).Limit(pageSize).All(&result)
	} else {
		err = mgr.reports.Find(query).All(&result)
	}
	return
}

// ReportGet returns a specific report given an id
func (mgr *Manager) ReportGet(id bson.ObjectId) (result types.Report, err error) {
	err = mgr.reports.Find(bson.M{"_id": id}).One(&result)
	return
}
