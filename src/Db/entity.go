package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model

	// User Basic data
	UserName    string `gorm:"type:text;"`
	PhoneNumber string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Gender      string `gorm:"type:text;"`
	Age         int

	// User Account data
	Exp                   int           `gorm:"not null;default:'0'"`
	Level                 int           `gorm:"not null;default:'1'"`
	HeadPortraitUrl       string        ``
	FollowingUsers        []*User       `gorm:"many2many:user_following_ships;association_jointable_foreignkey:following_user_id"`
	FollowerUsers         []*User       `gorm:"many2many:user_follower_ships;association_jointable_foreignkey:follower_user_id"`
	CollectedArticles     []Article     `gorm:"many2many:user_collected_article"`
	CollectedTopics       []Topic       `gorm:"many2many:user_collected_topic"`
	JoinedGroups          []*FriendGroup `gorm:"many2many:user_joined_group"`
	BloodRecords          []BloodRecord
	HealthRecords         []HealthRecord
	FamilyMembers         []FamilyMember
	UserPrivacySetting    UserPrivacySetting
	MessageStateInGroup   MessageStateInGroup
	MessageStateU2u       MessageStateU2u
	UserPrivacySettingID  uint
	MessageStateInGroupID uint
	MessageStateU2uID     uint

	// User healthy data
	Height float64
	Weight float64
	Area   string
	Job    string
}

type UserCheckIn struct {
	gorm.Model
	CheckTime time.Time `gorm:"type:date"`
	UserID    uint
}

type UserPrivacySetting struct {
	gorm.Model
	ShowPhoneNumber bool
	ShowGender      bool
	ShowAge         bool
	ShowHeight      bool
	ShowWeight      bool
	ShowArea        bool
	ShowJob         bool
}

type Article struct {
	gorm.Model
	Title           string `gorm:"not null"`
	Content         string `gorm:"not null;type:text;"`
	PlainContent    string `gorm:"not null;type:text;"`
	CoverImageUrl   string
	ReadCount       int `gorm:"not null;default:'0'"`
	ArticleComments []ArticleComment
}

type ArticleComment struct {
	gorm.Model
	Content       string `gorm:"not null;type:text"`
	ThumbsUpCount int    `gorm:"not null;default:'0'"`
	User          User
	UserID        int
	Article       Article
	ArticleID     int
}

type Topic struct {
	gorm.Model
	Content       string `gorm:"not null;type:text"`
	User          User
	UserID        int
	ThumbsUpCount int `gorm:"not null;default:'0'"`
	LordReplies   []TopicLordReply
}

type TopicLordReply struct {
	gorm.Model
	Content       string `gorm:"not null;type:text"`
	User          User
	UserID        int
	ThumbsUpCount int `gorm:"not null;default:'0'"`
	TopicID       uint
	LayerReplies  []TopicLayerReply
}

type TopicLayerReply struct {
	gorm.Model
	Content          string `gorm:"not null;type:text"`
	User             User
	UserID           int
	ThumbsUpCount    int `gorm:"not null;default:'0'"`
	TopicLordReplyID uint
}

type BloodRecord struct {
	gorm.Model
	Level      string
	RecordTime string
	RecordDate time.Time `gorm:"type:date;"`
	User       User
	UserID     uint
}

type HealthRecord struct {
	gorm.Model
	Insulin       string
	SportTime     string
	Weight        string
	BloodPressure string
	RecordTime    string
	RecordDate    time.Time `gorm:"type:date;"`
	User          User
	UserID        uint
}

type FamilyMember struct {
	gorm.Model
	PhoneNumber string
	CallName    string
	User        User
	UserID      uint
}

type MessageStateU2u struct {
	gorm.Model
	LastReadTime time.Time
}

type MessageU2u struct {
	gorm.Model
	SenderID uint
	Sender   User
	Content  string `gorm:"type:text"`
	TargetID uint
	Target   User
}

type FriendGroup struct {
	gorm.Model
	Name    string `gorm:"unique;not null;"`
	Members []*User `gorm:"many2many:user_joined_group"`
	User    User
	UserID  uint
}

type MessageStateInGroup struct {
	gorm.Model
	LastReadTime time.Time
}

type MessageInGroup struct {
	gorm.Model
	SenderID uint
	Sender   User
	GroupID  uint
	Content  string `gorm:"type:text"`
}

type UserReply struct {
	TopicLordReplyKey  uint      `gorm:"Column:topic_lord_reply_key"`
	TopicLayerReplyKey uint      `gorm:"Column:topic_layer_reply_key"`
	CreatedAt          time.Time `gorm:"Column:created_at"`
	Content            string    `gorm:"Column:content"`
	UserId             int       `gorm:"Column:user_id"`
	ThumbsUpCount      int       `gorm:"Column:thumbs_up_count"`
	FatherID           uint      `gorm:"Column:father_id"`
}
