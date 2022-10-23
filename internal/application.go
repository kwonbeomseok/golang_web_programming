package internal

import (
	"errors"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	// UserName, MembershipType 중 하나라도 없으면 실패
	if request.UserName != "" && request.MembershipType != "" {
		// MembershipType이 naver, toss, payco 인지 확인
		if request.MembershipType == "naver" || request.MembershipType == "toss" || request.MembershipType == "payco" {
			_, ok := app.repository.data[request.UserName] //기존에 등록된 멤버인지 확인
			if !ok {
				var user Membership
				// 기존에 등록된 멤버 몇명인지 확인, ID 숫자로 할당....그러면 중간 숫자 아이디 사라지면 새로운 멤버 등록시 ID 중복됨....
				// 숫자 대신 다시 이름으로 ID 할당하자
				//var lengthMembership int = len(app.repository.data)
				//user = Membership{ID: strconv.Itoa(lengthMembership + 1), UserName: request.UserName, MembershipType: request.MembershipType}
				user = Membership{ID: request.UserName, UserName: request.UserName, MembershipType: request.MembershipType}
				app.repository.data[user.ID] = user
				return CreateResponse{user.ID, user.MembershipType}, nil
			}
		} else {
			return CreateResponse{request.UserName, request.MembershipType}, errors.New("Please Check MembershipType")
		}
		return CreateResponse{"", ""}, errors.New("UserName Already exist")
	}
	return CreateResponse{"", ""}, errors.New("Please Insert UserName or MembershipType")

}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	// update할 새로운 membership 생성
	var updateUser Membership
	updateUser = Membership{ID: request.UserName, UserName: request.UserName, MembershipType: request.MembershipType}
	// 기존 membership data에 update 진행
	app.repository.data[request.ID] = updateUser
	return UpdateResponse{ID: updateUser.ID, UserName: updateUser.UserName, MembershipType: updateUser.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}
