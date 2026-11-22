package main

import "time"

type Poll struct {
	Question string
	Choices  []string
}

func getDayNum() int {
	start := time.Date(2026, time.October, 30, 0, 0, 0, 0, time.UTC)
	days := int(time.Since(start).Hours() / 24)

	pollIndex := days
	return pollIndex
}

func todayPoll() Poll {
	pollIndex := getDayNum()

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
	{
		Question: "Is cereal a soup?",
		Choices: []string{
			"Yes, cold soup",
			"No, it's breakfast",
			"Only with milk first",
			"Cereal is a salad",
		},
	},
	{
		Question: "Tabs or spaces?",
		Choices: []string{
			"Tabs",
			"Spaces",
			"Whatever gofmt says",
			"I use both to cause chaos",
		},
	},
	{
		Question: "Does pineapple belong on pizza?",
		Choices: []string{
			"Yes",
			"Absolutely not",
			"Only with ham",
			"Pizza belongs on pineapple",
		},
	},
	{
		Question: "How many holes does a straw have?",
		Choices: []string{
			"One",
			"Two",
			"Zero",
			"It is one long hole",
		},
	},
	{
		Question: "Is water wet?",
		Choices: []string{
			"Yes",
			"No, it makes things wet",
			"Only on Tuesdays",
			"Water is a lie",
		},
	},
	{
		Question: "Best way to exit vim?",
		Choices: []string{
			":q!",
			"Unplug the computer",
			"Throw the laptop",
			"I still live in vim",
		},
	},
	{
		Question: "Which is the best meal of the day?",
		Choices: []string{
			"Breakfast",
			"Lunch",
			"Dinner",
			"2 AM snack",
		},
	},
	{
		Question: "Would you rather fight one horse-sized duck or 100 duck-sized horses?",
		Choices: []string{
			"One horse-sized duck",
			"100 duck-sized horses",
			"Negotiate peace",
			"Run",
		},
	},
	{
		Question: "Is a pop-tart a ravioli?",
		Choices: []string{
			"Yes",
			"No",
			"It's a dessert calzone",
			"Don't ask me this",
		},
	},
	{
		Question: "Light mode or dark mode?",
		Choices: []string{
			"Dark mode",
			"Light mode",
			"Whatever the OS picks",
			"Monitor brightness at 0",
		},
	},
	{
		Question: "What's the best number?",
		Choices: []string{
			"7",
			"42",
			"0",
			"Infinity",
		},
	},
	{
		Question: "Is a taco a sandwich?",
		Choices: []string{
			"Yes",
			"No",
			"It's a folded pizza",
			"Tacos are above labels",
		},
	},
	{
		Question: "How do you pronounce GIF?",
		Choices: []string{
			"Gif (hard g)",
			"Jif (soft g)",
			"G-I-F",
			"I just say 'the moving picture'",
		},
	},
	{
		Question: "Would you rather be able to fly or be invisible?",
		Choices: []string{
			"Fly",
			"Invisible",
			"Both but only for 5 minutes",
			"Neither, I'm good",
		},
	},
	{
		Question: "What's the best season?",
		Choices: []string{
			"Spring",
			"Summer",
			"Fall",
			"Winter",
		},
	},
	{
		Question: "Does a hotdog count as a taco?",
		Choices: []string{
			"Yes",
			"No",
			"Only if folded",
			"It's a hot sandwich taco",
		},
	},
	{
		Question: "Is the S in SSH silent?",
		Choices: []string{
			"Yes",
			"No",
			"Only the second one",
			"There is no S",
		},
	},
	{
		Question: "Which came first?",
		Choices: []string{
			"The chicken",
			"The egg",
			"The rooster",
			"The frying pan",
		},
	},
	{
		Question: "Best programming language?",
		Choices: []string{
			"Go",
			"Python",
			"Rust",
			"HTML (yes, it counts)",
		},
	},
	{
		Question: "Is a hot pocket a sandwich?",
		Choices: []string{
			"Yes",
			"No",
			"It's a calzone",
			"It's lava in bread",
		},
	},
	{
		Question: "Would you rather have no internet or no AC?",
		Choices: []string{
			"No internet",
			"No AC",
			"No both, I'll be a hermit",
			"I'd move to Antarctica",
		},
	},
	{
		Question: "How should toilet paper hang?",
		Choices: []string{
			"Over",
			"Under",
			"On the floor",
			"I use the roll as a hat",
		},
	},
	{
		Question: "What's the best dinosaur?",
		Choices: []string{
			"T. rex",
			"Triceratops",
			"Velociraptor",
			"The chicken",
		},
	},
	{
		Question: "Is soup a drink?",
		Choices: []string{
			"Yes",
			"No",
			"Only if you sip it",
			"Soup is a meal you drink",
		},
	},
	{
		Question: "Would you rather know every language or play every instrument?",
		Choices: []string{
			"Every language",
			"Every instrument",
			"Every programming language",
			"Neither, I'm tired",
		},
	},
	{
		Question: "Is 'Die Hard' a Christmas movie?",
		Choices: []string{
			"Yes",
			"No",
			"Only in December",
			"Every movie is a Christmas movie",
		},
	},
	{
		Question: "Cats or dogs?",
		Choices: []string{
			"Cats",
			"Dogs",
			"Both",
			"Ducks",
		},
	},
	{
		Question: "What's the worst sound?",
		Choices: []string{
			"Nails on a chalkboard",
			"Fork on a plate",
			"Dial-up internet",
			"My alarm",
		},
	},
	{
		Question: "How many days of the week are actually good?",
		Choices: []string{
			"All 7",
			"Only weekends",
			"None",
			"Only Fridays",
		},
	},
	{
		Question: "Is a burrito a wrap?",
		Choices: []string{
			"Yes",
			"No",
			"A wrap is a sad burrito",
			"They're cousins",
		},
	},
	{
		Question: "Would you rather sneeze every time you laugh or laugh every time you sneeze?",
		Choices: []string{
			"Sneeze when I laugh",
			"Laugh when I sneeze",
			"Neither please",
			"Both at once",
		},
	},
	{
		Question: "What do you call the meal between breakfast and lunch?",
		Choices: []string{
			"Brunch",
			"Second breakfast",
			"Elevenses",
			"A snack",
		},
	},
	{
		Question: "Is a strawberry a berry?",
		Choices: []string{
			"Yes",
			"No",
			"It's a fake berry",
			"Botany is lying to us",
		},
	},
	{
		Question: "Is a banana a berry?",
		Choices: []string{
			"Yes",
			"No",
			"It's a herb",
			"It's a phone",
		},
	},
	{
		Question: "Would you rather live in the ocean or in space?",
		Choices: []string{
			"Ocean",
			"Space",
			"A cave",
			"My bed",
		},
	},
	{
		Question: "What's the best pizza topping?",
		Choices: []string{
			"Pepperoni",
			"Mushrooms",
			"Cheese only",
			"Pineapple",
		},
	},
	{
		Question: "How do you eat a Kit Kat?",
		Choices: []string{
			"Break it apart",
			"Bite it whole",
			"Eat the chocolate off first",
			"I don't eat Kit Kats",
		},
	},
	{
		Question: "Which is scarier?",
		Choices: []string{
			"Spiders",
			"Heights",
			"Public speaking",
			"Merge conflicts",
		},
	},
	{
		Question: "Is a cheeseburger a sandwich?",
		Choices: []string{
			"Yes",
			"No",
			"It's a burger, its own class",
			"Only without the patty",
		},
	},
	{
		Question: "Would you rather always be 10 minutes early or 10 minutes late?",
		Choices: []string{
			"10 minutes early",
			"10 minutes late",
			"Exactly on time",
			"I don't show up",
		},
	},
	{
		Question: "Best way to eat an Oreo?",
		Choices: []string{
			"Twist it",
			"Dunk it",
			"Bite it whole",
			"Just the cream",
		},
	},
	{
		Question: "Is it 'soda', 'pop', or 'coke'?",
		Choices: []string{
			"Soda",
			"Pop",
			"Coke",
			"Fizzy water",
		},
	},
	{
		Question: "What's the best emoji?",
		Choices: []string{
			"😂",
			"💀",
			"🔥",
			"🗿",
		},
	},
	{
		Question: "Is a lasagna a sandwich?",
		Choices: []string{
			"Yes",
			"No",
			"It's a pasta cake",
			"It's a stacked pizza",
		},
	},
	{
		Question: "Would you rather fight a goose or a raccoon?",
		Choices: []string{
			"Goose",
			"Raccoon",
			"Both at once",
			"Make friends with them",
		},
	},
	{
		Question: "What's the best breakfast food?",
		Choices: []string{
			"Pancakes",
			"Waffles",
			"Eggs",
			"Leftover pizza",
		},
	},
	{
		Question: "Should you put the milk or cereal first?",
		Choices: []string{
			"Cereal first",
			"Milk first",
			"Only cereal",
			"Only milk",
		},
	},
	{
		Question: "How many hours of sleep do you actually get?",
		Choices: []string{
			"8+",
			"6-7",
			"4-5",
			"What is sleep",
		},
	},
	{
		Question: "What's the best Pokemon starter type?",
		Choices: []string{
			"Fire",
			"Water",
			"Grass",
			"Pikachu",
		},
	},
	{
		Question: "Is a Pop-Tart a pastry or a cracker?",
		Choices: []string{
			"Pastry",
			"Cracker",
			"Sandwich",
			"Toaster dessert",
		},
	},
	{
		Question: "Would you rather have a rewind button or a pause button for your life?",
		Choices: []string{
			"Rewind",
			"Pause",
			"Fast forward",
			"Skip intro",
		},
	},
	{
		Question: "Best way to debug?",
		Choices: []string{
			"print statements",
			"A real debugger",
			"Rubber duck",
			"Restart the computer",
		},
	},
	{
		Question: "Is a mermaid more fish or human?",
		Choices: []string{
			"Fish",
			"Human",
			"Equally both",
			"Just a shark with hair",
		},
	},
	{
		Question: "Which is the worst homework?",
		Choices: []string{
			"Math",
			"Essays",
			"Group projects",
			"Reading 50 pages",
		},
	},
	{
		Question: "What's the best ice cream flavor?",
		Choices: []string{
			"Vanilla",
			"Chocolate",
			"Cookie dough",
			"Mint chip",
		},
	},
	{
		Question: "Is a quesadilla a grilled cheese?",
		Choices: []string{
			"Yes",
			"No",
			"It's a Mexican grilled cheese",
			"It's a flat taco",
		},
	},
	{
		Question: "Would you rather have teleportation or time travel?",
		Choices: []string{
			"Teleportation",
			"Time travel",
			"Both but random",
			"Neither",
		},
	},
	{
		Question: "How do you pronounce 'data'?",
		Choices: []string{
			"Day-ta",
			"Dah-ta",
			"Datum",
			"I just point at it",
		},
	},
	{
		Question: "What's the best day of the week?",
		Choices: []string{
			"Friday",
			"Saturday",
			"Sunday",
			"Wednesday",
		},
	},
	{
		Question: "Is a pigeon just a city dove?",
		Choices: []string{
			"Yes",
			"No",
			"Pigeons are robots",
			"Doves are rich pigeons",
		},
	},
	{
		Question: "What should a cat be named?",
		Choices: []string{
			"Whiskers",
			"Mr. Meow",
			"Steve",
			"Something human, like Kevin",
		},
	},
	{
		Question: "Is a ravioli a dumpling?",
		Choices: []string{
			"Yes",
			"No",
			"It's a pasta sandwich",
			"All filled food is a dumpling",
		},
	},
	{
		Question: "Would you rather only whisper or only shout?",
		Choices: []string{
			"Whisper",
			"Shout",
			"Sing everything",
			"Use sign language",
		},
	},
	{
		Question: "What's the best snack?",
		Choices: []string{
			"Chips",
			"Popcorn",
			"Fruit",
			"Crackers",
		},
	},
	{
		Question: "Do you put ketchup on eggs?",
		Choices: []string{
			"Yes",
			"No",
			"Only scrambled",
			"Hot sauce instead",
		},
	},
	{
		Question: "What's the best social app?",
		Choices: []string{
			"Discord",
			"Slack",
			"Instagram",
			"None, I talk to people",
		},
	},
	{
		Question: "Is a hot dog bun a taco shell?",
		Choices: []string{
			"Yes",
			"No",
			"It's bread soup",
			"It's a hinge",
		},
	},
	{
		Question: "How do you charge your phone?",
		Choices: []string{
			"At 1% only",
			"Whenever",
			"Always plugged in",
			"My phone is dead right now",
		},
	},
	{
		Question: "Would you rather have hands for feet or feet for hands?",
		Choices: []string{
			"Hands for feet",
			"Feet for hands",
			"Neither",
			"Tentacles instead",
		},
	},
	{
		Question: "What's the worst chore?",
		Choices: []string{
			"Dishes",
			"Laundry",
			"Vacuuming",
			"Taking out trash",
		},
	},
	{
		Question: "Is a crab a spider of the sea?",
		Choices: []string{
			"Yes",
			"No",
			"It's a sea tank",
			"It's a beach bug",
		},
	},
	{
		Question: "Best fast food?",
		Choices: []string{
			"Chick-fil-A",
			"McDonald's",
			"Taco Bell",
			"Whataburger",
		},
	},
	{
		Question: "How do you eat pizza?",
		Choices: []string{
			"Hand",
			"Fork and knife",
			"Fold it",
			"Crust first",
		},
	},
	{
		Question: "Would you rather be really tall or really short?",
		Choices: []string{
			"Really tall",
			"Really short",
			"Average",
			"Changes daily",
		},
	},
	{
		Question: "What's the best kind of weather?",
		Choices: []string{
			"Sunny",
			"Rainy",
			"Snowy",
			"Thunderstorms",
		},
	},
	{
		Question: "Is a donut a bagel?",
		Choices: []string{
			"Yes",
			"No",
			"It's a sweet bagel",
			"Bagels are bread donuts",
		},
	},
	{
		Question: "What would you do with a million dollars?",
		Choices: []string{
			"Save it",
			"Buy a house",
			"Travel",
			"Buy 1 million dollar's worth of snacks",
		},
	},
	{
		Question: "Should cats be allowed to vote?",
		Choices: []string{
			"Yes",
			"No",
			"Only in local elections",
			"They already run the government",
		},
	},
	{
		Question: "Is a nap a mini sleep?",
		Choices: []string{
			"Yes",
			"No",
			"It's a power-up",
			"Naps are time travel",
		},
	},
	{
		Question: "Which is the best planet?",
		Choices: []string{
			"Earth",
			"Saturn",
			"Mars",
			"Pluto (still a planet)",
		},
	},
	{
		Question: "What's the best way to cook eggs?",
		Choices: []string{
			"Scrambled",
			"Fried",
			"Boiled",
			"Raw (Rocky style)",
		},
	},
	{
		Question: "Would you rather talk to animals or speak every human language?",
		Choices: []string{
			"Talk to animals",
			"Every human language",
			"Both but only for an hour",
			"Neither",
		},
	},
	{
		Question: "Is a shark a fish?",
		Choices: []string{
			"Yes",
			"No",
			"It's a sea wolf",
			"It's an ocean dog",
		},
	},
	{
		Question: "What's the best video game genre?",
		Choices: []string{
			"RPG",
			"Shooter",
			"Puzzle",
			"Farming sim",
		},
	},
	{
		Question: "How should you hold a pencil?",
		Choices: []string{
			"The normal way",
			"Fist grip",
			"Between two fingers",
			"With your mouth",
		},
	},
	{
		Question: "Is a hammock a bed?",
		Choices: []string{
			"Yes",
			"No",
			"It's a hanging couch",
			"It's a fishing net for humans",
		},
	},
	{
		Question: "What's the most useless invention?",
		Choices: []string{
			"Selfie sticks",
			"Fidget spinners",
			"Smart fridges",
			"Printers",
		},
	},
	{
		Question: "Would you rather never use your phone again or never watch TV again?",
		Choices: []string{
			"Never use phone",
			"Never watch TV",
			"Neither, I'd cry",
			"I'd fake it",
		},
	},
	{
		Question: "What's the best drink?",
		Choices: []string{
			"Water",
			"Soda",
			"Juice",
			"Boba",
		},
	},
	{
		Question: "Is a toaster a robot?",
		Choices: []string{
			"Yes",
			"No",
			"It's a bread robot",
			"It's waiting for its moment",
		},
	},
	{
		Question: "Best way to spend a Saturday?",
		Choices: []string{
			"Sleeping",
			"Coding",
			"Outside",
			"Gaming",
		},
	},
	{
		Question: "What's the best color?",
		Choices: []string{
			"Blue",
			"Red",
			"Green",
			"Purple",
		},
	},
	{
		Question: "Would you rather be a wizard or a superhero?",
		Choices: []string{
			"Wizard",
			"Superhero",
			"Villain",
			"Normal person",
		},
	},
	{
		Question: "Is a pizza a pie?",
		Choices: []string{
			"Yes",
			"No",
			"It's an open sandwich",
			"It's a flat cake",
		},
	},
	{
		Question: "What's the best fruit?",
		Choices: []string{
			"Mango",
			"Apple",
			"Banana",
			"Watermelon",
		},
	},
	{
		Question: "Should you say hi to the bus driver?",
		Choices: []string{
			"Always",
			"Never",
			"Only on Mondays",
			"I am the bus driver",
		},
	},
	{
		Question: "Did you vote yesterday too?",
		Choices: []string{
			"Yes, I'm a regular",
			"No, first time",
			"I vote every day",
			"Who's asking?",
		},
	},
}
