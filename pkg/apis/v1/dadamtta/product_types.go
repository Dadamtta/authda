package dadamtta

type ProductRegisterFormRequest struct {
	CategoryCode string `json:"category_code"`
	Label        string `json:"label"`
	Price        uint32 `json:"price"`
	Description  string `json:"description"`
	Content      string `json:"content"`
}
