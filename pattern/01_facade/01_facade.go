package main

import "fmt"

/*
Паттерн «Фасад» предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы. Фасад определяет
интерфейс более высокого уровня, который упрощает использование подсистемы.

Основная идея фасада - скрыть сложность внутренней подсистемы и предоставить более простой интерфейс.
*/

// RegistrationSubsystem - регистрация нового пользователя. Тут вымышленные поля у структуры которые потом отправляются в вымышленную бд.
type RegistrationSubsystem struct{}

func (rs *RegistrationSubsystem) Register(username, password string) error {
	/* Логика регистрации нового пользователя. P.S. Очень сложная логика, тут вымышленно проверяем пароль на корректность
	(не был простым). Хэшируем пароль сохраняем в отдельную таблицу, другие данные указанные при регистрации тоже сохраняем
	в отдельную таблицу(Мы же не делетанты, логин/пароль не будем хранить с обычными данными(имя, фамилия и т.д.)),
	тут лучше не вникать ОЧИНЬ СЛОЖНА

	*/
	fmt.Printf("Регистрация пользователя %s\n", username)
	return nil
}

// AuthenticationSubsystem - проверка учетных данных
type AuthenticationSubsystem struct{}

func (as *AuthenticationSubsystem) Authenticate(username, password string) error {
	/* Логика проверки учетных данных. P.S. Тут очень сложно проверяем кеши с логинами и думаем разрешать входить или нет. Мы
	тут it-шники, это наша территория, так что по приколу можем еще кодик из смс попросить*/
	fmt.Printf("Проверка учетных данных для пользователя %s\n", username)
	return nil
}

// LogoutSubsystem - выход из системы
type LogoutSubsystem struct{}

func (ls *LogoutSubsystem) Logout(username string) error {
	// Логика выхода из системы P.S. Тут ничего сложного, здесь вымышлено выходим из системы у пользователя(как говорится люди приходят и уходят)
	fmt.Printf("Выход из системы для пользователя %s\n", username)
	return nil
}

// AuthFacade - фасад для авторизации
type AuthFacade struct {
	RegistrationSubsystem
	AuthenticationSubsystem
	LogoutSubsystem
}

// NewAuthFacade - создание нового экземпляра фасада
func NewAuthFacade() *AuthFacade {
	return &AuthFacade{}
}

// AuthUser - метод фасада для регистрации и аутентификации пользователя
func (af *AuthFacade) AuthUser(username, password string) error {
	err := af.Register(username, password)
	if err != nil {
		return err
	}

	erro := af.Authenticate(username, password)
	if erro != nil {
		return err
	}

	return nil
}

// LogoutUser - метод фасада для выхода пользователя из системы
func (af *AuthFacade) LogoutUser(username string) error {
	return af.Logout(username)
}

func main() {
	authFacade := NewAuthFacade()

	// Использование фасада для регистрации, аутентификации и выхода пользователя
	err := authFacade.AuthUser("asd123", "123")
	if err != nil {
		fmt.Println("Ошибка при аутентификации:", err)
		return
	}

	err = authFacade.LogoutUser("123")
	if err != nil {
		fmt.Println("Ошибка при выходе из системы:", err)
		return
	}
}
