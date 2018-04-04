package main

import (
	"fmt"
	"sync"
)

// Dbase creates a database interface
type Dbase interface {
	add(entry string) int
	retrieve(key int) string
}

// DB is the database
type DB struct {
	data []string
}

func (db *DB) add(entry string) int {
	return 0
}

func (db *DB) retrieve(key int) string {
	return "none"
}

func add(itf Dbase, entry string) int {
	return itf.add(entry)
}

var mydbase Dbase

type shape interface {
	draw()
	named() string
}

type circle struct {
	x      int
	y      int
	radius int
}

type square struct {
	x    int
	y    int
	side int
}

type mockShape struct {
	shape
	mutex sync.Mutex
}

func newMockShape() shape {
	return &mockShape{}
}

func (m *mockShape) draw() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println("mockShape calling draw")
}

func (m *mockShape) named() string {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println("mockShape calling named")
	return "mockShape:named"
}

func draw(s shape) {
	s.draw()
}

func named(s shape) string {
	return s.named()
}

var _ shape = (*mockShape)(nil)

type context struct {
	s shape
}

type config func() shape

func initContext(configs ...config) *context {
	ctx := &context{}
	for _, c := range configs {
		ctx.s = c()
	}
	return ctx
}

func main() {
	//var x = &DB{}
	//fmt.Println(add((Dbase)(x), "me"))

	//mocka := &mockShape{}
	//draw(mocka)
	//named(mocka)

	contexto := initContext(newMockShape)
	draw(contexto.s)
	named((*contexto).s)
}
