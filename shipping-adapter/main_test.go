package main

import (
	"fmt"
	"testing"

	"github.com/linkpoolio/bridges"
	"github.com/stretchr/testify/assert"
)

func TestGasStation_Run(t *testing.T) {
	wa := EAShipping{}

	h := bridges.NewHelper(&bridges.JSON{})
	val, err := wa.Run(h)

	//log.Info("TESTING ... %v", val)
	fmt.Printf("Payload %v", val)

	assert.Nil(t, err)
}

func TestGasStation_Opts(t *testing.T) {
	cc := EAShipping{}
	opts := cc.Opts()
	assert.Equal(t, opts.Name, "CustomEA")
	assert.True(t, opts.Lambda)
}
