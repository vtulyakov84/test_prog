package main

import (
	"log"
	"path/filepath"

	"your-project/windows/auth"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Инициализация GTK (обязательно в main)
	gtk.Init(nil)

	// Создаём окно авторизации
	gladePath := filepath.Join("windows", "auth", "auth.glade")
	authWin, err := auth.NewAuthWindow(gladePath)
	if err != nil {
		log.Fatalf("Ошибка загрузки окна: %v", err)
	}

	// Показываем окно
	authWin.Show()

	// Запускаем главный цикл GTK
	gtk.Main()
}
