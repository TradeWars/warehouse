package goss

import "time"

// Account contains the core information for a player's account.
// This does not include character information such as inventory.
type Account struct {
	Name         string    `redis:"name,omitempty"`
	Pass         string    `redis:"pass"`
	Ipv4         string    `redis:"ipv4"`
	Alive        bool      `redis:"alive"`
	Registration time.Time `redis:"regdate"`
	LastLogin    time.Time `redis:"lastlog"`
	LastSpawn    time.Time `redis:"spawntime,omitempty"`
	TotalSpawns  int       `redis:"spawns"`
	Warnings     int       `redis:"warnings"`
	Gpci         string    `redis:"gpci"`
	Active       bool      `redis:"active"`
	Banned       bool      `redis:"banned"`
	AdminLevel   int       `redis:"admin"`
	Whitelisted  bool      `redis:"whitelisted"`
	Reported     bool      `redis:"reported"`
}
