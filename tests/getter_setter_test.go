package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer_GetName(t *testing.T) {
	player := Player{Name: "Leo"}
	fmt.Println(player)
	assert.Equal(t, "Leo", player.Name)
	assert.Equal(t, "Leo", player.GetName())
}

func TestPlayer_SetName(t *testing.T) {
	player := Player{Name: "Leo"}
	fmt.Println(player)
	err := player.SetName("Evonne")
	assert.Nil(t, err)
	assert.Equal(t, "Evonne", player.GetName())
}

func TestPlayer_GetMoney(t *testing.T) {
	player := Player{Name: "Leo", money: 100}
	fmt.Println(player)
	// 原來在同一個目錄下都讀得到，應該是這個意思吧！？
	assert.Equal(t, 100, player.money)
	assert.Equal(t, 100, player.GetMoney())
}

func TestPlayer_SetMoney(t *testing.T) {
	player := Player{Name: "Leo"}
	fmt.Println(player)
	err := player.SetMoney(200)
	assert.Nil(t, err)
	assert.Equal(t, 200, player.GetMoney())
}
