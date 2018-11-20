package sio

type Phase struct {
	name  string
	limit int
}

type Timeline struct {
	phases []Phase
}

// PushByDiff pushes a phase by difference.
func (t *Timeline) PushByDiff(name string, diff int) {
	total := t.phases[len(t.phases)-1].limit
	t.phases = append(t.phases, Phase{name, total + diff})
}

func (t Timeline) Phase(count int) string {
	for _, p := range t.phases {
		if count < p.limit {
			return p.name
		}
	}
	return ""
}
