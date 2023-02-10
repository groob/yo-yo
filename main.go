package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/groob/plist"
)

var notificationDir = flag.String("n", "", "path to Yo notification files")

var yoPath = flag.String("p", "/usr/local/bin/yo", "path to Yo executable")

type yoOpts struct {
	Title          string `json:"title" plist:"title"`
	Subtitle       string `json:"subtitle" plist:"subtitle"`
	Info           string `json:"info" plist:"info"`
	ActionBtn      string `json:"action_button" plist:"action_button"`
	ActionPath     string `json:"action_path" plist:"action_path"`
	BashAction     string `json:"bash_action" plist:"bash_action"`
	OtherBtn       string `json:"other_button" plist:"other_button"`
	Icon           string `json:"icon" plist:"icon"`
	ContentImage   string `json:"content_image" plist:"content_image"`
	DeliverySound  string `json:"delivery_sound" plist:"delivery_sound"`
	IgnoreDND      bool   `json:"ignore_dnd" plist:"ignore_dnd"`
	LockScreenOnly bool   `json:"lockscreen_only" plist:"lockscreen_only"`
	Poof           bool   `json:"poof_on_cancel" plist:"poof_on_cancel"`
	BannerMode     bool   `json:"banner_mode" plist:"banner_mode"`
}

func (yo yoOpts) toStringArray() []string {
	arr := []string{}

	if yo.Title != "" {
		arr = append(arr, "--title", yo.Title)
	}

	if yo.Subtitle != "" {
		arr = append(arr, "--subtitle", yo.Subtitle)
	}

	if yo.Info != "" {
		arr = append(arr, "--info", yo.Info)
	}

	if yo.ActionBtn != "" {
		arr = append(arr, "--action-btn", yo.ActionBtn)
	}

	if yo.ActionPath != "" {
		arr = append(arr, "--action-path", yo.ActionPath)
	}

	if yo.BashAction != "" {
		arr = append(arr, "--bash-action", yo.BashAction)
	}

	if yo.OtherBtn != "" {
		arr = append(arr, "--other-btn", yo.OtherBtn)
	}

	if yo.Icon != "" {
		arr = append(arr, "--icon", yo.Icon)
	}

	if yo.ContentImage != "" {
		arr = append(arr, "--content-image", yo.ContentImage)
	}

	if yo.DeliverySound != "" {
		arr = append(arr, "--delivery-sound", yo.DeliverySound)
	}

	if yo.IgnoreDND {
		arr = append(arr, "--ignores-do-not-disturb")
	}

	if yo.LockScreenOnly {
		arr = append(arr, "--lockscreen-only")
	}

	if yo.Poof {
		arr = append(arr, "--poofs-on-cancel")
	}

	if yo.BannerMode {
		arr = append(arr, "--banner-mode")
	}

	return arr
}

func (yo yoOpts) run() error {

	cmd := exec.Command(*yoPath, yo.toStringArray()...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(stdout.String())
		return err
	}
	fmt.Printf(stderr.String())
	return nil
}

func fromJson(path string) (*yoOpts, error) {
	yo := &yoOpts{}
	file, err := os.Open(path)
	if err != nil {
		return yo, err
	}
	defer file.Close()

	return yo, json.NewDecoder(file).Decode(&yo)
}

func fromPlist(path string) (*yoOpts, error) {
	yo := &yoOpts{}
	file, err := os.Open(path)
	if err != nil {
		return yo, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return yo, err
	}
	return yo, plist.Unmarshal(data, &yo)
}

func init() {
	flag.Parse()
	if *notificationDir == "" {
		log.Println(errors.New("yo-trigger needs a path to notification directory"))
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	files, err := ioutil.ReadDir(*notificationDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		path := filepath.Join(*notificationDir, file.Name())
		yo := &yoOpts{}
		// file is json
		if filepath.Ext(file.Name()) == ".json" {
			yo, err = fromJson(path)
			if err != nil {
				log.Println(err)
				return
			}
		}

		// file is a plist
		if filepath.Ext(file.Name()) == ".plist" {
			yo, err = fromPlist(path)
			if err != nil {
				log.Println(err)
				return
			}
		}
		// run yo
		err = yo.run()
		if err != nil {
			log.Println(err)
			return
		}

		// on success remove file
		err = os.Remove(path)
		if err != nil {
			log.Println(err)
		}
	}
}
