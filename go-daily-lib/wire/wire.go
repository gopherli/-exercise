//+build wireinject

package main

import "github.com/google/wire"

func InitEndingA(m MonsterParam, p PlayerParam) (EndingA, error) {
	wire.Build(NewMonster, NewPlayer, NewEndingA)
	return EndingA{}, nil
}

func InitEndingB(m MonsterParam, p PlayerParam) (EndingB, error) {
	wire.Build(NewMonster, NewPlayer, NewEndingB)
	return EndingB{}, nil
}
