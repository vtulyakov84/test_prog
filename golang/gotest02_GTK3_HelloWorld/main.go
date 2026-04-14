package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Инициализация GTK
	gtk.Init(&os.Args)

	// Создание главного окна
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Ошибка создания окна:", err)
	}

	// Настройка окна
	win.SetTitle("Hello World на Go + GTK")
	win.SetDefaultSize(400, 300)

	// Обработка закрытия окна
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Создание текстовой метки
	label, err := gtk.LabelNew("Привет, Мир! 👋")
	if err != nil {
		log.Fatal("Ошибка создания метки:", err)
	}

	// Создание кнопки
	button, err := gtk.ButtonNewWithLabel("Нажми меня!")
	if err != nil {
		log.Fatal("Ошибка создания кнопки:", err)
	}

	// Обработка нажатия кнопки
	button.Connect("clicked", func() {
		label.SetLabel("Кнопка нажата! 🎉")
	})

	// Контейнер для вертикального размещения элементов
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Ошибка создания контейнера:", err)
	}
	box.SetMarginStart(20)
	box.SetMarginEnd(20)
	box.SetMarginTop(20)
	box.SetMarginBottom(20)

	// Добавление элементов в контейнер
	box.PackStart(label, false, false, 0)
	box.PackStart(button, false, false, 0)

	// Добавление контейнера в окно
	win.Add(box)

	// Отображение всех элементов
	win.ShowAll()

	// Запуск главного цикла обработки событий
	gtk.Main()
}
