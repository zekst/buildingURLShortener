package storeHouse

import (
	"fmt" 
	"context" 
	 "github.com/go-redis/redis/v8" 
	 "time"
)

	 var (
		storeService = &StorageService{}
		ctx = context.Background()
	 )

	 const CacheDuration = 6*time.Hour;

	 // wrapping around redis client as "Storage Service"
	 type StorageService struct {
		redisClient *redis.Client
	 }

	 func iniatializeStore() *StorageService {

		return storeService
	 }

	 func saveURLMapping(shortURL string, originalURL string, userId string){

	 }

	 func retrieveInitialURL(shortURL string) string{

		return string("initalURL")
	 }



