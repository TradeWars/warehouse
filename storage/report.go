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
func (mgr *Manager) ReportGetList(pageSize, page int, archived, noRead bool, by, of bson.ObjectId, from, to *time.Time) (result []types.Report, err error) {
	var and []bson.M

	if archived {
		and = append(and, bson.M{"archived": true})
	} else {
		and = append(and, bson.M{"archived": false})
	}
	if noRead {
		and = append(and, bson.M{"read": false})
	}
	if by != "" {
		and = append(and, bson.M{"by_player_id": by})
	}
	if of != "" {
		and = append(and, bson.M{"of_player_id": of})
	}
	if from != nil {
		and = append(and, bson.M{"_id": bson.M{"$gte": bson.NewObjectIdWithTime(*from)}})
	}
	if to != nil {
		and = append(and, bson.M{"_id": bson.M{"$lte": bson.NewObjectIdWithTime(*to)}})
	}

	query := bson.M{}

	if len(and) == 1 {
		query = and[0]
	} else if len(and) > 1 {
		query = bson.M{"$and": and}
	}

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
