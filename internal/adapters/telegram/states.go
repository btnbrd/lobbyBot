package telegram

type UserState int

const (
	StateNone UserState = iota
	StateJoinWait
	StateShuffleWait
	StateAddWait
)

func (h *Handler) setState(userName string, state UserState) {
	h.userStates[userName] = state
}

func (h *Handler) getState(userName string) UserState {
	return h.userStates[userName]
}

func (h *Handler) clearState(userName string) {
	delete(h.userStates, userName)
}
