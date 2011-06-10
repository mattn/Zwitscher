package main

import (
	"github.com/mattn/go-gtk/gdkpixbuf"
	"http"
	"json"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"gotter"
	//"github.com/garyburd/twister/oauth"
	"log"
)

type Tweet struct {
	Text       string
	Identifier string "id_str"
	Source     string
	CreatedAt  string "created_at"
	User       struct {
		Name               string
		ScreenName         string "screen_name"
		FollowersCount     int    "followers_count"
		ProfileImageURL    string "profile_image_url"
		ProfileImagePixbuf *gdkpixbuf.GdkPixbuf
	}
	Place *struct {
		Id       string
		FullName string "full_name"
	}
	Entities struct {
		HashTags []struct {
			Indices [2]int
			Text    string
		}
		UserMentions []struct {
			Indices    [2]int
			ScreenName string "screen_name"
		}    "user_mentions"
		Urls []struct {
			Indices [2]int
			Url     string
		}
	}
}

func Connect() Accounts{
	var account Accounts
	file, config := gotter.GetConfig()
	config["ClientToken"] = "lhCgJRAE1ECQzwVXfs5NQ"
	config["ClientSecret"] = "qk9i30vuzWHspsRttKsYrnoKSw9XBmWHdsis76z4"
	token, authorized, err := gotter.GetAccessToken(config)
	if err != nil {
		log.Fatal("faild to get access token:", err)
	}
	if authorized {
		b, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal("failed to store file:", err)
		}
		err = ioutil.WriteFile(file, b, 0700)
		if err != nil {
			log.Fatal("failed to store file:", err)
		}
	}
	account.Credentials = token
	return account
}

func UpdatePublicTimeline(callback func(tweet *Tweet)) {
	go func() {
		r, err := http.Get("http://twitter.com/statuses/public_timeline.json")
		if err == nil {
			var b []byte
			if r.ContentLength == -1 {
				b, err = ioutil.ReadAll(r.Body)
			} else {
				b = make([]byte, r.ContentLength)
				_, err = io.ReadFull(r.Body, b)
			}
			if err != nil {
				println(err.String())
				return
			}
			var tweets []Tweet
			json.NewDecoder(bytes.NewBuffer(b)).Decode(&tweets)
			for _, tweet := range tweets {
				tweet.User.ProfileImagePixbuf = url2pixbuf(tweet.User.ProfileImageURL)
				callback(&tweet)
			}
		}
	}()
}

func url2pixbuf(url string) *gdkpixbuf.GdkPixbuf {
	r, err := http.Get(url)
	if err != nil {
		return nil
	}
	t := r.Header.Get("Content-Type")
	b := make([]byte, r.ContentLength)
	if _, err = io.ReadFull(r.Body, b); err != nil {
		return nil
	}
	var loader *gdkpixbuf.GdkPixbufLoader
	if strings.Index(t, "jpeg") >= 0 {
		loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/jpeg")
	} else {
		loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/png")
	}
	loader.SetSize(24, 24)
	loader.Write(b)
	loader.Close()
	return loader.GetPixbuf()
}

func SendTweet(text string) {
	gotter.PostTweet(accounts.Credentials, "https://api.twitter.com/1/statuses/update.json", map[string]string{"status": text})
}

