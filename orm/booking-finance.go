package orm

import "time"

type BookingFinance struct {
	// Loan Section
	PaymentMethodID   string  `json:"payment_method_id" query:"payment_method_id" gorm:"type:varchar(100)"`
	FinanceID         string  `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)"`
	DownAmount        float64 `json:"down_amount" query:"down_amount" gorm:"type:decimal(14,2)"`
	LoanAmount        float64 `json:"loan_amount" query:"loan_amount" gorm:"type:decimal(14,2)"`
	Interest          float64 `json:"interest" query:"interest" gorm:"type:decimal(14,2)"`
	FinanceMonth      string  `json:"finance_month" query:"finance_month" gorm:"type:varchar(2)"`
	InstallmentAmount float64 `json:"installment_amount" query:"installment_amount" gorm:"type:decimal(14,2)"`

	// เพิ่ม 1. เซ้นใบสมัคร 2.ไฟแน้นมารับใบสมัคร 3.เอกสารครบ 4.อนุมัติ
	IsRegisterFinanceOnline bool       `json:"is_register_finance_online" query:"is_register_finance_online" gorm:"type:bool"`
	FinanceRegisterDate     *time.Time `json:"finance_register_date" query:"finance_register_date"`
	FinanceRegisterBy       string     `json:"finance_register_by" query:"finance_register_by" gorm:"type:varchar(100)"`
	FinanceReceiveDate      *time.Time `json:"finance_receive_date" query:"finance_receive_date"`
	FinanceReceiveBy        string     `json:"finance_receive_by" query:"finance_receive_by" gorm:"type:varchar(100)"`
	FinanceDocDoneDate      *time.Time `json:"finance_doc_done_date" query:"finance_doc_done_date"`
	FinanceDocDoneBy        string     `json:"finance_doc_done_by" query:"finance_doc_done_by" gorm:"type:varchar(100)"`
	FinanceApproveDate      *time.Time `json:"finance_approve_date" query:"finance_approve_date"`
	FinanceApproveBy        string     `json:"finance_approve_by" query:"finance_approve_by" gorm:"type:varchar(100)"`
	FinanceApproveNo        string     `json:"finance_approve_no" query:"finance_approve_no" gorm:"type:varchar(100)"`
	// FinanceAgentID      string     `json:"finance_agent_id" query:"finance_agent_id" gorm:"type:varchar(36)"`
	FinanceAgentName   string `json:"finance_agent_name" query:"finance_agent_name" gorm:"type:varchar(100)"`
	FinanceAgentMobile string `json:"finance_agent_mobile" query:"finance_agent_mobile" gorm:"type:varchar(20)"`

	// For Booking Report
	SentCaseBy           string     `json:"sent_case_by" query:"sent_case_by" gorm:"type:varchar(100)"`
	SentCaseDate         *time.Time `json:"sent_case_date" query:"sent_case_date"`
	SignContractBy       string     `json:"sign_contract_by" query:"sign_contract_by" gorm:"type:varchar(100)"`
	SignContractDate     *time.Time `json:"sign_contract_date" query:"sign_contract_date"`
	DocCompleteBy        string     `json:"doc_complete_by" query:"doc_complete_by" gorm:"type:varchar(100)"`
	DocCompleteDate      *time.Time `json:"doc_complete_date" query:"doc_complete_date"`
	FinanceRejectBy      string     `json:"finance_reject_by" query:"finance_reject_by" gorm:"type:varchar(100)"`
	FinanceRejectDate    *time.Time `json:"finance_reject_date" query:"finance_reject_date"`
	FinanceRejectReason  string     `json:"finance_reject_reason" query:"finance_reject_reason" gorm:"type:text"`
	FinanceRejectComment string     `json:"finance_reject_comment" query:"finance_reject_comment" gorm:"type:text"`
}