package server

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

var (
	playerID1 bson.ObjectId
	playerID2 bson.ObjectId
	playerID3 bson.ObjectId
)

func TestAdminPre(t *testing.T) {
	var err error
	playerID1, err = app.store.PlayerCreate(types.Player{
		Name:         "playerID1",
		Pass:         "74dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
		Ipv4:         1544996175,
		Alive:        &[]bool{true}[0],
		Registration: time.Now(),
		LastLogin:    time.Now(),
		TotalSpawns:  &[]int32{0}[0],
		Warnings:     &[]int32{0}[0],
		Gpci:         "b801a9f9553b892c4cda9219171a4f6d8c8b299a",
	})
	assert.NoError(t, err)
	playerID2, err = app.store.PlayerCreate(types.Player{
		Name:         "playerID2",
		Pass:         "74dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
		Ipv4:         1544996175,
		Alive:        &[]bool{true}[0],
		Registration: time.Now(),
		LastLogin:    time.Now(),
		TotalSpawns:  &[]int32{0}[0],
		Warnings:     &[]int32{0}[0],
		Gpci:         "b801a9f9553b892c4cda9219171a4f6d8c8b299a",
	})
	assert.NoError(t, err)
	playerID3, err = app.store.PlayerCreate(types.Player{
		Name:         "playerID3",
		Pass:         "74dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
		Ipv4:         1544996175,
		Alive:        &[]bool{true}[0],
		Registration: time.Now(),
		LastLogin:    time.Now(),
		TotalSpawns:  &[]int32{0}[0],
		Warnings:     &[]int32{0}[0],
		Gpci:         "b801a9f9553b892c4cda9219171a4f6d8c8b299a",
	})
	assert.NoError(t, err)
}

func TestApp_adminSet(t *testing.T) {
	tests := []struct {
		name       string
		body       types.Admin
		wantStatus types.Status
	}{
		{"v create 1", types.Admin{
			PlayerID: playerID1,
			Level:    &[]int32{1}[0],
		}, types.Status{
			Success: true,
			Message: "",
		}},
		{"v create 2", types.Admin{
			PlayerID: playerID2,
			Level:    &[]int32{2}[0],
		}, types.Status{
			Success: true,
			Message: "",
		}},
		{"v create 3", types.Admin{
			PlayerID: playerID3,
			Level:    &[]int32{3}[0],
		}, types.Status{
			Success: true,
			Message: "",
		}},
		{"v set 3 0", types.Admin{
			PlayerID: playerID3,
			Level:    &[]int32{0}[0],
		}, types.Status{
			Success: true,
			Message: "",
		}},
		{"v set 2 3", types.Admin{
			PlayerID: playerID2,
			Level:    &[]int32{3}[0],
		}, types.Status{
			Success: true,
			Message: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var status types.Status

			resp, err := client.R().
				SetBody(tt.body).
				SetResult(&status).
				Post("/store/adminSet")

			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, http.StatusOK, resp.StatusCode())

			if status.Success {
				assert.Equal(t, tt.wantStatus.Success, true)
				assert.Empty(t, status.Result)
				assert.Empty(t, status.Message)
			} else {
				assert.Equal(t, tt.wantStatus.Success, false)
				assert.Equal(t, tt.wantStatus.Message, status.Message)
				assert.Equal(t, tt.wantStatus.Result, status.Result)
			}
		})
	}
}

func TestApp_adminGetList(t *testing.T) {
	var status types.Status

	resp, err := client.R().
		SetResult(&status).
		Get("/store/adminGetList")

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, 200, resp.StatusCode())

	var admins []types.Admin
	b, _ := json.Marshal(status.Result)
	assert.NoError(t, json.Unmarshal(b, &admins))

	wantAdmins := []types.Admin{
		types.Admin{
			PlayerID: playerID1,
			Level:    &[]int32{1}[0],
		},
		types.Admin{
			PlayerID: playerID2,
			Level:    &[]int32{3}[0],
		},
	}

	assert.Equal(t, len(wantAdmins), len(admins))

	for i := range admins {
		assert.Equal(t, wantAdmins[i].PlayerID, admins[i].PlayerID)
		assert.Equal(t, wantAdmins[i].Level, admins[i].Level)
	}
}

func TestAdminPost(t *testing.T) {
	var err error

	err = app.store.PlayerRemove(playerID1)
	assert.NoError(t, err)
	err = app.store.PlayerRemove(playerID2)
	assert.NoError(t, err)
	err = app.store.PlayerRemove(playerID3)
	assert.NoError(t, err)
}
