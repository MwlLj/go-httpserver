package httpserver

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var _ = fmt.Println

type urlParse struct {
}

func (*urlParse) findMatch(topic *string, subscribes *sync.Map, param *urlParam) (bool, interface{}) {
	var isExist bool = false
	var resultValue interface{} = nil
	inTopics := strings.Split(*topic, "/")
	inTopicLen := len(inTopics)
	f := func(k, v interface{}) bool {
		subscribeTopic := k.(string)
		topics := strings.Split(subscribeTopic, "/")
		topicLen := len(topics)
		if inTopicLen != topicLen {
			return true
		}
		for i, t := range topics {
			if t == "" {
				continue
			}
			exp, err := regexp.Compile(":(.*)?")
			if err != nil {
				continue
			}
			r := exp.FindStringSubmatch(t)
			rLen := len(r)
			inT := inTopics[i]
			if rLen == 2 {
				param.add(&(r[1]), &inT)
			} else {
				if t != inT {
					return true
				}
			}
		}
		isExist = true
		resultValue = v
		return false
	}
	subscribes.Range(f)
	if isExist == false {
		return false, nil
	}
	return true, resultValue
}
