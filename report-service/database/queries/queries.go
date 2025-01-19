package queries

import _ "embed"

var (
	//go:embed sql/get_report_item_by_factory_and_date.sql
	GetFactoryReportItems string

	//go:embed sql/get_report_item_by_user_and_date.sql
	GetUserReportItems string

	//go:embed sql/add_report_Item.sql
	AddReportItems string
)
