package helpers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"strconv"
)

func GetID(ctx *ext.Context) (int64, error) {
	msg := ctx.Message
	args := ctx.Args()
	if len(args) == 1 {
		if msg.ReplyToMessage != nil {
			return msg.ReplyToMessage.From.Id, nil
		} else {
			return msg.From.Id, nil
		}
	} else {
		r, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
}
