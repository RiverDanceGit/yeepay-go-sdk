package response

type NccashierapiApiPayResponse struct {
	Result struct {
		Code       string `json:"code"`
		Message    string `json:"message"`
		PayTool    string `json:"payTool"`
		PayType    string `json:"payType"`
		MerchantNo string `json:"merchantNo"`
		Token      string `json:"token"`
		ResultType string `json:"resultType"`
		ResultData string `json:"resultData"`
	} `json:"result"`
}

type NccashierapiApiPayResponseResultData struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (resp *NccashierapiApiPayResponse) IsSuccess() bool {
	if "CAS00000" == resp.Result.Code {
		return true
	}
	return false
}
