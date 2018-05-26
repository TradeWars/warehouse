package server

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

var (
	playerIDs = make(map[string]bson.ObjectId)
	timeTruth = time.Now().Truncate(time.Millisecond).UTC()
)

func Test_playerCreate(t *testing.T) {
	tests := []struct {
		name       string
		body       types.Player
		wantStatus types.Status
	}{
		{"v create 1", types.Player{
			Name:         "John",
			Pass:         "74dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
			Ipv4:         1544996175,
			Alive:        &[]bool{true}[0],
			Registration: timeTruth,
			LastLogin:    timeTruth,
			TotalSpawns:  &[]int32{0}[0],
			Warnings:     &[]int32{0}[0],
			Gpci:         "b801a9f9553b892c4cda9219171a4f6d8c8b299a",
		}, types.Status{
			Success: true,
			Message: "",
		}},
		{"v create 2", types.Player{
			Name:         "Alice",
			Pass:         "84dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
			Ipv4:         2544996175,
			Alive:        &[]bool{true}[0],
			Registration: timeTruth,
			LastLogin:    timeTruth,
			TotalSpawns:  &[]int32{0}[0],
			Warnings:     &[]int32{0}[0],
			Gpci:         "c801a9f9553b892c4cda9219171a4f6d8c8b299a",
		}, types.Status{
			Success: true,
			Message: "",
		}},
		{"i create dup", types.Player{
			Name:         "Alice",
			Pass:         "94dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
			Ipv4:         3544996175,
			Alive:        &[]bool{true}[0],
			Registration: timeTruth,
			LastLogin:    timeTruth,
			TotalSpawns:  &[]int32{0}[0],
			Warnings:     &[]int32{0}[0],
			Gpci:         "d801a9f9553b892c4cda9219171a4f6d8c8b299a",
		}, types.Status{
			Success: false,
			Message: "player name already registered",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var status types.Status

			resp, err := client.R().
				SetBody(tt.body).
				SetResult(&status).
				Post("/store/playerCreate")

			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, 200, resp.StatusCode())

			if status.Success {
				assert.Equal(t, tt.wantStatus.Success, true)
				assert.Len(t, status.Result, 24)
				assert.Empty(t, status.Message)

				playerIDs[tt.body.Name] = bson.ObjectIdHex(status.Result.(string))
			} else {
				assert.Equal(t, tt.wantStatus.Success, false)
				assert.Equal(t, tt.wantStatus.Message, status.Message)
				assert.Equal(t, tt.wantStatus.Result, status.Result)
			}
		})
	}
}

func Test_playerGetByName(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus types.Status
	}{
		{"v get 1", "John", types.Status{
			Result: types.Player{
				Name:         "John",
				Pass:         "74dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
				Ipv4:         1544996175,
				Alive:        &[]bool{true}[0],
				Registration: timeTruth,
				LastLogin:    timeTruth,
				TotalSpawns:  &[]int32{0}[0],
				Warnings:     &[]int32{0}[0],
				Gpci:         "b801a9f9553b892c4cda9219171a4f6d8c8b299a",
			},
			Success: true,
		}},
		{"v get 2", "Alice", types.Status{
			Result: types.Player{
				Name:         "Alice",
				Pass:         "84dfc2b27acfa364da55f93a5caee29ccad3557247eda238831b3e9bd931b01d77fe994e4f12b9d4cfa92a124461d2065197d8cf7f33fc88566da2db2a4d6eae",
				Ipv4:         2544996175,
				Alive:        &[]bool{true}[0],
				Registration: timeTruth,
				LastLogin:    timeTruth,
				TotalSpawns:  &[]int32{0}[0],
				Warnings:     &[]int32{0}[0],
				Gpci:         "c801a9f9553b892c4cda9219171a4f6d8c8b299a",
			},
			Success: true,
		}},
		{"i get 3", "Nobody", types.Status{
			Success: false,
			Message: "not found",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var status types.Status

			resp, err := client.R().
				SetQueryParam("name", tt.body).
				SetResult(&status).
				Get("/store/playerGet")

			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, 200, resp.StatusCode())

			if status.Success {
				assert.Equal(t, tt.wantStatus.Success, true)

				// can't simply infer type so quick JSON hack to coerce it
				var result types.Player
				b, _ := json.Marshal(status.Result)
				assert.NoError(t, json.Unmarshal(b, &result))

				// we can't know the ID beforehand because it's generated
				// so simply copy the ID from the creation map to the wantResult
				wantResult := tt.wantStatus.Result.(types.Player)
				wantResult.ID = playerIDs[tt.body]

				assert.Equal(t, wantResult, result)
				assert.Empty(t, status.Message)
			} else {
				assert.Equal(t, tt.wantStatus.Success, false)
				assert.Equal(t, tt.wantStatus.Message, status.Message)
				assert.Equal(t, tt.wantStatus.Result, status.Result)
			}
		})
	}
}

func Test_playerGetByID(t *testing.T) {

}

func Test_playerUpdate(t *testing.T) {

}
