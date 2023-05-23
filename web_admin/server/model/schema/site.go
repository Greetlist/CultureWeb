package schema

type Site struct {
    SiteID int `gorm:"column:site_id;primary_key;not null;autoIncrement" json:"site_id"`
    SiteName string `gorm:"column:site_name;unique;not null" json:"site_name"`
    Location string `gorm:"column:location;not null" json:"location"`
    PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
    ContactName string `gorm:"column:contact_name" json:"contact_name"`
    Reservations []*Reservation `gorm:"many2many:reservation_sites;" json:"reservations"`
}

func (l Site) TableName() string {
    return "site"
}
