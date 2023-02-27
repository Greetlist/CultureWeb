package schema

type Interest struct {
    ID uint `gorm:"primary_key;AUTO_INCREMENT;"`
    UserID uint `gorm:"column:user_id;"`
    ArticleID uint `gorm:"article_id;"`
}
