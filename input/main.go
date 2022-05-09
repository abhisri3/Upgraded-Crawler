package main

import (
	dbconfig "crawl/dbconfig"
	model "crawl/models"
)

func main() {
	startSerach := model.Search{}

	startSerach.Url = "https://www.aegonlife.com/"

	startSerach.Container = "ul.c-menu-online-plans__submenu"

	startSerach.Patterns = []string{"insurance-plans", "saral"}

	titles := []string{"h1.node_title > span", "div.text-block > h2"}
	startSerach.Title = titles

	descs := []string{"div.star-insurance-plan > div.desc-block > p", "div.desc-block > p", "div.text-block > p >mark"}
	startSerach.Desc = descs

	benefits := []string{"div.carousel-content > div > p", "div.carousel-content > p", "li > div >p"} //"li > div >p"}
	startSerach.Benefits = benefits

	eligibility := []string{"di.MsoTableGrid > tr > td", "div.carousel-content > div > p", "div.carousel-content > p"}
	startSerach.Eligibility = eligibility

	dbconfig.Insert(startSerach)

	// startSerach2 := model.Search{}

	// startSerach2.Url = "https://www.bhartiaxa.com/"

	// titles2 := []string{"h1.banner-main-title"}
	// startSerach2.Title = titles2

	// descs2 := []string{"div > p.titel_desc"}
	// startSerach2.Desc = descs2

	// benefits2 := []string{"div.rowImageWrapper > p.rowImageTitle"}
	// startSerach2.Benefits = benefits2

	// eligibility2 := []string{}
	// startSerach2.Eligibility = eligibility2

	// dbconfig.Insert(startSerach2)

	// startSerach3 := model.Search{}

	// startSerach3.Url = "https://www.avivaindia.com/"

	// startSerach3.Container = "div.newhome"

	// startSerach3.Patterns = []string{"aviva-"}

	// titles3 := []string{"div.tc_termInsu > div.headinSection > h1.headText"}
	// startSerach3.Title = titles3

	// descs3 := []string{"div.tc_termInsu > div.headinSection > p"}
	// startSerach3.Desc = descs3

	// benefits3 := []string{"div.benPara > p, p.reasonPara"} //"li > div >p"}
	// startSerach3.Benefits = benefits3

	// eligibility3 := []string{}
	// startSerach3.Eligibility = eligibility3

	// dbconfig.Insert(startSerach3)

}
