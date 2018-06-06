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

func init() {
	OkComputer = Disc{
		Title:  "OK Computer",
		Artist: "Radiohead",
		Year:   1997,
		Id:     "1",
	}
	TheQueenIsDead = Disc{
		Title:  "The Queen is dead",
		Artist: "The Smiths",
		Year:   1986,
		Id:     "2",
	}
	BeHereNow = Disc{
		Title:  "Be Here Now",
		Artist: "Oasis",
		Year:   1997,
		Id:     "3",
	}
	AppetiteForDest = Disc{
		Title:  "Appetite for Destruction",
		Artist: "Guns N' Roses ",
		Year:   1987,
		Id:     "4",
	}
	BackToBlack = Disc{
		Title:  "Back To Black",
		Artist: "Amy Winehouse",
		Year:   2006,
		Id:     "5",
	}
	HotelCal = Disc{
		Title:  "Hotel California",
		Artist: "Eagles",
		Year:   1976,
		Id:     "6",
	}

	DiscData = map[int]Disc{
		1: OkComputer,
		2: TheQueenIsDead,
		3: BeHereNow,
		4: AppetiteForDest,
		5: BackToBlack,
		6: HotelCal,
	}

	discType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Disc",
		Description: "A set of songs from one or many artists.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The Identifier of the disc.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Id, nil
					}
					return nil, nil
				},
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "The Title of the album.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Title, nil
					}
					return nil, nil
				},
			},
			"artist": &graphql.Field{
				Type:        graphql.String,
				Description: "The Artist of the album.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Artist, nil
					}
					return nil, nil
				},
			},
			"year": &graphql.Field{
				Type:        graphql.Int,
				Description: "The release year.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Year, nil
					}
					return nil, nil
				},
			},
		},
	})
}

//GetAllDiscs return all the dummy data
func GetAllDiscs() []Disc {
	discs := []Disc{}
	for _, disc := range DiscData {
		discs = append(discs, disc)
	}
	return discs
}
