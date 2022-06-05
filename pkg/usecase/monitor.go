package usecase

import "app/pkg/domain"

func MonitorPrompt() *domain.Prompt {

	// retrun部分はそのうち分けて置き換えれるように
	return &domain.Prompt{
		Override: false,
		LastSimple: &domain.Simple{
			Speech: "接続完了しました",
			Text:   "接続完了",
		},
	}
}
