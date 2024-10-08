package main

import (
	"bitcask"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var db *bitcask.DB

func init() {
	// 初始化 DB 实例
	var err error
	options := bitcask.DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-http")
	options.DirPath = dir
	db, err = bitcask.Open(options)
	if err != nil {
		panic(fmt.Sprintf("failed to open db: %v", err))
	}
}

func handlePut(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	for key, value := range data {
		if err := db.Put([]byte(key), []byte(value)); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Printf("failed to put value in db: %v\n", err)
			return
		}
	}
}

func handleGet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := request.URL.Query().Get("key")
	value, err := db.Get([]byte(key))
	if err != nil && err != bitcask.ErrKeyNotFound {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to get value in db: %v\n", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(string(value))
}

func handleDelete(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodDelete {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := request.URL.Query().Get("key")
	err := db.Delete([]byte(key))
	if err != nil && err != bitcask.ErrKeyNotFound {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to delete value in db: %v\n", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(string("Delete successfully"))
}

func handleListKeys(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	keys := db.ListKeys()
	writer.Header().Set("Content-Type", "application/json")
	var result []string
	for _, key := range keys {
		result = append(result, string(key))
	}
	_ = json.NewEncoder(writer).Encode(result)
}

func handleStat(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stat := db.Stat()
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(stat)
}

func handleMerge(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := db.Merge()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to merge db: %v\n", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write([]byte("Merge successfully"))
}

func handleBackup(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dir := request.URL.Query().Get("dir")
	if dir == "" {
		http.Error(writer, "dir is empty", http.StatusBadRequest)
		return
	}

	err := db.Backup(dir)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to backup db: %v\n", err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write([]byte("Backup successfully"))
}

func main() {
	// 注册处理方法
	http.HandleFunc("/bitcask/put", handlePut)
	http.HandleFunc("/bitcask/get", handleGet)
	http.HandleFunc("/bitcask/delete", handleDelete)
	http.HandleFunc("/bitcask/listkeys", handleListKeys)
	http.HandleFunc("/bitcask/stat", handleStat)
	http.HandleFunc("/bitcask/merge", handleMerge)
	http.HandleFunc("/bitcask/backup", handleBackup)

	// 启动 HTTP 服务
	http.ListenAndServe("localhost:8080", nil)
}
