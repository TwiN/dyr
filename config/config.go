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
	// Banner to include before
	Banner string `yaml:"banner"`
}

func LoadConfiguration() error {
	// Find home directory
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	// Configure Viper
	viper.New()
	viper.AddConfigPath(fmt.Sprintf("%s/%s", home, ConfigurationDirPathRelativeToUserHome))
	viper.SetConfigName("dyr")
	viper.SetConfigType("yaml")

	// Read the existing configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file missing or corrupted: ", err)
		fmt.Println("Recreating configuration file")

		// Create default configuration
		config := buildDefaultConfiguration()
		viper.SetDefault("banner", config.Banner)

		// Check if ~/.config exists and creates it if it doesn't
		err = createDirIfNotExists(fmt.Sprintf("%s/%s", home, ConfigurationDirPathRelativeToUserHome))
		if err != nil {
			return err
		}

		// Save current configuration to file only if the file does not exist
		viper.SafeWriteConfig()
		if err := viper.SafeWriteConfigAs(fmt.Sprintf("%s/%s/dyr.yaml", home, ConfigurationDirPathRelativeToUserHome)); err != nil {
			if os.IsNotExist(err) {
				err = viper.WriteConfigAs(fmt.Sprintf("%s/%s/dyr.yaml", home, ConfigurationDirPathRelativeToUserHome))
			}
		}
	}

	// If there are new properties, (i.e. when the user updates rubicon), write these to the file
	viper.WriteConfig()
	return err
}

func buildDefaultConfiguration() *Configuration {
	return &Configuration{}
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
