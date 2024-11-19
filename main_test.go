package main

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	clientID := 1

	cl, err := selectClient(db, clientID)

	require.NoError(t, err)
	require.NotEmpty(t, cl)

	assert.NotEmpty(t, cl.Birthday)
	assert.NotEmpty(t, cl.Email)
	assert.NotEmpty(t, cl.FIO)
	assert.NotEmpty(t, cl.Login)
	assert.Equal(t, clientID, cl.ID)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	clientID := -1

	cl, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)
	//require.Empty(t, cl)
	assert.Empty(t, cl.Email)
	assert.Empty(t, cl.FIO)
	assert.Empty(t, cl.Birthday)
	assert.Empty(t, cl.Login)
	assert.Empty(t, cl.ID)

	// напиши тест здесь
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	clientID, err := insertClient(db, cl)

	cl.ID = clientID

	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	clientSelect, err := selectClient(db, clientID)
	require.NoError(t, err)
	//require.Empty(t, cl)
	assert.Equal(t, cl.Email, clientSelect.Email)
	assert.Equal(t, cl.FIO, clientSelect.FIO)
	assert.Equal(t, cl.Birthday, clientSelect.Birthday)
	assert.Equal(t, cl.Login, clientSelect.Login)
	assert.Equal(t, cl.ID, clientSelect.ID)

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	clientID, err := insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, clientID)

	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)

	err = deleteClient(db, client.ID)
	require.NoError(t, err)

	_, err = selectClient(db, cl.ID)
	require.Equal(t, sql.ErrNoRows, err)
}
