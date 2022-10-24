package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)

		req = CreateRequest{"jenny", "payco"}
		res, err = app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)

	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{MembershipType: "naver"}
		_, err := app.Create(req)
		assert.Nil(t, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{UserName: "jenny"}
		_, err := app.Create(req)
		assert.Nil(t, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{UserName: "jenny", MembershipType: "kakao"}
		_, err := app.Create(req)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("멤버십 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		updateReq := UpdateRequest{ID: "1", UserName: "jenny", MembershipType: "payco"}
		updateRes, err := app.Update(updateReq)
		assert.Nil(t, err)
		assert.NotEmpty(t, updateRes.ID)
		assert.Equal(t, updateReq.MembershipType, updateRes.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)
		req = CreateRequest{"beomseok", "payco"}
		_, _ = app.Create(req)

		updateReq := UpdateRequest{ID: "1", UserName: "beomseok", MembershipType: "payco"}
		_, err := app.Update(updateReq)
		assert.Nil(t, err)

	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		updateReq := UpdateRequest{UserName: "jenny", MembershipType: "payco"}
		_, err := app.Update(updateReq)
		assert.Nil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		updateReq := UpdateRequest{ID: "1", MembershipType: "payco"}
		_, err := app.Update(updateReq)
		assert.Nil(t, err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		updateReq := UpdateRequest{ID: "1", UserName: "jenny"}
		_, err := app.Update(updateReq)
		assert.Nil(t, err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)
		updateReq := UpdateRequest{ID: "1", UserName: "jenny", MembershipType: "kakao"}
		_, err := app.Update(updateReq)
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)
		err := app.Delete("1")
		assert.Nil(t, err)

	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		err := app.Delete("")
		assert.Nil(t, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		err := app.Delete("2")
		assert.Nil(t, err)
	})
}

func TestCheck(t *testing.T) {
	t.Run("멤버십 정보를 확인한다", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		res, err := app.Check("1")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		fmt.Println(res)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)

		res, err := app.Check("2")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})
}
