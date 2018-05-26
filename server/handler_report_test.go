package server

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

var reportIDs = make(map[string]bson.ObjectId)

func Test_reportCreate(t *testing.T) {
	tests := []struct {
		name       string
		body       types.Report
		wantStatus types.Status
	}{
		{"v create 1", types.Report{
			Name:     "John",
			Reason:   "Hacking",
			By:       playerIDs["Alice"],
			Date:     timeTruth,
			Read:     &[]bool{false}[0],
			Type:     "AC",
			Position: types.Geo{PosX: 20.0, PosY: 55.0, PosZ: 12.0},
			Metadata: "135h",
			Archived: &[]bool{false}[0],
		}, types.Status{
			Success: true,
		}},
		{"v create 2", types.Report{
			Name:     "Steve",
			Reason:   "Ban evasion",
			By:       playerIDs["Alice"],
			Date:     timeTruth.Add(-time.Hour),
			Read:     &[]bool{true}[0],
			Type:     "PLY",
			Position: types.Geo{},
			Metadata: "",
			Archived: &[]bool{true}[0],
		}, types.Status{
			Success: true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var status types.Status

			resp, err := client.R().
				SetBody(tt.body).
				SetResult(&status).
				Post("/store/reportCreate")

			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, 200, resp.StatusCode())

			if status.Success {
				assert.Equal(t, tt.wantStatus.Success, true)
				assert.Len(t, status.Result, 24)
				assert.Empty(t, status.Message)

				reportIDs[tt.body.Name] = bson.ObjectIdHex(status.Result.(string))
			} else {
				assert.Equal(t, tt.wantStatus.Success, false)
				assert.Equal(t, tt.wantStatus.Message, status.Message)
				assert.Equal(t, tt.wantStatus.Result, status.Result)
			}
		})
	}
}

func Test_reportArchive(t *testing.T) {
	tests := []struct {
		name       string
		body       bson.ObjectId
		wantStatus types.Status
	}{
		{"v archive 1", reportIDs["John"], types.Status{
			Success: true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var status types.Status

			resp, err := client.R().
				SetQueryParam("id", tt.body.Hex()).
				SetQueryParam("archive", "true").
				SetResult(&status).
				Patch("/store/reportArchive")

			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, 200, resp.StatusCode())

			if status.Success {
				assert.Equal(t, tt.wantStatus.Success, true)
				assert.Empty(t, status.Result)
				assert.Empty(t, status.Message)

				resp, err := client.R().
					SetQueryParam("id", tt.body.Hex()).
					SetResult(&status).
					Get("/store/reportGet")
				if err != nil {
					t.Error(err)
					return
				}
				assert.Equal(t, 200, resp.StatusCode())
				assert.Equal(t, status.Success, true)

				var gotReport types.Report
				b, _ := json.Marshal(status.Result)
				assert.NoError(t, json.Unmarshal(b, &gotReport))
				assert.True(t, *gotReport.Archived)
			} else {
				assert.Equal(t, tt.wantStatus.Success, false)
				assert.Equal(t, tt.wantStatus.Message, status.Message)
				assert.Equal(t, tt.wantStatus.Result, status.Result)
			}
		})
	}
}

func Test_reportGetList(t *testing.T) {
	tests := []struct {
		name        string
		body        map[string]string
		wantStatus  types.Status
		wantReports []types.Report
	}{
		{"v archive 1", map[string]string{
			"archived": "true",
		}, types.Status{
			Success: true,
		}, []types.Report{
			types.Report{
				ID:       reportIDs["John"],
				Name:     "John",
				Reason:   "Hacking",
				By:       playerIDs["Alice"],
				Date:     timeTruth,
				Read:     &[]bool{false}[0],
				Type:     "AC",
				Position: types.Geo{PosX: 20.0, PosY: 55.0, PosZ: 12.0},
				Metadata: "135h",
				Archived: &[]bool{true}[0],
			},
			types.Report{
				ID:       reportIDs["Steve"],
				Name:     "Steve",
				Reason:   "Ban evasion",
				By:       playerIDs["Alice"],
				Date:     timeTruth.Add(-time.Hour),
				Read:     &[]bool{true}[0],
				Type:     "PLY",
				Position: types.Geo{},
				Metadata: "",
				Archived: &[]bool{true}[0],
			},
		}},
	}
	for _, tt := range tests {
		var status types.Status
		resp, err := client.R().
			SetQueryParams(tt.body).
			SetResult(&status).
			Get("/store/reportGetList")
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, 200, resp.StatusCode())

		var reports []types.Report
		b, _ := json.Marshal(status.Result)
		assert.NoError(t, json.Unmarshal(b, &reports))

		assert.Equal(t, tt.wantReports, reports)
	}
}

func Test_reportGet(t *testing.T) {

}
