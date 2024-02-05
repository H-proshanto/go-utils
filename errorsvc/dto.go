package errorsvc

type ErrorDetail struct {
	InternalCode string `json:"internalCode" bson:"internal_code"`
	MessageEn    string `json:"messageEn" bson:"message_en"`
	MessageBn    string `json:"messageBn" bson:"message_bn"`
}

type ErrorResponse struct {
	Timestamp   int64        `json:"timestamp"`
	Description string       `json:"description"`
	Error       *ErrorDetail `json:"error"`
}
