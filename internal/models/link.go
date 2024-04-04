package models

type Links struct {
	shortened_links map[string]string
	links           map[string]string
}

func (l *Links) Init() {
	l.shortened_links = make(map[string]string)
	l.links = make(map[string]string)
}

func (l *Links) InsertLink(link string, shortenedLink string) {
	l.shortened_links[shortenedLink] = link
	l.links[link] = shortenedLink
}

func (l *Links) GetLink(shortenedLink string) (string, bool) {
	link, ok := l.shortened_links[shortenedLink]
	return link, ok
}

func (l *Links) GetShortenedLink(link string) (string, bool) {
	shortened_link, ok := l.links[link]
	return shortened_link, ok
}
