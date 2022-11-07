package go_reepay

type Reepay struct {
	Key        string
	SuccessURL string
	CancelURL  string
}

func Init(key, success, cancel string) (reepay Reepay) {
	return Reepay{
		Key:        key,
		SuccessURL: success,
		CancelURL:  cancel,
	}
}
