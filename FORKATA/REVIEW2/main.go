package main

import "fmt"

type Service interface{
	DataHardTask()int
}

type DataBase struct{
	data int
}

func (d *DataBase) DataHardTask() int{
	return d.data * 1000
}

type ProxyDataBase struct{
	data DataBase
	cache int
}

func (p *ProxyDataBase) DataHardTask() int{
	if p.cache != 0{
		fmt.Println("USE CACHE")
		return p.cache
	} 
	data := p.data.DataHardTask()
	p.cache = data
	return data
}

type Scaner interface{
	Scan()int
}

type Printer interface{
	Print()
}


type Squere struct{
	sm int
}

func (s *Squere) Scan() int{
	return s.sm
}

type AdapterSquere struct{
	Scaner Scaner
}

func (a *AdapterSquere) Print(){
	fmt.Println(a.Scaner.Scan())
}


func main(){
	squ := &Squere{sm: 180}
	printer := &AdapterSquere{Scaner: squ}
	printer.Print()

	newData := &DataBase{data: 782}
	proxyData := &ProxyDataBase{data: *newData, cache: 0}
	data := proxyData.DataHardTask()
	fmt.Println(data)
	data = proxyData.DataHardTask()
	fmt.Println(data)
}

