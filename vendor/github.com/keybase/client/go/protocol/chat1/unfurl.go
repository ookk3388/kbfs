// Auto-generated by avdl-compiler v1.3.26 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/unfurl.avdl

package chat1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type UnfurlType int

const (
	UnfurlType_GENERIC UnfurlType = 0
	UnfurlType_YOUTUBE UnfurlType = 1
)

func (o UnfurlType) DeepCopy() UnfurlType { return o }

var UnfurlTypeMap = map[string]UnfurlType{
	"GENERIC": 0,
	"YOUTUBE": 1,
}

var UnfurlTypeRevMap = map[UnfurlType]string{
	0: "GENERIC",
	1: "YOUTUBE",
}

func (e UnfurlType) String() string {
	if v, ok := UnfurlTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type UnfurlGenericRaw struct {
	Title       string  `codec:"title" json:"title"`
	Url         string  `codec:"url" json:"url"`
	SiteName    string  `codec:"siteName" json:"siteName"`
	FaviconUrl  *string `codec:"faviconUrl,omitempty" json:"faviconUrl,omitempty"`
	ImageUrl    *string `codec:"imageUrl,omitempty" json:"imageUrl,omitempty"`
	PublishTime *int    `codec:"publishTime,omitempty" json:"publishTime,omitempty"`
	Description *string `codec:"description,omitempty" json:"description,omitempty"`
}

func (o UnfurlGenericRaw) DeepCopy() UnfurlGenericRaw {
	return UnfurlGenericRaw{
		Title:    o.Title,
		Url:      o.Url,
		SiteName: o.SiteName,
		FaviconUrl: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.FaviconUrl),
		ImageUrl: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ImageUrl),
		PublishTime: (func(x *int) *int {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.PublishTime),
		Description: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Description),
	}
}

type UnfurlYoutubeRaw struct {
}

func (o UnfurlYoutubeRaw) DeepCopy() UnfurlYoutubeRaw {
	return UnfurlYoutubeRaw{}
}

type UnfurlRaw struct {
	UnfurlType__ UnfurlType        `codec:"unfurlType" json:"unfurlType"`
	Generic__    *UnfurlGenericRaw `codec:"generic,omitempty" json:"generic,omitempty"`
	Youtube__    *UnfurlYoutubeRaw `codec:"youtube,omitempty" json:"youtube,omitempty"`
}

func (o *UnfurlRaw) UnfurlType() (ret UnfurlType, err error) {
	switch o.UnfurlType__ {
	case UnfurlType_GENERIC:
		if o.Generic__ == nil {
			err = errors.New("unexpected nil value for Generic__")
			return ret, err
		}
	case UnfurlType_YOUTUBE:
		if o.Youtube__ == nil {
			err = errors.New("unexpected nil value for Youtube__")
			return ret, err
		}
	}
	return o.UnfurlType__, nil
}

func (o UnfurlRaw) Generic() (res UnfurlGenericRaw) {
	if o.UnfurlType__ != UnfurlType_GENERIC {
		panic("wrong case accessed")
	}
	if o.Generic__ == nil {
		return
	}
	return *o.Generic__
}

func (o UnfurlRaw) Youtube() (res UnfurlYoutubeRaw) {
	if o.UnfurlType__ != UnfurlType_YOUTUBE {
		panic("wrong case accessed")
	}
	if o.Youtube__ == nil {
		return
	}
	return *o.Youtube__
}

func NewUnfurlRawWithGeneric(v UnfurlGenericRaw) UnfurlRaw {
	return UnfurlRaw{
		UnfurlType__: UnfurlType_GENERIC,
		Generic__:    &v,
	}
}

func NewUnfurlRawWithYoutube(v UnfurlYoutubeRaw) UnfurlRaw {
	return UnfurlRaw{
		UnfurlType__: UnfurlType_YOUTUBE,
		Youtube__:    &v,
	}
}

