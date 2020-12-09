package confParser

type Conf struct {
	maxKeyCnt  int
	driverType string
}

const defaultSize = 10

// undo read conf
func New(path string, name string) *Conf {
	c := &Conf{
		maxKeyCnt:  defaultSize,
		driverType: "Lru",
	}
	return c
}

func (c *Conf) GetDriverType() string {
	return c.driverType
}

func (c *Conf) GetMaxKeyCnt() int {
	return c.maxKeyCnt
}
