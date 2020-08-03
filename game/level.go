package game

import (
	"bufio"
	"image"
	"strconv"
	"strings"
)

var level1 = `100 100 176 87
100 100 151 155
151 155 236 151
176 87 236 151
100 100 236 151
100 100 152 226
176 87 152 226
`

func loadLevel() ([]rope, []*image.Point) {
	var ropes []rope
	var pins []*image.Point

	s := strings.NewReader(level1)
	scanner := bufio.NewScanner(s)

	for scanner.Scan() {
		text := scanner.Text()
		numbers := strings.Split(text, " ")

		if len(numbers) != 4 {
			panic("invalid count of numbers")
		}

		p1 := &image.Point{X: mustAtoi(numbers[0]), Y: mustAtoi(numbers[1])}
		if p := findPin(pins, p1); p != nil {
			p1 = p
		} else {
			pins = append(pins, p1)
		}

		p2 := &image.Point{X: mustAtoi(numbers[2]), Y: mustAtoi(numbers[3])}
		if p := findPin(pins, p2); p != nil {
			p2 = p
		} else {
			pins = append(pins, p2)
		}

		ropes = append(ropes, rope{a: p1, b: p2})
	}

	return ropes, pins
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func findPin(pins []*image.Point, p *image.Point) *image.Point {
	for _, pin := range pins {
		if pin.X == p.X && pin.Y == p.Y {
			return pin
		}
	}
	return nil
}
