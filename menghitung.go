package main

import “fmt”

var (

massa, volume float64

)

func main(){

fmt.Println(“masukkan massa benda :”)

fmt.Scanf(“%f”,&massa)

fmt.Println(“masukkan volume benda :”)

fmt.Scanf(“%f”,&volume)

massa_jenis:=massa/volume

fmt.Println(“massa jenis benda adalah :”,massa_jenis)

}