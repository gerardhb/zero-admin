// Code generated by goctl. DO NOT EDIT!
// Source: sys.proto

//go:generate mockgen -destination ./sys_mock.go -package sysclient -source $GOFILE

package sysclient

import (
	"context"

	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	LoginLogListData     = sys.LoginLogListData
	LoginLogDeleteResp   = sys.LoginLogDeleteResp
	SysLogDeleteReq      = sys.SysLogDeleteReq
	UserUpdateReq        = sys.UserUpdateReq
	MenuUpdateResp       = sys.MenuUpdateResp
	MenuListData         = sys.MenuListData
	DictUpdateResp       = sys.DictUpdateResp
	DeptListData         = sys.DeptListData
	UpdateRoleRoleReq    = sys.UpdateRoleRoleReq
	MenuAddReq           = sys.MenuAddReq
	UpdateRoleRoleResp   = sys.UpdateRoleRoleResp
	DictListData         = sys.DictListData
	DeptListReq          = sys.DeptListReq
	LoginLogListReq      = sys.LoginLogListReq
	SysLogDeleteResp     = sys.SysLogDeleteResp
	ConfigListReq        = sys.ConfigListReq
	UserDeleteReq        = sys.UserDeleteReq
	RoleAddResp          = sys.RoleAddResp
	ConfigListData       = sys.ConfigListData
	RoleListReq          = sys.RoleListReq
	DictListReq          = sys.DictListReq
	DictListResp         = sys.DictListResp
	DeptUpdateReq        = sys.DeptUpdateReq
	LoginLogListResp     = sys.LoginLogListResp
	UserAddResp          = sys.UserAddResp
	UserListResp         = sys.UserListResp
	DeptDeleteResp       = sys.DeptDeleteResp
	ConfigAddReq         = sys.ConfigAddReq
	InfoReq              = sys.InfoReq
	DeptListResp         = sys.DeptListResp
	MenuDeleteReq        = sys.MenuDeleteReq
	UpdateMenuRoleResp   = sys.UpdateMenuRoleResp
	DictAddReq           = sys.DictAddReq
	UpdateConfigRoleResp = sys.UpdateConfigRoleResp
	UpdateUserRoleReq    = sys.UpdateUserRoleReq
	MenuUpdateReq        = sys.MenuUpdateReq
	UpdateMenuRoleReq    = sys.UpdateMenuRoleReq
	LoginLogAddReq       = sys.LoginLogAddReq
	SysLogAddReq         = sys.SysLogAddReq
	SysLogListReq        = sys.SysLogListReq
	SysLogListResp       = sys.SysLogListResp
	ConfigUpdateReq      = sys.ConfigUpdateReq
	MenuListReq          = sys.MenuListReq
	MenuDeleteResp       = sys.MenuDeleteResp
	ConfigUpdateResp     = sys.ConfigUpdateResp
	DeptAddResp          = sys.DeptAddResp
	SysLogAddResp        = sys.SysLogAddResp
	RoleDeleteResp       = sys.RoleDeleteResp
	DictUpdateReq        = sys.DictUpdateReq
	DictAddResp          = sys.DictAddResp
	DeptAddReq           = sys.DeptAddReq
	DeptUpdateResp       = sys.DeptUpdateResp
	LoginLogDeleteReq    = sys.LoginLogDeleteReq
	SysLogListData       = sys.SysLogListData
	UpdateUserRoleResp   = sys.UpdateUserRoleResp
	RoleUpdateReq        = sys.RoleUpdateReq
	UserDeleteResp       = sys.UserDeleteResp
	UserStatusReq        = sys.UserStatusReq
	RoleDeleteReq        = sys.RoleDeleteReq
	UpdateConfigRoleReq  = sys.UpdateConfigRoleReq
	InfoResp             = sys.InfoResp
	UserUpdateResp       = sys.UserUpdateResp
	ReSetPasswordResp    = sys.ReSetPasswordResp
	RoleListResp         = sys.RoleListResp
	RoleUpdateResp       = sys.RoleUpdateResp
	ConfigListResp       = sys.ConfigListResp
	LoginResp            = sys.LoginResp
	UserListReq          = sys.UserListReq
	ReSetPasswordReq     = sys.ReSetPasswordReq
	MenuAddResp          = sys.MenuAddResp
	DeptDeleteReq        = sys.DeptDeleteReq
	ConfigDeleteResp     = sys.ConfigDeleteResp
	RoleAddReq           = sys.RoleAddReq
	RoleListData         = sys.RoleListData
	LoginReq             = sys.LoginReq
	MenuListResp         = sys.MenuListResp
	UserStatusResp       = sys.UserStatusResp
	DictDeleteReq        = sys.DictDeleteReq
	DictDeleteResp       = sys.DictDeleteResp
	LoginLogAddResp      = sys.LoginLogAddResp
	ConfigAddResp        = sys.ConfigAddResp
	ConfigDeleteReq      = sys.ConfigDeleteReq
	UserAddReq           = sys.UserAddReq
	UserListData         = sys.UserListData

	Sys interface {
		Login(ctx context.Context, in *LoginReq) (*LoginResp, error)
		UserInfo(ctx context.Context, in *InfoReq) (*InfoResp, error)
		UserAdd(ctx context.Context, in *UserAddReq) (*UserAddResp, error)
		UserList(ctx context.Context, in *UserListReq) (*UserListResp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq) (*UserUpdateResp, error)
		UserDelete(ctx context.Context, in *UserDeleteReq) (*UserDeleteResp, error)
		UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq) (*UpdateUserRoleResp, error)
		ReSetPassword(ctx context.Context, in *ReSetPasswordReq) (*ReSetPasswordResp, error)
		UpdateUserStatus(ctx context.Context, in *UserStatusReq) (*UserStatusResp, error)
		RoleAdd(ctx context.Context, in *RoleAddReq) (*RoleAddResp, error)
		RoleList(ctx context.Context, in *RoleListReq) (*RoleListResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq) (*RoleUpdateResp, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq) (*RoleDeleteResp, error)
		UpdateRoleRole(ctx context.Context, in *UpdateRoleRoleReq) (*UpdateRoleRoleResp, error)
		MenuAdd(ctx context.Context, in *MenuAddReq) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq) (*MenuUpdateResp, error)
		MenuDelete(ctx context.Context, in *MenuDeleteReq) (*MenuDeleteResp, error)
		UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq) (*UpdateMenuRoleResp, error)
		DictAdd(ctx context.Context, in *DictAddReq) (*DictAddResp, error)
		DictList(ctx context.Context, in *DictListReq) (*DictListResp, error)
		DictUpdate(ctx context.Context, in *DictUpdateReq) (*DictUpdateResp, error)
		DictDelete(ctx context.Context, in *DictDeleteReq) (*DictDeleteResp, error)
		DeptAdd(ctx context.Context, in *DeptAddReq) (*DeptAddResp, error)
		DeptList(ctx context.Context, in *DeptListReq) (*DeptListResp, error)
		DeptUpdate(ctx context.Context, in *DeptUpdateReq) (*DeptUpdateResp, error)
		DeptDelete(ctx context.Context, in *DeptDeleteReq) (*DeptDeleteResp, error)
		LoginLogAdd(ctx context.Context, in *LoginLogAddReq) (*LoginLogAddResp, error)
		LoginLogList(ctx context.Context, in *LoginLogListReq) (*LoginLogListResp, error)
		LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq) (*LoginLogDeleteResp, error)
		SysLogAdd(ctx context.Context, in *SysLogAddReq) (*SysLogAddResp, error)
		SysLogList(ctx context.Context, in *SysLogListReq) (*SysLogListResp, error)
		SysLogDelete(ctx context.Context, in *SysLogDeleteReq) (*SysLogDeleteResp, error)
		ConfigAdd(ctx context.Context, in *ConfigAddReq) (*ConfigAddResp, error)
		ConfigList(ctx context.Context, in *ConfigListReq) (*ConfigListResp, error)
		ConfigUpdate(ctx context.Context, in *ConfigUpdateReq) (*ConfigUpdateResp, error)
		ConfigDelete(ctx context.Context, in *ConfigDeleteReq) (*ConfigDeleteResp, error)
		UpdateConfigRole(ctx context.Context, in *UpdateConfigRoleReq) (*UpdateConfigRoleResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

func (m *defaultSys) Login(ctx context.Context, in *LoginReq) (*LoginResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.Login(ctx, in)
}

func (m *defaultSys) UserInfo(ctx context.Context, in *InfoReq) (*InfoResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserInfo(ctx, in)
}

func (m *defaultSys) UserAdd(ctx context.Context, in *UserAddReq) (*UserAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserAdd(ctx, in)
}

func (m *defaultSys) UserList(ctx context.Context, in *UserListReq) (*UserListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserList(ctx, in)
}

func (m *defaultSys) UserUpdate(ctx context.Context, in *UserUpdateReq) (*UserUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserUpdate(ctx, in)
}

func (m *defaultSys) UserDelete(ctx context.Context, in *UserDeleteReq) (*UserDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserDelete(ctx, in)
}

func (m *defaultSys) UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq) (*UpdateUserRoleResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UpdateUserRole(ctx, in)
}

