package main

import (
	"fmt"
	"log"
	"os/exec"
)

func p(print string) { 
	fmt.Println(print)
}

func hrun(cheat string) {
	chmodCmd := exec.Command("chmod", "a=rwx", cheat)
	if err := chmodCmd.Run(); err != nil {
		log.Printf("Ошибка изменения прав: %v", err)
		return
	}
	runCmd := exec.Command(cheat)
	if err := runCmd.Run(); err != nil {
		log.Printf("Ошибка выполнения: %v", err)
	}
}

func crun(cmd string) {
	 exec.Command("su", "-c", cmd).Run() 
}

func main() {
	print := "Выберите чит\n1.Empasis android\n2.Empasis emulator\n3.Eternity android\n4.Eternity pc\n5.eon internal\n6.eon external\n7.eon external pc\n8.Mureana\n9.Samurai hack\n10.melodium android\n11.melodium pc\n12.nigtmare\n13.pepel\n14.noobs\nВыберите что-то из списка в цифрах 1-12\n"; p(print)
	var v_c int; fmt.Scan(&v_c); print = "Вы выбрали:"; p(print)  
	switch v_c {
		case 1: hrun("./bin/empasis_android") 
		case 2: hrun("./bin/empasis_emulator")
		case 3: hrun("./bin/eternity_android")
		case 4: hrun("./bin/eternity_pc")
		case 5: hrun("./bin/eon_internal")
		case 6: hrun("./bin/eon_external")
		case 7: hrun("./bin/eon_external_pc")
		case 8: hrun("./bin/mureana")
		case 9: hrun("./bin/samurai_external")
		case 10: hrun("./bin/melodium_android")
		case 11: hrun("./bin/melodium_pc")
		case 12: hrun("./bin/nightmare_external")
		case 13: hrun("./bin/pepel")
	    case 14: hrun("./bin/noobs")
	    case 1488: print = "ПОСХАЛКО"; p(print)  
		default: print = "Неверный выбор!"; p(print)  
	}
}