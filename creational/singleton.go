package creational

type counter struct {
	value int
}

func (c *counter) Inc() {
	c.value++
}

func (c *counter) Value() int {
	return c.value
}

var Counter = counter{}
