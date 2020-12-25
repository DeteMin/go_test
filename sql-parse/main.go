package main

import (
	"fmt"

	//"github.com/xwb1989/sqlparser"
)

func main(){
	sql := "SELECT  a.production_id, a.production_link, a.production_title, a.spotlight," +
		"a.in_stock_time,a.category,a.level_two_category, a.status, c.campaign_name, c.id as campaign_id," +
		"c.status as campaign_status, a.brand, brand_logo,a.images," +
		"ct.category_name,lct.category_name as level_two_category_name," +
		"c.secret, c.is_deleted as campaign_deleted_status, c.start_time, c.end_time," +
		"a.owner_id, u.company_name, b.brand_name, t.rebate_id, sr.status as rebate_is_delete," +
		"sr.type as rebate_type,sr.min_point, sr.max_point, sr.source_point, sr.rebate_name," +
		"t.production_join_time,a.create_time, a.update_time" +
		" FROM campaign_info as a" +
		" LEFT JOIN brand_info as b ON a.brand = b.brand_id" +
		" LEFT JOIN sponsor_user_profile as u ON a.owner_id = u.id" +
		" LEFT JOIN production_spec as p ON a.production_id = p.production_id" +
		" LEFT JOIN production_campaign_sub_query as c ON a.production_id = c.production_id" +
		" LEFT JOIN production_rebate as t ON a.production_id = t.production_id" +
		" LEFT JOIN star_rebate as sr ON t.rebate_id = sr.id" +
		" LEFT JOIN category as ct ON a.category = ct.id" +
		" LEFT JOIN level_two_category as lct ON a.level_two_category = lct.id" +
		" GROUP BY a.production_id"

	//stmt, err := sqlparser.Parse(sql)
	//if err != nil {
	//	fmt.Println("=== err ===")
	//	return
	//}

	//1.解析select
	//expr := stmt.(*sqlparser.Select).SelectExprs
	//for _, v := range expr {
	//	buf := sqlparser.NewTrackedBuffer(nil)
	//	v.Format(buf)
	//	fmt.Println(fmt.Sprintf("===查询的字段:%v===", buf.String()))
	//}


	tables, columns, err := ParseQuery(sql)
	if err != nil {
		panic(err)
	}

	fmt.Println("Columns:", columns)
	fmt.Println("Tables:", tables)
}
