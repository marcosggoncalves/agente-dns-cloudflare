package message

const (
	FALHA_ATUALIZACAO_DNS  = "falhar ao realizar tentativa de atualização no DNS: %v"
	IP_ATUALIZADO          = "ip do DNS alterado com sucesso: %s\n"
	FALHA_CAPTURA_IP       = "não foi possivel capturar ip publico: %v\n"
	FALHA_JSON             = "erro ao converter para JSON: %v"
	FALHA_ENV              = "erro ao carregar o arquivo .env"
	FALHAR_OBTER_RECORD_ID = "erro ao obter record ID para ataulização do DNS"
)
