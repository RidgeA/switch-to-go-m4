package behavioral

type (
	command interface {
		execute()
	}

	door struct {
		isOpen bool
	}

	openDoorCommand struct {
		d door
	}

	closeDoorCommand struct {
		d door
	}
)

func (d *door) Open() {
	d.isOpen = true
}

func (d *door) Close() {
	d.isOpen = false
}

func (c openDoorCommand) execute() {
	c.d.Open()
}

func (c closeDoorCommand) execute() {
	c.d.Close()
}
