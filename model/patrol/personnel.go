// 自动生成模板Personnel
package patrol

import (
	"github.com/oldweipro/gin-admin/global"
)

// Personnel 结构体
type Personnel struct {
	global.Model
	PersonUrl                   string `json:"personUrl" form:"personUrl" gorm:"column:person_url;comment:人员URL;"`
	Specialty                   string `json:"specialty" form:"specialty" gorm:"column:specialty;comment:特长;"`
	UnitNo                      string `json:"unitNo" form:"unitNo" gorm:"column:unit_no;comment:单位编号;"`
	Nation                      string `json:"nation" form:"nation" gorm:"column:nation;comment:民族;"`
	CompanyName                 string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:单位名称;"`
	IsSecurityFocus             *int   `json:"isSecurityFocus" form:"isSecurityFocus" gorm:"column:is_security_focus;comment:是否重点人员;"`
	RoomId                      string `json:"roomId" form:"roomId" gorm:"column:room_id;comment:房间ID;"`
	FloorId                     string `json:"floorId" form:"floorId" gorm:"column:floor_id;comment:楼层ID;"`
	BuildingNo                  string `json:"buildingNo" form:"buildingNo" gorm:"column:building_no;comment:建筑物编号;"`
	Contact                     string `json:"contact" form:"contact" gorm:"column:contact;comment:联系方式;"`
	UserDefinedAddress          string `json:"userDefinedAddress" form:"userDefinedAddress" gorm:"column:user_defined_address;comment:自定义地址;"`
	HouseNo                     string `json:"houseNo" form:"houseNo" gorm:"column:house_no;comment:房屋编号;"`
	SecurityFocusDetail         string `json:"securityFocusDetail" form:"securityFocusDetail" gorm:"column:security_focus_detail;comment:重点人员详情;"`
	RelationshipWithHouseHolder string `json:"relationshipWithHouseHolder" form:"relationshipWithHouseHolder" gorm:"column:relationship_with_house_holder;comment:与户主关系;"`
	SpecialPersonMark           string `json:"specialPersonMark" form:"specialPersonMark" gorm:"column:special_person_mark;comment:特殊人员标记;"`
	BodyFeature                 string `json:"bodyFeature" form:"bodyFeature" gorm:"column:body_feature;comment:体格特征;"`
	Height                      *int   `json:"height" form:"height" gorm:"column:height;comment:身高;"`
	BloodTypeName               string `json:"bloodTypeName" form:"bloodTypeName" gorm:"column:blood_type_name;comment:血型名称;"`
	PictureUrl                  string `json:"pictureUrl" form:"pictureUrl" gorm:"column:picture_url;comment:人员图片URL;"`
	Birth                       string `json:"birth" form:"birth" gorm:"column:birth;comment:生日;"`
	IsUsualResidence            *int   `json:"isUsualResidence" form:"isUsualResidence" gorm:"column:is_usual_residence;comment:是否常住人口;"`
	BuildingName                string `json:"buildingName" form:"buildingName" gorm:"column:building_name;comment:建筑物名称;"`
	Nationality                 string `json:"nationality" form:"nationality" gorm:"column:nationality;comment:国籍;"`
	CardId                      string `json:"cardId" form:"cardId" gorm:"column:card_id;comment:身份证号;"`
	AuthExpireTime              string `json:"authExpireTime" form:"authExpireTime" gorm:"column:auth_expire_time;comment:认证过期时间;"`
	FloorNo                     string `json:"floorNo" form:"floorNo" gorm:"column:floor_no;comment:楼层编号;"`
	DataSource                  *int   `json:"dataSource" form:"dataSource" gorm:"column:data_source;comment:数据来源;"`
	MaritalStatus               string `json:"maritalStatus" form:"maritalStatus" gorm:"column:marital_status;comment:婚姻状况;"`
	Status                      string `json:"status" form:"status" gorm:"column:status;comment:状态;"`
	FailedReason                string `json:"failedReason" form:"failedReason" gorm:"column:failed_reason;comment:失败原因;"`
	DetailedCompanyName         string `json:"detailedCompanyName" form:"detailedCompanyName" gorm:"column:detailed_company_name;comment:详细单位名称;"`
	OtherContact                string `json:"otherContact" form:"otherContact" gorm:"column:other_contact;comment:其他联系方式;"`
	IsHouseHolder               *int   `json:"isHouseHolder" form:"isHouseHolder" gorm:"column:is_house_holder;comment:是否户主;"`
	CriminalRecordDetail        string `json:"criminalRecordDetail" form:"criminalRecordDetail" gorm:"column:criminal_record_detail;comment:刑事记录详情;"`
	CensusAddrDetailname        string `json:"censusAddrDetailname" form:"censusAddrDetailname" gorm:"column:census_addr_detailname;comment:户籍地详细地址;"`
	BloodType                   *int   `json:"bloodType" form:"bloodType" gorm:"column:blood_type;comment:血型;"`
	PrecinctId                  string `json:"precinctId" form:"precinctId" gorm:"column:precinct_id;comment:区域ID;"`
	InputerContact              string `json:"inputerContact" form:"inputerContact" gorm:"column:inputer_contact;comment:录入人联系方式;"`
	AliasNames                  string `json:"aliasNames" form:"aliasNames" gorm:"column:alias_names;comment:别名;"`
	UnitId                      string `json:"unitId" form:"unitId" gorm:"column:unit_id;comment:单位ID;"`
	PersonType                  string `json:"personType" form:"personType" gorm:"column:person_type;comment:人员类型;"`
	Address                     string `json:"address" form:"address" gorm:"column:address;comment:地址;"`
	Sex                         *int   `json:"sex" form:"sex" gorm:"column:sex;comment:性别;"`
	CardType                    *int   `json:"cardType" form:"cardType" gorm:"column:card_type;comment:证件类型;"`
	UpdateTime                  string `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`
	BuildingId                  string `json:"buildingId" form:"buildingId" gorm:"column:building_id;comment:建筑物ID;"`
	PersonName                  string `json:"personName" form:"personName" gorm:"column:person_name;comment:人员真实姓名;"`
	LiveRelationShip            string `json:"liveRelationShip" form:"liveRelationShip" gorm:"column:live_relation_ship;comment:居住关系;"`
	IsCriminalRecord            *int   `json:"isCriminalRecord" form:"isCriminalRecord" gorm:"column:is_criminal_record;comment:是否有刑事记录;"`
	CreateTime                  string `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
	Culture                     string `json:"culture" form:"culture" gorm:"column:culture;comment:文化程度;"`
	NotInRegisteredPlaceReason  string `json:"notInRegisteredPlaceReason" form:"notInRegisteredPlaceReason" gorm:"column:not_in_registered_place_reason;comment:不在注册地原因;"`
	RentTime                    string `json:"rentTime" form:"rentTime" gorm:"column:rent_time;comment:租赁时间;"`
	OriginName                  string `json:"originName" form:"originName" gorm:"column:origin_name;comment:老家名称来自哪里;"`
}

// TableName Personnel 表名
func (Personnel) TableName() string {
	return "personnel"
}
