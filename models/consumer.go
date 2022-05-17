package models

type Consumer struct {
	ID       string
	IsActive bool
}

type Consumers []Consumer

func (c Consumers) ActiveConsumer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.IsActive {
			resp = append(resp, v)
		}
	}
	return resp
}
