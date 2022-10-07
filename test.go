package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NPC struct {
	Hp, Atk int
}

func (npc_self *NPC) Attack(npc_2 *NPC) {
	npc_2.Hp -= npc_self.Atk
}

func main() {
	boss := NPC{
		Hp:  100,
		Atk: 20,
	}

	player := NPC{
		Hp:  30,
		Atk: 10,
	}

	rand.Seed(time.Now().UnixMicro())

	for true {
		fmt.Printf("\n玩家HP %d\n", player.Hp)
		fmt.Printf("\n魔王HP %d\n", boss.Hp)
		if boss.Hp <= 0 {
			fmt.Print("Player win!!")
			break
		}
		if player.Hp <= 0 {
			fmt.Print("Boss win!!")
			break
		}
		a, b := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%d + %d = ? : ", a, b)
		scan_ans := 1
		fmt.Scan(&scan_ans)
		ans := a + b
		if ans == scan_ans {
			player.Attack(&boss)
		} else {
			boss.Attack(&player)
		}
	}
}
