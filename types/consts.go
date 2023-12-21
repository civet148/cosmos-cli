package types

const (
	DEFAULT_DENON           = "uhby"
	DEFAULT_CHAIN_ID        = "hobby_9000-1"
	DEFAULT_KEY_PHRASE      = "88888888"
	DEFAULT_NODE_CMD        = "hobbyd"
	DEFAULT_CONFIG_FILE     = "config.yml"
	DEFAULT_KEYRING_BACKEND = KEYRING_BACKEND_FILE
)

const (
	EXEC_CMD_WHICH = "which"
	EXEC_CMD_COPY  = "cp"
	EXEC_CMD_SHELL = "sh"
	EXEC_CMD_MKDIR = "mkdir"
)

const (
	EXEC_SHELL_ARG = "-c"
)

const (
	COMMAND_NAME_EXPECT  = "expect"
	COMMAND_NAME_YUM     = "yum"
	COMMAND_NAME_APT_GET = "apt-get"
)

const (
	COSMOS_P2P_PORT      = "26656"
	COSMOS_RPC_PORT      = "26657"
	COSMOS_GRPC_PORT     = "9090"
	COSMOS_GRPC_WEB_PORT = "9091"
)

const (
	EXPORT_KEY_FILE      = "/tmp/account.key"
	KEYRING_BACKEND_TEST = "test"
	KEYRING_BACKEND_FILE = "file"
)

const (
	FILE_NAME_APP     = "app.toml"
	FILE_NAME_CONFIG  = "config.toml"
	FILE_NAME_GENESIS = "genesis.json"
	CONFIG_SUBPATH    = "config"
)
