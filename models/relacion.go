package models

type Relacion struct {
	UsuarioID        string `bson: "usuarioid" json:"usuarioId"`
	UsurioRelacionID string `bson: "usuariorelacionid" json:"usuarioRelacionId"`
}
