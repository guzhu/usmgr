package app

func (app *App) initUI() {
	app.initHead()
}
func (app *App) initHead() {
	// update basic information
	hostname, kernel, ostype := "ucloudstor","3.10.0-957.1.3.el7.x86_64","Linux"
	app.head.UpdateBasicInfo(hostname, kernel, ostype)

}