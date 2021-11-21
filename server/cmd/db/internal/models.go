package internal

// Customer - table
type Customer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name" gorm:"not null;size:50"`
	LastName  string `json:"last_name" gorm:"not null;size:50"`
	Email     string `json:"email" gorm:"unique;not null;size:50"`
	Password  string `gorm:"not null;unique"`
	CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime:nano"`
}

// Tariff - table
type Tariff struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" gorm:"not null;unique;size:50"`
	Cost      uint   `json:"cost" gorm:"not null;size:6"`
	Duration  uint   `json:"duration" gorm:"not null;size:20"`
	CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime:nano"`
}

// BillAccount - table
type BillAccount struct {
	ID          uint     `json:"id"`
	Balance     float32  `json:"balance" gorm:"not null;precision:3;size:9"`
	NextPayDate uint     `json:"next_pay_date" gorm:"not null;size:20"`
	CreatedAt   uint     `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt   uint     `json:"updated_at" gorm:"autoUpdateTime:nano"`
	TariffID    uint     `json:"tariff_id"`
	Tariff      Tariff   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CustID      uint     `json:"cust_id"`
	Customer    Customer `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Bill - table
type Bill struct {
	ID            uint        `json:"id"`
	IsPayed       bool        `json:"is_payed" gorm:"not null"`
	CreatedAt     uint        `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt     uint        `json:"updated_at" gorm:"autoUpdateTime:nano"`
	ExpiredAt     uint        `json:"expired_at" gorm:"not null;size:20"`
	TariffID      uint        `json:"tariff_id"`
	Tariff        Tariff      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BillAccountID uint        `json:"bill_account_id"`
	BillAccount   BillAccount `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReportID      uint
}

// MothlyReport - table
type MothlyReport struct {
	ID        uint   `json:"id"`
	CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:nano"`
	Total     uint   `json:"total" gorm:"not null;size:20"`
	Bills     []Bill `gorm:"foreignKey:ReportID"`
}

// Session - table
type Session struct {
	ID        string
	ExpiredAt uint
	CustID    uint
	Customer  Customer `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// ProjectStep - table
type ProjectStep struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" gorm:"not null;size:50"`
	Description string `json:"description" gorm:"not null;size:1000"`
	CreatedAt   uint   `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt   uint   `json:"updated_at" gorm:"autoUpdateTime:nano"`
}

// Project - table
type Project struct {
	ID                     uint                 `json:"id"`
	Name                   string               `json:"name" gorm:"not null;size:50"`
	CreatedAt              uint                 `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt              uint                 `json:"updated_at" gorm:"autoUpdateTime:nano"`
	SelectedCollectionIds  []SelectedCollection `json:"selected_collection_ids"`
	SuggestedCollectionIds []SelectedCollection `json:"suggested_collection_ids" gorm:"not null"`
	CustID                 uint                 `json:"cust_id"`
	Customer               Customer             `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StepID                 uint                 `json:"step_id"`
	ProjectStep            ProjectStep          `gorm:"not null;foreignKey:StepID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TemplateCollection - table, customer's personal, 'template' collection
type TemplateCollection struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name" gorm:"not null;size:50"`
	CreatedAt uint     `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt uint     `json:"updated_at" gorm:"autoUpdateTime:nano"`
	CustID    uint     `json:"cust_id"`
	Customer  Customer `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// ProjectCollection - table, order's, project's collection
type ProjectCollection struct {
	ID                   uint               `json:"id"`
	CreatedAt            uint               `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt            uint               `json:"updated_at" gorm:"autoUpdateTime:nano"`
	CustID               uint               `json:"cust_id"`
	Customer             Customer           `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID            uint               `json:"project_id"`
	Project              Project            `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TemplateCollectionID uint               `json:"template_collection_id"`
	TemplateCollection   TemplateCollection `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// SelectedCollection - table, show which collection selected/choosed
type SelectedCollection struct {
	ID                   uint               `json:"id"`
	IsClientChoosed      bool               `json:"is_client_choosed"`
	CreatedAt            uint               `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt            uint               `json:"updated_at" gorm:"autoUpdateTime:nano"`
	CustID               uint               `json:"cust_id"`
	Customer             Customer           `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID            uint               `json:"project_id"`
	TemplateCollectionID uint               `json:"template_collection_id"`
	TemplateCollection   TemplateCollection `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Photo - table
type Photo struct {
	ID                   uint               `json:"id"`
	Source               string             `json:"src" gorm:"not null;size:200"`
	Name                 string             `json:"filename" gorm:"not null;size:50"`
	IsPoster             bool               `json:"is_poster"`
	IsSelected           bool               `json:"is_selected"`
	CreatedAt            uint               `json:"created_at" gorm:"autoCreateTime:nano"`
	CustID               uint               `json:"cust_id"`
	Customer             Customer           `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TemplateCollectionID uint               `json:"template_collection_id"`
	TemplateCollection   TemplateCollection `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectCollectionID  uint               `json:"project_collection_id"`
	ProjectCollection    ProjectCollection  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Link - table
type Link struct {
	ID        uint     `json:"id"`
	Source    string   `json:"src" gorm:"not null;size:200"`
	CreatedAt uint     `json:"created_at" gorm:"autoCreateTime:nano"`
	CustID    uint     `json:"cust_id"`
	Customer  Customer `gorm:"not null;foreignKey:CustID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID uint     `json:"project_id"`
	Project   Project  `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
