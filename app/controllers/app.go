package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"io/ioutil"
	"regexp"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Upload(torrent []byte, folder string) revel.Result {
	c.Validation.Required(torrent)
	c.Validation.MaxSize(torrent, 2*MB).
		Message("File cannot be larger than 2MB")

	c.Validation.Required(folder)
	c.Validation.Match(folder, regexp.MustCompile("(movies|music|tv|misc)"))

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect((*App).Index)
	}

	err := ioutil.WriteFile(fmt.Sprintf("/watch/%s/%s", folder, c.Params.Files["torrent"][0].Filename), torrent, 0644)
	if err != nil {
        panic(err)
    }

	c.Flash.Success(fmt.Sprintf("Successfully uploaded to %s!", folder))
	return c.Redirect((*App).Index)
}
