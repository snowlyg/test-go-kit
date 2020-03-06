package service

import (
	"context"
	"testing"
	time2 "time"
)

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}

func TestStatus(t *testing.T) {
	srv, ctx := setup()
	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// 测试 status
	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestGet(t *testing.T) {
	srv, ctx := setup()
	d, err := srv.Get(ctx)
	if err != nil {
		t.Errorf("Error:%s", err)
	}

	time := time2.Now()
	today := time.Format("02/01/2006")

	// 测试今天的日期
	ok := today == d
	if !ok {
		t.Errorf("expected dates to be equal")
	}
}

func TestValidate(t *testing.T) {
	srv, ctx := setup()
	b, err := srv.Validate(ctx, "31/12/2019")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// 测试日期是否有效
	if !b {
		t.Errorf("date should be valid")
	}

	// 测试无效日期
	b, err = srv.Validate(ctx, "31/31/2019")
	if b {
		t.Errorf("date should be invalid")
	}

	// 测试美国日期
	b, err = srv.Validate(ctx, "12/31/2019")
	if b {
		t.Errorf("USA date should be invalid")
	}
}
