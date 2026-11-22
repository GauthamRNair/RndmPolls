package main

import "time"

type Poll struct {
	Question string
	Choices []string
}

func getDayIndex() int {
	start := time.Date(2026, time.October, 30, 0, 0, 0, 0, time.UTC)
	days := int(time.Since(start).Hours() / 24)

	pollIndex := days % len(polls)
	return pollIndex
}

func todayPoll() Poll {
	pollIndex := getDayIndex()

	return polls[pollIndex]
}

var polls = []Poll{
	{
		Question: "What does the fox say?",
		Choices: []string{
			"IDK",
			"*Bow wow*",
			"*Bark*",
			"*Chirp*",
		},
	},
	{
		Question: "Is a sea a soup?",
		Choices: []string{
			"Yes, it's salty broth",
			"No, soup needs to be hot",
			"Only after seasoning",
			"The ocean is a stew",
		},
	},
	{
		Question: "Are hotdogs sandwiches?",
		Choices: []string{
			"Yes",
			"No",
			"It's a taco",
			"It's its own thing",
		},
	},
}