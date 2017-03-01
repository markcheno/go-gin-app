package models

var tmpUserList []User
var tmpArticleList []Article

// This function is used to store the main lists into the temporary one for testing
func saveLists() {
	tmpUserList = UserList
	tmpArticleList = ArticleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	UserList = tmpUserList
	ArticleList = tmpArticleList
}
