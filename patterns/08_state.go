package patterns

import (
	"fmt"
	"log"
	"time"
)

type State interface {
	Order()
}

type SingleProduct struct {
	currentState State
	readyToSell  State
	sold         State
	noProduct    State
}

func (sp *SingleProduct) SetState(state State) {
	sp.currentState = state
}

func (sp *SingleProduct) Order() {
	sp.currentState.Order()
}

type ReadyToSellState struct {
	sp *SingleProduct
}

func (rts *ReadyToSellState) Order() {
	fmt.Println("A new order created! Check information on your email!")
	rts.sp.SetState(rts.sp.sold)
}

type SoldState struct {
	sp *SingleProduct
}

func (s *SoldState) Order() {
	fmt.Println("You have already ordered our product.")
}

type NoProductState struct {
	sp *SingleProduct
}

func (s *NoProductState) Order() {
	fmt.Println("There is no product, sorry :( Check our website later...")
	log.Println("Got new product...")
	time.Sleep(5 * time.Second)
	s.sp.SetState(s.sp.readyToSell)
}
