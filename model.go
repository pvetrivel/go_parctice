package test

import "time"

type damage struct {
	VehicleDamageID string `bson:"vehicleDamageID" json:"vehicleDamageID"`
	ImageURL        string `bson:"imageURL" json:"imageURL"`
	DamageType      string `bson:"damageType" json:"damageType"`
	Description     string `bson:"description" json:"description"`
	Priority        int    `bson:"priority" json:"priority"`
	Icon            string `bson:"icon" json:"icon"`
}

type address struct {
	DealerAddressID string `bson:"dealerAddressID" json:"dealerAddressID"`
	AddressType     int    `bson:"addressType"json:"addressType"`
	StreetAddress1  string `bson:"streetAddress1"json:"streetAddress1"`
	StreetAddress2  string `bson:"streetAddress2" json:"streetAddress2"`
	City            string `bson:"city" json:"city"`
	State           string `bson:"state" json:"state"`
	ZipCode         string `bson:"zipCode" json:"zipCode"`
	Country         string `bson:"country" json:"country"`
	County          string `bson:"county" json:"county"`
	IsActive        bool   `bson:"isActive" json:"isActive"`
	LocationURL     string `bson:"locationUrl" json:"locationUrl"`
	IsdCode         string `bson:"isdCode" json:"isdCode"`
}

type logo struct {
	Width   int    `bson:"width" json:"width"`
	Height  int    `bson:"height" json:"height"`
	Title   string `bson:"title" json:"title"`
	ImageID string `bson:"imageID" json:"imageID"`
}

type businessHours struct {
	OpeningTime string `bson:"openingTime" json:"openingTime"`
	ClosingTime string `bson:"closingTime" json:"closingTime"`
}

type dealerOperation struct {
	DealerOperationScheduleID string          `bson:"dealerOperationScheduleID" json:"dealerOperationScheduleID"`
	DealerOperationType       int             `bson:"dealerOperationType" json:"dealerOperationType"`
	BusinessHours             []businessHours `bson:"businessHours" json:"businessHours"`
}
type simpleFeature struct {
	ID        int      `bson:"id" json:"id"`
	Name      string   `bson:"name" json:"name"`
	IsEnabled bool     `bson:"is_enabled" json:"is_enabled"`
	Roles     []string `bson:"roles" json:"roles"`
}

type features struct {
	ID          int             `bson:"id" json:"id"`
	Name        string          `bson:"name" json:"name"`
	IsEnabled   bool            `bson:"is_enabled" json:"is_enabled"`
	SubFeatures []simpleFeature `bson:"subFeatures" json:"subFeatures,omitempty"`
}

type Dealer struct {
	ID                           string            `bson:"_id" json:"_id"`
	DealerName                   string            `bson:"dealerName" json:"dealerName"`
	MakeCode                     []string          `bson:"makeCode" json:"makes"`
	DealerDisplayName            string            `bson:"dealerDisplayName" json:"dealerDisplayName"`
	StateIssuedNumber            string            `bson:"stateIssuedNumber" json:"stateIssuedNumber"`
	ManufacturerIssuedNumber     string            `bson:"manufacturerIssuedNumber" json:"manufacturerIssuedNumber"`
	TenantID                     string            `bson:"tenantID" json:"tenantID"`
	Website                      string            `bson:"website" json:"website"`
	Phone                        string            `bson:"phone" json:"phone"`
	VehicleDamage                []damage          `bson:"vehicleDamage" json:"vehicleDamage"`
	TimeZone                     string            `bson:"timeZone" json:"timeZone"`
	Currency                     string            `bson:"currency" json:"currency"`
	DealershipCode               string            `bson:"dealershipCode" json:"dealershipCode"`
	DealerGroup                  []interface{}     `bson:"dealerGroup" json:"dealerGroup"`
	DealerAddress                []address         `bson:"dealerAddress" json:"dealerAddress"`
	DealerLogos                  []logo            `bson:"dealerLogos" json:"dealerLogos"`
	DealerOperationSchedule      []dealerOperation `bson:"dealerOperationSchedule" json:"dealerOperationSchedule"`
	DealerContact                []interface{}     `bson:"dealerContact" json:"dealerContact"`
	IsActive                     bool              `bson:"isActive" json:"isActive"`
	LastUpdatedByUser            string            `bson:"lastUpdatedByUser" json:"lastUpdatedByUser"`
	LastUpdatedByDisplayName     string            `bson:"lastUpdatedByDisplayName" json:"lastUpdatedByDisplayName"`
	CreateDateTime               time.Time         `bson:"createDateTime" json:"createDateTime"`
	LastUpdatedDateTime          time.Time         `bson:"lastUpdatedDateTime" json:"lastUpdatedDateTime"`
	DocumentVersion              float64           `bson:"documentVersion" json:"documentVersion"`
	DealerDoingBusinessAsName    string            `bson:"dealerDoingBusinessAsName" json:"dealerDoingBusinessAsName"`
	ApplicationCode              string            `bson:"applicationCode" json:"applicationCode"`
	LateAppointmentMins          int               `bson:"lateAppointmentMins" json:"lateAppointmentMins"`
	MissedAppointmentMins        int               `bson:"missedAppointmentMins" json:"missedAppointmentMins"`
	DefaultOperationCodePosition string            `bson:"defaultOperationCodePosition" json:"defaultOperationCodePosition"`
	DMSIntegration               string            `bson:"DMSIntegration" json:"DMSIntegration"`
	Features                     []features        `bson:"features" json:"features"`
}

