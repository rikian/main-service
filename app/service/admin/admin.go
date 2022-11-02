package admin

type AdminService interface{}

type adminImpl struct{}

func NewAdminService() AdminService {
	return &adminImpl{}
}
