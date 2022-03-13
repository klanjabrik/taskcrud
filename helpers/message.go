package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const langText = `{
	"id": {
		"empty_state":           "Data kosong",
		"empty_state_formatted": "%s kosong",
		"empty_notfound":           "Data tidak ditemukan",
		"empty_notfound_formatted": "%s tidak ditemukan",
		"datetime_not_valid": "Waktu tidak valid",
		"email_not_verified": "Alamat email Anda belum terverifikasi, silakan cem email Anda untuk instruksi selanjutnya."
	},
	"en": {
		"empty_state": "Data is empty",
		"empty_state_formatted": "%s is empty",
		"empty_notfound":           "Data is not found",
		"empty_notfound_formatted": "%s is not found",
		"datetime_not_valid": "Date time is not valid",
		"email_not_verified": "Alamat email Anda belum terverifikasi, silakan cem email Anda untuk instruksi selanjutnya."
	}
}`

func GetMessage(key string) string {
	return GetTextMessage(key)
}

func GetFormattedMessage(key string, a ...interface{}) string {
	return fmt.Sprintf(GetTextMessage(key), a...)
}

func GetTextMessage(key string) string {

	var text = map[string]map[string]string{}

	err := json.Unmarshal([]byte(langText), &text)
	if err != nil {
		log.Fatal(err)
	}

	var textMessage = text[GetLanguage()][key]

	if textMessage == "" {
		return key
	} else {
		return textMessage
	}
}

func GetLanguage() string {
	return os.Getenv("APP_LANGUAGE")
}
