package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// пустая строка по умолчанию
	var str string
	// со спец символами
	var hello string = "Привет\n\t"
	// без спец символов
	var world string = `Мир\n\t`

	// UTF-8 из коробки
	var helloWorld = "Привет, Мир!"
	hi := "1주차 - 언어 기초"

	// одинарные ковычки для байт
	var rawBinary byte = '\x27'

	// rune (uint32) для UTF-8 символов
	var someKorean rune = '초'

	helloWorld = "Привет Мир"

	// конкатенация строк
	andGoodMorning := helloWorld + " и доброе утро!"

	// строки неизменяемы!!
	//helloWorld[0] = 72

	// получение длины строки
	byteLen := len(helloWorld)                    // 19 байт
	symbols := utf8.RuneCountInString(helloWorld) // 10 рун

	// получение подстроки в байтах, не в символах!
	hello = helloWorld[:12] //0-11 байт
	H := helloWorld[0]      // byte, 72, не 'П'

	// Конвертация в слайс байт и обратно
	var byteString = []byte(helloWorld)
	helloWorld = string(byteString)

	fmt.Println(str, hello, world, hi, rawBinary, someKorean, andGoodMorning, byteLen, symbols, H)

}
