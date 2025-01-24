package tests

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 自定義類型模擬
type definesEquipID string
type definesJobType int
type definesGearID int

// Gear 結構定義
type Gear struct {
	JobType definesJobType `json:"jobType"` // 職業類型
	GearID  definesGearID  `json:"gearID"`  // 第幾套
}

// String 方法將 Gear 轉為字串
func (g Gear) String() string {
	return fmt.Sprintf("%d:%d", g.JobType, g.GearID)
}

// PlayerGear 玩家穿戴資料
type PlayerGear struct {
	PlayerID definesEquipID                   `json:"playerID"` // 玩家編號
	Data     map[Gear][]definesEquipID        `json:"-"`        // 穿戴列表，使用自定義序列化
	Active   map[definesJobType]definesGearID `json:"active"`   // 啟用中的裝備編組
	Elite    definesEquipID                   `json:"elite"`    // 目前使用的專武編號
}

// MarshalJSON 自定義序列化
func (pg PlayerGear) MarshalJSON() ([]byte, error) {
	// 定義 Alias，避免遞迴
	type Alias PlayerGear
	alias := Alias(pg)

	// 將 Gear key 轉換為字串型 map
	data := make(map[string][]definesEquipID)
	for k, v := range pg.Data {
		data[k.String()] = v
	}

	// 合併序列化
	return json.Marshal(&struct {
		Data map[string][]definesEquipID `json:"data"`
		Alias
	}{
		Data:  data,
		Alias: alias,
	})
}

// UnmarshalJSON 自定義反序列化
func (pg *PlayerGear) UnmarshalJSON(data []byte) error {
	// 定義 Alias，避免遞迴
	type Alias PlayerGear
	aux := &struct {
		Data map[string][]definesEquipID `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(pg),
	}

	// 解碼 JSON
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// 將字串 key 轉回 Gear
	pg.Data = make(map[Gear][]definesEquipID)
	for k, v := range aux.Data {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid key format: %s", k)
		}
		jobType, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("invalid jobType format: %s", parts[0])
		}
		gearID, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid gearID format: %s", parts[1])
		}
		pg.Data[Gear{
			JobType: definesJobType(jobType),
			GearID:  definesGearID(gearID),
		}] = v
	}
	return nil
}

func TestStructMap(t *testing.T) {
	// 測試資料
	playerGear := PlayerGear{
		PlayerID: "12345",
		Data: map[Gear][]definesEquipID{
			{JobType: 1, GearID: 101}: {"equip1", "equip2"},
			{JobType: 2, GearID: 202}: {"equip3"},
		},
		Active: map[definesJobType]definesGearID{
			1: 201,
			2: 202,
		},
		Elite: "elite1",
	}

	// 測試序列化
	jsonData, err := json.Marshal(playerGear)
	if err != nil {
		fmt.Println("Marshal Error:", err)
		return
	}
	fmt.Println("Serialized JSON:", string(jsonData))

	// 測試反序列化
	var decoded PlayerGear
	if err := json.Unmarshal(jsonData, &decoded); err != nil {
		fmt.Println("Unmarshal Error:", err)
		return
	}
	fmt.Printf("Decoded Struct: %+v\n", decoded)
}
