package dao

import (
	"database/sql"
	"fmt"
	"go-snark/utils"
	"time"
)

const WorkerTableName = "gpu_workers"

func QueryWorker(ipAddr string) (bool, error) {
	var id int
	err := DB.QueryRow(fmt.Sprintf(`SELECT id FROM %s WHERE worker_addr = "%s" LIMIT 1`, WorkerTableName, ipAddr)).Scan(&id)
	if nil != err && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

func ChangeWorkerStatus(ipAddr, field string, status int) error {
	_, err := DB.Exec(fmt.Sprintf(`UPDATE %s SET %s = %d WHERE worker_addr = "%s"`, WorkerTableName, field, status, ipAddr))
	if nil != err {
		return err
	}

	return nil
}

func InsertWorker(gpuType, ipAddr string) error {
	s := fmt.Sprintf(`INSERT INTO %s (gpu_type, worker_addr, is_allocate, is_online, is_up, create_time) VALUES("%s", "%s", %d, %d, %d, "%s")`,
		WorkerTableName, gpuType, ipAddr, 0, 1, 1, time.Now().Format(utils.TimeFmt))
	_, err := DB.Exec(s)

	return err
}

func ChangeTaskStatus(result, miner, ip_addr string, status int, sectorNumber uint64) error {
	tx, err := DB.Begin()
	if nil != err {
		return err
	}
	_, err = tx.Exec(fmt.Sprintf(`UPDATE %s_c2 SET c2_result="%s", c2_status=%d, end_time="%s" WHERE sector_number=%d`,
		miner, result, status, time.Now().Format(utils.TimeFmt), sectorNumber))
	if nil != err {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`UPDATE %s SET is_allocate=%d, update_time="%s" WHERE worker_addr = "%s"`,
		WorkerTableName, 0, time.Now().Format(utils.TimeFmt), ip_addr))
	if nil != err {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	return err
}
