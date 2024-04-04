package models

type Memory interface {
	Init()
	InsertLink(link string, shortenedLink string)
	GetLink(shortenedLink string) (string, bool)
	GetShortenedLink(link string) (string, bool)
}
