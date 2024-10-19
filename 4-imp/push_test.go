package imp

import (
	"context"
	"testing"
)

func Test_SendSendPush(t *testing.T) {
	p := NewPushRedisRepo()
	uid := "user-1"
	err := p.SendPush(context.TODO(), uid, "test message 1")
	if err != nil {
		t.Error(err)
	}

	err = p.SendPush(context.TODO(), uid, "test message 2")
	if err != nil {
		t.Error(err)
	}

	err = p.SendPush(context.TODO(), uid, "test message 3")
	if err != nil {
		t.Error(err)
	}

	err = p.SendPush(context.TODO(), uid, "test message 4")
	if err != nil {
		t.Error(err)
	}
}

func Test_ReceivePush(t *testing.T) {
	p := NewPushRedisRepo()
	uid := "user-1"

	messages, err := p.GetPushs(context.TODO(), uid, 2)
	if err != nil {
		t.Error(err)
	}

	for i, m := range messages {
		t.Log(i, m)
	}
}
