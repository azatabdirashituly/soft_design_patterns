package main

func main() {
	lightBulb := getLightBulbInstance()

	onCommand := &onCommand{}
	onCommand.SetDevice(lightBulb)

	offCommand := &offCommand{}
	offCommand.SetDevice(lightBulb)

	oncmd := &remote{}
	oncmd.SetCommand(onCommand)
	oncmd.pressButton()

	offcmd := &remote{}
	offcmd.SetCommand(offCommand)
	offcmd.pressButton()
}
