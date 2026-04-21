package auth

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

// AuthWindow - структура, представляющая окно авторизации
type AuthWindow struct {
	window      *gtk.Window
	userEntry   *gtk.Entry
	passEntry   *gtk.Entry
	loginButton *gtk.Button
}

// NewAuthWindow - конструктор окна авторизации
func NewAuthWindow(gladePath string) (*AuthWindow, error) {
	// Инициализация GTK (обычно делается в main, но можно и здесь)
	// gtk.Init(nil) // НЕ вызывайте здесь, иначе main не сможет повторно инициализировать

	// Загружаем Glade файл через Builder
	builder, err := gtk.BuilderNewFromFile(gladePath)
	if err != nil {
		return nil, err
	}

	// Получаем главное окно
	obj, err := builder.GetObject("auth_window")
	if err != nil {
		return nil, err
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal("Не удалось привести объект к типу *gtk.Window")
	}

	// Получаем поле для ввода имени пользователя
	obj, err = builder.GetObject("username_entry")
	if err != nil {
		return nil, err
	}
	userEntry, ok := obj.(*gtk.Entry)
	if !ok {
		log.Fatal("Не удалось привести объект к типу *gtk.Entry")
	}

	// Получаем поле для ввода пароля
	obj, err = builder.GetObject("password_entry")
	if err != nil {
		return nil, err
	}
	passEntry, ok := obj.(*gtk.Entry)
	if !ok {
		log.Fatal("Не удалось привести объект к типу *gtk.Entry")
	}

	// Получаем кнопку входа
	obj, err = builder.GetObject("login_button")
	if err != nil {
		return nil, err
	}
	loginButton, ok := obj.(*gtk.Button)
	if !ok {
		log.Fatal("Не удалось привести объект к типу *gtk.Button")
	}

	auth := &AuthWindow{
		window:      window,
		userEntry:   userEntry,
		passEntry:   passEntry,
		loginButton: loginButton,
	}

	// Настраиваем обработчики событий
	auth.loginButton.Connect("clicked", auth.onLoginClicked)

	// Обработка закрытия окна
	auth.window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return auth, nil
}

// onLoginClicked - приватный метод с логикой авторизации
func (a *AuthWindow) onLoginClicked() {
	username, err := a.userEntry.GetText()
	if err != nil {
		log.Printf("Ошибка получения имени пользователя: %v", err)
		return
	}

	password, err := a.passEntry.GetText()
	if err != nil {
		log.Printf("Ошибка получения пароля: %v", err)
		return
	}

	// Здесь ваша логика проверки (временная заглушка)
	if username == "admin" && password == "12345" {
		log.Println("Авторизация успешна!")
		a.window.Destroy()
		// Здесь можно открыть главное окно
	} else {
		log.Println("Ошибка авторизации!")
		a.showError("Неверное имя пользователя или пароль")
	}
}

// showError - приватный метод для показа диалога ошибки
func (a *AuthWindow) showError(message string) {
	dialog := gtk.MessageDialogNew(
		a.window,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_ERROR,
		gtk.BUTTONS_OK,
		message,
	)
	dialog.Run()
	dialog.Destroy()
}

// Show - публичный метод для отображения окна
func (a *AuthWindow) Show() {
	a.window.ShowAll()
}

// Hide - публичный метод для скрытия окна
func (a *AuthWindow) Hide() {
	a.window.Hide()
}

// GetWindow - если нужен доступ к самому окну (опционально)
func (a *AuthWindow) GetWindow() *gtk.Window {
	return a.window
}
