package main

import (
	"fmt"
)

type Animal struct {
	runSpeed  int
	iteration int
	name      string
	voice     string
}
type Place struct {
	winner Animal
	second Animal
	loser  Animal
}

type Tiger struct {
	Animal
	colorOfStripes string
	numOfFangs     int
}

type Turtle struct {
	Animal
	armorSize  float64
	armorColor string
}

type Fish struct {
	Animal
	numOfFins        int
	PresenceOfThorns bool
}

type Race struct {
	distance int
	Turtle   Turtle
	Tiger    Tiger
	Fish     Fish
	Place
}

func (a *Animal) WinnerSay() {
	fmt.Printf("%s says: %s\n", a.name, a.voice)
}

func (a *Animal) LoserSay() {
	fmt.Printf("%s says: Ohhhh noooo, I lost((((\n", a.name)
}

func (a *Animal) SecondSay() {
	fmt.Printf("%s says: I'm only Second((((((\n", a.name)
}

func (r *Race) CreateTeam(turtle Turtle, tiger Tiger, fish Fish) {
	r.Turtle = turtle
	r.Tiger = tiger
	r.Fish = fish
}

func (r *Race) FindWinner() {
	if r.Turtle.runSpeed >= r.distance {
		r.Place.winner = r.Turtle.Animal
	} else if r.Tiger.runSpeed >= r.distance {
		r.Place.winner = r.Tiger.Animal
	} else if r.Fish.runSpeed >= r.distance {
		r.Place.winner = r.Fish.Animal
	}

	fmt.Printf("%s won - the running time is %d iterations\n", r.Place.winner.name, r.Place.winner.iteration)
	r.Place.winner.WinnerSay()
}

func (r *Race) FindLoser() {
	if r.Turtle.runSpeed < r.Tiger.runSpeed && r.Turtle.runSpeed < r.Fish.runSpeed {
		r.Place.loser = r.Turtle.Animal
	}
	if r.Tiger.runSpeed < r.Turtle.runSpeed && r.Tiger.runSpeed < r.Fish.runSpeed {
		r.Place.loser = r.Tiger.Animal
	}
	if r.Tiger.runSpeed < r.Turtle.runSpeed && r.Tiger.runSpeed < r.Fish.runSpeed {
		r.Place.loser = r.Fish.Animal
	}

	fmt.Printf("%s last - the running time is %d iterations\n", r.Place.loser.name, r.Place.loser.iteration)
	r.Place.loser.LoserSay()
}

func (r *Race) FindSecondPlace() {
	winnerSpeed := r.Place.winner.runSpeed
	loserSpeed := r.Place.loser.runSpeed

	if r.Turtle.runSpeed < winnerSpeed && r.Turtle.runSpeed > loserSpeed {
		r.Place.second = r.Turtle.Animal
	}
	if r.Tiger.runSpeed < winnerSpeed && r.Tiger.runSpeed > loserSpeed {
		r.Place.second = r.Tiger.Animal
	}
	if r.Fish.runSpeed < winnerSpeed && r.Fish.runSpeed > loserSpeed {
		r.Place.second = r.Fish.Animal
	}

	fmt.Printf("%s finished second - the running time is %d iterations\n", r.Place.second.name, r.Place.second.iteration)
	r.Place.second.SecondSay()
}

func (r *Race) Start() {
	fmt.Println("The race has started!")
	fmt.Printf("Distance: %d\nTiger speed: %d\nFish speed: %d\nTurtle speed: %d\n", r.distance, r.Tiger.runSpeed,
		r.Fish.runSpeed, r.Turtle.runSpeed)

	increaseAmountTurtle := r.Turtle.runSpeed
	increaseAmountTiger := r.Tiger.runSpeed
	increaseAmountFish := r.Fish.runSpeed

	for r.Turtle.runSpeed < r.distance && r.Tiger.runSpeed < r.distance && r.Fish.runSpeed < r.distance {
		r.Turtle.iteration++
		r.Turtle.runSpeed += increaseAmountTurtle
		fmt.Printf("Turtle Distance left: %d\n", r.Turtle.runSpeed)

		r.Tiger.iteration++
		r.Tiger.runSpeed += increaseAmountTiger
		fmt.Printf("Tiger Distance left: %d\n", r.Tiger.runSpeed)

		r.Fish.iteration++
		r.Fish.runSpeed += increaseAmountFish
		fmt.Printf("Fish Distance left: %d\n", r.Fish.runSpeed)
	}
	fmt.Println("The race was finished!")

	if r.Turtle.runSpeed >= r.distance && r.Tiger.runSpeed >= r.distance && r.Fish.runSpeed >= r.distance {
		fmt.Println("It's a draw!")
	} else {
		r.FindWinner()
		r.FindSecondPlace()
		r.FindLoser()
	}
}

func main() {
	race := Race{
		distance: 1000,
	}
	race.CreateTeam(
		Turtle{Animal: Animal{iteration: 0, runSpeed: 100, name: "Turtle", voice: "I am a turtle!"}, armorSize: 15.5, armorColor: "Brown"},
		Tiger{Animal: Animal{iteration: 0, runSpeed: 150, name: "Tiger", voice: "Roar!"}, colorOfStripes: "Black", numOfFangs: 4},
		Fish{Animal: Animal{iteration: 0, runSpeed: 120, name: "Fish", voice: "Blub-blub!"}, numOfFins: 5, PresenceOfThorns: false},
	)
	race.Start()

}
