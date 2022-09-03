package cmd

import (
	"os"
)

type WebVtt struct {
	VttFile     *os.File      `json:"file"`
	VttElements []*VTTElement `json:"vtt_elements"`
	Header      *Header       `json:"header"`
}

func NewWebVtt() *WebVtt {
	return &WebVtt{}
}

type Header struct {
	Head string `json:"head"`
	Note string `json:"note"`
}

type VTTElement struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Position  string `json:"position"`
	Line      string `json:"line"`
	Text      string `json:"text"`
}

//Scan scan and bind one block of vtt file.
func (wv *WebVtt) Scan() *VTTElement {
	return nil
}

//AppendVttElement append VTTElement to WebVtt
func (wv *WebVtt) AppendVttElement(vtt *VTTElement) {
	wv.VttElements = append(wv.VttElements, vtt)
}

// NewVttElement initialize VTTElement
func (wv *WebVtt) NewVttElement() *VTTElement {
	return &VTTElement{}
}

//SkipHeader ignore header of vtt file.
func (wv *WebVtt) SkipHeader() *Header {
	return &Header{}
}
