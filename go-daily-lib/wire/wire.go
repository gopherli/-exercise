//+build wireinject

package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/wire"
)

var kitty = Monster{Name: "kitty"}

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

var endingA = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), wire.Value(kitty)))
var endingB = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB)))

func InitEndingA(m MonsterParam, p PlayerParam) (EndingA, error) {
	wire.Build(endingA)
	return EndingA{}, nil
}

func InitEndingB(m MonsterParam, p PlayerParam) (EndingB, error) {
	wire.Build(endingB, wire.FieldsOf(new(Mission)))
	return EndingB{}, nil
}
func NewPlayer(name string) (Player, func(), error) {
	cleanup := func() {
		fmt.Println("cleanup!")
	}
	if time.Now().Unix()%2 == 0 {
		return Player{}, cleanup, errors.New("player dead")
	}
	return Player{Name: name}, cleanup, nil
}

func main() {
	mission, cleanup, err := InitMission("dj")
	if err != nil {
		log.Fatal(err)
	}

	mission.Start()
	cleanup()
}
