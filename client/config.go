package main

type Version struct {
	Version string `json:"version"`
	Md5     string `json:"md5"`
	Sha1    string `json:"sha1"`
	Sha256  string `json:"sha256"`
	UpTime  string `json:"up_time"`
}
