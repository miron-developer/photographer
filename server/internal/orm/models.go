package orm

// Customer - table
type Customer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string
}

// Tariff - table
type Tariff struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Cost     uint   `json:"cost"`
	Duration uint   `json:"duration"`
}

// BillAccount - table
type BillAccount struct {
	ID          uint    `json:"id"`
	Balance     float32 `json:"balance"`
	NextPayDate uint    `json:"nextPayDate"`
	TariffID    uint    `json:"tariffID"`
	CustID      uint    `json:"custID"`
}

// Bill - table
type Bill struct {
	ID       uint  `json:"id"`
	IsPayed  uint8 `json:"isPayed"`
	CreateAt uint  `json:"createAt"`
	ExpireAt uint  `json:"expireAt"`
	TariffID uint  `json:"tariffID"`
	CustID   uint  `json:"custID"`
}

// MothlyReport - table
type MothlyReport struct {
	ID       uint `json:"id"`
	CreateAt uint `json:"createAt"`
	Total    uint `json:"total"`
}

// Session - table
type Session struct {
	ID     string
	Expire string
	CustID uint
}

// ProjectStep - table
type ProjectStep struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PhotoProject - table
type PhotoProject struct {
	ID                     uint   `json:"id"`
	CreateAt               uint   `json:"createAt"`
	Name                   string `json:"name"`
	SelectedCollectionIds  []uint `json:"selected"`
	SuggestedCollectionIds []uint `json:"suggested"`
	CustID                 uint   `json:"custID"`
	StepID                 uint   `json:"stepID"`
}

// PhotoCollection - table
type PhotoCollection struct {
	ID        uint   `json:"id"`
	CreateAt  uint   `json:"createAt"`
	EditAt    uint   `json:"editAt"`
	Name      string `json:"name"`
	CustID    uint   `json:"custID"`
	ProjectID uint   `json:"projectID"` // if belong to the certain project
}

// Photo - table
type Photo struct {
	ID           uint   `json:"id"`
	Source       string `json:"src"`
	Name         string `json:"filename"`
	IsPoster     uint8  `json:"isPoster"`
	IsSelected   uint8  `json:"isSelected"`
	CustID       uint   `json:"custID"`
	CollectionID uint   `json:"collectionID"`
}

// Link - table
type Link struct {
	ID        uint   `json:"id"`
	Source    string `json:"src"`
	CustID    uint   `json:"custID"`
	ProjectID uint   `json:"projectID"`
}