func (o UnfurlRaw) DeepCopy() UnfurlRaw {
	return UnfurlRaw{
		UnfurlType__: o.UnfurlType__.DeepCopy(),
		Generic__: (func(x *UnfurlGenericRaw) *UnfurlGenericRaw {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Generic__),
		Youtube__: (func(x *UnfurlYoutubeRaw) *UnfurlYoutubeRaw {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Youtube__),
	}
}

type UnfurlGeneric struct {
	Title       string  `codec:"title" json:"title"`
	Url         string  `codec:"url" json:"url"`
	SiteName    string  `codec:"siteName" json:"siteName"`
	Favicon     *Asset  `codec:"favicon,omitempty" json:"favicon,omitempty"`
	Image       *Asset  `codec:"image,omitempty" json:"image,omitempty"`
	PublishTime *int    `codec:"publishTime,omitempty" json:"publishTime,omitempty"`
	Description *string `codec:"description,omitempty" json:"description,omitempty"`
}

func (o UnfurlGeneric) DeepCopy() UnfurlGeneric {
	return UnfurlGeneric{
		Title:    o.Title,
		Url:      o.Url,
		SiteName: o.SiteName,
		Favicon: (func(x *Asset) *Asset {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Favicon),
		Image: (func(x *Asset) *Asset {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Image),
		PublishTime: (func(x *int) *int {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.PublishTime),
		Description: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Description),
	}
}

type UnfurlYoutube struct {
}

func (o UnfurlYoutube) DeepCopy() UnfurlYoutube {
	return UnfurlYoutube{}
}

type Unfurl struct {
	UnfurlType__ UnfurlType     `codec:"unfurlType" json:"unfurlType"`
	Generic__    *UnfurlGeneric `codec:"generic,omitempty" json:"generic,omitempty"`
	Youtube__    *UnfurlYoutube `codec:"youtube,omitempty" json:"youtube,omitempty"`
}

func (o *Unfurl) UnfurlType() (ret UnfurlType, err error) {
	switch o.UnfurlType__ {
	case UnfurlType_GENERIC:
		if o.Generic__ == nil {
			err = errors.New("unexpected nil value for Generic__")
			return ret, err
		}
	case UnfurlType_YOUTUBE:
		if o.Youtube__ == nil {
			err = errors.New("unexpected nil value for Youtube__")
			return ret, err
		}
	}
	return o.UnfurlType__, nil
}

func (o Unfurl) Generic() (res UnfurlGeneric) {
	if o.UnfurlType__ != UnfurlType_GENERIC {
		panic("wrong case accessed")
	}
	if o.Generic__ == nil {
		return
	}
	return *o.Generic__
}

func (o Unfurl) Youtube() (res UnfurlYoutube) {
	if o.UnfurlType__ != UnfurlType_YOUTUBE {
		panic("wrong case accessed")
	}
	if o.Youtube__ == nil {
		return
	}
	return *o.Youtube__
}

func NewUnfurlWithGeneric(v UnfurlGeneric) Unfurl {
	return Unfurl{
		UnfurlType__: UnfurlType_GENERIC,
		Generic__:    &v,
	}
}

func NewUnfurlWithYoutube(v UnfurlYoutube) Unfurl {
	return Unfurl{
		UnfurlType__: UnfurlType_YOUTUBE,
		Youtube__:    &v,
	}
}

func (o Unfurl) DeepCopy() Unfurl {
	return Unfurl{
		UnfurlType__: o.UnfurlType__.DeepCopy(),
		Generic__: (func(x *UnfurlGeneric) *UnfurlGeneric {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Generic__),
		Youtube__: (func(x *UnfurlYoutube) *UnfurlYoutube {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Youtube__),
	}
}

type UnfurlResult struct {
	Unfurl Unfurl `codec:"unfurl" json:"unfurl"`
	Url    string `codec:"url" json:"url"`
}

func (o UnfurlResult) DeepCopy() UnfurlResult {
	return UnfurlResult{
		Unfurl: o.Unfurl.DeepCopy(),
		Url:    o.Url,
	}
}

type UnfurlImageDisplay struct {
	Url    string `codec:"url" json:"url"`
	Height int    `codec:"height" json:"height"`
	Width  int    `codec:"width" json:"width"`
}

func (o UnfurlImageDisplay) DeepCopy() UnfurlImageDisplay {
	return UnfurlImageDisplay{
		Url:    o.Url,
		Height: o.Height,
		Width:  o.Width,
	}
}

type UnfurlGenericDisplay struct {
	Title       string              `codec:"title" json:"title"`
	Url         string              `codec:"url" json:"url"`
	SiteName    string              `codec:"siteName" json:"siteName"`
	Favicon     *UnfurlImageDisplay `codec:"favicon,omitempty" json:"favicon,omitempty"`
	Image       *UnfurlImageDisplay `codec:"image,omitempty" json:"image,omitempty"`
	PublishTime *int                `codec:"publishTime,omitempty" json:"publishTime,omitempty"`
	Description *string             `codec:"description,omitempty" json:"description,omitempty"`
}

func (o UnfurlGenericDisplay) DeepCopy() UnfurlGenericDisplay {
	return UnfurlGenericDisplay{
		Title:    o.Title,
		Url:      o.Url,
		SiteName: o.SiteName,
		Favicon: (func(x *UnfurlImageDisplay) *UnfurlImageDisplay {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Favicon),
		Image: (func(x *UnfurlImageDisplay) *UnfurlImageDisplay {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Image),
		PublishTime: (func(x *int) *int {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.PublishTime),
		Description: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Description),
	}
}

type UnfurlYoutubeDisplay struct {
}

func (o UnfurlYoutubeDisplay) DeepCopy() UnfurlYoutubeDisplay {
	return UnfurlYoutubeDisplay{}
}

type UnfurlDisplay struct {
	UnfurlType__ UnfurlType            `codec:"unfurlType" json:"unfurlType"`
	Generic__    *UnfurlGenericDisplay `codec:"generic,omitempty" json:"generic,omitempty"`
	Youtube__    *UnfurlYoutubeDisplay `codec:"youtube,omitempty" json:"youtube,omitempty"`
}

func (o *UnfurlDisplay) UnfurlType() (ret UnfurlType, err error) {
	switch o.UnfurlType__ {
	case UnfurlType_GENERIC:
		if o.Generic__ == nil {
			err = errors.New("unexpected nil value for Generic__")
			return ret, err
		}
	case UnfurlType_YOUTUBE:
		if o.Youtube__ == nil {
			err = errors.New("unexpected nil value for Youtube__")
			return ret, err
		}
	}
	return o.UnfurlType__, nil
}

func (o UnfurlDisplay) Generic() (res UnfurlGenericDisplay) {
	if o.UnfurlType__ != UnfurlType_GENERIC {
		panic("wrong case accessed")
	}
	if o.Generic__ == nil {
		return
	}
	return *o.Generic__
}

func (o UnfurlDisplay) Youtube() (res UnfurlYoutubeDisplay) {
	if o.UnfurlType__ != UnfurlType_YOUTUBE {
		panic("wrong case accessed")
	}
	if o.Youtube__ == nil {
		return
	}
	return *o.Youtube__
}

func NewUnfurlDisplayWithGeneric(v UnfurlGenericDisplay) UnfurlDisplay {
	return UnfurlDisplay{
		UnfurlType__: UnfurlType_GENERIC,
		Generic__:    &v,
	}
}

func NewUnfurlDisplayWithYoutube(v UnfurlYoutubeDisplay) UnfurlDisplay {
	return UnfurlDisplay{
		UnfurlType__: UnfurlType_YOUTUBE,
		Youtube__:    &v,
	}
}

func (o UnfurlDisplay) DeepCopy() UnfurlDisplay {
	return UnfurlDisplay{
		UnfurlType__: o.UnfurlType__.DeepCopy(),
		Generic__: (func(x *UnfurlGenericDisplay) *UnfurlGenericDisplay {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Generic__),
		Youtube__: (func(x *UnfurlYoutubeDisplay) *UnfurlYoutubeDisplay {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Youtube__),
	}
}

type UnfurlMode int

const (
	UnfurlMode_ALWAYS      UnfurlMode = 0
	UnfurlMode_NEVER       UnfurlMode = 1
	UnfurlMode_WHITELISTED UnfurlMode = 2
)

func (o UnfurlMode) DeepCopy() UnfurlMode { return o }

var UnfurlModeMap = map[string]UnfurlMode{
	"ALWAYS":      0,
	"NEVER":       1,
	"WHITELISTED": 2,
}

var UnfurlModeRevMap = map[UnfurlMode]string{
	0: "ALWAYS",
	1: "NEVER",
	2: "WHITELISTED",
}

func (e UnfurlMode) String() string {
	if v, ok := UnfurlModeRevMap[e]; ok {
		return v
	}
	return ""
}

type UnfurlSettings struct {
	Mode      UnfurlMode      `codec:"mode" json:"mode"`
	Whitelist map[string]bool `codec:"whitelist" json:"whitelist"`
}

func (o UnfurlSettings) DeepCopy() UnfurlSettings {
	return UnfurlSettings{
		Mode: o.Mode.DeepCopy(),
		Whitelist: (func(x map[string]bool) map[string]bool {
			if x == nil {
				return nil
			}
			ret := make(map[string]bool, len(x))
			for k, v := range x {
				kCopy := k
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Whitelist),
	}
}

type UnfurlInterface interface {
}

func UnfurlProtocol(i UnfurlInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.unfurl",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type UnfurlClient struct {
	Cli rpc.GenericClient
}
