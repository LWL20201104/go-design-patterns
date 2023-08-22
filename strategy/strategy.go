package strategy

import "fmt"

type Strategy interface {
	fun()
}

var StrategyMap map[string]Strategy

func init() {
	StrategyMap = map[string]Strategy{
		"1": &Strategy1{},
		"2": &Strategy2{},
	}
}

func NewStrategy(strategyType string) (Strategy, error) {
	if val, ok := StrategyMap[strategyType]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("No supported")
}

// 策略1
type Strategy1 struct{}

func (s *Strategy1) fun() {
	fmt.Println(1)
}

// 策略2
type Strategy2 struct{}

func (s *Strategy2) fun() {
	fmt.Println(2)
}
