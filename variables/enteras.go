package variables

import "fmt"

func MostrarEnteros() {

	//Delaclaracion de variable de tipo enetro

	var intcomun int
	/*resultado de int32 toma el tipo de dato de la asigancion¨*/
	intde32 := int32(10)
	intde64 := int64(100)

	//impreme las vaiable usando el paquete
	fmt.Println("intcomun =", intcomun)
	fmt.Println("intde32  =", intde32)
	fmt.Println("intde64  =", intde64)
}
