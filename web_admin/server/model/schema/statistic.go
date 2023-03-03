package schema

var StatisticTypeMap = map[uint]string {
    1: "request_success",
    2: "request_error",
    3: "request_number",
    4: "visit_number",
    5: "male_number",
    6: "female_number",
}

type Statistic struct {
    StatisticID uint `gorm:"column:statistic_id;primary_key;" json:"statistic_id"`
    StatisticName string `gorm:"column:statistic_name" json:"statistic_name"`
    StatisticCount uint `gorm:"column:statistic_count" json:"statistic_count"`
}

func (s Statistic) TableName() string {
    return "statistic"
}