func (m *defaultSys) ReSetPassword(ctx context.Context, in *ReSetPasswordReq) (*ReSetPasswordResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.ReSetPassword(ctx, in)
}

func (m *defaultSys) UpdateUserStatus(ctx context.Context, in *UserStatusReq) (*UserStatusResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in)
}

func (m *defaultSys) RoleAdd(ctx context.Context, in *RoleAddReq) (*RoleAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleAdd(ctx, in)
}

func (m *defaultSys) RoleList(ctx context.Context, in *RoleListReq) (*RoleListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleList(ctx, in)
}

func (m *defaultSys) RoleUpdate(ctx context.Context, in *RoleUpdateReq) (*RoleUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in)
}

func (m *defaultSys) RoleDelete(ctx context.Context, in *RoleDeleteReq) (*RoleDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleDelete(ctx, in)
}

func (m *defaultSys) UpdateRoleRole(ctx context.Context, in *UpdateRoleRoleReq) (*UpdateRoleRoleResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UpdateRoleRole(ctx, in)
}

func (m *defaultSys) MenuAdd(ctx context.Context, in *MenuAddReq) (*MenuAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuAdd(ctx, in)
}

func (m *defaultSys) MenuList(ctx context.Context, in *MenuListReq) (*MenuListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuList(ctx, in)
}

