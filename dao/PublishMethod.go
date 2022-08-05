package dao

import "log"

func (video *Video) IdFindVideo(id int64) {
	err := DB.Where("id=?", id).Find(&video).Error
	if err != nil { //record not found
		log.Println("method IdFindVideo: ", err)
	}
}
