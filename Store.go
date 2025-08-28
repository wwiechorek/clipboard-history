package main

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Store struct{ DB *sql.DB }

func NewStore(ctx context.Context) (*Store, error) {
	exePath, _ := os.Executable()
	appName := filepath.Base(exePath)
	// salve no diretório de dados do app (Wails)
	appDir, _ := AppDataDir(appName)
	dsn := "file:" + appDir + "/clips.db?_busy_timeout=5000&_fk=1"

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	// ajustes recomendados
	if _, err := db.Exec(`
		PRAGMA journal_mode = WAL;
		PRAGMA foreign_keys = ON;
		PRAGMA synchronous = NORMAL;
	`); err != nil {
		return nil, err
	}

	// esquema (opção ISO)
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS clips (
		  id       INTEGER PRIMARY KEY AUTOINCREMENT,
		  ts_iso   TEXT    NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now')),
		  content  TEXT    NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_clips_ts ON clips(ts_iso DESC);
	`); err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}

// Inserir texto (timestamp automático)
func (s *Store) Add(content string) (int64, error) {
	res, err := s.DB.Exec(`INSERT INTO clips(content) VALUES (?)`, content)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (e *Store) Remove(id int64) error {
	_, err := e.DB.Exec(`DELETE FROM clips WHERE id = ?`, id)
	return err
}

// Buscar últimos N
type Clip struct {
	ID      int64
	TSISO   string
	Content string
}

func (s *Store) Latest(sk, n int) ([]Clip, error) {
	rows, err := s.DB.Query(`
		SELECT id, ts_iso, content
		FROM clips
		ORDER BY ts_iso DESC
		LIMIT ?, ?`, sk, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Clip
	for rows.Next() {
		var c Clip
		if err := rows.Scan(&c.ID, &c.TSISO, &c.Content); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// Converter ts_iso para time.Time quando precisar
func (c Clip) Time() (time.Time, error) {
	// Com %fZ usamos precisão de nanos; RFC3339Nano lida bem:
	return time.Parse(time.RFC3339Nano, c.TSISO)
}
