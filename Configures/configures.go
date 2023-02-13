package configures

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	FilePath      string `yaml:"file_path"`
	Interface     string `yaml:"interface"`
	SubnetDefault string `yaml:"default_subnet"`
}

func getConfig() (*config, error) {
	data, err := os.ReadFile(DefaultSettingsPath)
	if err != nil {
		return nil, err
	}

	cfg := new(config)
	if err = yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// CheckUserCase функция, для анализа пользовательского ввода.
// Возвращает вводные данные, либо по умолчанию, либо, переданные вручную.
func CheckUserCase() (string, string, string) {
	var (
		netInterface   string // переменная содержит название сетевого интерфейса, в котором будет заменена подсеть
		subnet         string // пересенная содержит значение новой подсети, которую необходимо выставить в настройках интерфейса
		interfacesPath string // переменная содержит путь до файла с описанием и настройкой сетевых интерфейсов
	)

	cfg, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&netInterface, "n", cfg.Interface, NetInterfaceManualSettingsHelp)
	flag.StringVar(&subnet, "s", cfg.SubnetDefault, SubnetSettingsHelp)
	flag.StringVar(&interfacesPath, "p", cfg.FilePath, InterfacesPathManualHelp)
	flag.Parse()

	return subnet, netInterface, interfacesPath
}
