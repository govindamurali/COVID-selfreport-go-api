package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Report struct {
	ID                      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Loc                     string             `json:"loc" bson:"loc"`
	City                    string             `json:"city" bson:"city"`
	District                string             `json:"district" bson:"district"`
	State                   string             `json:"state" bson:"state"`
	Country                 string             `json:"country" bson:"country"`
	PinCode                 int                `json:"pinCode" bson:"pinCode"`
	IsSelf                  bool               `json:"isSelf" bson:"isSelf"`
	FullName                string             `json:"fullName" bson:"fullName"`
	PhoneNumber             string             `json:"phoneNumber" bson:"phoneNumber"`
	AlternatePhoneNumber    string             `json:"alternatePhoneNumber" bson:"alternatePhoneNumber"`
	PreferredLanguage       string             `json:"preferredLanguage" bson:"preferredLanguage"`
	Symptoms                map[string]string  `json:"symptoms" bson:"symptoms"`
	TravelToHighRiskCountry bool               `json:"travelToHighRiskCountry" bson:"travelToHighRiskCountry"`
	ContactWithPatient      bool               `json:"contactWithPatient" bson:"contactWithPatient"`
	ContactPatientInfo      bool               `json:"contactPatientInfo" bson:"contactPatientInfo"`
	HasChronicIllness       bool               `json:"hasChronicIllness" bson:"hasChronicIllness"`
	ChronicIllness          string             `json:"chronicIllness" bson:"chronicIllness"`
	Notes                   string             `json:"notes" bson:"notes"`
	IsSuspected             bool               `json:"isSuspected" bson:"isSuspected"`
	IsVerified              bool               `json:"isVerified" bson:"isVerified"`
	CreatedAt               time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt               time.Time          `json:"updatedAt"bson:"updatedAt"`
}
