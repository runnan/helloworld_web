package util
import (
    "github.com/creamdog/gonfig"
    "os"
)

func GetConfigurationValue(name string) string {
  f, _ := os.Open("config/myconfig.json")
  defer f.Close();
  config, _ := gonfig.FromJson(f)
  mydb, _ := config.GetString(name, nil)
  return mydb
}
