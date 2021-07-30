package main

import "time"

func initMembers() []Member {
	members := make([]Member, 0)
	members = append(members, Member{ID: "1", Firstname: "Sasitorn", Lastname: "Parnsomboon", Nickname: "Joy", Birthday: Date(time.Date(1982, 11, 19, 0, 0, 0, 0, time.UTC))})
	members = append(members, Member{ID: "2", Firstname: "Suwicha", Lastname: "Kaewla-iad", Nickname: "Nuk", Birthday: Date(time.Date(1982, 11, 19, 0, 0, 0, 0, time.UTC))})
	members = append(members, Member{ID: "3", Firstname: "Okad", Lastname: "Kaewla-iad", Nickname: "", Birthday: Date(time.Date(1982, 11, 19, 0, 0, 0, 0, time.UTC))})
	return members
}