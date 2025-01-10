package schema

type DNSResponse struct {
	Success bool       `json:"success"`
	Errors  []DNSError `json:"errors"`
}
