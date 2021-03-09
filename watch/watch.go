package watch

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

type ConfFunc func(string) error

var registerConfMap = make(map[string]ConfFunc)

func GetRegisterConfMap() map[string]ConfFunc {
	return registerConfMap
}

func RegisterConf(filePath string, f ConfFunc) error {
	if registerConfMap == nil {
		registerConfMap = make(map[string]ConfFunc)
	}
	if _, ok := registerConfMap[filePath]; ok {
		return fmt.Errorf("filePath %s exist", filePath)
	}
	if len(filePath) <= 0 || f == nil {
		return fmt.Errorf("param err")
	}
	return nil
}

func perfromConfigFunc(filePaths []string) error {
	for _, v := range filePaths {
		if f, ok := registerConfMap[v]; ok {
			err := f(v)
			if err != nil {
				return err
			}
		} else {
			log.Printf("file %v not exist", v)
		}
	}
	return nil
}

func InitConf(filePaths []string) error {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalf("recover err: %+v", err)
			}
		}()
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			panic(fmt.Sprintf("NewWatcher err:%+v", err))
		}
		defer watcher.Close()

		err = watcher.Add("conf")
		if err != nil {
			panic(fmt.Sprintf("watcher.Add err:%v", err))
		}
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("event:%v", event)
				// event is create and remove when config is in k8s
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					log.Printf("modified file:%v", event.Name)
					err := perfromConfigFunc(filePaths)
					if err != nil {
						log.Fatalf("perfromConfigFunc err:%+v", err)
						return
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Fatalf("error:%+v", err)
			}
		}
	}()
	return perfromConfigFunc(filePaths)
}
