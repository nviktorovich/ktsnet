package configures

const (
	IPPattern                      = `[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}`
	DefaultSettingsPath            = "/bin/ktsnet/Default/Settings.yaml"
	NetInterfaceManualSettingsHelp = "- ручной выбор сетевого интерфейса для замены подсети."
	SubnetSettingsHelp             = "- настройка подсети. Введите номер подсети (1-255)."
	InterfacesPathManualHelp       = "- абсолютный путь до файла с описанием и настройками сетевых интерфейсов."
	NetworkingServiceCall          = "/etc/init.d/networking"
	RestartArg                     = "restart"
)
