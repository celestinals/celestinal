/*
 * Copyright 2025 The Tickex Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package notify provide functions log by telegram bot
package notify

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"github.com/tickexvn/tickex/pkg/flag"
	"github.com/tickexvn/tickex/pkg/protobuf"
)

var _ IRobot = (*Robot)(nil)

// New Telegram bot
func New(conf *tickex.Config) (IRobot, error) {
	bot, err := tgbotapi.NewBotAPI(conf.GetBotToken())
	if err != nil {
		return &Robot{}, err
	}

	return &Robot{
		bot:  bot,
		conf: conf,
	}, nil
}

// IRobot telegram bot interface
type IRobot interface {
	Send(msg *tickex.TelegramMessage) error
}

// Robot Telegram bot
type Robot struct {
	bot  *tgbotapi.BotAPI
	conf *tickex.Config
}

// Send notify to group telegram
func (r *Robot) Send(msg *tickex.TelegramMessage) error {
	flags := flag.ParseEdge()
	if !flags.IsTurnOnBots {
		return nil
	}

	msgText := fmt.Sprintf(
		"*Tickex Message*\n\n"+
			"*Created At:* `%s`\n"+
			"*Author:* `%s`\n\n"+
			"`%s`\n\n"+
			"*body*\n```%s```\n\n"+
			"`%s`\n",
		protobuf.FromTime(msg.Metadata.CreatedAt).String(),
		msg.Metadata.Author,
		msg.Header,
		msg.Body,
		msg.Footer,
	)

	// Send to Telegram
	mdv2 := tgbotapi.NewMessage(r.conf.GetChatId(), msgText)
	mdv2.ParseMode = "MarkdownV2"

	_, err := r.bot.Send(mdv2)
	return err
}
