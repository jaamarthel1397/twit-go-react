package models

/* Tweet captura del Body, el mensaje que lleva */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
