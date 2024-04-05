package cmd

import (
	"os"        // Используется для взаимодействия с операционной системой
	"os/signal" // Пакет signal используется для обработки сигналов ОС
	"syscall"   // Пакет syscall предоставляет интерфейс к низкоуровневым системным вызовам

	"github.com/spf13/cobra" // Cobra - популярная библиотека для создания CLI-приложений
)

var (
	// cfgPath хранит путь к файлу конфигурации, значение по умолчанию - "config.yaml"
	cfgPath string

	// stopNotification используется для уведомления о необходимости остановки приложения
	stopNotification = make(chan struct{})

	// rootCmd описывает корневую команду CLI-приложения
	rootCmd = &cobra.Command{
		Use:           "scraper [command]",      // Как использовать команду
		Long:          "eda.ru scraper service", // Длинное описание команды
		SilenceUsage:  true,                     // Не показывать сообщение usage при возникновении ошибки
		SilenceErrors: true,                     // Не показывать ошибки, они будут обработаны самостоятельно
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// PersistentPreRunE выполняется до основного обработчика команды и во всех дочерних командах
			go func() {
				var c = make(chan os.Signal, 1) // Создание канала для сигналов ОС
				// Настройка приема системных сигналов SIGHUP, SIGINT, SIGTERM
				signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
				<-c                            // Ожидание сигнала
				stopNotification <- struct{}{} // Отправка уведомления о необходимости остановки
			}()
			return nil
		},
	}
)

// Execute запускает корневую команду приложения
func Execute() error {
	// Настройка флага "config" для корневой команды
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "config.yaml", "config file")
	if err := rootCmd.Execute(); err != nil {
		return err // Возврат ошибки, если выполнение команды завершилось неудачей
	}
	return nil
}
