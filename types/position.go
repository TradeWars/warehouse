package types

// Geo represents a set of geographical data to describe a point in the game
// world. This can be a player, vehicle, item or anything else. It contains
// fields that are very commonly grouped together such as interior and velocity.
type Geo struct {
	PosX     float32 `json:"posx"`
	PosY     float32 `json:"posy"`
	PosZ     float32 `json:"posz"`
	RotX     float32 `json:"rotx"`
	RotY     float32 `json:"roty"`
	RotZ     float32 `json:"rotz"`
	RotW     float32 `json:"rotw"`
	VelX     float32 `json:"velx"`
	VelY     float32 `json:"vely"`
	VelZ     float32 `json:"velz"`
	Interior int32   `json:"interior"`
	World    int32   `json:"world"`
}
