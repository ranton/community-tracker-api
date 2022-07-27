package models

type AddPeopleDetails struct {
	Gop                                 bool `gorm:"column:has_gop" json:"gop"`
	ExpSettMtg                          bool `gorm:"column:has_expectation_setting_mtg" json:"expectation_setting_mtg"`
	SignedExpSettDoc                    bool `gorm:"column:has_signed_expectation_setting_document" json:"signed_expectation_setting_document"`
	MonthlyTPIntroEmail                 bool `gorm:"column:has_monthly_touchpoint_intro_email" json:"monthly_touchpoint_intro_email"`
	PerfEvalRecurringMtg                bool `gorm:"column:has_perf_eval_recurring_mtg" json:"perf_eval_recurring_mtg"`
	CommunityComms                      bool `gorm:"column:included_in_community_communications" json:"included_in_community_communications"`
	GopInstEmail                        bool `gorm:"column:has_gop_instruction_email" json:"gop_instruction_email"`
	SecondMonthTPMtg                    bool `gorm:"column:has_second_month_touchpoint_mtg" json:"second_month_touchpoint_mtg"`
	ThirdMonthPerfEvalFromProjLead      bool `gorm:"column:has_third_month_perf_eval_from_proj_lead" json:"third_month_perf_eval_from_proj_lead"`
	ThirdMonthTPPerfEvalFromProjLead    bool `gorm:"column:has_third_month_touchpoint_perf_eval_from_proj_lead" json:"third_month_touchpoint_perf_eval_from_proj_lead"`
	ThirdMonthTPMtg                     bool `gorm:"column:has_third_touchpoint_mtg" json:"third_touchpoint_mtg"`
	FourthMonthTPMtg                    bool `gorm:"column:has_fourth_month_touchpoint_mtg" json:"fourth_month_touchpoint_mtg"`
	FifthTPMonthPerfEvalReqFromProjLead bool `gorm:"column:has_fifth_touchpoint_perf_eval_req_from_proj_lead" json:"fifth_touchpoint_perf_eval_req_from_proj_lead"`
	FifthMonthPerfEvalFromProjLead      bool `gorm:"column:has_fifth_month_perf_eval_from_proj_lead" json:"fifth_month_perf_eval_from_proj_lead"`
	FifthMonthTPMtg                     bool `gorm:"column:has_fifth_month_touchpoint_mtg" json:"fifth_month_touchpoint_mtg"`
}

func (AddPeopleDetails) TableName() string {
	return "peopledetails"
}


