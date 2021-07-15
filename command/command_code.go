package command

type CMD_CODE uint8
type CMD_SUBCODE uint8

const (
	CMDC_GEN CMD_CODE = iota
	CMDC_OVD
	CMDC_AUDIO
	CMDC_FGR
	CMDC_RMT
	CMDC_FOTA
	CMDC_NET
	CMDC_CON
	CMDC_HBAR
	CMDC_MCU
)

type CMDS_GEN CMD_SUBCODE

const (
	CMD_GEN_INFO CMDS_GEN = iota
	CMD_GEN_LED
	CMD_GEN_RTC
	CMD_GEN_ODO
	CMD_GEN_ANTITHIEF
	CMD_GEN_RPT_FLUSH
	CMD_GEN_RPT_BLOCK
)

type CMDS_OVD CMD_SUBCODE

const (
	CMD_OVD_STATE CMDS_OVD = iota
	CMD_OVD_RPT_INTERVAL
	CMD_OVD_RPT_FRAME
	CMD_OVD_RMT_SEAT
	CMD_OVD_RMT_ALARM
)

type CMDS_AUDIO CMD_SUBCODE

const (
	CMD_AUDIO_BEEP CMDS_AUDIO = iota
)

type CMDS_FGR CMD_SUBCODE

const (
	CMD_FGR_FETCH CMDS_FGR = iota
	CMD_FGR_ADD
	CMD_FGR_DEL
	CMD_FGR_RST
)

type CMDS_RMT CMD_SUBCODE

const (
	CMD_RMT_PAIRING CMDS_RMT = iota
)

type CMDS_FOTA CMD_SUBCODE

const (
	CMD_FOTA_VCU CMDS_FOTA = iota
	CMD_FOTA_HMI
)

type CMDS_NET CMD_SUBCODE

const (
	CMD_NET_SEND_USSD CMDS_NET = iota
	CMD_NET_READ_SMS
)

type CMDS_CON CMD_SUBCODE

const (
	CMD_CON_APN CMDS_CON = iota
	CMD_CON_FTP
	CMD_CON_MQTT
)

type CMDS_HBAR CMD_SUBCODE

const (
	CMD_HBAR_DRIVE CMDS_HBAR = iota
	CMD_HBAR_TRIP
	CMD_HBAR_AVG
	CMD_HBAR_REVERSE
)

type CMDS_MCU CMD_SUBCODE

const (
	CMD_MCU_SPEED_MAX CMDS_MCU = iota
	CMD_MCU_TEMPLATES
)
