package menu

import (
	"context"
	"errors"
	"fmt"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/dixonwille/wmenu/v5"
)

var (
	errExit = errors.New("exit signal")
)

type Menu struct {
	menu *wmenu.Menu

	accept   ports.AcceptManager
	show     ports.ShowManager
	alert    ports.AlertManager
	alertCfg ports.AlertConfigManager
}

func New(accept ports.AcceptManager, show ports.ShowManager, alert ports.AlertManager, alertCfg ports.AlertConfigManager) *Menu {
	return &Menu{
		menu:     wmenu.NewMenu("Выберите вариант: "),
		accept:   accept,
		show:     show,
		alert:    alert,
		alertCfg: alertCfg,
	}
}

func (m *Menu) Run(ctx context.Context) error {
	m.menu.Option("Создать группу", ctx, false, m.createGroup)
	m.menu.Option("Начать трейс", ctx, false, m.startTrace)
	m.menu.Option("Завершить трейс", ctx, false, m.endTrace)
	m.menu.Option("Отправить событие", ctx, false, m.sendEvent)
	m.menu.Option("Посмотреть список групп", ctx, false, m.groupList)
	m.menu.Option("Посмотреть список ошибок", ctx, false, m.eventList)
	m.menu.Option("Посмотреть событие", ctx, false, m.showEvent)
	m.menu.Option("Исправить событие", ctx, false, m.fixEvent)
	m.menu.Option("Добавить настройку уведомления", ctx, false, m.addAlertConfig)
	m.menu.Option("Посмотреть список настроек уведомлений", ctx, false, m.alertConfigList)
	m.menu.Option("Удалить настройку уведомлений", ctx, false, m.deleteAlertConfig)
	m.menu.Option("Посмотреть список уведомлений", ctx, false, m.alertList)
	m.menu.Option("Выход", ctx, false, func(_ wmenu.Opt) error {
		return errExit
	})

	for {
		err := m.menu.Run()
		if errors.Is(err, errExit) {
			break
		}
		if err != nil {
			fmt.Printf("Error: %s.\n\n", err)
			continue
		}

		fmt.Println("Success.")
		fmt.Println()
	}

	contextutils.Logger(ctx).Infow("exiting")

	return nil
}
