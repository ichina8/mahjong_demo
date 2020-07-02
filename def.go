package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	DONGFENG_MJ  = 1
	NANFENG_MJ   = 2
	XIFENG_MJ    = 3
	BEIFENG_MJ   = 4
	HONGZHONG_MJ = 5
	FACAI_MJ     = 6
	BAIBAN_MJ    = 7

	YITIAO_MJ  = 10
	ERTIAO_MJ  = 11
	BATIAO_MJ  = 17
	JIUTIAO_MJ = 18

	YIWAN_MJ  = 21
	ERWAN_MJ  = 22
	BAWAN_MJ  = 28
	JIUWAN_MJ = 29

	YITONG_MJ  = 32
	ERTONG_MJ  = 33
	BATONG_MJ  = 39
	JIUTONG_MJ = 40

	ANYTHING_MJ = 43
)

const MJ_MAX = 144
const PLAYER_MAX = 4
const SHOUPAI_MJ_MAX = 14
const WASH_MAX = 600
const ONE_MJ_MAX = 54

const (
	CHUPAI_OPT  = 100
	ZIMO_OPT    = 101
	DIANPAO_OPT = 102
)

type PLAYER struct {
	AILevel    int
	ShouPaiNum int
	ShouPai    []byte

	AIData AI_DATA
}
type DESK struct {
	PaiSet     []byte
	PaiLevel   byte
	PaiLaizip  byte
	PaiLaizi   []byte
	PaiSetFIdx int
	PaiSetTIdx int
	PlayerNum  int
	ZhuangJia  int
	Player     []PLAYER
}

type CANOPT struct {
	OptID  int
	OptPai []byte
}

func NewDesk() *DESK {
	desk := new(DESK)
	desk.PaiSet = make([]byte, 0, MJ_MAX)
	desk.PlayerNum = PLAYER_MAX
	desk.Player = make([]PLAYER, desk.PlayerNum, PLAYER_MAX)

	for i := 0; i < len(desk.Player); i++ {
		desk.Player[i].ShouPai = make([]byte, ONE_MJ_MAX)
	}
	return desk
}

func (desk *DESK) initPai(pai byte, count int) {
	for i := 0; i < count; i++ {
		desk.PaiSet = append(desk.PaiSet, pai)
		desk.PaiLevel++
	}
	return
}

func (desk *DESK) InitPaiSet() {
	for i := DONGFENG_MJ; i <= BEIFENG_MJ; i++ {
		desk.initPai(byte(i), 4)
	}

	for i := HONGZHONG_MJ; i <= BAIBAN_MJ; i++ {
		desk.initPai(byte(i), 4)
	}

	for i := YITIAO_MJ; i <= JIUTIAO_MJ; i++ {
		desk.initPai(byte(i), 4)
	}

	for i := YIWAN_MJ; i <= JIUWAN_MJ; i++ {
		desk.initPai(byte(i), 4)
	}

	for i := YITONG_MJ; i <= JIUTONG_MJ; i++ {
		desk.initPai(byte(i), 4)
	}
	return
}

func (desk *DESK) WashPai() {
	pailevel := len(desk.PaiSet)
	for i := 0; i < WASH_MAX; i++ {
		r1 := rand.Int() % pailevel
		r2 := rand.Int() % pailevel
		desk.PaiSet[r1], desk.PaiSet[r2] = desk.PaiSet[r2], desk.PaiSet[r1]
	}
	return
}

func (desk *DESK) IsLaizi(Pai byte) bool {
	for _, v := range desk.PaiLaizi {
		if v == Pai {
			return true
		}
	}
	return false
}

func (desk *DESK) FaOnePai(Siteid int) (byte, error) {
	if desk.IsHuangZhuang() {
		return 0, errors.New("huang zhuang le")

	}
	index := desk.PaiSetFIdx
	paival := desk.PaiSet[index]
	desk.Player[Siteid].ShouPai[paival]++
	desk.Player[Siteid].ShouPaiNum++
	desk.PaiSetFIdx++
	desk.PaiLevel--
	return paival, nil
}

func (desk *DESK) FaShouPai() {
	Siteid := desk.ZhuangJia
	for i := 0; i < SHOUPAI_MJ_MAX-1; i++ {
		for j := 0; j < desk.PlayerNum; j++ {
			Siteid = (Siteid + j) % desk.PlayerNum
			desk.FaOnePai(Siteid)
		}
	}
	//庄家多发一手牌
	Siteid = desk.ZhuangJia
	desk.FaOnePai(Siteid)
	return
}

func (desk *DESK) GetCheckHuOpt(SiteId int) []CANOPT {
	var optlist = make([]CANOPT, 0, ONE_MJ_MAX)
	if desk.Player[SiteId].ShouPaiNum%3 == 2 { //自摸
		opt := CANOPT{}
		opt.OptID = ZIMO_OPT
		opt.OptPai = []byte{0} //lastPai
		optlist = append(optlist, opt)
	} else { //点炮
		opt := CANOPT{}
		opt.OptID = DIANPAO_OPT
		opt.OptPai = []byte{0} //lastPai
		optlist = append(optlist, opt)
	}

	return optlist
}

func (desk *DESK) GetCanOpt(SiteId int) []CANOPT {
	optlist := make([]CANOPT, 0, ONE_MJ_MAX)

	for i := 0; i < ONE_MJ_MAX; i++ {
		if desk.Player[SiteId].ShouPai[i] > 0 {
			opt := CANOPT{}
			opt.OptID = CHUPAI_OPT //假设此处为出牌
			opt.OptPai = []byte{byte(i)}
			optlist = append(optlist, opt)
		}
	}
	return optlist
}

func (desk *DESK) IsHuangZhuang() bool {
	if desk.PaiLevel == 0 {
		return true
	}
	return false
}

func (desk *DESK) PrintShouPaiString(ShouPai []byte) {
	for i := 0; i < ONE_MJ_MAX; i++ {
		if ShouPai[i] > 0 {
			fmt.Printf("[%d]%d ", i, ShouPai[i])
		}
	}
	fmt.Printf("\n")
}

func (desk *DESK) PrintShouPai(SiteId int) {
	for i := 0; i < ONE_MJ_MAX; i++ {
		if desk.Player[SiteId].ShouPai[i] > 0 {
			fmt.Printf("[%d]%d ", i, desk.Player[SiteId].ShouPai[i])
		}
	}
	fmt.Printf("\n")
}

//仅用于模拟测试
func (desk *DESK) AutoHuPai(SiteId int) (bool, error) {
	index := 1
	desk.PrintShouPai(SiteId)
	for {
		ishu := desk.AICheckHu(SiteId)
		if ishu {
			desk.PrintShouPai(SiteId)
			fmt.Printf("HuPai le, PaiLevel= %d\n", desk.PaiLevel)
			return true, nil
		} else {
			Opt := desk.AIBestDoOpt(SiteId, desk.GetCanOpt)
			if Opt.OptID == CHUPAI_OPT {
				Pai := Opt.OptPai[0]
				fmt.Printf("[%d]ChuPai:%d\n", index, Pai)
				index++
				desk.Player[SiteId].ShouPai[Pai]--
				paival, err := desk.FaOnePai(SiteId)
				if err == nil {
					fmt.Printf("[%d]FAPai:%d\n", index, paival)
				} else {
					fmt.Println(err)
					return false, err
				}
			}
		}
	}
}
