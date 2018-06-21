package types

// Geo represents a set of geographical data to describe a point in the game
// world. This can be a player, vehicle, item or anything else. It contains
// fields that are very commonly grouped together such as interior and velocity.
type Geo struct {
	PosX     float32 `json:"posx,omitempty" bson:"posx,omitempty"`
	PosY     float32 `json:"posy,omitempty" bson:"posy,omitempty"`
	PosZ     float32 `json:"posz,omitempty" bson:"posz,omitempty"`
	RotX     float32 `json:"rotx,omitempty" bson:"rotx,omitempty"`
	RotY     float32 `json:"roty,omitempty" bson:"roty,omitempty"`
	RotZ     float32 `json:"rotz,omitempty" bson:"rotz,omitempty"`
	RotW     float32 `json:"rotw,omitempty" bson:"rotw,omitempty"`
	VelX     float32 `json:"velx,omitempty" bson:"velx,omitempty"`
	VelY     float32 `json:"vely,omitempty" bson:"vely,omitempty"`
	VelZ     float32 `json:"velz,omitempty" bson:"velz,omitempty"`
	Interior int32   `json:"interior,omitempty" bson:"interior,omitempty"`
	World    int32   `json:"world,omitempty" bson:"world,omitempty"`
}
