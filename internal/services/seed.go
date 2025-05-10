package services

import "log"

func SeedDefault() {
	defaultName := "komal"
	defaultEmail := "komal@example.com"
	defaultPassword := "komal123"

	exists, err := isEmailTaken(defaultEmail)
	if err != nil {
		log.Printf("❌ Failed to check existing user: %v", err)
		return
	}
	if exists {
		log.Println("ℹ️ Default komal user already exists.")
		return
	}

	_, err = RegisterUser(defaultName, defaultEmail, defaultPassword)
	if err != nil {
		log.Printf("❌ Failed to seed default user: %v", err)
		return
	}

	log.Println("✅ Default komal user created: komal@example.com")

}
