package paginator

import (
	"math"
)

type Paginator struct {
	page    int
	perPage int
	count   int

	pages int
	many  int

	next        int
	hasNext     bool
	hasNextMany bool

	prev        int
	hasPrev     bool
	hasPrevMany bool
}

func (p *Paginator) SetPage(val int) {
	p.page = val
	p.reload()
}

func (p *Paginator) SetPerPage(val int) {
	p.perPage = val
	p.calcPages()
}

func (p *Paginator) SetCount(val int) {
	p.count = val
	p.calcPages()
}

func (p *Paginator) SetMany(val int) {
	p.many = val
	p.reload()
}

func (p *Paginator) calcPages() {
	pages := float64(p.count) / float64(p.perPage)
	p.pages = int(math.Ceil(pages))
	p.reload()
}

func (p *Paginator) reload() {

	p.prev = 1 // defaults to 1
	p.hasPrev = p.page > 1
	if p.hasPrev {
		p.prev = p.page - 1
	}
	p.hasPrevMany = p.page-1 > p.many

	p.next = p.pages // defaults to the last page
	p.hasNext = p.page < p.pages
	if p.hasNext {
		p.next = p.page + 1
	}
	p.hasNextMany = p.pages-p.page > p.many

	return
}

func (p Paginator) Page() int {
	return p.page
}

func (p Paginator) PerPage() int {
	return p.perPage
}

func (p Paginator) Count() int {
	return p.count
}

func (p Paginator) Pages() int {
	return p.pages
}

func (p Paginator) Many() int {
	return p.many
}

func (p Paginator) Next() int {
	return p.next
}

func (p Paginator) HasNext() bool {
	return p.hasNext
}

func (p Paginator) HasNextMany() bool {
	return p.hasNextMany
}

func (p Paginator) Prev() int {
	return p.prev
}

func (p Paginator) HasPrev() bool {
	return p.hasPrev
}

func (p Paginator) HasPrevMany() bool {
	return p.hasPrevMany
}
