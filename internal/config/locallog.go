package config

import "github.com/cihub/seelog"

const (
	seelogConfigFile = "conf/seelog.xml"
)

func initLocalLog(file string) error {

	fileName := seelogConfigFile

	if file != empty {
		fileName = file
	}

	logger, err := seelog.LoggerFromConfigAsFile(fileName)
	if err != nil {
		return err
	}

	err = seelog.UseLogger(logger)
	seelog.Flush()
	return nil
}

func init() {
	err := initLocalLog(empty)
	if err != nil {
		seelog.Infof("init seelog error %v", err)
	}
}
