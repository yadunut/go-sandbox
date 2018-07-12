package main

import "testing"

func BenchmarkStrcmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strcmp("alalskjdfakjsdflksajfs", "ldfdl;fjla;lkasjdfa;sfj")
	}

}
