package main

import (
	"fmt"
	"sync"
)

type Command interface {
	execute()
}

type Device interface {
	on()
	off()
}

/* ----------------- on command ----------------- */
type onCommand struct {
	device Device
}

func (r *onCommand) SetDevice(device Device) {
	r.device = device
}
func (loc onCommand) execute() {
	loc.device.on()
}

/* ----------------- off command ----------------- */
type offCommand struct {
	device Device
}

func (r *offCommand) SetDevice(device Device) {
	r.device = device
}
func (loc offCommand) execute() {
	loc.device.off()
}

/* ----------------- concrete object ----------------- */
type lightBulb struct{}

func (lb lightBulb) on() {
	fmt.Println("Light bulb on")
}
func (lb lightBulb) off() {
	fmt.Println("Light bulb off")
}

var instance *lightBulb
var once sync.Once

func getLightBulbInstance() *lightBulb {
	once.Do(func() {
		instance = &lightBulb{}
	})
	return instance
}

/* ----------------- concrete object implementation ----------------- */
type remote struct {
	cmd Command
}

func (r *remote) SetCommand(cmd Command) {
	r.cmd = cmd
}
func (r *remote) pressButton() {
	r.cmd.execute()
}
