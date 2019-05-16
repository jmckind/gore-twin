// Copyright 2019 gtrchive authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/url"
	"os"
	"runtime"

	"github.com/ChimeraCoder/anaconda"
)

var log = anaconda.BasicLogger

func main() {
	log.Debugf("Go Version: %s", runtime.Version())
	log.Debugf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)

	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("GTR_TWITTER_ACCESS_TOKEN"),
		os.Getenv("GTR_TWITTER_ACCESS_SECRET"),
		os.Getenv("GTR_TWITTER_CONSUMER_KEY"),
		os.Getenv("GTR_TWITTER_CONSUMER_SECRET"),
	)
	api.Log = log

	params := url.Values{
		"track": {os.Getenv("GTR_TWITTER_TRACK")},
	}

	rdbHost := os.Getenv("GTR_RETHINKDB_HOST")
	rdbPort := os.Getenv("GTR_RETHINKDB_PORT")
	rdbUsername := os.Getenv("GTR_RETHINKDB_USERNAME")
	rdbPassword := os.Getenv("GTR_RETHINKDB_PASSWORD")

	log.Debugf("RethinkDB Host: %s", rdbHost)
	log.Debugf("RethinkDB Port: %s", rdbPort)
	log.Debugf("RethinkDB Username: %s", rdbUsername)
	log.Debugf("RethinkDB Password: %s", rdbPassword)

	log.Debugf("Streaming tweets using params: %v", params)
	stream := api.PublicStreamFilter(params)

	for obj := range stream.C {
		switch o := obj.(type) {
		case anaconda.Tweet:
			log.Debugf("%-15s: %s", o.User.ScreenName, o.Text)
		}
	}
}
