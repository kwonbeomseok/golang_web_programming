package internal

import (
	"errors"
	"strconv"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) CheckMemberExist(UserName string) bool {
	for _, val := range app.repository.data {
		if val.UserName == UserName {
			return true
		}
	}
	return false
}

func (app *Application) CheckIDExist(id string) bool {
	for _, val := range app.repository.data {
		if val.ID == id {
			return true
		}
	}
	return false
}

func (app *Application) CheckIDAndName(UserID string, UserName string) bool {
	for _, val := range app.repository.data {
		if (val.UserName == UserName) && (val.ID != UserID) {
			return true
		}
	}
	return false
}

func (app *Application) MakeNewID() string {
	if len(app.repository.data) == 0 {
		return "1"
	}

	var lastID string
	var newID int
	for _, val := range app.repository.data {
		lastID = val.ID
	}
	newID, _ = strconv.Atoi(lastID)
	newID = newID + 1
	return strconv.Itoa(newID)
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	// UserName, MembershipType 중 하나라도 없으면 실패
	if request.UserName == "" || request.MembershipType == "" {
		return CreateResponse{"", ""}, errors.New("Please Insert UserName or MembershipType")
	}

	// MembershipType이 naver, toss, payco 인지 확인
	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return CreateResponse{request.UserName, request.MembershipType}, errors.New("Please Check MembershipType")
	}

	// 기존에 등록된 멤버인지 확인
	check := app.CheckMemberExist(request.UserName)
	// 기존에 등록된 멤버이면 실패
	if check {
		return CreateResponse{"", ""}, errors.New("기존에 존재하는 멤버")
	}

	// 기존에 없는 멤버이면 등록
	var user Membership
	newID := app.MakeNewID() // 마지막 멤버 ID + 1
	user = Membership{ID: newID, UserName: request.UserName, MembershipType: request.MembershipType}
	app.repository.Create(user)
	return CreateResponse{user.ID, user.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	// 멤버십 아이디 or 사용자 이름 or 멤버십 타입 미입력 -> 에러
	if request.ID == "" || request.UserName == "" || request.MembershipType == "" {
		return UpdateResponse{}, errors.New("멤버십 아이디, 사용자 이름, 멤버십타입을 모두 입력해주세요")
	}

	// MembershipType이 naver, toss, payco 인지 확인
	if request.MembershipType != "naver" && request.MembershipType != "toss" && request.MembershipType != "payco" {
		return UpdateResponse{}, errors.New("존재하지 않는 멤버십타입입니다.")
	}

	// 기존에 존재하는 멤버이름인지 확인
	check := app.CheckIDAndName(request.ID, request.UserName)

	// 기존에 존재하는 멤버이름 -> 에러
	if check {
		return UpdateResponse{}, errors.New("기존에 존재하는 멤버 이름입니다")
	}

	// 기존에 존재하지 않는 멤버이름 -> 수정 진행
	var updateUser Membership
	updateUser = Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType}
	app.repository.Update(updateUser)
	return UpdateResponse{ID: app.repository.data[request.ID].ID,
			UserName:       app.repository.data[request.ID].UserName,
			MembershipType: app.repository.data[request.ID].MembershipType},
		nil
}

func (app *Application) Delete(id string) error {
	// ID 미입력시 -> 에러
	if id == "" {
		return errors.New("ID를 입력하세요")
	}

	// 존재하는 ID인지 확인
	check := app.CheckIDExist(id)

	// 존재하지 않는 ID -> 에러
	if !check {
		return errors.New("ID를 확인해주세요")
	}

	app.repository.Delete(id)
	return nil
}

func (app *Application) Check(id string) (Membership, error) {
	// 존재하는 ID인지 확인
	check := app.CheckIDExist(id)

	// 존재하지 않는 ID -> 에러
	if !check {
		return Membership{}, errors.New("존재하지 않는 ID입니다.")
	}
	return app.repository.Check(id), nil

}
