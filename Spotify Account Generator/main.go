package main

import (
	generator "Telegram/Core/Generator"
	hotmailbox "Telegram/Core/Hotmailbox"
	logger "Telegram/Core/Log"
	types "Telegram/Core/Types"
	utils "Telegram/Core/Utils"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	Proxies     *utils.Iterator
	Config      = &types.Config{}
	Mutex       = &sync.Mutex{}
	Wg          = &sync.WaitGroup{}
	Generated   = []string{}
	SavedEmails = utils.Readlines("Output/full_access_emails.txt")
	Ticker      = time.NewTicker(time.Millisecond * 10)
)

func init() {
	utils.Clear()
	types.LoadJson("Data/Config.json", Config)
	go types.WaitForChanges(Config, "Data/Config.json")

	Proxies = utils.NewFromFile("Data/proxies.txt")

	utils.ShuffleArray(&Proxies.Data)
}

func GenerateAccount(success *int32) {
	var EmailPassword string

	instance, err := generator.NewGeneratorInstance(Config.Captcha.APIKey, Proxies.Next(), Config.Captcha.MaxRetries)

	if err != nil {
		logger.LogError(logger.GetStackTrace(), "Could not create a new generator instance, Error: %v", err.Error())
		return
	}

	instance.SetDisplayName(Config.Generator.DisplayName, Config.Generator.DisplayNameSuffix)
	instance.SetPassword(Config.Generator.CustomPassword, Config.Generator.PasswordSuffix)
	instance.Config = Config

	if !Config.Generator.UseHotmailbox {
		instance.SetEmail()
	} else {
		if len(SavedEmails) <= 0 {
			emailInstance := hotmailbox.NewClient(Config.Hotmailbox.APIKey)

			emails, err := emailInstance.GetNewEmails(10, strings.Contains(Config.Hotmailbox.Domain, "outlook"))

			if err != nil {
				logger.LogError(logger.GetStackTrace(), "Could not get emails, Error: %v", err)
				return
			}

			for _, v := range emails {
				utils.AppendLine("Output/full_access_emails.txt", v, Mutex)
			}

			SavedEmails = utils.Readlines("Output/full_access_emails.txt")

			logger.LogInfo(logger.GetStackTrace(), "Successfully generated 10 emails.")
		}

		Mutex.Lock()
		combo := SavedEmails[0]
		SavedEmails = SavedEmails[1:]
		Mutex.Unlock()

		EmailPassword = strings.Split(combo, ":")[1]

		instance.SetEmail(strings.Split(combo, ":")[0])
	}

	err = instance.GenerateAccount()

	if err != nil {
		logger.LogError(logger.GetStackTrace(), "Could not generate a new account, Error: %v", err.Error())
		return
	}

	atomic.AddInt32(success, 1)

	utils.AppendLine("Output/generated_accounts.txt", fmt.Sprintf("%v:%v", instance.Email, instance.Password), Mutex)
	logger.LogInfo(logger.GetStackTrace(), "Successfully generated a Spotify Account: [%v:%v]", instance.Email, instance.Password)
	Generated = append(Generated, fmt.Sprintf("%v:%v", instance.Email, instance.Password))

	if EmailPassword != "" {
		utils.AppendLine("Output/saved_accounts.txt", fmt.Sprintf("%v:%v|%v:%v", instance.Email, instance.Password, instance.Email, EmailPassword), Mutex)
		utils.DeleteLine("Output/full_access_emails.txt", instance.Email)
	}
}

func main() {
	successes := new(int32)
	started := time.Now()

	go func() {
		for {
			utils.SetTitle("https://github.com/Aran404/Spotify-Account-Creator | Created: %d/%d | Time Elapsed: %vs", *successes, Config.Threads.AmountOfAccounts, int(time.Since(started).Seconds()))
			<-Ticker.C
		}
	}()

	guard := make(chan struct{}, Config.Threads.MaxThreads)

	for i := 0; i < Config.Threads.AmountOfAccounts; i++ {
		Wg.Add(1)

		guard <- struct{}{}
		go func() {
			defer Wg.Done()

			GenerateAccount(successes)
			<-guard
		}()
	}

	Wg.Wait()

	logger.LogInfo(logger.GetStackTrace(), "Successfully generated %d/%d Spotify Accounts. [%v]", *successes, Config.Threads.AmountOfAccounts, time.Since(started).Seconds())

}
