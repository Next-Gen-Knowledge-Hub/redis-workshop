package imp

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Widget struct {
	client *redis.Client
}

func NewWidget() *Widget {
	w := Widget{
		client: redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   4,
		}),
	}

	return &w
}

func (w *Widget) SendWidget(ctx context.Context, widget *WidgetModel) (err error) {
	wid := uuid.NewString()[:7]

	pip := w.client.TxPipeline()

	widgetKey := fmt.Sprintf("wid:%s", wid)
	err = pip.Set(ctx, widgetKey, widget.MarshalString(), 0).Err()
	if err != nil {
		return err
	}

	messageKey := fmt.Sprintf("msgs:%s", widget.UserId)
	err = pip.ZAdd(ctx, messageKey, redis.Z{
		Score:  float64(widget.Priority),
		Member: wid,
	}).Err()
	if err != nil {
		return err
	}

	_, err = pip.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (w *Widget) GetWidgets(ctx context.Context, uid string, count int) (widgets []WidgetModel, err error) {
	widgets = make([]WidgetModel, 0)

	messageKey := fmt.Sprintf("msgs:%s", uid)
	messages, err := w.client.ZPopMin(ctx, messageKey, int64(count)).Result()
	if err != nil {
		return nil, err
	}

	for _, m := range messages {
		widgetKey := fmt.Sprintf("wid:%s", m.Member.(string))
		var v string
		v, err = w.client.Get(ctx, widgetKey).Result()
		if err != nil && err != redis.Nil {
			return nil, err
		}

		widget := FromString(v)
		if widget == nil {
			continue
		}

		widgets = append(widgets, *widget)

		_ = w.client.Del(ctx, widgetKey)
	}

	return widgets, nil
}

type WidgetModel struct {
	UserId   string `json:"uid"`
	Priority int    `json:"im"`
	Message  string `json:"msg"`
}

func (w *WidgetModel) MarshalString() string {
	return fmt.Sprintf("%s;%d;%s", w.UserId, w.Priority, w.Message)
}

func FromString(v string) *WidgetModel {
	p := strings.Split(v, ";")
	if len(p) < 3 {
		return nil
	}

	pr, _ := strconv.Atoi(p[1])

	w := WidgetModel{
		UserId:   p[0],
		Priority: pr,
		Message:  p[2],
	}

	return &w
}
