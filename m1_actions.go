package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const  DEFAULT_CONFIG_PATH string = ".actions"

type menu struct {
  configPath string
  actions    map[string]string
}

func (menu *menu) loadMenu() {
  if menu == nil {
    return
  }
  config, err := os.Open(menu.configPath)
  if err != nil {
      log.Fatal(err)
  }
  defer config.Close()

  menu.actions = make(map[string]string)
  scanner := bufio.NewScanner(config)
  for scanner.Scan() {
      line := strings.Split(scanner.Text(),":")
      menu.actions[line[0]] = line[1]
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}

func main() {
  menu := menu{ configPath:DEFAULT_CONFIG_PATH }
  menu.loadMenu()
  if len(os.Args) == 1 {
    for name := range menu.actions {
      fmt.Println(name)
    }
  } else {
    target := os.Args[1]
    fmt.Println(strings.TrimSpace(menu.actions[target]))
  }
}
