package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/utrace-ltd/changelogger"
	"github.com/voxelbrain/goptions"
	"regexp"
)

func main()  {
	var opt Options
	goptions.ParseAndFail(&opt)

	logger := changelogger.NewChangeLogger()
	ver, err := logger.GetVersionChangeLog(opt.Version)
	if err != nil {
		panic(err)
	}

	var message string
	re := regexp.MustCompile(opt.Find)

	message = fmt.Sprintf("*%s - new version released %s*\n", opt.AppName, ver.Tag.Name)
	for _, gr := range ver.CommitGroups {
		message = message + fmt.Sprintf("_%s:_\n", gr.Title)
		for _, com := range gr.Commits {
			if opt.Find != "" && opt.Replace != "" {
				message = message + fmt.Sprintf("%s\n", re.ReplaceAll([]byte(com.Header), []byte(opt.Replace)))
			} else {
				message = message + com.Header + "\n"
			}
		}
	}

	msg := slack.WebhookMessage{
		Text: message,
	}

	err = slack.PostWebhook(opt.SlackHookUrl.String(), &msg)
	if err != nil {
		panic(err)
	}
}
