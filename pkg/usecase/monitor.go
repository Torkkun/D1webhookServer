package usecase

import "app/pkg/domain"

func MonitorSuccessPrompt() *domain.Prompt {

	// retrun部分はそのうち分けて置き換えれるように
	return &domain.Prompt{
		Override: false,
		LastSimple: &domain.Simple{
			Speech: "接続完了しました",
			Text:   "成功",
		},
	}
}

func MonitorFailedPrompt() *domain.Prompt {
	return &domain.Prompt{
		Override: true,
		LastSimple: &domain.Simple{
			Speech: "接続に失敗しました。もう一度お試しください",
			Text:   "Internal Server Error",
		},
	}
}
