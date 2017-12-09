package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const  DEFAULT_CONFIG_PATH string = ".m1"

type menu struct {
  configPath string
  actions    []action
}

type action struct {
  name    string
  command string
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

  menu.actions = make([]action, 0)
  scanner := bufio.NewScanner(config)
  for scanner.Scan() {
      line := strings.Split(scanner.Text(),":")
      menu.actions = append(menu.actions, action{
        name:    line[0],
        command: line[1] })
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}

func main() {
  menu := menu{configPath:DEFAULT_CONFIG_PATH}
  menu.loadMenu()
  switch os.Args[1] {
  case "actions":
    for _, action := range menu.actions {
      fmt.Println(action.name)
    }
  case "run":
    target := os.Args[2]
    for _, action := range menu.actions {
      if action.name == target {
        fmt.Println(action.command)
      }
    }
  }
}
