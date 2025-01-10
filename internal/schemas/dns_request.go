package schema

type DNSRequest struct {
	Content   string   `json:"content"`
	Name      string   `json:"name"`
	Proxiable bool     `json:"proxiable"`
	Proxied   bool     `json:"proxied"`
	TTL       int      `json:"ttl"`
	Type      string   `json:"type"`
	ZoneID    string   `json:"zone_id"`
	ZoneName  string   `json:"zone_name"`
	Settings  struct{} `json:"settings"`
	Tags      []string `json:"tags"`
	ID        string   `json:"id"`
}
