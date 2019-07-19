## gin-basicauth
Gin web frameword's basic authorization middleware.Support Custom User Model Based on ORM Framework Gorm.

## Usage
See the source code for specific usage.
### define user model 
The user model must define the methods of  interface type IBasicAuth.
```go
// IBasicAuth basic auth interface
type IBasicAuth interface {
	IsActived() bool
	CheckPassword(string) bool
	UsernameColumnName() string
}

// UserProfile model
type UserProfile struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Username    string    `gorm:"type:varchar(150);unique_index:uidx_name" json:"username,omitempty"`
	Password    string    `gorm:"type:varchar(128);default:'666'"`
	IsSuperUser bool      `gorm:"column:is_superuser;default:false"  json:"-"`
	IsStuff     bool      `gorm:"column:is_stuff;default:false"  json:"-"`
	IsActive    bool      `gorm:"column:is_active;default:false"  json:"-"`
	FirstName   string    `gorm:"column:first_name;type:varchar(30)" json:"first_name,omitempty"`
	LastName    string    `gorm:"column:last_name;type:varchar(150)" json:"last_name,omitempty"`
	Email       string    `gorm:"type:varchar(254)"`
	DateJoined  time.Time `gorm:"column:date_joined"  json:"-"`
	LastLogin   time.Time `gorm:"column:last_login"  json:"-"`
}

// TableName Set UserProfile's table name
func (UserProfile) TableName() string {
	return "users_userprofile"
}

// IsActived return true if user is active,otherwise false
func (u UserProfile) IsActived() bool {

	return u.IsActive
}

// CheckPassword Check whether the user's password is the same as the given password
func (u UserProfile) CheckPassword(pw string) bool {

	return u.Password == pw
}
```
### Use in Gin
```go
import "gin-basicauth/basicauth"

app := gin.Default()
 //input *gorm.DB and user model
app.Use(basicauth.BasicAuth(db, &UserProfile{}))
```

### In Gin Conatroller
```go
func getUserHandler(ctx *gin.Context) {

	user, exists := ctx.Get(basicauth.AuthUserKey) 
	if !exists {
		ctx.JSON(401, "user not authenticated")
		return
	}
	u, ok := user.(*UserProfile)//important
	if ok {
		ctx.JSON(200, u)
		return
	}
	ctx.JSON(401, "user not authenticated")
}
```


