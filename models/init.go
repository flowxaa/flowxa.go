package models

import (
	"fmt"
	"os"
)

var (
	categoryJSONFile    = "posts/category.json"
	topicMarkdownFolder = "posts"

	//Topics store all the topic
	Topics []*Topic
	//TopicsGroupByMonth store the topic by month
	TopicsGroupByMonth []*MonthList
	//TopicsGroupByCategory store all the category
	TopicsGroupByCategory []*TopicCategory
)

func init() {
	if err := InitTopicCategoryList(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := InitTopicList(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
