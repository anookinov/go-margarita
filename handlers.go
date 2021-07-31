package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func getMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, m := range members {
		if m.ID == params["id"] {
			json.NewEncoder(w).Encode(m)
			return
		}
	}
	json.NewEncoder(w).Encode(&Member{})
}

func createMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Member
	_ = json.NewDecoder(r.Body).Decode(&m)
	m.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	members = append(members, m)
	json.NewEncoder(w).Encode(m)
}

func updateMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, m := range members {
		if m.ID == params["id"] {
			var member Member
			_ = json.NewDecoder(r.Body).Decode(&member)
			members[i] = member
			json.NewEncoder(w).Encode(member)
			return
		}
	}
	json.NewEncoder(w).Encode(members)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, m := range members {
		if m.ID == params["id"] {
			members = append(members[:i], members[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(members)
}
