package locker


type LockOperator interface {
	ExecuteLock() (err error)
}
