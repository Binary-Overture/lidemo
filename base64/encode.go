package main

import (
	"log"
)

func Conversion(Origin string) (string, error) {
	originNumber := len(Origin)
	answer := ""
	how := 0
	//fmt.Println(originNumber)
	if originNumber%3 == 0 {
		three := ""
		for i := 0; i < originNumber; i += 3 {
			three += string(Origin[i])
			three += string(Origin[i+1])
			three += string(Origin[i+2])
			smf, err := emtToSmf(how, three)
			if err != nil {
				log.Fatal("wrong test~")
				return "", err
			}
			answer += smf
		}
	} else {
		i := 0
		if originNumber > 3 {
			for ; i < originNumber/3*3; i += 3 {
				three := ""
				three += string(Origin[i])
				three += string(Origin[i+1])
				three += string(Origin[i+2])
				how = 0
				smf, err := emtToSmf(how, three)
				if err != nil {
					return "", err
				}
				answer += smf
			}
		}
		three := ""
		for i < originNumber {
			three += string(Origin[i])
			i++
		}

		if i%3 == 1 {
			how = 1
		} else if i%3 == 2 {
			how = 2
		}
		smf, err := emtToSmf(how, three)
		if err != nil {
			return "", err
		}
		answer += smf
	}
	return answer, nil
}

func emtToSmf(how int, emt string) (string, error) {
	//1、首先默认只有三个字符
	//fmt.Println(len(emt))
	var character1 = emt
	result := ""
	result += string(base64EncodeChars[((character1[0]&0xff)>>2)%64])
	if how == 0 {
		result += string(base64EncodeChars[(character1[0]&0xff&3<<4+character1[1]&0xff>>4)%64])
		result += string(base64EncodeChars[(character1[1]&0xff&0xf<<2+character1[2]&0xff>>6)%64])
		result += string(base64EncodeChars[(character1[2]&0xff)%64])
	} else if how == 1 {
		result += string(base64EncodeChars[(character1[0]&0xff&3<<4)%64])
		result += "=="
	} else if how == 2 {
		result += string(base64EncodeChars[(character1[0]&0xff&3<<4+character1[1]&0xff>>4)%64])
		result += string(base64EncodeChars[(character1[1]&0xff&0xf<<2)%64])
		result += "="
	}
	return result, nil
}
