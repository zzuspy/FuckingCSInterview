package driver

import (
	conf "../confParser"
)

type Value struct {
	key       interface{}
	value     interface{}
}

type Driver interface {
	InitDriver(conf conf.Conf)
	
	
	Set(key interface{}, value interface{})

	// get val by key
	// return 
	// bool: data valid or not
	// interface{}:val
	Get(key interface{}) (interface{}, bool)
}