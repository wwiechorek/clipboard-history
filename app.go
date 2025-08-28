package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/atotto/clipboard"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"

	"golang.design/x/hotkey"
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

var icon []byte

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {

	a.ctx = ctx

	go hotKeyListener(a.ctx)
	// wailsRuntime.WindowSetPosition(a.ctx, 100, 100)

	trayMenu := menu.NewMenu()
	trayMenu.Append(menu.Text("Abrir janela", keys.CmdOrCtrl("O"), func(_ *menu.CallbackData) {
		wailsRuntime.WindowShow(ctx)
	}))
	trayMenu.Append(menu.Separator())
	trayMenu.Append(menu.Text("Sair", keys.CmdOrCtrl("Q"), func(_ *menu.CallbackData) {
		wailsRuntime.Quit(ctx)
	}))

	applyAccessoryPolicy()
}

func hotKeyListener(ctx context.Context) {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyV)

	err := hk.Register()
	if err != nil {
		panic(err)
	}

	for {
		<-hk.Keydown()
		wailsRuntime.WindowShow(ctx)

		time.Sleep(1 * time.Millisecond)
		CenterWindowOnMouseMonitor(ctx)
		// wailsRuntime.WindowSetPosition(ctx, 100, 100)
	}
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

func (a *App) HideApplication() {
	wailsRuntime.WindowHide(a.ctx)
}
