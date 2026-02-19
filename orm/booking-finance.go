package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type BookingFinance struct {
	ID        string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	BookingID string `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	model.GormModel
	// Loan Section
	FinanceID         string  `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)"`
	FinanceStatusID   string  `json:"finance_status_id" query:"finance_status_id" gorm:"type:varchar(10)"`
	DownAmount        float64 `json:"down_amount" query:"down_amount" gorm:"type:decimal(14,2)"`
	LoanAmount        float64 `json:"loan_amount" query:"loan_amount" gorm:"type:decimal(14,2)"`
	Interest          float64 `json:"interest" query:"interest" gorm:"type:decimal(14,2)"`
	InstallmentMonth  string  `json:"installment_month" query:"installment_month" gorm:"type:varchar(2)"`
	InstallmentAmount float64 `json:"installment_amount" query:"installment_amount" gorm:"type:decimal(14,2)"`

	CoLoanFirstName string `json:"co_loan_first_name" query:"co_loan_first_name" gorm:"type:varchar(256)"`
	CoLoanLastName  string `json:"co_loan_last_name" query:"co_loan_last_name" gorm:"type:varchar(256)"`
	CoLoanMobile    string `json:"co_loan_mobile" query:"co_loan_mobile" gorm:"type:varchar(20)"`
	CoLoanRelation  string `json:"co_loan_relation" query:"co_loan_relation" gorm:"type:varchar(256)"`

	CoLoanFirstName2 string `json:"co_loan_first_name_2" query:"co_loan_first_name_2" gorm:"column:co_loan_first_name_2;type:varchar(256)"`
	CoLoanLastName2  string `json:"co_loan_last_name_2" query:"co_loan_last_name_2" gorm:"column:co_loan_last_name_2;type:varchar(256)"`
	CoLoanMobile2    string `json:"co_loan_mobile_2" query:"co_loan_mobile_2" gorm:"column:co_loan_mobile_2;type:varchar(20)"`
	CoLoanRelation2  string `json:"co_loan_relation_2" query:"co_loan_relation_2" gorm:"column:co_loan_relation_2;type:varchar(256)"`

	FileFinance  string `json:"file_finance" query:"file_finance" gorm:"type:varchar(500)"`
	FileFinance1 string `json:"file_finance_1" query:"file_finance_1" gorm:"column:file_finance_1;type:varchar(500)"`
	FileFinance2 string `json:"file_finance_2" query:"file_finance_2" gorm:"column:file_finance_2;type:varchar(500)"`
	FileFinance3 string `json:"file_finance_3" query:"file_finance_3" gorm:"column:file_finance_3;type:varchar(500)"`
	FileFinance4 string `json:"file_finance_4" query:"file_finance_4" gorm:"column:file_finance_4;type:varchar(500)"`
	FileFinance5 string `json:"file_finance_5" query:"file_finance_5" gorm:"column:file_finance_5;type:varchar(500)"`

	// เพิ่ม 1. เซ้นใบสมัคร 2.ไฟแน้นมารับใบสมัคร 3.เอกสารครบ 4.อนุมัติ
	IsRegisterFinanceOnline bool       `json:"is_register_finance_online" query:"is_register_finance_online" gorm:"type:bool"`
	RegisterDate            *time.Time `json:"register_date" query:"register_date"`
	RegisterBy              string     `json:"register_by" query:"register_by" gorm:"type:varchar(100)"`
	ReceiveDate             *time.Time `json:"receive_date" query:"receive_date"`
	ReceiveBy               string     `json:"receive_by" query:"receive_by" gorm:"type:varchar(100)"`
	DocDoneDate             *time.Time `json:"doc_done_date" query:"doc_done_date"`
	DocDoneBy               string     `json:"doc_done_by" query:"doc_done_by" gorm:"type:varchar(100)"`
	IsFinanceApprove        bool       `json:"is_finance_approve" query:"is_finance_approve" gorm:"type:bool"`
	ApproveDate             *time.Time `json:"approve_date" query:"approve_date"`
	ApproveBy               string     `json:"approve_by" query:"approve_by" gorm:"type:varchar(100)"`
	ApproveNo               string     `json:"approve_no" query:"approve_no" gorm:"type:varchar(100)"`
	// AgentID      string     `json:"agent_id" query:"agent_id" gorm:"type:varchar(36)"`
	AgentName   string `json:"agent_name" query:"agent_name" gorm:"type:varchar(100)"`
	AgentMobile string `json:"agent_mobile" query:"agent_mobile" gorm:"type:varchar(20)"`

	// For Booking Report
	SentCaseBy       string     `json:"sent_case_by" query:"sent_case_by" gorm:"type:varchar(100)"`
	SentCaseDate     *time.Time `json:"sent_case_date" query:"sent_case_date"`
	SignContractBy   string     `json:"sign_contract_by" query:"sign_contract_by" gorm:"type:varchar(100)"` // เซ็นสัญญาโดย
	SignContractDate *time.Time `json:"sign_contract_date" query:"sign_contract_date"`                      // วันที่เซ็นสัญญา
	DocCompleteBy    string     `json:"doc_complete_by" query:"doc_complete_by" gorm:"type:varchar(100)"`
	DocCompleteDate  *time.Time `json:"doc_complete_date" query:"doc_complete_date"`
	RejectBy         string     `json:"reject_by" query:"reject_by" gorm:"type:varchar(100)"`
	RejectDate       *time.Time `json:"reject_date" query:"reject_date"`
	RejectReason     string     `json:"reject_reason" query:"reject_reason" gorm:"type:text"`
	RejectComment    string     `json:"reject_comment" query:"reject_comment" gorm:"type:text"`
}
