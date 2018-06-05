package musicutil

import (
	"github.com/graphql-go/graphql"
)

var (
	OkComputer      Disc
	TheQueenIsDead  Disc
	BeHereNow       Disc
	AppetiteForDest Disc
	BackToBlack     Disc
	HotelCal        Disc

	DiscData map[int]Disc

	discType *graphql.Object

	MusicSchema graphql.Schema
)

type Disc struct {
	Id     string
	Title  string
	Artist string
	Year   int
}
