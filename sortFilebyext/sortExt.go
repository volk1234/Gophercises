package main

import (
	"fmt"
	"path/filepath"
	"io/fs"
	"log"
	"sort"
)

func main() {
	dirPath := "/home/teth/go1225"
	stats := make(map[string] struct{
		count int
		size int64
	})

	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		if ext == "" {
			ext = "unknown"
		}
		if ext == ".jpeg" {
			ext = ".jpg"
		}
		st := stats[ext]
		st.count++
		st.size += info.Size()
		stats[ext] = st

		fmt.Println(stats)
		return nil
	})

	if err != nil {
		log.Fatalf("error walking dir: %v", err)
	}

	type statsEntry struct {
		ext string
		count int
		size int64
	}
	var entries []statsEntry

	for ext, st := range stats {	
		entries = append(entries, statsEntry{ext, st.count, st.size})
		fmt.Println(entries)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].size > entries[j].size
	})

	for _, entry := range entries {
		fmt.Printf("[%s] %d files, %d Kb\n", entry.ext, entry.count, entry.size/1024)
	}
}

