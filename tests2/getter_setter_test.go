package tests2

import (
	"fmt"
	"github.com/fishingboy/GoTest/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer_GetName(t *testing.T) {
	// 原來要加 tests 才行
	player := tests.Player{Name: "Leo"}
	fmt.Println(player)
	assert.Equal(t, "Leo", player.Name)
	assert.Equal(t, "Leo", player.GetName())
}

func TestPlayer_SetName(t *testing.T) {
	player := tests.Player{Name: "Leo"}
	fmt.Println(player)
	err := player.SetName("Evonne")
	assert.Nil(t, err)
	assert.Equal(t, "Evonne", player.GetName())
}

func TestPlayer_GetMoney(t *testing.T) {
	player := tests.Player{Name: "Leo"}
	fmt.Println(player)
	err := player.SetMoney(100)
	assert.Nil(t, err)
	assert.Equal(t, 100, player.GetMoney())
}

func TestPlayer_SetMoney(t *testing.T) {
	player := tests.Player{Name: "Leo"}
	fmt.Println(player)
	err := player.SetMoney(200)
	assert.Nil(t, err)
	assert.Equal(t, 200, player.GetMoney())
}
