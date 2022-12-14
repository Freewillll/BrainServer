package models

import (
	"time"
)

type TAnotation struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	Name      string    `xorm:"not null unique VARCHAR(100)"`
	Soma      string    `xorm:"not null index VARCHAR(100)"`
	Owner     string    `xorm:"not null index VARCHAR(100)"`
	Isdeleted int       `xorm:"not null default 0 INT"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated"`
}

type TArbor struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	Name      string    `xorm:"not null unique VARCHAR(100)"`
	Somaid    string    `xorm:"not null unique(t_arbor_SomaId_X_Y_Z_uindex) VARCHAR(100)"`
	Image     string    `xorm:"not null index VARCHAR(100)"`
	X         float64   `xorm:"not null unique(t_arbor_SomaId_X_Y_Z_uindex) DECIMAL(10,3)"`
	Y         float64   `xorm:"not null unique(t_arbor_SomaId_X_Y_Z_uindex) DECIMAL(10,3)"`
	Z         float64   `xorm:"not null unique(t_arbor_SomaId_X_Y_Z_uindex) DECIMAL(10)"`
	Status    int       `xorm:"not null default 0 index(t_arbor_Is_deleted_Status_index) INT"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated"`
	Isdeleted int       `xorm:"not null default 0 index(t_arbor_Is_deleted_Status_index) INT"`
}

type TArbordetail struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	Arborid   int       `xorm:"not null index(t_arbordetail_ArborId_Isdeleted_index) unique(t_arbordetail_ArborId_X_Y_Z_Type_uindex) INT"`
	X         float64   `xorm:"not null unique(t_arbordetail_ArborId_X_Y_Z_Type_uindex) DECIMAL(10,3)"`
	Y         float64   `xorm:"not null unique(t_arbordetail_ArborId_X_Y_Z_Type_uindex) DECIMAL(10,3)"`
	Z         float64   `xorm:"not null unique(t_arbordetail_ArborId_X_Y_Z_Type_uindex) DECIMAL(10,3)"`
	Type      int       `xorm:"not null unique(t_arbordetail_ArborId_X_Y_Z_Type_uindex) INT"`
	Owner     string    `xorm:"not null index VARCHAR(200)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated"`
	Isdeleted int       `xorm:"not null default 0 index(t_arbordetail_ArborId_Isdeleted_index) INT"`
}

type TArborresult struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	Arborid   int       `xorm:"not null index(t_arborresult_ArborId_Iddeleted_index) INT"`
	Result    int       `xorm:"not null comment('?????????????????????') INT"`
	Form      int       `xorm:"not null comment('?????????????????????????????????') index INT"`
	Owner     string    `xorm:"not null index VARCHAR(200)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated"`
	Isdeleted int       `xorm:"not null default 0 index(t_arborresult_ArborId_Iddeleted_index) INT"`
}

type TEffectSoma struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	Name      string    `xorm:"not null unique VARCHAR(100)"`
	X         float64   `xorm:"not null unique(t_effect_soma_ImageId_X_Y_Z_uindex) DECIMAL(10,3)"`
	Y         float64   `xorm:"not null unique(t_effect_soma_ImageId_X_Y_Z_uindex) DECIMAL(10,3)"`
	Z         float64   `xorm:"not null unique(t_effect_soma_ImageId_X_Y_Z_uindex) DECIMAL(10,3)"`
	Image     string    `xorm:"not null unique(t_effect_soma_ImageId_X_Y_Z_uindex) index VARCHAR(100)"`
	From      int       `xorm:"not null default 0 comment('?????? 0:??????t_somainfo') INT"`
	Isdeleted int       `xorm:"not null default 0 INT"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated"`
}

type TImage struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	Name      string    `xorm:"not null unique VARCHAR(100)"`
	Detail    string    `xorm:"not null JSON"`
	Isdeleted int       `xorm:"not null default 0 INT"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated"`
}

type TPotentialsomalocation struct {
	Id        int       `xorm:"not null pk autoincr comment('??????') INT"`
	Image     string    `xorm:"not null comment('??????????????????') index unique(t_potentialsomalocation_loc) VARCHAR(100)"`
	X         int       `xorm:"not null unique(t_potentialsomalocation_loc) INT"`
	Y         int       `xorm:"not null unique(t_potentialsomalocation_loc) INT"`
	Z         int       `xorm:"not null unique(t_potentialsomalocation_loc) INT"`
	Type      int       `xorm:"default 0 comment('????????????????????? -1:???????????? 0:????????? 1:???soma 2:???soma') INT"`
	Dataset   int       `xorm:"default 0 comment('??????????????? 0:soma') INT"`
	Owner     string    `xorm:"not null default '' comment('?????????????????????') index VARCHAR(100)"`
	Isdeleted int       `xorm:"not null default 0 comment('???????????????0???????????????0??????') unique(t_potentialsomalocation_loc) INT"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('????????????') TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('????????????????????????') TIMESTAMP updated"`
}

type TSomainfo struct {
	Id        int       `xorm:"not null pk autoincr comment('??????') INT"`
	Name      string    `xorm:"not null comment('soma?????????????????????e.18454_00001') unique VARCHAR(100)"`
	Image     string    `xorm:"not null comment('?????????????????????') index unique(t_somainfo_loc) VARCHAR(100)"`
	X         float64   `xorm:"not null unique(t_somainfo_loc) DECIMAL(10,3)"`
	Y         float64   `xorm:"not null unique(t_somainfo_loc) DECIMAL(10,3)"`
	Z         float64   `xorm:"not null unique(t_somainfo_loc) DECIMAL(10,3)"`
	Location  int       `xorm:"not null comment('???????????????????????????') index INT"`
	Client    int       `xorm:"default 0 comment('???????????? 0:Hi5') INT"`
	Owner     string    `xorm:"not null comment('??????????????????') index VARCHAR(100)"`
	Updater   string    `xorm:"not null default '' comment('????????????') VARCHAR(100)"`
	Color     string    `xorm:"not null default '0000ff' comment('????????????????????????') VARCHAR(100)"`
	Status    int       `xorm:"not null default 0 comment('?????? 0:????????? 1:??????????????? 2:???????????????') index INT"`
	Methods   string    `xorm:"comment('refine???????????????') JSON"`
	Refinedid int       `xorm:"not null default 0 comment('refine?????????id') INT"`
	Isdeleted int       `xorm:"not null default 0 comment('???????????????0???????????????0??????') unique(t_somainfo_loc) INT"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('????????????') TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('????????????') TIMESTAMP updated"`
}

type TUserinfo struct {
	Id       int    `xorm:"not null pk autoincr INT"`
	Name     string `xorm:"not null unique VARCHAR(100)"`
	Email    string `xorm:"not null unique VARCHAR(100)"`
	Nickname string `xorm:"not null VARCHAR(100)"`
	Passwd   string `xorm:"not null VARCHAR(100)"`
	Score    int    `xorm:"not null default 0 index INT"`
	Appkey   string `xorm:"not null default '' comment('????????????appkey
') VARCHAR(100)"`
	Isdeleted int       `xorm:"not null default 0 INT"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated"`
}

type TGameUserinfo struct {
	Id        int       `xorm:"not null pk autoincr unique INT"`
	Name      string    `xorm:"not null unique VARCHAR(100)"`
	Email     string    `xorm:"not null unique VARCHAR(100)"`
	Passwd    string    `xorm:"not null VARCHAR(100)"`
	Score     int       `xorm:"not null default 0 comment('???????????????????????????') INT"`
	Isdeleted int       `xorm:"not null default 0 INT"`
	Ctime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('????????????') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('????????????') TIMESTAMP"`
}
