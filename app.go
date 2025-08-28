package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/atotto/clipboard"
)

// App struct
type App struct {
	ctx context.Context
}

var store *Store

// NewApp creates a new App application struct
func NewApp() *App {
	go clipboardListener()

	var err error
	store, err = NewStore(context.Background())

	if err != nil {
		log.Fatal("Erro ao abrir banco de dados:", err)
	}

	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetCurrentTime returns the current time as a formatted string
func (a *App) GetCurrentTime() string {
	return fmt.Sprintf("Hora atual: %s", time.Now().Format("15:04:05"))
}

var lastClipboardText string

func clipboardListener() {
	for {
		// lê texto do clipboard
		text, err := clipboard.ReadAll()
		if err != nil {
			log.Println("Erro ao ler clipboard:", err)
			continue
		}

		// grava com timestamp
		if text != "" && text != lastClipboardText {
			lastClipboardText = text
			id, err := store.Add(text)
			if err != nil {
				log.Println("Erro ao salvar no banco:", err)
				continue
			}
			fmt.Println("Novo texto copiado:", text, "id:", id)
		}

		time.Sleep(1 * time.Second) // verifica a cada 2s
	}
}

// GetClipboarText
func (a *App) GetClipboarText() string {
	return lastClipboardText
}

func (a *App) SetClipboarText(text string) error {
	return clipboard.WriteAll(text)
}

func (a *App) SaveClip(text string) (int64, error) {
	id, err := store.Add(text)
	fmt.Println("Salvo id:", id, "erro:", err)
	return id, err
}

func (a *App) RemoveClip(id int64) error {
	return store.Remove(id)
}

func (a *App) GetLatestClips(skip, n int) ([]Clip, error) {
	return store.Latest(skip, n)
}

func (a *App) GetClipsAfter(ts string) ([]Clip, error) {
	return store.After(ts)
}
