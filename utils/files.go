package utils

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	projectsMedia = "projects_media"
	deliveryMedia = "delivery_media"
)

type MediaDirs string

const (
	BaseProjectsMediaDir MediaDirs = "projects_media"
	DeliveryMediaDir     MediaDirs = "delivery_media"
	CommentsDir          MediaDirs = "invoice_comments"
)

// HomeDir returns the home directory, if found
func HomeDir() (hd string, err error) {
	usr, err := user.Current()
	if err != nil {
		return
	}
	hd = usr.HomeDir
	return hd, nil
}

// MediaDir returns the media directory
func MediaDir() (md string, err error) {
	homeDir, err := HomeDir()
	if err != nil {
		return
	}
	md = filepath.Join(homeDir, projectsMedia, deliveryMedia)
	err = os.MkdirAll(md, 0700)
	if err != nil {
		return
	}
	return
}

// GetDirectory gets the full path to the required directory
func GetDirectory(dirType MediaDirs) (dir string, err error) {
	hm, err := HomeDir()
	if err != nil {
		return
	}
	dir = filepath.Join(hm, projectsMedia, deliveryMedia, string(dirType))
	err = os.MkdirAll(dir, os.ModePerm)
	return
}

// MediaPath is the path to the actual file
func MediaPath(file string) (pth string, err error) {
	mDir, err := MediaDir()
	if err != nil {
		return
	}
	pth = filepath.Join(mDir, file)
	_, err = os.Stat(pth)
	if os.IsNotExist(err) {
		return "", err
	}
	return pth, nil
}
