package maps

type Response struct {
	Response GeoObjectCollection `json:"response"`
}

type GeoObjectCollection struct {
	MetaDataProperty GeoObjectCollectionMetaDataProperty `json:"metaDataProperty"`
	FeatureMember    []FeatureMember                     `json:"featureMember"`
}

type GeoObjectCollectionMetaDataProperty struct {
	GeocoderResponseMetaData GeocoderResponseMetaData `json:"GeocoderResponseMetaData"`
}

type GeocoderResponseMetaData struct {
	Request string `json:"request"`
	Results string `json:"results"`
	Found   string `json:"found"`
}

type FeatureMember struct {
	GeoObject GeoObject `json:"GeoObject"`
}

type GeoObject struct {
	MetaDataProperty GeoObjectMetaDataProperty `json:"metaDataProperty"`
	Name             string                    `json:"name"`
	Description      string                    `json:"description"`
	Point            Point                     `json:"Point"`
}

type GeoObjectMetaDataProperty struct {
	GeocoderMetaData GeocoderMetaData `json:"GeocoderMetaData"`
}

type GeocoderMetaData struct {
	Precision      string         `json:"precision"`
	Text           string         `json:"text"`
	Kind           string         `json:"kind"`
	Address        Address        `json:"Address"`
	AddressDetails AddressDetails `json:"AddressDetails"`
}

type Address struct {
	CountryCode string      `json:"country_code"`
	Formatted   string      `json:"formatted"`
	Components  []Component `json:"Components"`
}

type Component struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type AddressDetails struct {
	Country Country `json:"Country"`
}

type Country struct {
	AddressLine        string             `json:"AddressLine"`
	CountryNameCode    string             `json:"CountryNameCode"`
	CountryName        string             `json:"CountryName"`
	AdministrativeArea AdministrativeArea `json:"AdministrativeArea"`
}

type AdministrativeArea struct {
	AdministrativeAreaName string                `json:"AdministrativeAreaName"`
	SubAdministrativeArea  SubAdministrativeArea `json:"SubAdministrativeArea"`
}

type SubAdministrativeArea struct {
	SubAdministrativeAreaName string   `json:"SubAdministrativeAreaName"`
	Locality                  Locality `json:"Locality"`
}

type Locality struct {
	DependentLocality DependentLocality `json:"DependentLocality"`
}

type DependentLocality struct {
	DependentLocalityName string             `json:"DependentLocalityName"`
	DependentLocality     *DependentLocality `json:"DependentLocality,omitempty"`
	Thoroughfare          Thoroughfare       `json:"Thoroughfare,omitempty"`
}

type Thoroughfare struct {
	ThoroughfareName string  `json:"ThoroughfareName"`
	Premise          Premise `json:"Premise"`
}

type Premise struct {
	PremiseNumber string `json:"PremiseNumber"`
}

type BoundedBy struct {
	Envelope Envelope `json:"Envelope"`
}

type Envelope struct {
	LowerCorner string `json:"lowerCorner"`
	UpperCorner string `json:"upperCorner"`
}

type Point struct {
	Pos string `json:"pos"`
}
