package conf

import (
	"errors"
	"github.com/imroc/req"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/log"
	"net/url"
	"path"
	"strings"
)

func initConfigFile(temp Template, core *Core) error {
	//default config load
	if temp.AppYamlPath != "" {
		if i, err := initLoadViper(temp.AppYamlPath, core.viper); err != nil {
			return errors.New("Viper Load Error For: " + temp.AppYamlPath + " " + err.Error())
		} else {
			core.viper = i
			log.StdOut("Conf", "Viper.Main.Loaded", temp.AppYamlPath)
		}
	}
	//configs load
	if temp.AppYamlPathChildren != nil && len(temp.AppYamlPathChildren) > 0 {
		for k, v := range temp.AppYamlPathChildren {
			if i, err := initLoadViper(v, nil); err != nil {
				return errors.New("Viper Load Error For: " + v + " " + err.Error())
			} else {
				if core.viperChildren == nil {
					core.viperChildren = make(map[string]*viper.Viper, 0)
				}
				core.viperChildren[k] = i
				log.StdOut("Conf", "Viper.Child.Loaded", v)
			}
		}
	}
	return nil
}

func initLoadViper(yamlFilePath string, vip *viper.Viper) (*viper.Viper, error) {
	if vip == nil {
		vip = viper.New()
	}
	vip.SetConfigType("yaml") //set config file content type

	//remote content mode
	if uri, _ := url.Parse(yamlFilePath); uri != nil && uri.Scheme != "" {
		resp, err := req.Get(yamlFilePath)
		if err != nil {
			return vip, err
		}
		if err := vip.ReadConfig(resp.Response().Body); err != nil {
			return vip, err
		}
		return vip, nil
	}

	//local content mode
	if ext := path.Ext(yamlFilePath); ext != "" {
		dir, file := path.Split(yamlFilePath)
		if dir == "" {
			dir = "."
		}
		vip.AddConfigPath(dir)
		vip.SetConfigName(strings.Split(file, ext)[0])
	} else {
		vip.AddConfigPath(yamlFilePath)
		vip.SetConfigName("application")
	}
	if err := vip.ReadInConfig(); err != nil {
		return nil, err
	}
	return vip, nil
}
