package db

// Page represents a Wiki page
type Page struct {
	Tx       *Tx
	Name     []byte
	Text     []byte
	Created  []byte
	Modified []byte
}

func (p *Page) bucket() []byte {
	return []byte("pages")
}

func (p *Page) createdbucket() []byte {
	return []byte("created")
}

func (p *Page) modifiedbucket() []byte {
	return []byte("modified")
}

func (p *Page) get() ([]byte, error) {
	text := p.Tx.Bucket(p.bucket()).Get(p.Name)
	if text == nil {
		return nil, &Error{"page not found", nil}
	}
	return text, nil
}

func (p *Page) getCreated() ([]byte, error) {
	times := p.Tx.Bucket(p.createdbucket()).Get(p.Name)
	if times == nil {
		return nil, &Error{"page not found", nil}
	}
	return times, nil
}

func (p *Page) getModified() ([]byte, error) {
	times := p.Tx.Bucket(p.modifiedbucket()).Get(p.Name)
	if times == nil {
		return nil, &Error{"page not found", nil}
	}
	return times, nil
}

// Load retrieves a page from the database.
func (p *Page) Load() error {
	text, err := p.get()
	if err != nil {
		return err
	}

	p.Text = text

	created, err := p.getCreated()
	if err != nil {
		return err
	}

	p.Created = created

	modified, err := p.getModified()
	if err != nil {
		return err
	}

	p.Modified = modified

	return nil
}

// Save commits the Page to the database.
func (p *Page) Save() error {
	assert(len(p.Name) != 0, "uninitialized page cannot be saved")
	p.Tx.Bucket(p.createdbucket()).Put(p.Name, p.Created)
	p.Tx.Bucket(p.modifiedbucket()).Put(p.Name, p.Modified)
	return p.Tx.Bucket(p.bucket()).Put(p.Name, p.Text)
}
