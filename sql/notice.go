package sql

type Content struct {
	Nid int `json:"nid"`
	Title string `json:"title"`
	Content string `json:"content"`
	Open int `json:"open"`
}

func GetNotice(nid int)Content  {
	var content Content
	DB.Raw("select * from content where nid =? ", nid).Scan(&content)
	return content
}

func UpdateNotice(text string,nid int) {
	DB.Exec("update content set content = ? ,open = 1 where nid =? ", text, nid)
}

func StopAnnounce()  {
	DB.Exec("update content set open =0 where nid <4")
}

func GetContent(nid int)Content  {
	var content Content
	DB.Raw("select * from content where nid =?",nid).Scan(&content)
	return content
}

