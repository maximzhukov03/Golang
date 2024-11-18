/*Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power,
с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов,
но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo
(тогда Ammo уменьшается на единицу, а метод возвращает true),
если его нет, то метод вернет false. Метод RideBike работает также,
но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно,
вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main,
в дальнейшем программа проверит результат.*/

package main

import (
	"fmt"
)

type Device struct{
	On bool
	Ammo int
	Power int
}

func (d *Device) Shoot() bool{
	if !d.On{
		return false
	}
	if d.Ammo > 0{
		d.Ammo--
		return true
	}
	return false
}

func (d *Device) RideBike() bool{
	if !d.On{
		return false
	}
	if d.Power > 0{
		d.Power--
		return true
	}
	return false
}

func test3(){
	fmt.Printf("ЗАПУСК ТЕСТА №3\n")
	testStruct := &Device{
		On: true,
		Ammo: 5,
		Power: 5,
	}

	fmt.Println("Проверка на вшивость: ", testStruct.Ammo)
	fmt.Println("Проверка на вшивость: ", testStruct.Shoot())
	fmt.Println("Проверка на вшивость: ", testStruct.Power)
	fmt.Println("Проверка на вшивость: ", testStruct.RideBike())

	testStruct.On = false
	fmt.Println("Проверка на вшивость: ", testStruct.Ammo)
	fmt.Println("Проверка на вшивость: ", testStruct.Shoot())
	fmt.Println("Проверка на вшивость: ", testStruct.Power)
	fmt.Println("Проверка на вшивость: ", testStruct.RideBike())
}