type DealerInput struct {
	ID                       string   `bson:"_id" json:"dealerId"`
	DealerName               string   `bson:"dealerName" json:"dealerName"`
	MakeCode                 []string `bson:"makeCode" json:"makes"`
	StateIssuedNumber        string   `bson:"stateIssuedNumber" json:"stateIssuedNumber"`
	ManufacturerIssuedNumber string   `bson:"manufacturerIssuedNumber" json:"manufacturerIssuedNumber"`
	TenantID                 string   `bson:"tenantID" json:"tenantID"`
	Website                  string   `bson:"website" json:"website"`
	Phone                    string   `bson:"phone" json:"phone"`
	TimeZone                 string   `bson:"timeZone" json:"timeZone"`
	Currency                 string   `bson:"currency" json:"currency"`
	DMSIntegration           string   `bson:"DMSIntegration" json:"DMSIntegration"`
	TenantCode               string   `bson:"tenantCode" json:"tenantCode"`
	MFOpening                string   `bson:"MFOpening" json:"MFOpening"`
	MFClosing                string   `bson:"MFClosing" json:"MFClosing"`
	SatOpening               string   `bson:"SatOpening" json:"SatOpening"`
	SatClosing               string   `bson:"SatClosing" json:"SatClosing"`
	StreetAddress1           string   `bson:"streetAddress1"json:"streetAddress1"`
	StreetAddress2           string   `bson:"streetAddress2" json:"streetAddress2"`
	City                     string   `bson:"city" json:"city"`
	State                    string   `bson:"state" json:"state"`
	ZipCode                  string   `bson:"zipCode" json:"zipCode"`
	Country                  string   `bson:"country" json:"country"`
	County                   string   `bson:"county" json:"county"`
	LocationURL              string   `bson:"locationUrl" json:"locationUrl"`
	TenantName               string   `bson:"tenantName" json:"tenantName"`
	ApplicationCode          string   `bson:"applicationCode" json:"applicationCode"`
	DealershipCode           string   `bson:"dealershipCode" json:"dealershipCode"`
	Twilio                   string   `bson:"twilio" json:"twilio"`
	BDCService               string   `bson:"bdcService" json:"bdcService"`
}

//StreetAddress1  string `bson:"streetAddress1"json:"streetAddress1"`
//StreetAddress2  string `bson:"streetAddress2" json:"streetAddress2"`
//City            string `bson:"city" json:"city"`
//State           string `bson:"state" json:"state"`
//ZipCode         string `bson:"zipCode" json:"zipCode"`
//Country         string `bson:"country" json:"country"`
//County          string `bson:"county" json:"county"`
//LocationURL     string `bson:"locationUrl" json:"locationUrl"`

type phone struct {
	Service string `bson:"service" json:"service"`
}

type defaultAppointment struct {
	AppointmentReminderTextTemplate    string `bson:"appointmentReminderTextTemplate" json:"appointmentReminderTextTemplate"`
	AppointmentReminderEmailTemplate   string `bson:"appointmentReminderEmailTemplate" json:"appointmentReminderEmailTemplate"`
	AppointmentReminderTimeFormat      string `bson:"appointmentReminderTimeFormat" json:"appointmentReminderTimeFormat"`
	DealerAddressURLSample             string `bson:"dealerAddressURLSample" json:"dealerAddressURLSample"`
	ApplicationURLSample               string `bson:"applicationURLSample" json:"applicationURLSample"`
	DealerNamePlaceholderSampleOption1 string `bson:"dealerNamePlaceholderSampleOption1" json:"dealerNamePlaceholderSampleOption1"`
	DealerNamePlaceholderSampleOption2 string `bson:"dealerNamePlaceholderSampleOption2" json:"dealerNamePlaceholderSampleOption2"`
	DealerNamePlaceholderSampleOption3 string `bson:"dealerNamePlaceholderSampleOption3" json:"dealerNamePlaceholderSampleOption3"`
}

//DealerCommunication
type Communication struct {
	ID                                     string             `bson:"_id" json:"_id"`
	DealerID                               string             `bson:"dealerID" json:"dealerID"`
	PhoneNumber                            phone              `bson:"phoneNumber" json:"phoneNumber"`
	BdcPhoneNumber                         phone              `bson:"bdcPhoneNumber" json:"bdcPhoneNumber"`
	AutoReplyMessage                       string             `bson:"autoReplyMessage" json:"autoReplyMessage"`
	DefaultAppointmentReminderPlaceholders defaultAppointment `bson:"defaultAppointmentReminderPlaceholders" json:"defaultAppointmentReminderPlaceholders"`
	AllowedNotificationSettingTypes        []string           `bson:"allowedNotificationSettingTypes" json:"allowedNotificationSettingTypes"`
}

type Sequence struct {
	ID           string `json:"_id"`
	SequenceName string `json:"sequenceName"`
	Next         int64  `json:"next"`
}
