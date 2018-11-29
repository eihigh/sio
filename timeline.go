package sio

type phase struct {
	end   int
	state int
}

type Timeline struct {
	phases []phase
	state  State
}

func (t *Timeline) AppendByDiff(state, diff int) {
	total := 0
	if len(t.phases) > 0 {
		total = t.phases[len(t.phases)-1].end
	}
	t.phases = append(t.phases, phase{
		end:   total + diff,
		state: state,
	})
}

func (t *Timeline) Current() int {
	for _, p := range t.phases {
		if t.state.Count() < p.end {
			return p.state
		}
	}
	return 0
}
