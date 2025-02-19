package redisdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// RedisClient struct
 type RedisClient struct {
	client *redis.Client
 }

// NewRedisClient initializes a new Redis client
func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisClient{client: client}
}

// AddProducts stores product data in Redis
func (r *RedisClient) AddProducts() {
	products := map[string]map[string]string{
		"1": {"name": "Laptop", "price": "1000", "stock": "10"},
		"2": {"name": "Smartphone", "price": "700", "stock": "15"},
		"3": {"name": "Headphones", "price": "100", "stock": "50"},
		"4": {"name": "Smartwatch", "price": "200", "stock": "20"},
		"5": {"name": "Tablet", "price": "500", "stock": "8"},
		"6": {"name": "Camera", "price": "800", "stock": "12"},
		"7": {"name": "Wireless Earbuds", "price": "150", "stock": "30"},
		"8": {"name": "Fitness Tracker", "price": "120", "stock": "25"},
		"9": {"name": "External Hard Drive", "price": "300", "stock": "18"},
		"10": {"name": "The River", "price": "420", "stock": "10"},
		"11": {"name": "Bluetooth Speaker", "price": "180", "stock": "22"},

	}

	for pid, pdata := range products {
		r.client.HSet(ctx, "product:"+pid, pdata)
	}
}

// AddUsers stores user data in Redis
func (r *RedisClient) AddUsers() {
	users := map[string]map[string]string{
		"1": {"name": "Alice", "email": "alice@example.com"},
		"2": {"name": "Bob", "email": "bob@example.com"},
		"3": {"name": "Charlie", "email": "charlie@example.com"},
		"4": {"name": "David", "email": "david@example.com"},
		"5": {"name": "Eve", "email": "eve@example.com"},
		"6": {"name": "Frank", "email": "frank@example.com"},
		"7": {"name": "Grace", "email": "grace@example.com"},
		"8": {"name": "Hank", "email": "hank@example.com"},
		"9": {"name": "Ivy", "email": "ivy@example.com"},
		"10": {"name": "Jack", "email": "jack@example.com"},
	}

	for uid, udata := range users {
		r.client.HSet(ctx, "user:"+uid, udata)
	}
}

// AddOrders stores order data in Redis
func (r *RedisClient) AddOrders() {
	orders := map[string]map[string]string{
		"1": {"user_id": "1", "product_id": "2", "quantity": "1", "status": "shipped"},
		"2": {"user_id": "2", "product_id": "1", "quantity": "1", "status": "pending"},
		"3": {"user_id": "3", "product_id": "3", "quantity": "2", "status": "shipped"},
		"4": {"user_id": "4", "product_id": "4", "quantity": "1", "status": "pending"},
		"5": {"user_id": "5", "product_id": "1", "quantity": "1", "status": "shipped"},
		"6": {"user_id": "6", "product_id": "2", "quantity": "2", "status": "pending"},
		"7": {"user_id": "7", "product_id": "3", "quantity": "1", "status": "shipped"},	
		"8": {"user_id": "8", "product_id": "4", "quantity": "1", "status": "pending"},
		"9": {"user_id": "9", "product_id": "1", "quantity": "2", "status": "shipped"},
		"10": {"user_id": "10", "product_id": "2", "quantity": "1", "status": "pending"},
		"11": {"user_id": "1", "product_id": "3", "quantity": "1", "status": "shipped"},
		"12": {"user_id": "2", "product_id": "4", "quantity": "2", "status": "pending"},
		"13": {"user_id": "3", "product_id": "1", "quantity": "1", "status": "shipped"},
		"14": {"user_id": "4", "product_id": "2", "quantity": "1", "status": "pending"},
		"15": {"user_id": "5", "product_id": "3", "quantity": "2", "status": "shipped"},
		"16": {"user_id": "6", "product_id": "4", "quantity": "1", "status": "pending"},
		"17": {"user_id": "7", "product_id": "1", "quantity": "1", "status": "shipped"},
	}

	for oid, odata := range orders {
		r.client.HSet(ctx, "order:"+oid, odata)
	}
}

// AddCategories stores product categories in Redis
func (r *RedisClient) AddCategories() {
	r.client.SAdd(ctx, "categories", "Electronics", "Fashion", "Home Appliances", "Books", "Sports", 
	"Toys", "Beauty", "Automotive", "Jewelry", "Musical Instruments")
}

// AddRecentlyViewed stores recently viewed products in Redis
func (r *RedisClient) AddRecentlyViewed() {
	r.client.LPush(ctx, "recently_viewed:1", "2", "3", "1")
}

// PrintData retrieves and prints stored data from Redis
func (r *RedisClient) PrintData() {
	fmt.Println("Products:")
	keys, _ := r.client.Keys(ctx, "product:*").Result()
	for _, key := range keys {
		data, _ := r.client.HGetAll(ctx, key).Result()
		fmt.Println(data)
	}

	fmt.Println("\nUsers:")
	keys, _ = r.client.Keys(ctx, "user:*").Result()
	for _, key := range keys {
		data, _ := r.client.HGetAll(ctx, key).Result()
		fmt.Println(data)
	}

	fmt.Println("\nOrders:")
	keys, _ = r.client.Keys(ctx, "order:*").Result()
	for _, key := range keys {
		data, _ := r.client.HGetAll(ctx, key).Result()
		fmt.Println(data)
	}

	fmt.Println("\nCategories:", r.client.SMembers(ctx, "categories").Val())
	fmt.Println("\nRecently Viewed by User 1:", r.client.LRange(ctx, "recently_viewed:1", 0, -1).Val())
}
