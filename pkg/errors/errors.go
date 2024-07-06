/*
 * go-template
 * errors.go
 * This file is part of go-template.
 * Copyright (c) 2023.
 * Last modified at Sun, 24 Dec 2023 21:02:16 -0500 by nick.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any constants, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (that includes, but not limited to, loss of data, data being rendered inaccurate, or losses
 * sustained by you or third parties, or a failure of the software to operate with any other programs), even if the
 * author has been advised of the possibility of such damages.
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions
 * set forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute,
 * or otherwise exploit this software.
 */

package errors

// base is a basic struct for errors
// It implements the Unwrap() method from the errors package
// It is used as a base struct for all internal errors.
type base struct { //nolint:unused // It is used as a base struct for all internal errors.
	code CodeStatus
	err  error `exhaustruct:"optional"`
}

func (e *base) Unwrap() error { //nolint:unused // It is used to unwrap the error.
	return e.err
}
