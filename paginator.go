package paginator

func NewByCount(page, perPage, count int) (Paginator, error) {
	p := Paginator{
		page:    page,
		perPage: perPage,
		count:   count,
	}

	p.calcPages()
	p.reload()

	return p, nil
}

func New(options ...func(*Paginator) error) (Paginator, error) {
	p := Paginator{
		page:    1,
		perPage: 10,
		count:   1,
		pages:   1,
		many:    3,
	}

	for _, option := range options {
		if err := option(&p); err != nil {
			return p, err
		}
	}

	p.reload()

	return p, nil
}

func Page(val int) func(p *Paginator) error {
	return func(p *Paginator) error {
		p.SetPage(val)
		return nil
	}
}

func PerPage(val int) func(p *Paginator) error {
	return func(p *Paginator) error {
		p.SetPerPage(val)
		return nil
	}
}

func Count(val int) func(p *Paginator) error {
	return func(p *Paginator) error {
		p.SetCount(val)
		return nil
	}
}

func Many(val int) func(p *Paginator) error {
	return func(p *Paginator) error {
		p.SetMany(val)
		return nil
	}
}
