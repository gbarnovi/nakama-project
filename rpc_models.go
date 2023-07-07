package main

type Request struct {
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

type Response struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}
