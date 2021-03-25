package main

import (
	"errors"
	"fmt"
	"time"
)

type PlayerParam string
type MonsterParam string

type Monster struct {
	Name string
}

type Player struct {
	Name string
}

type EndingA struct {
	Monster Monster
	Player  Player
}

type EndingB struct {
	Monster Monster
	Player  Player
}

func main() {
	endingA, err := InitEndingA("lizhi", "liziwan")
	if err != nil {
		fmt.Println(err)
	}
	endingA.Start()

	endingB, err := InitEndingB("lizhi", "liziwan")
	if err != nil {
		fmt.Println(err)
	}
	endingB.Start()
}

func NewMonster(name MonsterParam) Monster {
	return Monster{Name: string(name)}
}

func NewPlayer(name PlayerParam) (Player, error) {
	if time.Now().Unix()%2 == 0 {
		return Player{}, errors.New("player dead")
	}
	return Player{Name: string(name)}, nil
}

func NewEndingA(p Player, m Monster) EndingA {
	return EndingA{m, p}
}

func (m EndingA) Start() {
	fmt.Printf("%s is beating %s", m.Player, m.Monster)
}

func NewEndingB(p Player, m Monster) EndingB {
	return EndingB{m, p}
}

func (m EndingB) Start() {
	fmt.Printf("%s is beating %s", m.Monster, m.Player)
}
