package main

type Cacher interface {
	Get(int) (string, bool)
	Set(int, string) error
	Remove(int) error
}
