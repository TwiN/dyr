package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

const (
	ConfigurationDirPathRelativeToUserHome = ".config/dyr"
)

type Configuration struct {
	// Banner to include before printing a note
	Banner string `yaml:"banner"`
}

func LoadConfiguration() (err error) {
	// Configure Viper
	viper.New()
	viper.AddConfigPath(GetConfigDir())
	viper.SetConfigName("dyr")
	viper.SetConfigType("yaml")

	// Read the existing configuration
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Config file missing or corrupted: ", err)
		fmt.Println("Recreating configuration file")

		// Create default configuration
		config := buildDefaultConfiguration()
		viper.SetDefault("banner", config.Banner)

		// Check if ~/.config exists and creates it if it doesn't
		err = createDirIfNotExists(GetConfigDir())
		if err != nil {
			return err
		}

		// Save current configuration to file only if the file does not exist
		viper.SafeWriteConfig()
		if err := viper.SafeWriteConfigAs(fmt.Sprintf("%s/dyr.yaml", GetConfigDir())); err != nil {
			if os.IsNotExist(err) {
				err = viper.WriteConfigAs(fmt.Sprintf("%s/dyr.yaml", GetConfigDir()))
			}
		}
	}

	// If there are new properties, (i.e. when the user updates rubicon), write these to the file
	viper.WriteConfig()
	return err
}

func buildDefaultConfiguration() *Configuration {
	return &Configuration{
		Banner: "",
	}
}

func createDirIfNotExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create %s folder: %s", path, err.Error())
		}
	}
	return err
}

func GetConfigDir() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s", home, ConfigurationDirPathRelativeToUserHome)
}
