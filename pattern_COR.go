package main

import "fmt"

/// создаем клиента
type Client struct {
	money    int
	allTime float64
	sizeOfWheel int
	typeOfInjector string
	modelYear int
	mileage int
}

/// создаем интерфейс
type WorkStation interface {
	Execute(c *Client)
	SetNext(w WorkStation)
}

/// имплементация "Шиномонтаж"
type WheelSwap struct {
	NextWorkStation WorkStation
}
func (ws *WheelSwap) SetNext(w WorkStation) {
	ws.NextWorkStation = w
}
func (ws *WheelSwap) Execute(c *Client) {
	if c.money < 1500 || c.sizeOfWheel > 21 {
		if c.sizeOfWheel > 21 {
			fmt.Println("Шиномонтаж данного размера колес невозможен")
		} else if c.money < 1500  {
			fmt.Println("Недостаточно средств для выполнения шиномонтажа")
		}
		if ws.NextWorkStation != nil {
		   ws.NextWorkStation.Execute(c)
		}
		return
	}
	fmt.Println("Шиномонтаж завершен успешно")
	c.money -= 1500
	c.allTime += 1.5

	if ws.NextWorkStation != nil {
		ws.NextWorkStation.Execute(c)
	}
}

/// имплементация "Чистка форсунок"
type CleanInjector struct {
	NextWorkStation WorkStation
}
func (ci *CleanInjector) SetNext(w WorkStation) {
	ci.NextWorkStation = w
}
func (ci *CleanInjector) Execute(c *Client) {
	if c.typeOfInjector == "ceramic"{
		fmt.Println("Чистка керамических форсунок ультразвуком невозможна")
	} else if c.money < 2000  {
		fmt.Println("Недостаточно средств для выполнения чистки форсунок")
	}
	if c.money < 2000 || c.typeOfInjector == "ceramic" {
		if ci.NextWorkStation != nil {
			ci.NextWorkStation.Execute((c))
		}
		return
	}
	fmt.Println("Чистка форсунок завершена успешно")
	c.money -= 2000
	c.allTime += 2.0

	if ci.NextWorkStation != nil {
		ci.NextWorkStation.Execute(c)
	}
}

/// имплементация "Замена масла"
type ChangeOil struct {
	NextWorkStation WorkStation
}
func (co *ChangeOil) SetNext(w WorkStation) {
	co.NextWorkStation = w
}
func (co *ChangeOil) Execute(c *Client) {
	if c.money < 500 || c.mileage < 5000 {
		if c.mileage < 5000 {
			fmt.Println("Межсервисный пробег слишком мал для замены масла")
		} else if c.money < 500 {
			fmt.Println("Недостаточно средств для замены масла")
		}
		if co.NextWorkStation != nil {
			co.NextWorkStation.Execute(c)
		}
		return
	}
	fmt.Println("Замена масла выполнена")
	c.money -= 500
	c.allTime += 0.5

	if co.NextWorkStation != nil {
		co.NextWorkStation.Execute(c)
	}
}

/// имплементация "Сброс ошибок"
type ResetErrCodes struct {
	NextWorkStation WorkStation
}
func (rec *ResetErrCodes) SetNext (w WorkStation) {
	rec.NextWorkStation = w
}
func (rec *ResetErrCodes) Execute (c *Client) {
	if c.money < 200 || c.modelYear < 2005 {
		if c.modelYear < 2005{
			fmt.Println("Ваш автомобиль не оснащен системой оповещения об ошибках ДВС")
		} else if  c.money < 200 {
			fmt.Println("Недостаточно средств для выполнения сброса кодов ошибок")
		}
		if rec.NextWorkStation != nil {
			rec.NextWorkStation.Execute(c)
		}
		return
	}
	fmt.Println("Коды ошибок удалены")
	c.money -= 200
	c.allTime += 0.1

	if rec.NextWorkStation != nil {
		rec.NextWorkStation.Execute(c)
	}
}

func AllVariant () WorkStation {

	WheelSwap := WheelSwap{}
	ChangeOil := ChangeOil{}
	CleanInjector := CleanInjector{}
	ResetErrCodes := ResetErrCodes{}

	ChangeOil.SetNext(&CleanInjector)
	CleanInjector.SetNext(&ResetErrCodes)
	ResetErrCodes.SetNext(&WheelSwap)

	return &ChangeOil
}


func TOVariant () WorkStation {
	ChangeOil := ChangeOil{}
	CleanInjector := CleanInjector{}
	ResetErrCodes := ResetErrCodes{}

	ChangeOil.SetNext(&CleanInjector)
	CleanInjector.SetNext(&ResetErrCodes)

	return &ChangeOil
}


func main() {
	var sum int
	var tim float64
	var wheel int
	var year int
	var mile int
	var Name string
	var variant string

	fmt.Println("Как к вам обращаться?")
	fmt.Scan(&Name)

	fmt.Println("На какую сумму ремонта рассчитываете?")
	fmt.Scan(&sum)

	fmt.Println("Какого года ваш автомобиль?")
	fmt.Scan(&year)

	fmt.Println("Какой размер колес вашего авто в дюймах?")
	fmt.Scan(&wheel)

	fmt.Println("Каков пробег с вашего прошлого ТО?")
	fmt.Scan(&mile)

	fmt.Println("Что предполагается сделать с машиной? (1 - полный комплект обслуживания, 2 - минимальное обслуживание)")
	fmt.Scan(&variant)

	c := Client{
		money: sum,
		allTime: tim,
		sizeOfWheel: wheel,
		modelYear: year,
		typeOfInjector: "non-ceramic",
		mileage: mile,
	}

	if variant == "1" {
		AllVariant().Execute(&c)
	} else if variant == "2" {
		TOVariant().Execute(&c)
	}
	
	fmt.Printf("%s, ваша сдача составила %d рублей", Name, c.money)
	fmt.Println("")
	fmt.Printf("Всего затрачено %.1f ч.", c.allTime)
}
