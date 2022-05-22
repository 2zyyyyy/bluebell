package snowflake

import (
	"time"

	"go.uber.org/zap"

	sf "github.com/sony/sonyflake"
)

var (
	sonyFlake     *sf.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init 需传入当前的机器ID
func Init(startTime string, machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", startTime)
	setting := sf.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sf.NewSonyflake(setting)
	return
}

// GenID 返回生成的id
func GenID() (id uint64, err error) {
	if sonyFlake == nil {
		zap.L().Error("sony flake not init")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
