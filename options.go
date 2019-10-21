package main

import (
	"github.com/voxelbrain/goptions"
	"net/url"
)

type Options struct {
	AppName string `goptions:"-a, --app-name, obligatory, description='Application name'"`
	Version string `goptions:"-v, --version, obligatory, description='Version from get change log'"`
	SlackHookUrl *url.URL `goptions:"-s, --slack-hook-url, obligatory, description='Slack hook URL'"`
	Find string `goptions:"-f, --find, description='Find string like'"`
	Replace string `goptions:"-r, --replace, description='Replace string like'"`
	Help goptions.Help `goptions:"-h, --help, description='Show this help'"`
}
