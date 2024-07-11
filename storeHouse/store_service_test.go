package storeHouse

import (
	"testing"
	"github.com/stretchr/testify/assert"
	
)

var testStoreService = &StorageService{}

func init(){
	testStoreService = iniatializeStore()
}

func TestStoreInit(testin *testing.T){
	assert.True(testin, testStoreService.redisClient!=nil)
}

func TestInsertionAndRetrieval(testin *testing.T){
	initialLink :="https://www.kachkuyadav/sellingwater/insteadofmilk/getttinglotofmoney"
	//userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	saveURLMapping(shortURL, initialLink)
	
	retrievedURL := retrieveInitialURL(shortURL)

	assert.Equal(testin, initialLink, retrievedURL)
}
