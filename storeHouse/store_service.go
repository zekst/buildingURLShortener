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

		//Initializing redis client 

		redisClient := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		})

		pong, error := redisClient.Ping(ctx).Result()

		if error!=nil{
			panic(fmt.Sprintf("dikkat hogyi hai %v", error))
		}

		fmt.Printf("Redis server started successfully with pong message = {%s}", pong)

		storeService.redisClient = redisClient

		return storeService
	 }

	 func saveURLMapping(shortURL string, originalURL string){

		error := storeService.redisClient.Set(ctx, shortURL, originalURL, CacheDuration).Err()

		if(error!=nil){
			panic(fmt.Sprintf("Failed at saving key URL i am so sorry for that; short URL {%s}, origianal URL {%s}, error {%v}", shortURL, originalURL, error))
		}

	 }

	 func retrieveInitialURL(shortURL string) string{

		retrievedURL, error := storeService.redisClient.Get(ctx, shortURL).Result()

		if(error!=nil){
			panic(fmt.Sprintf("can not access original URL of short url {%s} and error {%v} happened", shortURL, error))		
		}

		return retrievedURL
	 }



