package main

//import "fmt"

type AI_WUSER struct {
	Weight int
	IsTing bool
	IsHu   bool
}

type AI_INFO struct {
	KS_MAX     byte
	SZ_Count   byte
	KZ_Count   byte
	DZ11_Count byte
	DZ12_Count byte
	DZ23_Count byte
	DZ13_Count byte
	DZ01_Count byte
	OptPaiVal  byte
}

type AI_DATA struct {
	BZ_Count int
	MS_Count int
	ML_Count int

	WR_Value int
	DR_Value int

	OpIdx  int
	Weight int
	Info   AI_INFO
	WUser  AI_WUSER
}

const (
	SHUNZI_WEIGHT   int = 0
	KEZI_WEIGHT     int = -1
	DUIZI_WEIGHT    int = -10
	DAZI_23_WEIGHT  int = -11
	DAZI_13_WEIGHT  int = -12
	DAZI_12_WEIGHT  int = -13
	DANZHANG_WEIGHT int = -16
)

func (desk *DESK) CutKZ(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {
	beign, end := DONGFENG_MJ, JIUTONG_MJ
	for i := beign; i <= end; i++ {
		if CardString[i] >= 3 {
			CardString[i] -= 3
			currentData.Info.KZ_Count++
			desk.AIScan(Siteid, CardString, currentData, Shoupai)
			CardString[i] += 3
			currentData.Info.KZ_Count--
		}
	}
}

func (desk *DESK) CutSZ(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {
	begin, end := YITIAO_MJ, JIUTONG_MJ
	for i := begin; i < end; i++ {
		if CardString[i] != 0 && CardString[i+1] != 0 && CardString[i+2] != 0 {
			CardString[i]--
			CardString[i+1]--
			CardString[i+2]--
			currentData.Info.SZ_Count++
			desk.AIScan(Siteid, CardString, currentData, Shoupai)
			currentData.Info.SZ_Count--
			CardString[i]++
			CardString[i+1]++
			CardString[i+2]++
		}
	}
}

func (desk *DESK) CutDZ11(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {
	beign, end := DONGFENG_MJ, JIUTONG_MJ
	for i := beign; i <= end; i++ {
		if CardString[i] >= 2 {
			CardString[i] -= 2
			currentData.Info.DZ11_Count++
			desk.AIScan(Siteid, CardString, currentData, Shoupai)
			CardString[i] += 2
			currentData.Info.DZ11_Count--
		}
	}
}

func (desk *DESK) CutDZ12(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {
	begin, end := YITIAO_MJ, JIUTONG_MJ
	for i := begin; i < end; i++ {
		if !(i == YITIAO_MJ || i == YIWAN_MJ || i == YITONG_MJ || i == BAWAN_MJ || i == BATIAO_MJ || i == BATONG_MJ) {
			continue
		}
		if CardString[i] != 0 && CardString[i+1] != 0 {
			CardString[i]--
			CardString[i+1]--
			currentData.Info.DZ12_Count++
			desk.AIScan(Siteid, CardString, currentData, Shoupai)
			currentData.Info.DZ12_Count--
			CardString[i]++
			CardString[i+1]++
		}
	}
}

func (desk *DESK) CutDZ13(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {
	begin, end := YITIAO_MJ, JIUTONG_MJ
	for i := begin; i < end; i++ {
		if !(i == YITIAO_MJ || i == YIWAN_MJ || i == YITONG_MJ || i == BAWAN_MJ || i == BATIAO_MJ || i == BATONG_MJ) {
			continue
		}
		if CardString[i] != 0 && CardString[i+2] != 0 {
			CardString[i]--
			CardString[i+2]--
			currentData.Info.DZ13_Count++
			desk.AIScan(Siteid, CardString, currentData, Shoupai)
			currentData.Info.DZ13_Count--
			CardString[i]++
			CardString[i+2]++
		}
	}
}

func (desk *DESK) CutDZ23(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {
	begin, end := YITIAO_MJ, JIUTONG_MJ
	for i := begin; i < end; i++ {
		if i == YITIAO_MJ || i == YIWAN_MJ || i == YITONG_MJ || i == BAWAN_MJ || i == BATIAO_MJ || i == BATONG_MJ {
			continue
		}
		if CardString[i] != 0 && CardString[i+1] != 0 {
			CardString[i]--
			CardString[i+1]--
			currentData.Info.DZ23_Count++
			desk.AIScan(Siteid, CardString, currentData, Shoupai)
			currentData.Info.DZ23_Count--
			CardString[i]++
			CardString[i+1]++
		}
	}
}

func (desk *DESK) CutDZ(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {

	desk.CutDZ23(Siteid, CardString, currentData, Shoupai)
	desk.CutDZ12(Siteid, CardString, currentData, Shoupai)
	desk.CutDZ13(Siteid, CardString, currentData, Shoupai)
	//desk.CutDZ11(Siteid, CardString, currentData, Shoupai)

}

func (desk *DESK) AICalKSWeight(Siteid int, CardString []byte, currentData *AI_DATA) int {

	sum := 0
	sum = KEZI_WEIGHT*int(currentData.Info.KZ_Count) +
		SHUNZI_WEIGHT*int(currentData.Info.SZ_Count) +
		DUIZI_WEIGHT*int(currentData.Info.DZ11_Count) +
		DAZI_23_WEIGHT*int(currentData.Info.DZ23_Count) +
		DAZI_12_WEIGHT*int(currentData.Info.DZ12_Count) +
		DAZI_13_WEIGHT*int(currentData.Info.DZ13_Count) +
		DANZHANG_WEIGHT*int(currentData.Info.DZ01_Count)
	return sum
}

func (desk *DESK) AICalBZWeight(Siteid int, CardString []byte, currentData *AI_DATA, shoupai []byte) int {
	g_WeightPai := []int{
		0, 1, 1, 1, 1, 1, 1, 1, 0, 0, //
		3, 4, 5, 5, 5, 5, 5, 4, 3, 0, 0, //
		3, 4, 5, 5, 5, 5, 5, 4, 3, 0, 0, //
		3, 4, 5, 5, 5, 5, 5, 4, 3, 0, 0, //
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	totalBuzhang := 0
	currentChupaiVal := currentData.Info.OptPaiVal
	valWeight := make([]int, ONE_MJ_MAX)

	copy(valWeight, g_WeightPai)

	for i := DONGFENG_MJ; i <= JIUTONG_MJ; i++ {
		if shoupai[i] > 0 && !desk.IsLaizi(byte(i)) {
			totalBuzhang += valWeight[i]
		}
	}
	if currentChupaiVal >= YITIAO_MJ && totalBuzhang >= currentData.BZ_Count {

		if shoupai[currentChupaiVal-2] > 0 && valWeight[currentChupaiVal-2] >= 3 {
			totalBuzhang -= 5
		}

		if shoupai[currentChupaiVal-1] > 0 && valWeight[currentChupaiVal-1] >= 3 {
			totalBuzhang -= 10
		}

		if shoupai[currentChupaiVal] > 0 && valWeight[currentChupaiVal] >= 1 {
			totalBuzhang -= 10
			if CardString[currentChupaiVal] > 1 {
				totalBuzhang -= 5
			}
		}

		if shoupai[currentChupaiVal+1] > 0 && valWeight[currentChupaiVal+1] >= 3 {
			totalBuzhang -= 10
		}

		if shoupai[currentChupaiVal+2] > 0 && valWeight[currentChupaiVal+2] >= 3 {
			totalBuzhang -= 5
		}
	}

	return totalBuzhang
}

func (desk *DESK) AIScan(Siteid int, CardString []byte, currentData *AI_DATA, Shoupai []byte) {
	desk.CutKZ(Siteid, CardString, currentData, Shoupai)
	desk.CutSZ(Siteid, CardString, currentData, Shoupai)
	if currentData.Info.KZ_Count+currentData.Info.SZ_Count < currentData.Info.KS_MAX {
		return
	}
	currentData.Info.KS_MAX = currentData.Info.KZ_Count + currentData.Info.SZ_Count
	desk.CutDZ(Siteid, CardString, currentData, Shoupai)

	var Singal_Count byte = 0
	var Duizi_Count byte = 0
	for i := DONGFENG_MJ; i <= JIUTONG_MJ; i++ {

		if CardString[i] == 2 {
			Duizi_Count++
		}

		if CardString[i] == 1 {
			Singal_Count++
		}

	}
	currentData.Info.DZ01_Count = Singal_Count
	currentData.Info.DZ11_Count = Duizi_Count

	lacknum := desk.AICalLackNum(Siteid, CardString, currentData, Shoupai)
	sum := desk.AICalKSWeight(Siteid, CardString, currentData)

	totalBuzhang := desk.AICalBZWeight(Siteid, CardString, currentData, Shoupai)

	if lacknum < currentData.ML_Count ||
		((lacknum == currentData.ML_Count) &&
			((sum > currentData.MS_Count) ||
				(sum == currentData.MS_Count && totalBuzhang >= currentData.BZ_Count))) {
		currentData.ML_Count = lacknum
		currentData.MS_Count = sum
		currentData.BZ_Count = totalBuzhang
		currentData.WUser.IsHu = desk.AIIsHu(Siteid, CardString, currentData, Shoupai)
		/*
			AIData := currentData
			fmt.Printf("ML_Count = %d MS_Count = %d, BZ_Count = %d IsHu = %v  %d %d %d %d %d %d %d\r\n",
				AIData.ML_Count, AIData.MS_Count, AIData.BZ_Count, AIData.WUser.IsHu,
				AIData.Info.SZ_Count, AIData.Info.KZ_Count, AIData.Info.DZ11_Count, AIData.Info.DZ12_Count,
				AIData.Info.DZ13_Count, AIData.Info.DZ23_Count, AIData.Info.DZ01_Count)
		*/

	}
}

func (desk *DESK) AICalLackNum(SiteId int, CardString []byte, AIData *AI_DATA, Shoupai []byte) int {
	lacknum := AIData.Info.DZ11_Count +
		AIData.Info.DZ12_Count +
		AIData.Info.DZ13_Count +
		AIData.Info.DZ23_Count +
		AIData.Info.DZ01_Count*2
	lacknum -= Shoupai[ANYTHING_MJ]
	return int(lacknum)
}

func (desk *DESK) AIIsHu(SiteId int, CardString []byte, AIData *AI_DATA, Shoupai []byte) bool {
	if desk.Player[SiteId].ShouPaiNum%3 == 2 {
		if Shoupai[ANYTHING_MJ] == 0 { //仅处理无赖情况
			if AIData.ML_Count == 1 && AIData.Info.DZ11_Count == 1 {
				return true
			}
		} else {
			if AIData.ML_Count < 0 {
				return true
			} else if Shoupai[ANYTHING_MJ]-AIData.Info.DZ12_Count-AIData.Info.DZ13_Count-AIData.Info.DZ23_Count >= 0 {
				return true
			}
		}
	}
	return false
}

func (desk *DESK) AIInitData(AIData *AI_DATA) {
	AIData.MS_Count = -10000
	AIData.BZ_Count = -10000
	AIData.ML_Count = 10000
	AIData.WUser.IsHu = false
	AIData.WUser.IsTing = false
	return
}

func (desk *DESK) AIDeal(SiteId int, CardString []byte, AIData *AI_DATA, ShouPai []byte) {
	desk.AIInitData(AIData)
	desk.AIScan(SiteId, CardString, AIData, ShouPai)
	return
}

func (desk *DESK) AITryDoOpt(SiteId int, CardString []byte, Opt CANOPT) {
	switch Opt.OptID {
	case 100: //chupai
		CardString[Opt.OptPai[0]] -= 1
		break
	case 101:
		break
	case 102:
		break
	default:
		break

	}
	return
}

func (desk *DESK) AICheckData(src *AI_DATA, desc *AI_DATA) bool {
	if desc.Weight < src.Weight {
		return false
	} else if desc.Weight > src.Weight {
		return true
	}

	if desc.ML_Count > src.ML_Count {
		return false
	} else if desc.ML_Count < src.ML_Count {
		return true
	}

	if desc.MS_Count < src.MS_Count {
		return false
	} else if desc.MS_Count > src.MS_Count {
		return true
	}

	if desc.BZ_Count < src.BZ_Count {
		return false
	} else if desc.BZ_Count > src.BZ_Count {
		return true
	}
	//其它情况
	return false
}

func (desk *DESK) AICalWeightRatio(SiteId int, AIData *AI_DATA) {
	if AIData.WUser.IsHu == true {
		AIData.Weight += 10
	} else if AIData.WUser.IsTing == true {
		AIData.Weight += 5
	}
	return
}

func (desk *DESK) AICheckHu(SiteId int) bool {
	desk.AIBestDoOpt(SiteId, desk.GetCheckHuOpt)
	AIData := desk.Player[SiteId].AIData
	if AIData.ML_Count == 1 && AIData.Info.DZ11_Count == 1 {
		return true
	}
	return false
}

func (desk *DESK) AIBestDoOpt(SiteId int, GetCanOpt func(SiteId int) []CANOPT) CANOPT {
	Opt := CANOPT{}

	CardString := make([]byte, ONE_MJ_MAX)
	ShouPai := make([]byte, ONE_MJ_MAX)
	OptList := GetCanOpt(SiteId)
	desk.Player[SiteId].AIData = AI_DATA{}
	desk.AIInitData(&desk.Player[SiteId].AIData)
	for i, v := range OptList {
		copy(CardString, desk.Player[SiteId].ShouPai)
		desk.AITryDoOpt(SiteId, CardString, v)
		copy(ShouPai, CardString)
		AIData := AI_DATA{OpIdx: i}
		if v.OptID == 100 {
			AIData.Info.OptPaiVal = v.OptPai[0]
		}
		desk.AIDeal(SiteId, CardString, &AIData, ShouPai)
		desk.AICalWeightRatio(SiteId, &AIData)

		if desk.AICheckData(&desk.Player[SiteId].AIData, &AIData) {
			desk.Player[SiteId].AIData = AIData
			Opt = v
		}
		/*
			fmt.Printf("Id = %d pai = %d  ML_Count = %d MS_Count = %d, BZ_Count = %d IsHu = %v\r\n",
				i, v.OptPai[0], AIData.ML_Count, AIData.MS_Count, AIData.BZ_Count, AIData.WUser.IsHu)
			fmt.Printf("OpIdx = %d ML_Count = %d MS_Count = %d, BZ_Count = %d IsHu = %v\r\n",
				desk.Player[SiteId].AIData.OpIdx, desk.Player[SiteId].AIData.ML_Count,
				desk.Player[SiteId].AIData.MS_Count, desk.Player[SiteId].AIData.MS_Count,
				desk.Player[SiteId].AIData.WUser.IsHu)

			fmt.Printf("\n\n")
		*/
	}
	return Opt
}
