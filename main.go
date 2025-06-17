package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Done      bool               `bson:"done" json:"done"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

var notesCollection *mongo.Collection
var ctx = context.Background()

func main() {
	// Koneksi ke MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	notesCollection = client.Database("todolist").Collection("notes")

	// Routing
	http.HandleFunc("/notes", corsMiddleware(notesHandler))
	http.HandleFunc("/notes/", corsMiddleware(noteDeleteHandler))

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set header CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNotes(w, r)
	case http.MethodPost:
		createNote(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	cursor, err := notesCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Gagal ambil data", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var notes []Note
	for cursor.Next(ctx) {
		var note Note
		cursor.Decode(&note)
		notes = append(notes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	note.CreatedAt = time.Now()
	note.Done = false

	result, err := notesCollection.InsertOne(ctx, note)
	if err != nil {
		http.Error(w, "Gagal simpan data", http.StatusInternalServerError)
		return
	}

	note.ID = result.InsertedID.(primitive.ObjectID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func noteDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/notes/")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	result, err := notesCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil || result.DeletedCount == 0 {
		http.Error(w, "Note tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Note berhasil dihapus"})
}
