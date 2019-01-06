package httpserver

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var _ = fmt.Println

type urlParsePro struct {
	m_topicMap sync.Map
}

type subscribeInfo struct {
	singleFields []string
	isExistPound bool
	matchUrl     string
}

func (this *urlParsePro) regisnterUrl(topic string) error {
	singleMatch, err := regexp.Compile(":([^/]*)?")
	if err != nil {
		return err
	}
	fields := singleMatch.FindAllStringSubmatch(topic, -1)
	var fieldList []string
	for _, field := range fields {
		if len(field) > 1 {
			fieldList = append(fieldList, field[1])
		}
	}
	str := singleMatch.ReplaceAll([]byte(topic), []byte("([^/]*)?"))
	match, err := regexp.Compile("#")
	if err != nil {
		return err
	}
	matchUrl := match.ReplaceAll(str, []byte("(.*)?"))
	matchUrlStr := string(matchUrl)
	isExistPound := match.Match(str)
	///////
	ts := strings.Split(topic, "/")
	tsLast := ts[len(ts)-1]
	if singleMatch.Match([]byte(tsLast)) && !isExistPound {
		matchUrlStr += "$"
	}
	///////
	_, ok := this.m_topicMap.Load(matchUrlStr)
	if ok {
		return errors.New("exist register")
	}
	info := subscribeInfo{
		singleFields: fieldList,
		isExistPound: isExistPound,
		matchUrl:     matchUrlStr,
	}
	this.m_topicMap.Store(topic, info)
	return nil
}

func (this *urlParsePro) findMatch(topic *string) (isFind bool, findTopic *string, params *map[string]string) {
	isfind := false
	findUrl := ""
	paramMap := make(map[string]string)
	f := func(k, v interface{}) bool {
		info := v.(subscribeInfo)
		match, err := regexp.Compile(string(info.matchUrl))
		if err != nil {
			return true
		}
		matches := match.FindStringSubmatch(*topic)
		matchLen := len(matches) - 1
		if matchLen == 0 && *topic == matches[0] {
			// full equal
			isfind = true
			findUrl = k.(string)
			return false
		}
		if matchLen > 0 {
			// find
			isfind = true
			findUrl = k.(string)
			if !info.isExistPound {
				for i := 1; i <= matchLen; i++ {
					paramMap[info.singleFields[i-1]] = matches[i]
				}
			} else {
				for i := 1; i < matchLen; i++ {
					paramMap[info.singleFields[i-1]] = matches[i]
				}
				paramMap["#"] = matches[matchLen]
			}
			return false
		}
		return true
	}
	this.m_topicMap.Range(f)
	return isfind, &findUrl, &paramMap
}
