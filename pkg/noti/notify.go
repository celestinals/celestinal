// Copyright 2025 The Celestinal Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cestnoti provide functions log by telegram bot
package cestnoti

import (
	"fmt"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	cestflag "github.com/celestinals/celestinal/pkg/flag"
	cestpb "github.com/celestinals/celestinal/pkg/protobuf"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ INoti = (*Noti)(nil)

// New Telegram bot
func New(conf *celestinal.Config) (INoti, error) {
	bot, err := tgbotapi.NewBotAPI(conf.GetTelegramToken())
	if err != nil {
		return &Noti{}, err
	}

	return &Noti{
		bot:  bot,
		conf: conf,
	}, nil
}

// INoti telegram bot interface
type INoti interface {
	Send(msg *celestinal.TelegramMessage) error
}

// Noti Telegram bot
type Noti struct {
	bot  *tgbotapi.BotAPI
	conf *celestinal.Config
}

// Send cestnoti to group telegram
func (r *Noti) Send(msg *celestinal.TelegramMessage) error {
	flags := cestflag.ParseEdge()
	if !flags.Telegram {
		return nil
	}

	msgText := fmt.Sprintf(
		"*Celestinal Message*\n\n"+
			"*Created At:* `%s`\n"+
			"*Author:* `%s`\n\n"+
			"`%s`\n\n"+
			"*body*\n```%s```\n\n"+
			"`%s`\n",
		cestpb.FromTime(msg.Metadata.CreatedAt).String(),
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