func (m *defaultSys) MenuUpdate(ctx context.Context, in *MenuUpdateReq) (*MenuUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in)
}

func (m *defaultSys) MenuDelete(ctx context.Context, in *MenuDeleteReq) (*MenuDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuDelete(ctx, in)
}

func (m *defaultSys) UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq) (*UpdateMenuRoleResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UpdateMenuRole(ctx, in)
}

func (m *defaultSys) DictAdd(ctx context.Context, in *DictAddReq) (*DictAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictAdd(ctx, in)
}

func (m *defaultSys) DictList(ctx context.Context, in *DictListReq) (*DictListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictList(ctx, in)
}

func (m *defaultSys) DictUpdate(ctx context.Context, in *DictUpdateReq) (*DictUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictUpdate(ctx, in)
}

func (m *defaultSys) DictDelete(ctx context.Context, in *DictDeleteReq) (*DictDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictDelete(ctx, in)
}

func (m *defaultSys) DeptAdd(ctx context.Context, in *DeptAddReq) (*DeptAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptAdd(ctx, in)
}

func (m *defaultSys) DeptList(ctx context.Context, in *DeptListReq) (*DeptListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptList(ctx, in)
}

func (m *defaultSys) DeptUpdate(ctx context.Context, in *DeptUpdateReq) (*DeptUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptUpdate(ctx, in)
}

func (m *defaultSys) DeptDelete(ctx context.Context, in *DeptDeleteReq) (*DeptDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptDelete(ctx, in)
}

func (m *defaultSys) LoginLogAdd(ctx context.Context, in *LoginLogAddReq) (*LoginLogAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogAdd(ctx, in)
}

func (m *defaultSys) LoginLogList(ctx context.Context, in *LoginLogListReq) (*LoginLogListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogList(ctx, in)
}

func (m *defaultSys) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq) (*LoginLogDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogDelete(ctx, in)
}

func (m *defaultSys) SysLogAdd(ctx context.Context, in *SysLogAddReq) (*SysLogAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysLogAdd(ctx, in)
}

func (m *defaultSys) SysLogList(ctx context.Context, in *SysLogListReq) (*SysLogListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysLogList(ctx, in)
}

func (m *defaultSys) SysLogDelete(ctx context.Context, in *SysLogDeleteReq) (*SysLogDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysLogDelete(ctx, in)
}

func (m *defaultSys) ConfigAdd(ctx context.Context, in *ConfigAddReq) (*ConfigAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.ConfigAdd(ctx, in)
}

func (m *defaultSys) ConfigList(ctx context.Context, in *ConfigListReq) (*ConfigListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.ConfigList(ctx, in)
}

func (m *defaultSys) ConfigUpdate(ctx context.Context, in *ConfigUpdateReq) (*ConfigUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.ConfigUpdate(ctx, in)
}

func (m *defaultSys) ConfigDelete(ctx context.Context, in *ConfigDeleteReq) (*ConfigDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.ConfigDelete(ctx, in)
}

func (m *defaultSys) UpdateConfigRole(ctx context.Context, in *UpdateConfigRoleReq) (*UpdateConfigRoleResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UpdateConfigRole(ctx, in)
}
