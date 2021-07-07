package main

// count=24 // количество результатов на страницу max=24
// categories= animals-pets // группа по которой производится поиск
// in_store=1 // поиск только в магазине
// processing_status=succeeded // уточнить
// sort_by=-relevance // тип сортировки
// type=models // тип работы
// cursor=bz0yNA // метка запроса, указывает на предыдущую/следующую страницу для поиска

/*
https://sketchfab.com/i/search?

//Sort by
sort_by ( -relevance -likeCount -viewCount -publishedAt )


*/

var finder = map[string]string{

	"cursor":            "",
	"count":             "",
	"q":                 "",
	"categories":        "",
	"in_store":          "",
	"processing_status": "",

	"type": "",

	// filters
	//Features
	"downloadable": "", // null 1
	"animated":     "", // null 1
	"staffpicked":  "", // null 1
	"sound":        "", // null 1

	//Sort by
	"sort_by": "", // -relevance -likeCount -viewCount -publishedAt

	//Date
	"date": "", // none(all) 1(day) 7(week) 31(month)

	"restricted": "", // 1(show restricted
}
