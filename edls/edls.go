package main

import "time"

// file types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

// file extension
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)

// File struct
type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

// Style File
type styleFileType struct {
	icon   string
	color  string
	symbol string
}

//Mapa como llave el tipo de archivo para retornar una estructura

var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📜"},
	fileDirectory:  {icon: "📂", color: "BLUE", symbol: "/"},
	fileExecutable: {icon: "🚀", color: "GREEN", symbol: "*"},
	fileCompress:   {icon: "📦", color: "RED"},
	fileImage:      {icon: "📸", color: "MAGENTA"},
	fileLink:       {icon: "🔗", color: "CYAN"},
}

//me quedé en el minuto 9:40
