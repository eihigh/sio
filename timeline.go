package sio

type Timeline struct {
	phases []phase
	index  int
	count  int
}

type phase struct {
	name   string
	length int
}

func (t *Timeline) Current() string {
	return t.phases[t.index].name
}

func (t *Timeline) Append(name string, length int) {
	t.phases = append(t.phases, phase{
		name:   name,
		length: length,
	})
}

func (t *Timeline) Update() {
	p := t.phases[t.index]
	if p.length == -1 {
		return
	}
	if t.count > p.length {
		t.index++
		t.count = 0
	}
	t.count++
}
