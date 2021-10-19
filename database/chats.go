package database

import "errors"

type Chat struct {
	ChatID     int64  `gorm:"primaryKey" gorm:"column:chat_id"`
	AutoBan    bool   `gorm:"column:auto_ban"`
	SpamDetect bool   `gorm:"column:spam_detect"`
	SpamAction string `gorm:"column:spam_action"`
}

func (c *Chat) DoesAutoBan() bool {
	if c == nil {
		return false
	}
	return c.AutoBan
}

func (c *Chat) DetectSpam() bool {
	if c == nil {
		return false
	}
	return c.SpamDetect
}

func InsertChat(ChatID int64, AutoBan bool, SpamDetect bool, SpamAction string) {
	tx := SESSION.Begin()
	chat := &Chat{ChatID: ChatID, AutoBan: AutoBan, SpamDetect: SpamDetect, SpamAction: SpamAction}
	tx.Save(chat)
	tx.Commit()
}

func GetChat(ChatID int64) (*Chat, error) {
	if SESSION == nil {
		return nil, errors.New("cannot access to SESSION " +
			"of db, because it's nil")
	}

	p := Chat{}
	SESSION.Where("chat_id = ?", ChatID).Take(&p)
	return &p, nil

}
