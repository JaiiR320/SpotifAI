package model

import "github.com/JaiiR320/SpotifAI/utils"

func AddTag(tag string) {
	Tags = append(Tags, tag)
}

func DeleteTag(tag string) error {
	return utils.DeleteFromSlice(&Tags, tag)
}
