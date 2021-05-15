package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Pair struct {
	Key string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func rankByWordCount(wordFrequencies map[string]float64) PairList{
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, float64(v)}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func BruteForce(msg string, kLen int) bruteforceTemplate {
	Logs := make(map[string]float64)

	f, err := os.Open("dictionary.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if len(scanner.Text()) == kLen {
			word := Sanitize(scanner.Text())
			decoded := Decipher(msg, word)
			res := FrequencyAnalysis(decoded)
			Logs[word] = res
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	pl := rankByWordCount(Logs)
	var Resp bruteforceTemplate
	for i := 0; i < 20; i++ {
		var record resultBoard
		//fmt.Printf("%s : %v\n", pl[i].Key, pl[i].Value)
		record.Id = i+1
		record.Key = pl[i].Key
		x := math.Round(pl[i].Value*100)/100
		record.Percent = fmt.Sprint(x)
		Resp.Result = append(Resp.Result, record)
	}
	return Resp
}

func FrequencyAnalysis(msg string) float64 {

	result := 0.0

	FrequencyTable := map[string]float64{
		"E": 12.02,
		"T": 9.10,
		"A": 8.12,
		"O": 7.68,
		"I": 7.31,
		"N": 6.95,
		"S": 6.28,
		"R": 6.02,
		"H": 5.92,
		"D": 4.32,
		"L": 3.98,
		"U": 2.88,
		"C": 2.71,
		"M": 2.61,
		"F": 2.30,
		"Y": 2.11,
		"W": 2.09,
		"G": 2.03,
		"P": 1.82,
		"B": 1.49,
		"V": 1.11,
		"K": 0.69,
		"X": 0.17,
		"Q": 0.11,
		"J": 0.10,
		"Z": 0.07,
	}

	MessageFrequency := make(map[string]float64)

	EnglishMessage := make(map[string]int)
	//EnglishOriginal := make(map[string]int)
	Alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//fill map with 0 values
	for _, v := range Alphabet {
		EnglishMessage[string(v)] = 0
		//EnglishOriginal[string(v)] = 0
	}

	//increment map values
	for _, v := range msg {
		EnglishMessage[string(v)] += 1
	}

	lenMsg := float64(len(msg))

	//fill MessageFrequency
	for k, v := range EnglishMessage {
		MessageFrequency[k] = float64(v) * 100.0/lenMsg
	}

	//fmt.Println(FrequencyTable)
	//fmt.Println(MessageFrequency)

	for k, v := range MessageFrequency {
		if v >= FrequencyTable[k] - 1.1 && v <= FrequencyTable[k] + 1.1 {
			result += 3.846
		}
	}
	return result
}

// Round returns near int value.
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func Sanitize(in string) string {
	var out []rune
	for _, v := range in {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return string(out)
}

func EncodePair(a, b rune) rune {
	return (((a - 'A') + (b - 'A')) % 26) + 'A'
}

func DecodePair(a, b rune) rune {
	return (((((a - 'A') - (b - 'A')) + 26) % 26) + 'A')
}

func Encipher(msg, key string) string {
	smsg, skey := Sanitize(msg), Sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, EncodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func Decipher(msg, key string) string {
	smsg, skey := Sanitize(msg), Sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, DecodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}