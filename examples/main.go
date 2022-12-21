package main

import (
	stderr "errors"
	"fmt"

	"github.com/zhongjin616/errors"

	code "github.com/zhongjin616/sample-code"
)

func main() {
	if err := bindUser(); err != nil {
		// %s: Returns the user-safe error string mapped to the error code or the error message if none is specified.
		fmt.Println("====================> %s <====================")
		fmt.Printf("%s\n\n", err)

		// %v: Alias for %s.
		fmt.Println("====================> %v <====================")
		fmt.Printf("%v\n\n", err)

		// %-v: Output caller details, useful for troubleshooting.
		fmt.Println("====================> %-v <====================")
		fmt.Printf("%-v\n\n", err)

		// %+v: Output full error stack details, useful for debugging.
		fmt.Println("====================> %+v <====================")
		fmt.Printf("%+v\n\n", err)

		// %#-v: Output caller details, useful for troubleshooting with JSON formatted output.
		fmt.Println("====================> %#-v <====================")
		fmt.Printf("%#-v\n\n", err)

		// %#+v: Output full error stack details, useful for debugging with JSON formatted output.
		fmt.Println("====================> %#+v <====================")
		fmt.Printf("%#+v\n\n", err)

		// print code and http_status
		fmt.Printf("err_code: %d, http_status: %d\n", errors.Code(err), errors.HTTPStatus(err))

		// do some business process based on the error type
		if !errors.IsCode(err, code.ErrBadRequest) {
			fmt.Println("this is NOT ErrBadRequest error")
		}

		if errors.IsCode(err, code.ErrUserNotFound) {
			fmt.Println("this is ErrUserNotFound error")
		}

		if errors.IsCode(err, code.ErrNoRow) {
			fmt.Println("this is ErrNoRow error")
		}

		// we can also find the cause error
		fmt.Println("root cause: ", errors.Cause(err).Error())
	}
}

func bindUser() error {
	if err := findUser(); err != nil {
		// 在靠近用户侧，追加更多的错误描述
		errors.WithMessage(err, "check your data")
		return err
	}

	return nil
}

func findUser() error {
	if err := getUser(); err != nil {
		errors.WithMessage(err, "staff xxx not found")
		return err
	}

	return nil
}

func getUser() error {
	if err := queryDatabase(); err != nil {
		// 转换为其他业务错误码
		return errors.Spawn(err, code.ErrUserNotFound, "spawn ErrUserNotFound")
	}

	return nil
}

func queryDatabase() error {
	if err := doQuery(); err != nil {
		// 自己的代码报错
		// return errors.New(code.ErrDatabase, "ErrDatabase")

		// 依赖的第三方库报错
		err = errors.Wrap(err, code.ErrNoRow)
		return err
	}

	return nil
}

func doQuery() error {
	return stderr.New("no rows")
}
