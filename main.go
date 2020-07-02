package main

import "fmt"
import "math/rand"
import "time"

func (desk *DESK) ZuoPai() {
	pai := make([]byte, ONE_MJ_MAX)
	pai[DONGFENG_MJ] = 3

	pai[XIFENG_MJ] = 3
	pai[BEIFENG_MJ] = 3
	pai[YITIAO_MJ] = 2
	pai[ERTIAO_MJ] = 1
	pai[SANTIAO_MJ] = 2

	copy(desk.Player[0].ShouPai, pai)

}

func AITest() {
	desk := NewDesk()
	desk.InitPaiSet()
	desk.WashPai()
	desk.FaShouPai()

	//desk.ZuoPai()

	SiteId := 0
	desk.AutoHuPai(SiteId)
}

func main() {
	fmt.Println("Start...")
	rand.Seed(time.Now().UnixNano())
	AITest()
	return
}
