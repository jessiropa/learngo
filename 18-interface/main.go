package main

import (
	"fmt"
	"math"
)

// deklarasi interface
type hitung interface {
	luas() float64
	keliling() float64
}

// membuat struct bangun datar lingkaran
type lingkaran struct {
	diameter float64
}

// membuat method dari struct lingkaran
func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(),2)
}

func (l lingkaran) keliling() float64{
	return math.Pi * l.diameter
}

// membuat struct persegi 
type persegi struct{
	sisi float64
}

func (p persegi) luas()float64{
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64{
	return p.sisi * 4
}

// embedded interface
type hitung2d interface{
	luas() float64
	keliling() float64
}

type hitung3d interface{
	volume() float64
}

type hitung1 interface{
	hitung2d
	hitung3d
}

// membuat struct kubus
type kubus struct {
	sisi float64
}

// membuat method dari struct kubus 
func (k *kubus) volume() float64{
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64{
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64{
	return k.sisi * 12
}
func main(){
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== Persegi")
	fmt.Println("luas       :", bangunDatar.luas())
	fmt.Println("keliling   :", bangunDatar.keliling())

	fmt.Println("")

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== Lingkaran")
	fmt.Println("luas        :", bangunDatar.luas())
	fmt.Println("keliling    :", bangunDatar.keliling())
	fmt.Println("jari-jari   :", bangunDatar.(lingkaran).jariJari())

	var bangunDaatar hitung = lingkaran{14.0}
	var bangunLingkaran lingkaran = bangunDaatar.(lingkaran)

	
	fmt.Println("jari-jari   :", bangunLingkaran.jariJari())

	var bangunRuang hitung1 = &kubus{4}

	fmt.Println("===== Kubus")
	fmt.Println("luas       :", bangunRuang.luas())
	fmt.Println("keliling   :", bangunRuang.keliling())
	fmt.Println("volume     :", bangunRuang.volume())
}