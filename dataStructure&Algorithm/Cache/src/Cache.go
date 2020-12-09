package Cache

import (
	driver "./driver"
	conf "./confParser"
)


type Cache struct {
	driver  driver.Driver
	conf    *conf.Conf
}

func New(confPath string, Name string) *Cache {
	c := &Cache{}
	c.conf = conf.New(confPath, Name)
	driverType := c.conf.GetDriverType()
	switch driverType {
	case "Lru":
		c.driver = &driver.LruDriver{}
	default:
		c.driver = &driver.LruDriver{}
	}
	return c
}
