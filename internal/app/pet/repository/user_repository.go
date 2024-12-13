package repository

import (
	"database/sql"
	"log"
	petInterface "onvet/internal/app/pet/interfaces"
	"onvet/internal/app/user/interfaces"
	"onvet/internal/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func GetAll(fiberCtx *fiber.Ctx) ([]interfaces.UserDTO, error) {
	rows, err := db.Query(`
		SELECT id, email, name, surname
		FROM pets;
	`)
	if err != nil {
		log.Println("failed to select pets", err)
		return nil, err
	}
	defer rows.Close()

	var userList = make([]interfaces.UserDTO, 0)
	for rows.Next() {
		var userItem interfaces.UserDTO

		err = rows.Scan(&userItem.ID, &userItem.Email, &userItem.Name, &userItem.Surname)

		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}

		userList = append(userList, userItem)
	}

	return userList, nil
}

func GetById(fiberCtx *fiber.Ctx, userId string) (interfaces.UserDTO, error) {
	query := `
        SELECT id, email, name, surname
        FROM users
		WHERE id = ?
    `

	rows, err := db.Query(query,
		userId,
	)
	if err != nil {
		log.Println("failed to select user: ", err)
		return interfaces.UserDTO{}, err
	}
	defer rows.Close()

	var userItem interfaces.UserDTO
	if rows.Next() {
		var surname *string
		err = rows.Scan(&userItem.ID, &userItem.Email, &userItem.Name, &surname)
		if err != nil {
			log.Println("failed to scan row: ", err)
			return interfaces.UserDTO{}, err
		}
		userItem.Surname = surname
	} else {
		return interfaces.UserDTO{}, sql.ErrNoRows
	}

	return userItem, nil
}

func Create(fiberCtx *fiber.Ctx, payload interfaces.CreateUserPayload) error {
	query := `
        INSERT INTO users (email, name, surname) 
        VALUES (?, ?, ?)
    `

	_, err := db.Exec(query,
		payload.Email,
		payload.Name,
		payload.Surname,
	)

	if err != nil {
		log.Println("failed to insert user: ", err)
		return err
	}

	return nil
}

func CreateUserAddress(fiberCtx *fiber.Ctx, payload interfaces.CreateUserAddressPayload) error {
	query := `
		INSERT INTO users_addresses (user_id, cep, street, neighborhood, uf, city, number, complement)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := db.Exec(query,
		payload.UserId,
		payload.Cep,
		payload.Street,
		payload.Neighborhood,
		payload.Uf,
		payload.City,
		payload.Number,
		payload.Complement,
	)

	if err != nil {
		log.Println("failed to insert user address: ", err)
		return err
	}

	return nil
}

func GetUserAddresses(fiberCtx *fiber.Ctx, userId string) ([]interfaces.UserAddressDTO, error) {
	query := `
		SELECT id, user_id, cep, street, neighborhood, uf, city, number, complement
		FROM users_addresses
		WHERE user_id = ?
	`

	rows, err := db.Query(query, userId)
	if err != nil {
		log.Println("failed to select user addrresses: ", err)
		return nil, err
	}
	defer rows.Close()

	var userAddressList []interfaces.UserAddressDTO
	for rows.Next() {
		var userItem interfaces.UserAddressDTO

		err = rows.Scan(
			&userItem.Id,
			&userItem.UserId,
			&userItem.Cep,
			&userItem.Street,
			&userItem.Neighborhood,
			&userItem.Uf,
			&userItem.City,
			&userItem.Number,
			&userItem.Complement,
		)

		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}

		userAddressList = append(userAddressList, userItem)
	}

	return userAddressList, nil
}

func CreateUserPet(fiberCtx *fiber.Ctx, payload petInterface.CreatePetPayload) error {
	query := `
		INSERT INTO pets (user_id, breed_id, type_id, name, gender, castrated, weight, profile_image)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		payload.UserId,
		payload.BreedId,
		payload.TypeId,
		payload.Name,
		payload.Gender,
		payload.Castrated,
		payload.Weight,
		payload.ProfileImage,
	)

	if err != nil {
		log.Println("failed to insert pet: ", err)
		return err
	}

	return nil
}
