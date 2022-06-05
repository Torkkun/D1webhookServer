package usecase

import "app/pkg/domain"

func SuccessPrompt() *domain.Prompt {
	return &domain.Prompt{
		Override: false,
		LastSimple: &domain.Simple{
			Speech: "接続完了しました",
			Text:   "成功",
		},
	}
}

// ちゃんと表示されるのか未検証
func FailedPrompt() *domain.Prompt {
	return &domain.Prompt{
		Override: true,
		LastSimple: &domain.Simple{
			Speech: "接続に失敗しました。もう一度お試しください",
			Text:   "Internal Server Error",
		},
	}
}
