package requester

type CtxRequesterKeyType string

const CtxRequesterKey CtxRequesterKeyType = "context_requester"

type CtxRequester interface {
	GetUserId() int64
	GetUserRole() int64
}

var _ CtxRequester = (*ginCtxRequester)(nil)

type ginCtxRequester struct {
	UserId int64
	RoleId int64
}

func NewGinCtxRequester(id int64, role int64) CtxRequester {
	return ginCtxRequester{
		UserId: id,
		RoleId: role,
	}
}

func (u ginCtxRequester) GetUserId() int64 {
	return u.UserId
}

func (u ginCtxRequester) GetUserRole() int64 {
	return u.RoleId
}
