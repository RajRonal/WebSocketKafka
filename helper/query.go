package helper

import (
	"Kafka-web-Socket/database"
	"github.com/sirupsen/logrus"
)

func InsertMessage(senderId, receiverId, message string) error {
	SQL := `INSERT INTO chat_app (sender_id,receiver_id,message) 	
			VALUES ($1, $2,$3)`
	_, err := database.DB.Exec(SQL, senderId, receiverId, message)
	if err != nil {
		logrus.Error("InsertMessage:Error in Creating Message history %v", err)
		return err
	}

	return nil
}
