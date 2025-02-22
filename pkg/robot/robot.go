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

// Package robot provide functions log by telegram bot
package robot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/pkg/logger"
)

var bot *tgbotapi.BotAPI

// Init telegram bot api
func Init(conf *types.Config) {
	if bot != nil {
		return
	}

	var err error
	bot, err = tgbotapi.NewBotAPI(conf.GetBotToken())

	if err != nil {
		logger.Errorf("Failed to connect to Telegram: %v", err)
	}

}
