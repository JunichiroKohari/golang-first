package main

import "fmt"

// mainは最初に呼び出される関数
func main() {
	fmt.Print("火星の表面で、私の体重は、")
	fmt.Print(53.0 * 0.3783)
	fmt.Print("ポンド、年齢は、")
	fmt.Print(24 * 365 / 687)
	fmt.Println("歳になるでしょう。")
	fmt.Println("printlnは改行")
	fmt.Printf("火星の表面で、私の体重は、%v ポンド", 149.0 * 0.3783)
	fmt.Printf("年齢は、%v 歳になるでしょう。\n", 41 * 365 / 687)
	fmt.Printf("私の体重は、%vの表面で%vポンドです\n", "Earth", 149.0)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}